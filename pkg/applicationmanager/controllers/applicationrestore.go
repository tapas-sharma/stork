package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/libopenstorage/stork/drivers/volume"
	"github.com/libopenstorage/stork/pkg/apis/stork"
	storkapi "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"github.com/libopenstorage/stork/pkg/controllers"
	"github.com/libopenstorage/stork/pkg/k8sutils"
	"github.com/libopenstorage/stork/pkg/log"
	"github.com/libopenstorage/stork/pkg/resourcecollector"
	restorepackage "github.com/libopenstorage/stork/pkg/restoreresources"
	"github.com/libopenstorage/stork/pkg/version"
	"github.com/portworx/sched-ops/k8s/apiextensions"
	"github.com/portworx/sched-ops/k8s/core"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// NewApplicationRestore creates a new instance of ApplicationRestoreController.
func NewApplicationRestore(mgr manager.Manager, r record.EventRecorder, rc resourcecollector.ResourceCollector) *ApplicationRestoreController {
	
	return &ApplicationRestoreController{
		client:            mgr.GetClient(),
		recorder:          r,
		resourceCollector: rc,
	}
}

// ApplicationRestoreController reconciles applicationrestore objects
type ApplicationRestoreController struct {
	client runtimeclient.Client

	recorder              record.EventRecorder
	resourceCollector     resourcecollector.ResourceCollector
	dynamicInterface      dynamic.Interface
	restoreAdminNamespace string
	restoreController     *restorepackage.RestoreController
}

// Init Initialize the application restore controller
func (a *ApplicationRestoreController) Init(mgr manager.Manager, restoreAdminNamespace string) error {
	err := a.createCRD()
	if err != nil {
		return err
	}

	a.restoreAdminNamespace = restoreAdminNamespace

	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("error getting cluster config: %v", err)
	}

	a.dynamicInterface, err = dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	a.restoreController = restorepackage.NewRestoreController(a.client, a.recorder, a.resourceCollector, a.dynamicInterface, restoreAdminNamespace)
	return controllers.RegisterTo(mgr, "application-restore-controller", a, &storkapi.ApplicationRestore{})
}

func (a *ApplicationRestoreController) setDefaults(restore *storkapi.ApplicationRestore) error {
	if restore.Spec.ReplacePolicy == "" {
		restore.Spec.ReplacePolicy = storkapi.ApplicationRestoreReplacePolicyRetain
	}
	// If no namespaces mappings are provided add mappings for all of them
	if len(restore.Spec.NamespaceMapping) == 0 {
		backup, err := storkops.Instance().GetApplicationBackup(restore.Spec.BackupName, restore.Namespace)
		if err != nil {
			return fmt.Errorf("error getting backup: %v", err)
		}
		if restore.Spec.NamespaceMapping == nil {
			restore.Spec.NamespaceMapping = make(map[string]string)
		}
		for _, ns := range backup.Spec.Namespaces {
			restore.Spec.NamespaceMapping[ns] = ns
		}
	}
	return nil
}

func (a *ApplicationRestoreController) verifyNamespaces(restore *storkapi.ApplicationRestore) error {
	// Check whether namespace is allowed to be restored to before each stage
	// Restrict restores to only the namespace that the object belongs
	// except for the namespace designated by the admin
	if !a.namespaceRestoreAllowed(restore) {
		return fmt.Errorf("Spec.Namespaces should only contain the current namespace")
	}
	backup, err := storkops.Instance().GetApplicationBackup(restore.Spec.BackupName, restore.Namespace)
	if err != nil {
		log.ApplicationRestoreLog(restore).Errorf("Error getting backup: %v", err)
		return err
	}
	return a.createNamespaces(backup, restore.Spec.BackupLocation, restore)
}

func (a *ApplicationRestoreController) createNamespaces(backup *storkapi.ApplicationBackup,
	backupLocation string,
	restore *storkapi.ApplicationRestore) error {
	var namespaces []*v1.Namespace

	nsData, err := a.restoreController.DownloadObject(backup, backupLocation, restore.Namespace, nsObjectName, true)
	if err != nil {
		return err
	}
	if nsData != nil {
		if err = json.Unmarshal(nsData, &namespaces); err != nil {
			return err
		}
		for _, ns := range namespaces {
			if restoreNS, ok := restore.Spec.NamespaceMapping[ns.Name]; ok {
				ns.Name = restoreNS
			} else {
				// Skip namespaces we aren't restoring
				continue
			}
			// create mapped restore namespace with metadata of backed up
			// namespace
			_, err := core.Instance().CreateNamespace(&v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        ns.Name,
					Labels:      ns.Labels,
					Annotations: ns.GetAnnotations(),
				},
			})
			log.ApplicationRestoreLog(restore).Tracef("Creating dest namespace %v", ns.Name)
			if err != nil {
				if errors.IsAlreadyExists(err) {
					oldNS, err := core.Instance().GetNamespace(ns.GetName())
					if err != nil {
						return err
					}
					annotations := make(map[string]string)
					labels := make(map[string]string)
					if restore.Spec.ReplacePolicy == storkapi.ApplicationRestoreReplacePolicyDelete {
						// overwrite all annotation in case of replace policy set to delete
						annotations = ns.GetAnnotations()
						labels = ns.GetLabels()
					} else if restore.Spec.ReplacePolicy == storkapi.ApplicationRestoreReplacePolicyRetain {
						// only add new annotation,labels in case of replace policy is set to retain
						annotations = oldNS.GetAnnotations()
						if annotations == nil {
							annotations = make(map[string]string)
						}
						for k, v := range ns.GetAnnotations() {
							if _, ok := annotations[k]; !ok {
								annotations[k] = v
							}
						}
						labels = oldNS.GetLabels()
						if labels == nil {
							labels = make(map[string]string)
						}
						for k, v := range ns.GetLabels() {
							if _, ok := labels[k]; !ok {
								labels[k] = v
							}
						}
					}
					log.ApplicationRestoreLog(restore).Tracef("Namespace already exists, updating dest namespace %v", ns.Name)
					_, err = core.Instance().UpdateNamespace(&v1.Namespace{
						ObjectMeta: metav1.ObjectMeta{
							Name:        ns.Name,
							Labels:      labels,
							Annotations: annotations,
						},
					})
					if err != nil {
						return err
					}
					continue
				}
				return err
			}
		}
		return nil
	}
	for _, namespace := range restore.Spec.NamespaceMapping {
		if ns, err := core.Instance().GetNamespace(namespace); err != nil {
			if errors.IsNotFound(err) {
				if _, err := core.Instance().CreateNamespace(&v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name:        ns.Name,
						Labels:      ns.Labels,
						Annotations: ns.GetAnnotations(),
					},
				}); err != nil {
					return err
				}
			}
			return err
		}
	}
	return nil
}

// Reconcile updates for ApplicationRestore objects.
func (a *ApplicationRestoreController) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	logrus.Infof("Reconciling ApplicationRestore %s/%s", request.Namespace, request.Name)

	// Fetch the ApplicationBackup instance
	restore := &storkapi.ApplicationRestore{}
	err := a.client.Get(context.TODO(), request.NamespacedName, restore)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{RequeueAfter: controllers.DefaultRequeueError}, err
	}

	if !controllers.ContainsFinalizer(restore, controllers.FinalizerCleanup) {
		controllers.SetFinalizer(restore, controllers.FinalizerCleanup)
		return reconcile.Result{Requeue: true}, a.client.Update(context.TODO(), restore)
	}

	if err = a.handle(context.TODO(), restore); err != nil && err != errResourceBusy {
		logrus.Errorf("%s: %s/%s: %s", reflect.TypeOf(a), restore.Namespace, restore.Name, err)
		return reconcile.Result{RequeueAfter: controllers.DefaultRequeueError}, err
	}

	return reconcile.Result{RequeueAfter: controllers.DefaultRequeue}, nil
}

// Handle updates for ApplicationRestore objects
func (a *ApplicationRestoreController) handle(ctx context.Context, restore *storkapi.ApplicationRestore) error {
	if restore.DeletionTimestamp != nil {
		if controllers.ContainsFinalizer(restore, controllers.FinalizerCleanup) {
			if err := a.cleanupRestore(restore); err != nil {
				logrus.Errorf("%s: cleanup: %s", reflect.TypeOf(a), err)
			}
		}

		if restore.GetFinalizers() != nil {
			controllers.RemoveFinalizer(restore, controllers.FinalizerCleanup)
			return a.client.Update(ctx, restore)
		}

		return nil
	}

	err := a.setDefaults(restore)
	if err != nil {
		log.ApplicationRestoreLog(restore).Errorf(err.Error())
		a.recorder.Event(restore,
			v1.EventTypeWarning,
			string(storkapi.ApplicationRestoreStatusFailed),
			err.Error())
		return nil
	}

	err = a.verifyNamespaces(restore)
	if err != nil {
		log.ApplicationRestoreLog(restore).Errorf(err.Error())
		a.recorder.Event(restore,
			v1.EventTypeWarning,
			string(storkapi.ApplicationRestoreStatusFailed),
			err.Error())
		return nil
	}

	switch restore.Status.Stage {
	case storkapi.ApplicationRestoreStageInitial:
		// Make sure the namespaces exist
		fallthrough
	case storkapi.ApplicationRestoreStageVolumes:
		err := a.restoreVolumes(restore)
		if err != nil {
			message := fmt.Sprintf("Error restoring volumes: %v", err)
			log.ApplicationRestoreLog(restore).Errorf(message)
			a.recorder.Event(restore,
				v1.EventTypeWarning,
				string(storkapi.ApplicationRestoreStatusFailed),
				message)
			if _, ok := err.(*volume.ErrStorageProviderBusy); ok {
				return errResourceBusy
			}
			return nil
		}
	case storkapi.ApplicationRestoreStageApplications:
		err := a.restoreResources(restore)
		if err != nil {
			message := fmt.Sprintf("Error restoring resources: %v", err)
			log.ApplicationRestoreLog(restore).Errorf(message)
			a.recorder.Event(restore,
				v1.EventTypeWarning,
				string(storkapi.ApplicationRestoreStatusFailed),
				message)
			return nil
		}

	case storkapi.ApplicationRestoreStageFinal:
		// DoNothing
		return nil
	default:
		log.ApplicationRestoreLog(restore).Errorf("Invalid stage for restore: %v", restore.Status.Stage)
	}

	return nil
}

func (a *ApplicationRestoreController) namespaceRestoreAllowed(restore *storkapi.ApplicationRestore) bool {
	// Restrict restores to only the namespace that the object belongs
	// except for the namespace designated by the admin
	if restore.Namespace != a.restoreAdminNamespace {
		for _, ns := range restore.Spec.NamespaceMapping {
			if ns != restore.Namespace {
				return false
			}
		}
	}
	return true
}

func (a *ApplicationRestoreController) getDriversForRestore(restore *storkapi.ApplicationRestore) map[string]bool {
	drivers := make(map[string]bool)
	for _, volumeInfo := range restore.Status.Volumes {
		drivers[volumeInfo.DriverName] = true
	}
	return drivers
}

func (a *ApplicationRestoreController) getNamespacedObjectsToDelete(restore *storkapi.ApplicationRestore, objects []runtime.Unstructured) ([]runtime.Unstructured, error) {
	tempObjects := make([]runtime.Unstructured, 0)
	for _, o := range objects {
		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return nil, err
		}

		// Skip PVs, we will let the PVC handle PV deletion where needed
		if objectType.GetKind() != "PersistentVolume" {
			tempObjects = append(tempObjects, o)
		}
	}

	return tempObjects, nil
}

func (a *ApplicationRestoreController) updateRestoreCRInVolumeStage(
	namespacedName types.NamespacedName,
	status storkapi.ApplicationRestoreStatusType,
	stage storkapi.ApplicationRestoreStageType,
	reason string,
	volumeInfos []*storkapi.ApplicationRestoreVolumeInfo,
) (*storkapi.ApplicationRestore, error) {
	restore := &storkapi.ApplicationRestore{}
	var err error
	for i := 0; i < maxRetry; i++ {
		err := a.client.Get(context.TODO(), namespacedName, restore)
		if err != nil {
			time.Sleep(retrySleep)
			continue
		}
		if restore.Status.Stage == storkapi.ApplicationRestoreStageFinal ||
			restore.Status.Stage == storkapi.ApplicationRestoreStageApplications {
			// updated timestamp for failed restores
			if restore.Status.Status == storkapi.ApplicationRestoreStatusFailed {
				restore.Status.FinishTimestamp = metav1.Now()
				restore.Status.LastUpdateTimestamp = metav1.Now()
				restore.Status.Reason = reason
			}
			return restore, nil
		}
		logrus.Infof("Updating restore  %s/%s in stage/stagus: %s/%s to volume stage", restore.Namespace, restore.Name, restore.Status.Stage, restore.Status.Status)
		restore.Status.Status = status
		restore.Status.Stage = stage
		restore.Status.Reason = reason
		restore.Status.LastUpdateTimestamp = metav1.Now()
		if volumeInfos != nil {
			restore.Status.Volumes = append(restore.Status.Volumes, volumeInfos...)
		}
		err = a.client.Update(context.TODO(), restore)
		if err != nil {
			time.Sleep(retrySleep)
			continue
		} else {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return restore, nil
}

func (a *ApplicationRestoreController) restoreVolumes(restore *storkapi.ApplicationRestore) error {
	restore.Status.Stage = storkapi.ApplicationRestoreStageVolumes

	backup, err := storkops.Instance().GetApplicationBackup(restore.Spec.BackupName, restore.Namespace)
	if err != nil {
		return fmt.Errorf("error getting backup spec for restore: %v", err)
	}
	objectMap := storkapi.CreateObjectsMap(restore.Spec.IncludeResources)
	info := storkapi.ObjectInfo{
		GroupVersionKind: metav1.GroupVersionKind{
			Group:   "core",
			Version: "v1",
			Kind:    "PersistentVolumeClaim",
		},
	}

	pvcCount := 0
	restoreDone := 0
	backupVolumeInfoMappings := make(map[string][]*storkapi.ApplicationBackupVolumeInfo)
	for _, namespace := range backup.Spec.Namespaces {
		if _, ok := restore.Spec.NamespaceMapping[namespace]; !ok {
			continue
		}
		for _, volumeBackup := range backup.Status.Volumes {
			if volumeBackup.Namespace != namespace {
				continue
			}
			// If a list of resources was specified during restore check if
			// this PVC was included
			info.Name = volumeBackup.PersistentVolumeClaim
			info.Namespace = volumeBackup.Namespace
			if len(objectMap) != 0 {
				if val, present := objectMap[info]; !present || !val {
					continue
				}
			}

			pvcCount++
			isVolRestoreDone := false
			for _, statusVolume := range restore.Status.Volumes {
				if statusVolume.SourceVolume == volumeBackup.Volume {
					isVolRestoreDone = true
					break
				}
			}
			if isVolRestoreDone {
				restoreDone++
				continue
			}

			if volumeBackup.DriverName == "" {
				volumeBackup.DriverName = volume.GetDefaultDriverName()
			}
			if backupVolumeInfoMappings[volumeBackup.DriverName] == nil {
				backupVolumeInfoMappings[volumeBackup.DriverName] = make([]*storkapi.ApplicationBackupVolumeInfo, 0)
			}
			backupVolumeInfoMappings[volumeBackup.DriverName] = append(backupVolumeInfoMappings[volumeBackup.DriverName], volumeBackup)
		}
	}
	if restore.Status.Volumes == nil {
		restore.Status.Volumes = make([]*storkapi.ApplicationRestoreVolumeInfo, 0)
	}
	namespacedName := types.NamespacedName{}
	namespacedName.Namespace = restore.Namespace
	namespacedName.Name = restore.Name
	restoreCompleteList := make([]*storkapi.ApplicationRestoreVolumeInfo, 0)
	if len(restore.Status.Volumes) != pvcCount {
		for driverName, vInfos := range backupVolumeInfoMappings {
			backupVolInfos := vInfos
			existingRestoreVolInfos := make([]*storkapi.ApplicationRestoreVolumeInfo, 0)
			driver, err := volume.Get(driverName)
			if err != nil {
				return err
			}

			// For each driver, check if it needs any additional resources to be
			// restored before starting the volume restore
			objects, err := a.restoreController.DownloadResources(backup, restore.Spec.BackupLocation, restore.Namespace)
			if err != nil {
				log.ApplicationRestoreLog(restore).Errorf("Error downloading resources: %v", err)
				return err
			}

			// Skip pv/pvc if replacepolicy is set to retain to avoid creating
			if restore.Spec.ReplacePolicy == storkapi.ApplicationRestoreReplacePolicyRetain {
				backupVolInfos, existingRestoreVolInfos, err = a.skipVolumesFromRestoreList(restore, objects, driver, vInfos)
				if err != nil {
					log.ApplicationRestoreLog(restore).Errorf("Error while checking pvcs: %v", err)
					return err
				}
			}

			preRestoreObjects, err := driver.GetPreRestoreResources(backup, restore, objects)
			if err != nil {
				log.ApplicationRestoreLog(restore).Errorf("Error getting PreRestore Resources: %v", err)
				return err
			}

			// pvc creation is not part of kdmp
			if driverName != "kdmp" {
				if err := a.restoreController.ApplyResources(restore, preRestoreObjects); err != nil {
					return err
				}
			}

			// Pre-delete resources for CSI driver
			if (driverName == "csi" || driverName == "kdmp") && restore.Spec.ReplacePolicy == storkapi.ApplicationRestoreReplacePolicyDelete {
				objectMap := storkapi.CreateObjectsMap(restore.Spec.IncludeResources)
				objectBasedOnIncludeResources := make([]runtime.Unstructured, 0)
				for _, o := range objects {
					skip, err := a.resourceCollector.PrepareResourceForApply(
						o,
						objects,
						objectMap,
						restore.Spec.NamespaceMapping,
						nil, // no need to set storage class mappings at this stage
						nil,
						restore.Spec.IncludeOptionalResourceTypes,
						nil,
					)
					if err != nil {
						return err
					}
					if !skip {
						objectBasedOnIncludeResources = append(
							objectBasedOnIncludeResources,
							o,
						)
					}
				}
				tempObjects, err := a.getNamespacedObjectsToDelete(
					restore,
					objectBasedOnIncludeResources,
				)
				if err != nil {
					return err
				}
				err = a.resourceCollector.DeleteResources(
					a.dynamicInterface,
					tempObjects)
				if err != nil {
					return err
				}
			}

			restoreCompleteList = append(restoreCompleteList, existingRestoreVolInfos...)
			restoreVolumeInfos, err := driver.StartRestore(restore, backupVolInfos, preRestoreObjects)
			if err != nil {
				message := fmt.Sprintf("Error starting Application Restore for volumes: %v", err)
				log.ApplicationRestoreLog(restore).Errorf(message)
				if _, ok := err.(*volume.ErrStorageProviderBusy); ok {
					msg := fmt.Sprintf("Volume restores are in progress. Restores are failing for some volumes"+
						" since the storage provider is busy. Restore will be retried. Error: %v", err)
					a.recorder.Event(restore,
						v1.EventTypeWarning,
						string(storkapi.ApplicationRestoreStatusInProgress),
						msg)

					log.ApplicationRestoreLog(restore).Errorf(msg)
					// Update the restore status even for failed restores when storage is busy
					_, updateErr := a.updateRestoreCRInVolumeStage(
						namespacedName,
						storkapi.ApplicationRestoreStatusInProgress,
						storkapi.ApplicationRestoreStageVolumes,
						msg,
						restoreVolumeInfos,
					)
					if updateErr != nil {
						logrus.Warnf("failed to update restore status: %v", updateErr)
					}
					return err
				}
				a.recorder.Event(restore,
					v1.EventTypeWarning,
					string(storkapi.ApplicationRestoreStatusFailed),
					message)
				_, err = a.updateRestoreCRInVolumeStage(namespacedName, storkapi.ApplicationRestoreStatusFailed, storkapi.ApplicationRestoreStageFinal, message, nil)
				return err
			}
			restoreCompleteList = append(restoreCompleteList, restoreVolumeInfos...)
		}
		restore, err = a.updateRestoreCRInVolumeStage(
			namespacedName,
			storkapi.ApplicationRestoreStatusInProgress,
			storkapi.ApplicationRestoreStageVolumes,
			"Volume restores are in progress",
			restoreCompleteList,
		)
		if err != nil {
			return err
		}
	}
	inProgress := false
	// Skip checking status if no volumes are being restored
	if len(restore.Status.Volumes) != 0 {
		drivers := a.getDriversForRestore(restore)
		volumeInfos := make([]*storkapi.ApplicationRestoreVolumeInfo, 0)

		var err error
		for driverName := range drivers {
			driver, err := volume.Get(driverName)
			if err != nil {
				return err
			}

			status, err := driver.GetRestoreStatus(restore)
			if err != nil {
				return fmt.Errorf("error getting restore status for driver %v: %v", driverName, err)
			}
			volumeInfos = append(volumeInfos, status...)
		}

		restore.Status.Volumes = volumeInfos
		restore.Status.LastUpdateTimestamp = metav1.Now()
		// Store the new status
		err = a.client.Update(context.TODO(), restore)
		if err != nil {
			return err
		}

		// Now check if there is any failure or success
		// TODO: On failure of one volume cancel other restores?
		for _, vInfo := range volumeInfos {
			if vInfo.Status == storkapi.ApplicationRestoreStatusInProgress || vInfo.Status == storkapi.ApplicationRestoreStatusInitial ||
				vInfo.Status == storkapi.ApplicationRestoreStatusPending {
				log.ApplicationRestoreLog(restore).Infof("Volume restore still in progress: %v->%v", vInfo.SourceVolume, vInfo.RestoreVolume)
				inProgress = true
			} else if vInfo.Status == storkapi.ApplicationRestoreStatusFailed {
				a.recorder.Event(restore,
					v1.EventTypeWarning,
					string(vInfo.Status),
					fmt.Sprintf("Error restoring volume %v->%v: %v", vInfo.SourceVolume, vInfo.RestoreVolume, vInfo.Reason))
				restore.Status.Stage = storkapi.ApplicationRestoreStageFinal
				restore.Status.FinishTimestamp = metav1.Now()
				restore.Status.Status = storkapi.ApplicationRestoreStatusFailed
				restore.Status.Reason = vInfo.Reason
				break
			} else if vInfo.Status == storkapi.ApplicationRestoreStatusSuccessful {
				a.recorder.Event(restore,
					v1.EventTypeNormal,
					string(vInfo.Status),
					fmt.Sprintf("Volume %v->%v restored successfully", vInfo.SourceVolume, vInfo.RestoreVolume))
			}
		}
	}

	// Return if we have any volume restores still in progress
	if inProgress || len(restore.Status.Volumes) != pvcCount {
		return nil
	}

	// If the restore hasn't failed move on to the next stage.
	if restore.Status.Status != storkapi.ApplicationRestoreStatusFailed {
		restore.Status.Stage = storkapi.ApplicationRestoreStageApplications
		restore.Status.Status = storkapi.ApplicationRestoreStatusInProgress
		restore.Status.Reason = "Application resources restore is in progress"
		restore.Status.LastUpdateTimestamp = metav1.Now()
		// Update the current state and then move on to restoring resources
		err := a.client.Update(context.TODO(), restore)
		if err != nil {
			return err
		}
		err = a.restoreResources(restore)
		if err != nil {
			log.ApplicationRestoreLog(restore).Errorf("Error restoring resources: %v", err)
			return err
		}
	}

	restore.Status.LastUpdateTimestamp = metav1.Now()
	// Only on success compute the total restore size
	for _, vInfo := range restore.Status.Volumes {
		restore.Status.TotalSize += vInfo.TotalSize
	}

	err = a.client.Update(context.TODO(), restore)
	if err != nil {
		return err
	}
	return nil
}

func (a *ApplicationRestoreController) skipVolumesFromRestoreList(
	restore *storkapi.ApplicationRestore,
	objects []runtime.Unstructured,
	driver volume.Driver,
	volInfo []*storkapi.ApplicationBackupVolumeInfo,
) ([]*storkapi.ApplicationBackupVolumeInfo, []*storkapi.ApplicationRestoreVolumeInfo, error) {
	existingInfos := make([]*storkapi.ApplicationRestoreVolumeInfo, 0)
	newVolInfos := make([]*storkapi.ApplicationBackupVolumeInfo, 0)
	for _, bkupVolInfo := range volInfo {
		restoreVolInfo := &storkapi.ApplicationRestoreVolumeInfo{}
		val, ok := restore.Spec.NamespaceMapping[bkupVolInfo.Namespace]
		if !ok {
			logrus.Infof("skipping namespace %s for restore", bkupVolInfo.Namespace)
			continue
		}

		// get corresponding pvc object from objects list
		pvcObject, err := volume.GetPVCFromObjects(objects, bkupVolInfo)
		if err != nil {
			return newVolInfos, existingInfos, err
		}

		ns := val
		pvc, err := core.Instance().GetPersistentVolumeClaim(pvcObject.Name, ns)
		if err != nil {
			if errors.IsNotFound(err) {
				newVolInfos = append(newVolInfos, bkupVolInfo)
				continue
			}
			return newVolInfos, existingInfos, fmt.Errorf("erorr getting pvc %s/%s: %v", ns, pvcObject.Name, err)
		}

		restoreVolInfo.PersistentVolumeClaim = bkupVolInfo.PersistentVolumeClaim
		restoreVolInfo.PersistentVolumeClaimUID = bkupVolInfo.PersistentVolumeClaimUID
		restoreVolInfo.SourceNamespace = bkupVolInfo.Namespace
		restoreVolInfo.SourceVolume = bkupVolInfo.Volume
		restoreVolInfo.DriverName = driver.String()
		restoreVolInfo.Status = storkapi.ApplicationRestoreStatusRetained
		restoreVolInfo.RestoreVolume = pvc.Spec.VolumeName
		restoreVolInfo.TotalSize = bkupVolInfo.TotalSize
		restoreVolInfo.Reason = fmt.Sprintf("Skipped from volume restore as policy is set to %s and pvc already exists", storkapi.ApplicationRestoreReplacePolicyRetain)
		existingInfos = append(existingInfos, restoreVolInfo)
	}

	return newVolInfos, existingInfos, nil
}

func (a *ApplicationRestoreController) restoreResources(
	restore *storkapi.ApplicationRestore,
) error {
	// check if the backuplocaion is nfs
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
		r := restorepackage.NewRestoreController(a.client, a.recorder, a.resourceCollector, a.dynamicInterface, a.restoreAdminNamespace)

		err = r.DownloadApplyResources(restore, nil)
		if err != nil {
			return err
		}
	} else {
		logrus.Infof("Job related implementation need to be done here")
	}

	if restoreLocation.Location.Type == storkapi.BackupLocationNFS {
		logrus.Infof("Updating resource status to application restore")
		//updateStatusInRestoreCRFromResourceExportCR()
	}

	// Before  updating to final stage, cleanup generic backup CRs, if any.
	err = a.cleanupResources(restore)
	if err != nil {
		return err
	}

	// Add all CSI PVCs and PVs back into resources.
	// CSI PVs are dynamically generated by the CSI controller for restore,
	// so we need to get the new PV name after restore volumes finishes
	if err := a.addCSIVolumeResources(restore); err != nil {
		return err
	}

	restore.Status.Stage = storkapi.ApplicationRestoreStageFinal
	restore.Status.FinishTimestamp = metav1.Now()
	restore.Status.Status = storkapi.ApplicationRestoreStatusSuccessful
	restore.Status.Reason = "Volumes and resources were restored up successfully"
	for _, resource := range restore.Status.Resources {
		if resource.Status != storkapi.ApplicationRestoreStatusSuccessful {
			restore.Status.Status = storkapi.ApplicationRestoreStatusPartialSuccess
			restore.Status.Reason = "Volumes were restored successfully. Some existing resources were not replaced"
			break
		}
	}

	restore.Status.LastUpdateTimestamp = metav1.Now()
	if err := a.client.Update(context.TODO(), restore); err != nil {
		return err
	}

	return nil
}

func (a *ApplicationRestoreController) addCSIVolumeResources(restore *storkapi.ApplicationRestore) error {
	for _, vrInfo := range restore.Status.Volumes {
		if vrInfo.DriverName != "csi" && vrInfo.DriverName != "kdmp" {
			continue
		}

		// Update PV resource for this volume
		pv, err := core.Instance().GetPersistentVolume(vrInfo.RestoreVolume)
		if err != nil {
			return fmt.Errorf("failed to get PV %s: %v", vrInfo.RestoreVolume, err)
		}
		pvContent, err := runtime.DefaultUnstructuredConverter.ToUnstructured(pv)
		if err != nil {
			return fmt.Errorf("failed to convert PV %s to unstructured: %v", vrInfo.RestoreVolume, err)
		}
		pvObj := &unstructured.Unstructured{}
		pvObj.SetUnstructuredContent(pvContent)
		pvObj.SetGroupVersionKind(schema.GroupVersionKind{
			Kind:    "PersistentVolume",
			Version: "v1",
			Group:   "core",
		})
		if err := a.restoreController.UpdateResourceStatus(
			restore,
			pvObj,
			vrInfo.Status,
			"Resource restored successfully",
			restore); err != nil {
			return err
		}

		// Update PVC resource for this volume
		ns, ok := restore.Spec.NamespaceMapping[vrInfo.SourceNamespace]
		if !ok {
			ns = vrInfo.SourceNamespace
		}
		pvc, err := core.Instance().GetPersistentVolumeClaim(vrInfo.PersistentVolumeClaim, ns)
		if err != nil {
			return fmt.Errorf("failed to get PVC %s/%s: %v", ns, vrInfo.PersistentVolumeClaim, err)
		}
		pvcContent, err := runtime.DefaultUnstructuredConverter.ToUnstructured(pvc)
		if err != nil {
			return fmt.Errorf("failed to convert PVC %s to unstructured: %v", vrInfo.RestoreVolume, err)
		}
		pvcObj := &unstructured.Unstructured{}
		pvcObj.SetUnstructuredContent(pvcContent)
		pvcObj.SetGroupVersionKind(schema.GroupVersionKind{
			Kind:    "PersistentVolumeClaim",
			Version: "v1",
			Group:   "core",
		})

		if err := a.restoreController.UpdateResourceStatus(
			restore,
			pvcObj,
			vrInfo.Status,
			"Resource restored successfully",
			restore); err != nil {
			return err
		}
	}

	return nil
}

func (a *ApplicationRestoreController) cleanupRestore(restore *storkapi.ApplicationRestore) error {
	drivers := a.getDriversForRestore(restore)
	for driverName := range drivers {
		driver, err := volume.Get(driverName)
		if err != nil {
			return fmt.Errorf("get %s driver: %s", driverName, err)
		}
		if err = driver.CancelRestore(restore); err != nil {
			return fmt.Errorf("cancel restore: %s", err)
		}
	}
	return nil
}

func (a *ApplicationRestoreController) createCRD() error {
	resource := apiextensions.CustomResource{
		Name:    storkapi.ApplicationRestoreResourceName,
		Plural:  storkapi.ApplicationRestoreResourcePlural,
		Group:   stork.GroupName,
		Version: storkapi.SchemeGroupVersion.Version,
		Scope:   apiextensionsv1beta1.NamespaceScoped,
		Kind:    reflect.TypeOf(storkapi.ApplicationRestore{}).Name(),
	}
	ok, err := version.RequiresV1Registration()
	if err != nil {
		return err
	}
	if ok {
		err := k8sutils.CreateCRD(resource)
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return apiextensions.Instance().ValidateCRD(resource.Plural+"."+resource.Group, validateCRDTimeout, validateCRDInterval)
	}
	err = apiextensions.Instance().CreateCRDV1beta1(resource)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	return apiextensions.Instance().ValidateCRDV1beta1(resource, validateCRDTimeout, validateCRDInterval)
}

func (a *ApplicationRestoreController) cleanupResources(restore *storkapi.ApplicationRestore) error {
	drivers := a.getDriversForRestore(restore)
	for driverName := range drivers {

		driver, err := volume.Get(driverName)
		if err != nil {
			return err
		}
		if err := driver.CleanupRestoreResources(restore); err != nil {
			logrus.Errorf("unable to cleanup post restore resources, err: %v", err)
		}
	}
	return nil
}
