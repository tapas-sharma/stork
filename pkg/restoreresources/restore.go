package restorepackage

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/libopenstorage/stork/drivers/volume"
	"github.com/libopenstorage/stork/drivers/volume/kdmp"
	storkapi "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"github.com/libopenstorage/stork/pkg/crypto"
	"github.com/libopenstorage/stork/pkg/k8sutils"
	"github.com/libopenstorage/stork/pkg/log"
	"github.com/libopenstorage/stork/pkg/objectstore"
	"github.com/libopenstorage/stork/pkg/resourcecollector"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"

	"k8s.io/client-go/dynamic"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	validateCRDInterval time.Duration = 5 * time.Second
	validateCRDTimeout  time.Duration = 1 * time.Minute

	resourceObjectName = "resources.json"
	crdObjectName      = "crds.json"
	nsObjectName       = "namespaces.json"
	metadataObjectName = "metadata.json"

	backupCancelBackoffInitialDelay = 5 * time.Second
	backupCancelBackoffFactor       = 1
	backupCancelBackoffSteps        = math.MaxInt32

	allNamespacesSpecifier        = "*"
	backupVolumeBatchCountEnvVar  = "BACKUP-VOLUME-BATCH-COUNT"
	defaultBackupVolumeBatchCount = 3
	backupResourcesBatchCount     = 15
	maxRetry                      = 10
	retrySleep                    = 10 * time.Second
	genericBackupKey              = "BACKUP_TYPE"
)

// NewApplicationRestore creates a new instance of ApplicationRestoreController.
func NewRestoreController(
	client runtimeclient.Client,
	recorder record.EventRecorder,
	resourceCollector resourcecollector.ResourceCollector,
	dynamicInterface dynamic.Interface,
	restoreAdminNamespace string) *RestoreController {
	return &RestoreController{
		client:                client,
		recorder:              recorder,
		resourceCollector:     resourceCollector,
		dynamicInterface:      dynamicInterface,
		restoreAdminNamespace: restoreAdminNamespace,
	}
}

// ApplicationRestoreController reconciles applicationrestore objects
type RestoreController struct {
	client                runtimeclient.Client
	recorder              record.EventRecorder
	resourceCollector     resourcecollector.ResourceCollector
	dynamicInterface      dynamic.Interface
	restoreAdminNamespace string
}

func (r *RestoreController) DownloadApplyResources(restore *storkapi.ApplicationRestore, statusUpdateCR interface{}) error {
	backup, err := storkops.Instance().GetApplicationBackup(restore.Spec.BackupName, restore.Namespace)
	if err != nil {
		log.ApplicationRestoreLog(restore).Errorf("Error getting backup: %v", err)
		return err
	}
	restoreLocation, err := storkops.Instance().GetBackupLocation(backup.Spec.BackupLocation, restore.Namespace)
	if err != nil {
		return err
	}

	if restoreLocation.Location.Type != storkapi.BackupLocationNFS {
		objects, err := r.downloadResources(backup, restore.Spec.BackupLocation, restore.Namespace)
		if err != nil {
			log.ApplicationRestoreLog(restore).Errorf("Error downloading resources: %v", err)
			return err
		}
		if err := r.applyResources(restore, objects); err != nil {
			return err
		}
	} else {
		logrus.Errorf("nfs not implemented yet")
		return fmt.Errorf("nfs not implemented yet")
	}
	return err
}

func (r *RestoreController) downloadResources(
	backup *storkapi.ApplicationBackup,
	backupLocation string,
	namespace string,
) ([]runtime.Unstructured, error) {
	// create CRD resource first
	if err := r.downloadCRD(backup, backupLocation, namespace); err != nil {
		return nil, fmt.Errorf("error downloading CRDs: %v", err)
	}
	data, err := r.downloadObject(backup, backupLocation, namespace, resourceObjectName, false)
	if err != nil {
		return nil, err
	}

	objects := make([]*unstructured.Unstructured, 0)
	if err = json.Unmarshal(data, &objects); err != nil {
		return nil, err
	}
	runtimeObjects := make([]runtime.Unstructured, 0)
	for _, o := range objects {
		runtimeObjects = append(runtimeObjects, o)
	}
	return runtimeObjects, nil
}

func (r *RestoreController) downloadCRD(
	backup *storkapi.ApplicationBackup,
	backupLocation string,
	namespace string,
) error {
	var crds []*apiextensionsv1beta1.CustomResourceDefinition
	var crdsV1 []*apiextensionsv1.CustomResourceDefinition
	crdData, err := r.downloadObject(backup, backupLocation, namespace, crdObjectName, true)
	if err != nil {
		return err
	}
	// No CRDs were uploaded
	if crdData == nil {
		return nil
	}
	if err = json.Unmarshal(crdData, &crds); err != nil {
		return err
	}
	if err = json.Unmarshal(crdData, &crdsV1); err != nil {
		return err
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("error getting cluster config: %v", err)
	}

	client, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return err
	}

	regCrd := make(map[string]bool)
	for _, crd := range crds {
		crd.ResourceVersion = ""
		regCrd[crd.GetName()] = false
		if _, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{}); err != nil && !errors.IsAlreadyExists(err) {
			regCrd[crd.GetName()] = true
			logrus.Warnf("error registering crds v1beta1 %v,%v", crd.GetName(), err)
			continue
		}
		// wait for crd to be ready
		if err := k8sutils.ValidateCRD(client, crd.GetName()); err != nil {
			logrus.Warnf("Unable to validate crds v1beta1 %v,%v", crd.GetName(), err)
		}
	}

	for _, crd := range crdsV1 {
		if val, ok := regCrd[crd.GetName()]; ok && val {
			crd.ResourceVersion = ""
			var updatedVersions []apiextensionsv1.CustomResourceDefinitionVersion
			// try to apply as v1 crd
			var err error
			if _, err = client.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{}); err == nil || errors.IsAlreadyExists(err) {
				logrus.Infof("registered v1 crds %v,", crd.GetName())
				continue
			}
			// updated fields
			crd.Spec.PreserveUnknownFields = false
			for _, version := range crd.Spec.Versions {
				isTrue := true
				if version.Schema == nil {
					openAPISchema := &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{XPreserveUnknownFields: &isTrue},
					}
					version.Schema = openAPISchema
				} else {
					version.Schema.OpenAPIV3Schema.XPreserveUnknownFields = &isTrue
				}
				updatedVersions = append(updatedVersions, version)
			}
			crd.Spec.Versions = updatedVersions

			if _, err := client.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{}); err != nil && !errors.IsAlreadyExists(err) {
				logrus.Warnf("error registering crdsv1 %v,%v", crd.GetName(), err)
				continue
			}
			// wait for crd to be ready
			if err := k8sutils.ValidateCRDV1(client, crd.GetName()); err != nil {
				logrus.Warnf("Unable to validate crdsv1 %v,%v", crd.GetName(), err)
			}

		}
	}

	return nil
}

func (r *RestoreController) downloadObject(
	backup *storkapi.ApplicationBackup,
	backupLocation string,
	namespace string,
	objectName string,
	skipIfNotPresent bool,
) ([]byte, error) {
	restoreLocation, err := storkops.Instance().GetBackupLocation(backup.Spec.BackupLocation, namespace)
	if err != nil {
		return nil, err
	}
	bucket, err := objectstore.GetBucket(restoreLocation)
	if err != nil {
		return nil, err
	}

	objectPath := backup.Status.BackupPath
	if skipIfNotPresent {
		exists, err := bucket.Exists(context.TODO(), filepath.Join(objectPath, objectName))
		if err != nil || !exists {
			return nil, nil
		}
	}

	data, err := bucket.ReadAll(context.TODO(), filepath.Join(objectPath, objectName))
	if err != nil {
		return nil, err
	}
	if restoreLocation.Location.EncryptionKey != "" {
		return nil, fmt.Errorf("EncryptionKey is deprecated, use EncryptionKeyV2 instead")
	}
	if restoreLocation.Location.EncryptionV2Key != "" {
		var decryptData []byte
		if decryptData, err = crypto.Decrypt(data, restoreLocation.Location.EncryptionV2Key); err != nil {
			logrus.Errorf("RestoreController/downloadObject: decrypt failed :%v, returing data direclty", err)
			return data, nil
		}
		return decryptData, nil
	}

	return data, nil
}

func (r *RestoreController) applyResources(
	restore *storkapi.ApplicationRestore,
	objects []runtime.Unstructured,
) error {
	pvNameMappings, err := r.getPVNameMappings(restore, objects)
	if err != nil {
		return err
	}
	objectMap := storkapi.CreateObjectsMap(restore.Spec.IncludeResources)
	tempObjects := make([]runtime.Unstructured, 0)
	for _, o := range objects {
		skip, err := r.resourceCollector.PrepareResourceForApply(
			o,
			objects,
			objectMap,
			restore.Spec.NamespaceMapping,
			restore.Spec.StorageClassMapping,
			pvNameMappings,
			restore.Spec.IncludeOptionalResourceTypes,
			restore.Status.Volumes,
		)
		if err != nil {
			return err
		}
		if !skip {
			tempObjects = append(tempObjects, o)
		}
	}
	objects = tempObjects

	// skip CSI PV/PVCs before applying
	objects, err = r.removeCSIVolumesBeforeApply(restore, objects)
	if err != nil {
		return err
	}
	// First delete the existing objects if they exist and replace policy is set
	// to Delete
	if restore.Spec.ReplacePolicy == storkapi.ApplicationRestoreReplacePolicyDelete {
		err = r.resourceCollector.DeleteResources(
			r.dynamicInterface,
			objects)
		if err != nil {
			return err
		}
	}

	for _, o := range objects {
		metadata, err := meta.Accessor(o)
		if err != nil {
			return err
		}
		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return err
		}

		log.ApplicationRestoreLog(restore).Infof("Applying %v %v/%v", objectType.GetKind(), metadata.GetNamespace(), metadata.GetName())
		retained := false

		err = r.resourceCollector.ApplyResource(
			r.dynamicInterface,
			o)
		if err != nil && errors.IsAlreadyExists(err) {
			switch restore.Spec.ReplacePolicy {
			case storkapi.ApplicationRestoreReplacePolicyDelete:
				log.ApplicationRestoreLog(restore).Errorf("Error deleting %v %v during restore: %v", objectType.GetKind(), metadata.GetName(), err)
			case storkapi.ApplicationRestoreReplacePolicyRetain:
				log.ApplicationRestoreLog(restore).Warningf("Error deleting %v %v during restore, ReplacePolicy set to Retain: %v", objectType.GetKind(), metadata.GetName(), err)
				retained = true
				err = nil
			}
		}

		if err != nil {
			if err := r.updateResourceStatus(
				restore,
				o,
				storkapi.ApplicationRestoreStatusFailed,
				fmt.Sprintf("Error applying resource: %v", err),
				restore); err != nil {
				return err
			}
		} else if retained {
			if err := r.updateResourceStatus(
				restore,
				o,
				storkapi.ApplicationRestoreStatusRetained,
				"Resource restore skipped as it was already present and ReplacePolicy is set to Retain",
				restore); err != nil {
				return err
			}
		} else {
			if err := r.updateResourceStatus(
				restore,
				o,
				storkapi.ApplicationRestoreStatusSuccessful,
				"Resource restored successfully",
				restore); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RestoreController) getPVNameMappings(
	restore *storkapi.ApplicationRestore,
	objects []runtime.Unstructured,
) (map[string]string, error) {
	pvNameMappings := make(map[string]string)
	for _, vInfo := range restore.Status.Volumes {
		if vInfo.SourceVolume == "" {
			return nil, fmt.Errorf("SourceVolume missing for restore")
		}
		if vInfo.RestoreVolume == "" {
			return nil, fmt.Errorf("RestoreVolume missing for restore")
		}
		pvNameMappings[vInfo.SourceVolume] = vInfo.RestoreVolume
	}
	return pvNameMappings, nil
}

func (a *RestoreController) removeCSIVolumesBeforeApply(
	restore *storkapi.ApplicationRestore,
	objects []runtime.Unstructured,
) ([]runtime.Unstructured, error) {
	tempObjects := make([]runtime.Unstructured, 0)
	// Get PVC to PV mapping first for checking if a PVC is bound to a generic CSI PV
	pvcToPVMapping, err := getPVCToPVMapping(objects)
	if err != nil {
		return nil, fmt.Errorf("failed to get PVC to PV mapping: %v", err)
	}
	for _, o := range objects {
		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return nil, err
		}

		switch objectType.GetKind() {
		case "PersistentVolume":
			// check if this PV is a generic CSI one
			var pv v1.PersistentVolume
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), &pv); err != nil {
				return nil, fmt.Errorf("error converting to persistent volume: %v", err)
			}

			// Check if this PV is a generic CSI one
			isGenericCSIPVC, err := isGenericCSIPersistentVolume(&pv)
			if err != nil {
				return nil, fmt.Errorf("failed to check if PV was provisioned by a CSI driver: %v", err)
			}
			isGenericDriverPV, err := isGenericPersistentVolume(&pv, restore.Status.Volumes)
			if err != nil {
				return nil, err
			}
			// Only add this object if it's not a generic CSI PV
			if !isGenericCSIPVC && !isGenericDriverPV {
				tempObjects = append(tempObjects, o)
			} else {
				log.ApplicationRestoreLog(restore).Debugf("skipping CSI PV in restore: %s", pv.Name)
			}

		case "PersistentVolumeClaim":
			// check if this PVC is a generic CSI one
			var pvc v1.PersistentVolumeClaim
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), &pvc); err != nil {
				return nil, fmt.Errorf("error converting PVC object: %v: %v", o, err)
			}

			// Find the matching PV for this PVC
			pv, ok := pvcToPVMapping[getNamespacedPVCLocation(&pvc)]
			if !ok {
				log.ApplicationRestoreLog(restore).Debugf("failed to find PV for PVC %s during CSI volume skip. Will not skip volume", pvc.Name)
				tempObjects = append(tempObjects, o)
				continue
			}

			// We have found a PV for this PVC. Check if it is a generic CSI PV
			// that we do not already have native volume driver support for.
			isGenericCSIPVC, err := isGenericCSIPersistentVolume(pv)
			if err != nil {
				return nil, err
			}
			isGenericDriverPVC, err := isGenericCSIPersistentVolumeClaim(&pvc, restore.Status.Volumes)
			if err != nil {
				return nil, err
			}

			// Only add this object if it's not a generic CSI PVC
			if !isGenericCSIPVC && !isGenericDriverPVC {
				tempObjects = append(tempObjects, o)
			} else {
				log.ApplicationRestoreLog(restore).Debugf("skipping PVC in restore: %s", pvc.Name)
			}

		default:
			// add all other objects
			tempObjects = append(tempObjects, o)
		}
	}

	return tempObjects, nil
}

func getNamespacedPVCLocation(pvc *v1.PersistentVolumeClaim) string {
	return fmt.Sprintf("%s/%s", pvc.Namespace, pvc.Name)
}

func isGenericCSIPersistentVolume(pv *v1.PersistentVolume) (bool, error) {
	driverName, err := volume.GetPVDriver(pv)
	if err != nil {
		return false, err
	}
	if driverName == "csi" {
		return true, nil
	}
	return false, nil
}
func isGenericPersistentVolume(pv *v1.PersistentVolume, volInfos []*storkapi.ApplicationRestoreVolumeInfo) (bool, error) {
	for _, vol := range volInfos {
		if vol.DriverName == kdmp.GetGenericDriverName() && vol.RestoreVolume == pv.Name {
			return true, nil
		}
	}
	return false, nil
}

func isGenericCSIPersistentVolumeClaim(pvc *v1.PersistentVolumeClaim, volInfos []*storkapi.ApplicationRestoreVolumeInfo) (bool, error) {
	for _, vol := range volInfos {
		if vol.DriverName == kdmp.GetGenericDriverName() && vol.PersistentVolumeClaim == pvc.Name {
			return true, nil
		}
	}
	return false, nil
}

// getPVCToPVMapping constructs a mapping of PVC name/namespace to PV objects
func getPVCToPVMapping(allObjects []runtime.Unstructured) (map[string]*v1.PersistentVolume, error) {

	// Get mapping of PVC name to PV name
	pvNameToPVCName := make(map[string]string)
	for _, o := range allObjects {
		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return nil, err
		}

		// If a PV, assign it to the mapping based on the claimRef UID
		if objectType.GetKind() == "PersistentVolumeClaim" {
			pvc := &v1.PersistentVolumeClaim{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), pvc); err != nil {
				return nil, fmt.Errorf("error converting to persistent volume: %v", err)
			}

			pvNameToPVCName[pvc.Spec.VolumeName] = getNamespacedPVCLocation(pvc)
		}
	}

	// Get actual mapping of PVC name to PV object
	pvcNameToPV := make(map[string]*v1.PersistentVolume)
	for _, o := range allObjects {
		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return nil, err
		}

		// If a PV, assign it to the mapping based on the claimRef UID
		if objectType.GetKind() == "PersistentVolume" {
			pv := &v1.PersistentVolume{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), pv); err != nil {
				return nil, fmt.Errorf("error converting to persistent volume: %v", err)
			}

			pvcName := pvNameToPVCName[pv.Name]

			// add this PVC name/PV obj mapping
			pvcNameToPV[pvcName] = pv
		}
	}

	return pvcNameToPV, nil
}

func (a *RestoreController) updateResourceStatus(
	restore *storkapi.ApplicationRestore,
	object runtime.Unstructured,
	status storkapi.ApplicationRestoreStatusType,
	reason string,
	updateCR interface{},
) error {
	var updatedResource *storkapi.ApplicationRestoreResourceInfo
	gkv := object.GetObjectKind().GroupVersionKind()
	metadata, err := meta.Accessor(object)
	if err != nil {
		log.ApplicationRestoreLog(restore).Errorf("Error getting metadata for object %v %v", object, err)
		return err
	}

	switch v := updateCR.(type) {
	case *storkapi.ApplicationRestore:
		for _, resource := range v.Status.Resources {
			if resource.Name == metadata.GetName() &&
				resource.Namespace == metadata.GetNamespace() &&
				(resource.Group == gkv.Group || (resource.Group == "core" && gkv.Group == "")) &&
				resource.Version == gkv.Version &&
				resource.Kind == gkv.Kind {
				updatedResource = resource
				break
			}
		}
		if updatedResource == nil {
			updatedResource = &storkapi.ApplicationRestoreResourceInfo{
				ObjectInfo: storkapi.ObjectInfo{
					Name:      metadata.GetName(),
					Namespace: metadata.GetNamespace(),
					GroupVersionKind: metav1.GroupVersionKind{
						Group:   gkv.Group,
						Version: gkv.Version,
						Kind:    gkv.Kind,
					},
				},
			}
			v.Status.Resources = append(v.Status.Resources, updatedResource)
		}
	default:
		log.ApplicationRestoreLog(restore).Errorf("Invalid CR for resource updating %v", updateCR)
		return fmt.Errorf("invalid CR for resource updating %v", updateCR)

	}

	updatedResource.Status = status
	updatedResource.Reason = reason
	eventType := v1.EventTypeNormal
	if status == storkapi.ApplicationRestoreStatusFailed {
		eventType = v1.EventTypeWarning
	}
	eventMessage := fmt.Sprintf("%v %v/%v: %v",
		gkv,
		updatedResource.Namespace,
		updatedResource.Name,
		reason)
	a.recorder.Event(restore, eventType, string(status), eventMessage)
	return nil
}
