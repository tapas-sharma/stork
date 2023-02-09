package resourcecollector

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/inflect"
	"github.com/heptio/ark/pkg/discovery"
	"github.com/libopenstorage/stork/drivers/volume"
	stork_api "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	storkcache "github.com/libopenstorage/stork/pkg/cache"
	"github.com/portworx/sched-ops/k8s/core"
	"github.com/portworx/sched-ops/k8s/rbac"
	"github.com/portworx/sched-ops/k8s/storage"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	"k8s.io/kubernetes/pkg/util/slice"
)

const (
	// Annotation to use when the resource shouldn't be collected
	skipResourceAnnotationDeprecated = "stork.libopenstorage.org/skipresource"
	skipResourceAnnotation           = "stork.libopenstorage.org/skip-resource"
	skipOwnerRefCheckAnnotation      = "stork.libopenstorage.org/skip-owner-ref-check"
	skipModifyResources              = "stork.libopenstorage.org/skip-modify-resource"
	// StorkResourceHash is the annotation used to keep track of resources
	// updates
	StorkResourceHash = "stork.libopenstorage.org/resource-hash"
	// SkipModifyResources is the annotation used to skip update of resources
	SkipModifyResources = "stork.libopenstorage.org/skip-modify-resource"
	// ProjectMappingsOption is a resource collector option to provide rancher project mappings
	ProjectMappingsOption = "ProjectMappings"
	// IncludeResources to not skip resources of specific type
	IncludeResources = "stork.libopenstorage.org/include-resource"
	// TransformedResourceName is the annotation used to check if resource has been updated
	// as per transformation rules
	TransformedResourceName = "stork.libopenstorage.org/resourcetransformation-name"
	// CurrentStorageClassName is the annotation used to store the current storage class of the PV before
	// taking backup as we will reset it to empty.
	CurrentStorageClassName = "stork.libopenstorage.org/current-storage-class-name"

	// ServiceKind for k8s service resources
	ServiceKind = "Service"
	// NetworkPolicyKind for network policy resources
	NetworkPolicyKind    = "NetworkPolicy"
	deletedMaxRetries    = 12
	deletedRetryInterval = 10 * time.Second
)

// ResourceCollector is used to collect and process unstructured objects in namespaces and using label selectors
type ResourceCollector struct {
	Driver           volume.Driver
	QPS              float32
	Burst            int
	discoveryHelper  discovery.Helper
	dynamicInterface dynamic.Interface
	coreOps          core.Ops
	rbacOps          rbac.Ops
	storkOps         storkops.Ops
	storageOps       storage.Ops
}

// Options are the options passed to the ResourceCollector APIs that dictate how k8s
// resources should be collected
type Options struct {
	// SkipServices if set will skip the service objects from collection
	SkipServices bool
	// IncludeAllNetworkPolicies if set will include all network policies
	// even if they have CIDRs set on them
	IncludeAllNetworkPolicies bool
	// RancherProjectMappings provides the rancher project mappings which allows the
	// resource collector to perform transformations on certain k8s resources.
	// TODO: temporary change required to handle project related transformations
	RancherProjectMappings map[string]string
}

// Objects Collection of objects
type Objects struct {
	Items []runtime.Unstructured
	// Map to prevent collection of duplicate objects
	resourceMap map[types.UID]bool
}

// Init initializes the resource collector
func (r *ResourceCollector) Init(config *restclient.Config) error {
	var err error
	if config == nil {
		config, err = restclient.InClusterConfig()
		if err != nil {
			return fmt.Errorf("error getting cluster config: %v", err)
		}
	}
	if r.QPS > 0 {
		config.QPS = r.QPS
	}
	if r.Burst > 0 {
		config.Burst = r.Burst
	}
	aeclient, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error getting apiextension client, %v", err)
	}

	discoveryClient := aeclient.Discovery()
	r.discoveryHelper, err = discovery.NewHelper(discoveryClient, logrus.New())
	if err != nil {
		return err
	}

	r.dynamicInterface, err = dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	// reset k8s instance to given cluster config
	r.coreOps, err = core.NewForConfig(config)
	if err != nil {
		return err
	}
	r.rbacOps, err = rbac.NewForConfig(config)
	if err != nil {
		return err
	}
	r.storkOps, err = storkops.NewForConfig(config)
	if err != nil {
		return err
	}
	r.storageOps, err = storage.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

func resourceToBeCollected(resource metav1.APIResource, grp schema.GroupVersion, crdKinds []metav1.GroupVersionKind, optionalResourceTypes []string) bool {
	// Ignore CSI Snapshot object
	if resource.Kind == "VolumeSnapshot" {
		return false
	}

	// Include all namespaced CRDs
	for _, res := range crdKinds {
		if res.Kind == resource.Kind &&
			res.Group == grp.Group && res.Version == grp.Version && resource.Namespaced {
			return true
		}
	}

	return GetSupportedK8SResources(resource.Kind, optionalResourceTypes)
}

// GetSupportedK8SResources returns supported k8s resources by resource collector
// pkgs, this can be used to validate list of resources supported by different stork
// controller like migration, backup, clone etc
func GetSupportedK8SResources(kind string, optionalResourceTypes []string) bool {
	switch kind {
	case "PersistentVolumeClaim",
		"PersistentVolume",
		"Deployment",
		"DeploymentConfig",
		"StatefulSet",
		"ConfigMap",
		"Service",
		"Secret",
		"DaemonSet",
		"ServiceAccount",
		"Role",
		"RoleBinding",
		"ClusterRole",
		"ClusterRoleBinding",
		"ImageStream",
		"Ingress",
		"Route",
		"Template",
		"CronJob",
		"ResourceQuota",
		"ReplicaSet",
		"LimitRange",
		"NetworkPolicy",
		"PodDisruptionBudget",
		"Endpoints",
		"ValidatingWebhookConfiguration",
		"MutatingWebhookConfiguration":
		return true
	case "Job":
		return slice.ContainsString(optionalResourceTypes, "job", strings.ToLower) ||
			slice.ContainsString(optionalResourceTypes, "jobs", strings.ToLower)
	default:
		return false
	}
}

// GetResourceTypes returns all the supported resource types by the collector
func (r *ResourceCollector) GetResourceTypes(
	optionalResourceTypes []string,
	allDrivers bool) ([]metav1.APIResource, error) {
	resourceTypes := make([]metav1.APIResource, 0)
	err := r.discoveryHelper.Refresh()
	if err != nil {
		return nil, err
	}
	var crdResources []metav1.GroupVersionKind
	var crdList *stork_api.ApplicationRegistrationList
	storkcache.Instance()
	if !reflect.ValueOf(storkcache.Instance()).IsNil() {
		crdList, err = storkcache.Instance().ListApplicationRegistrations()
	} else {
		crdList, err = r.storkOps.ListApplicationRegistrations()
	}
	if err != nil {
		logrus.Warnf("Unable to get registered crds, err %v", err)
	} else {
		if crdList != nil {
			for _, crd := range crdList.Items {
				for _, kind := range crd.Resources {
					crdResources = append(crdResources, kind.GroupVersionKind)
				}
			}
		}
	}
	for _, group := range r.discoveryHelper.Resources() {
		groupVersion, err := schema.ParseGroupVersion(group.GroupVersion)
		if err != nil {
			return nil, err
		}

		for _, resource := range group.APIResources {
			if !resourceToBeCollected(resource, groupVersion, crdResources, optionalResourceTypes) {
				continue
			}
			resource.Group = groupVersion.Group
			resource.Version = groupVersion.Version
			resourceTypes = append(resourceTypes, resource)
		}
	}
	return resourceTypes, nil
}

// GetResourcesForType gets all the resources for the given type
// and all PVCs which have an owner reference set.
func (r *ResourceCollector) GetResourcesForType(
	resource metav1.APIResource,
	objects *Objects,
	namespaces []string,
	labelSelectors map[string]string,
	includeObjects map[stork_api.ObjectInfo]bool,
	allDrivers bool,
	opts Options,
) (*Objects, []v1.PersistentVolumeClaim, error) {

	if objects == nil {
		objects = &Objects{
			Items: make([]runtime.Unstructured, 0),
		}
	}
	if objects.resourceMap == nil {
		objects.resourceMap = make(map[types.UID]bool)
	}

	gvr := schema.GroupVersionResource{
		Group:    resource.Group,
		Version:  resource.Version,
		Resource: resource.Name,
	}

	crbs, err := r.rbacOps.ListClusterRoleBindings()
	if err != nil {
		if !apierrors.IsForbidden(err) {
			return nil, nil, err
		}
	}
	for _, ns := range namespaces {
		var dynamicClient dynamic.ResourceInterface
		if !resource.Namespaced {
			dynamicClient = r.dynamicInterface.Resource(gvr)
		} else {
			dynamicClient = r.dynamicInterface.Resource(gvr).Namespace(ns)
		}

		var selectors string
		// PVs don't get the labels from their PVCs, so don't use the label selector
		switch resource.Kind {
		case "PersistentVolume":
		default:
			selectors = labels.Set(labelSelectors).String()
		}
		objectsList, err := dynamicClient.List(context.TODO(), metav1.ListOptions{
			LabelSelector: selectors,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("error listing objects for %v: %v", gvr, err)
		}
		resourceObjects, err := meta.ExtractList(objectsList)
		if err != nil {
			return nil, nil, err
		}
		for _, o := range resourceObjects {
			runtimeObject, ok := o.(runtime.Unstructured)
			if !ok {
				return nil, nil, fmt.Errorf("error casting object: %v", o)
			}

			collect, err := r.objectToBeCollected(includeObjects, labelSelectors, objects.resourceMap, runtimeObject, ns, allDrivers, opts, crbs)
			if err != nil {
				return nil, nil, fmt.Errorf("error processing object %v: %v", runtimeObject, err)
			}
			if !collect {
				continue
			}
			metadata, err := meta.Accessor(runtimeObject)
			if err != nil {
				return nil, nil, err
			}
			objects.Items = append(objects.Items, runtimeObject)
			objects.resourceMap[metadata.GetUID()] = true
		}
	}

	modObjects, pvcObjectsWithOwnerRef, err := r.pruneOwnedResources(objects.Items, objects.resourceMap)
	if err != nil {
		return nil, nil, err
	}
	var crdList *stork_api.ApplicationRegistrationList
	if !reflect.ValueOf(storkcache.Instance()).IsNil() {
		crdList, err = storkcache.Instance().ListApplicationRegistrations()
	} else {
		crdList, err = r.storkOps.ListApplicationRegistrations()
	}
	if err != nil {
		logrus.Warnf("Unable to get registered crds, err %v", err)
	}

	// Creating a list of PVCs with owner reference before calling prepareResourcesForCollection
	// prepareResourcesForCollection can update the PVC metadata which is required when updating
	// owner references on the destination PVC objects
	var pvcsWithOwnerReference []v1.PersistentVolumeClaim
	var pvc v1.PersistentVolumeClaim
	for _, o := range pvcObjectsWithOwnerRef {
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), &pvc); err != nil {
			logrus.Warnf("unable to cast pvcs with owner reference: %v", err)
		}
		pvcsWithOwnerReference = append(pvcsWithOwnerReference, pvc)
	}

	err = r.prepareResourcesForCollection(modObjects, namespaces, opts, crdList)
	if err != nil {
		return nil, nil, err
	}
	objects.Items = modObjects
	return objects, pvcsWithOwnerReference, nil
}

// GetResources gets all the resources in the given list of namespaces which match the labelSelectors
// and all PVCs which have an owner reference set
func (r *ResourceCollector) GetResources(
	namespaces []string,
	labelSelectors map[string]string,
	includeObjects map[stork_api.ObjectInfo]bool,
	optionalResourceTypes []string,
	allDrivers bool,
	opts Options,
) ([]runtime.Unstructured, []v1.PersistentVolumeClaim, error) {
	err := r.discoveryHelper.Refresh()
	if err != nil {
		return nil, nil, err
	}
	allObjects := make([]runtime.Unstructured, 0)
	// Map to prevent collection of duplicate objects
	resourceMap := make(map[types.UID]bool)
	var crdResources []metav1.GroupVersionKind
	var crdList *stork_api.ApplicationRegistrationList
	if !reflect.ValueOf(storkcache.Instance()).IsNil() {
		crdList, err = storkcache.Instance().ListApplicationRegistrations()
	} else {
		crdList, err = r.storkOps.ListApplicationRegistrations()
	}
	if err != nil {
		logrus.Warnf("Unable to get registered crds, err %v", err)
	} else {
		if crdList != nil {
			for _, crd := range crdList.Items {
				for _, kind := range crd.Resources {
					crdResources = append(crdResources, kind.GroupVersionKind)
				}
			}
		}
	}

	crbs, err := r.rbacOps.ListClusterRoleBindings()
	if err != nil {
		if !apierrors.IsForbidden(err) {
			return nil, nil, err
		}
	}

	for _, group := range r.discoveryHelper.Resources() {
		groupVersion, err := schema.ParseGroupVersion(group.GroupVersion)
		if err != nil {
			return nil, nil, err
		}

		for _, resource := range group.APIResources {
			if !resourceToBeCollected(resource, groupVersion, crdResources, optionalResourceTypes) {
				continue
			}
			for _, ns := range namespaces {
				var dynamicClient dynamic.ResourceInterface
				if !resource.Namespaced {
					dynamicClient = r.dynamicInterface.Resource(groupVersion.WithResource(resource.Name))
				} else {
					dynamicClient = r.dynamicInterface.Resource(groupVersion.WithResource(resource.Name)).Namespace(ns)
				}

				var objectToInclude map[stork_api.ObjectInfo]bool
				if !IsNsPresentInIncludeResource(includeObjects, ns) {
					objectToInclude = make(map[stork_api.ObjectInfo]bool)
				} else {
					objectToInclude = includeObjects
				}

				var selectors string
				// PVs don't get the labels from their PVCs, so don't use the label selector
				switch resource.Kind {
				case "PersistentVolume":
				default:
					selectors = labels.Set(labelSelectors).String()
				}
				objectsList, err := dynamicClient.List(context.TODO(), metav1.ListOptions{
					LabelSelector: selectors,
				})
				if err != nil {
					if apierrors.IsForbidden(err) {
						continue
					}
					return nil, nil, err
				}
				objects, err := meta.ExtractList(objectsList)
				if err != nil {
					return nil, nil, err
				}
				for _, o := range objects {
					runtimeObject, ok := o.(runtime.Unstructured)
					if !ok {
						return nil, nil, fmt.Errorf("error casting object: %v", o)
					}

					var err error
					var collect bool
					// If a namespace is present in both namespace list and IncludeResource Object,
					// IncludeResource takes priority and only those resources are backed up.
					// If a ns is only present in namespace list, all resources in that ns
					// is backed up.
					// With this now a user can choose to backup all resources in a ns and some
					// selected resources from different ns

					collect, err = r.objectToBeCollected(objectToInclude, labelSelectors, resourceMap, runtimeObject, ns, allDrivers, opts, crbs)
					if err != nil {
						if apierrors.IsForbidden(err) {
							continue
						}
						return nil, nil, fmt.Errorf("error processing object %v: %v", runtimeObject, err)
					}
					if !collect {
						continue
					}
					metadata, err := meta.Accessor(runtimeObject)
					if err != nil {
						return nil, nil, err
					}
					allObjects = append(allObjects, runtimeObject)
					resourceMap[metadata.GetUID()] = true
				}
			}
		}
	}

	allObjects, pvcObjectsWithOwnerRef, err := r.pruneOwnedResources(allObjects, resourceMap)
	if err != nil {
		return nil, nil, err
	}

	// Creating a list of PVCs with owner reference before calling prepareResourcesForCollection
	// prepareResourcesForCollection can update the PVC metadata which is required when updating
	// owner references on the destination PVC objects
	var pvcsWithOwnerReference []v1.PersistentVolumeClaim
	var pvc v1.PersistentVolumeClaim
	for _, o := range pvcObjectsWithOwnerRef {
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), &pvc); err != nil {
			logrus.Warnf("unable to cast pvcs with owner reference: %v", err)
		}
		pvcsWithOwnerReference = append(pvcsWithOwnerReference, pvc)
	}

	err = r.prepareResourcesForCollection(allObjects, namespaces, opts, crdList)
	if err != nil {
		return nil, nil, err
	}

	return allObjects, pvcsWithOwnerReference, nil
}

// IsNsPresentInIncludeResource checks if a given ns is present in the IncludeResource object
func IsNsPresentInIncludeResource(includeObjects map[stork_api.ObjectInfo]bool, namespace string) bool {
	for obj := range includeObjects {
		if obj.Namespace == namespace {
			return true
		}
	}

	return false
}

// SkipResource returns whether the annotations of the object require it to be
// skipped
func SkipResource(annotations map[string]string) bool {
	if value, present := annotations[skipResourceAnnotation]; present {
		if skip, err := strconv.ParseBool(value); err == nil && skip {
			return true
		}
		return false
	}
	if value, present := annotations[skipResourceAnnotationDeprecated]; present {
		if skip, err := strconv.ParseBool(value); err == nil && skip {
			return true
		}
		return false
	}
	return false
}

// skipOwnerRefCheck returns whether the object should be collected even if it
// has an owner reference
func skipOwnerRefCheck(annotations map[string]string) bool {
	if value, present := annotations[skipOwnerRefCheckAnnotation]; present {
		if skip, err := strconv.ParseBool(value); err == nil && skip {
			return true
		}
	}
	return false
}

// Returns whether an object should be collected or not for the requested
// namespace
func (r *ResourceCollector) objectToBeCollected(
	includeObjects map[stork_api.ObjectInfo]bool,
	labelSelectors map[string]string,
	resourceMap map[types.UID]bool,
	object runtime.Unstructured,
	namespace string,
	allDrivers bool,
	opts Options,
	crbs *rbacv1.ClusterRoleBindingList,
) (bool, error) {
	metadata, err := meta.Accessor(object)
	if err != nil {
		return false, err
	}

	if SkipResource(metadata.GetAnnotations()) {
		return false, err
	}

	// Skip if we've already processed this object
	if _, ok := resourceMap[metadata.GetUID()]; ok {
		return false, nil
	}

	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return false, err
	}

	if include, err := r.includeObject(object, includeObjects, namespace); err != nil {
		return false, err
	} else if !include {
		return false, nil
	}
	switch objectType.GetKind() {
	case "Service":
		return r.serviceToBeCollected(object)
	case "PersistentVolumeClaim":
		return r.pvcToBeCollected(object, namespace, allDrivers)
	case "PersistentVolume":
		return r.pvToBeCollected(includeObjects, labelSelectors, object, namespace, allDrivers)
	case "ClusterRoleBinding":
		return r.clusterRoleBindingToBeCollected(labelSelectors, object, namespace)
	case "ClusterRole":
		return r.clusterRoleToBeCollected(labelSelectors, object, crbs, namespace)
	case "ServiceAccount":
		return r.serviceAccountToBeCollected(object)
	case "Secret":
		return r.secretToBeCollected(object)
	case "Role":
		return r.roleToBeCollected(object)
	case "RoleBinding":
		return r.roleBindingToBeCollected(object)
	case "Ingress":
		return r.ingressToBeCollected(object)
	case "ConfigMap":
		return r.configmapToBeCollected(object)
	case "ResourceQuota":
		return r.resourceQuotaToBeCollected(object)
	case "NetworkPolicy":
		return r.networkPolicyToBeCollected(object, opts)
	case "DataVolume":
		return r.dataVolumesToBeCollected(object)
	case "VirtualMachineInstance":
		return r.virtualMachineInstanceToBeCollected(object)
	case "VirtualMachineInstanceMigration":
		return r.virtualMachineInstanceMigrationToBeCollected(object)
	case "Endpoints":
		return r.endpointsToBeCollected(object)
	case "MutatingWebhookConfiguration":
		return r.mutatingWebHookToBeCollected(object, namespace)
	case "ValidatingWebhookConfiguration":
		return r.validatingWebHookToBeCollected(object, namespace)
	}

	return true, nil
}

// Prune objects that are owned by a CRD if we are also collecting the CR
// since they will be recreated by the operator when the CR is created too.
// For objects that have an ownerRef but we aren't collecting it's owner
// remove the ownerRef so that the object doesn't get automatically deleted.
// Collect PVCs with owner reference separately as we do need to set owner
// references for them after migrating parent CR
func (r *ResourceCollector) pruneOwnedResources(
	objects []runtime.Unstructured,
	resourceMap map[types.UID]bool,
) ([]runtime.Unstructured, []runtime.Unstructured, error) {
	updatedObjects := make([]runtime.Unstructured, 0)
	pvcObjectsWithOwnerRef := make([]runtime.Unstructured, 0)

	for _, o := range objects {
		metadata, err := meta.Accessor(o)
		if err != nil {
			return nil, nil, err
		}

		collect := true

		objectType, err := meta.TypeAccessor(o)
		if err != nil {
			return nil, nil, err
		}
		owners := metadata.GetOwnerReferences()
		if len(owners) != 0 {
			if objectType.GetKind() == "PersistentVolumeClaim" {
				for _, owner := range owners {
					// Collect PVC objects separately. If OwnerReferences is available
					// but the corresponding resource is not being collected, then
					// no need to collect resource and we will not manually patch those PVCs
					if _, exists := resourceMap[owner.UID]; exists {
						pvcObjectsWithOwnerRef = append(pvcObjectsWithOwnerRef, o)
						break
					}
				}
			} else {
				if !skipOwnerRefCheck(metadata.GetAnnotations()) {
					for _, owner := range owners {
						// We don't collect pods, there might be some leader
						// election objects that could have pods as the owner, so
						// don't collect those objects
						if owner.Kind == "Pod" {
							collect = false
							break
						}
						if objectType.GetKind() != "Deployment" &&
							objectType.GetKind() != "StatefulSet" &&
							objectType.GetKind() != "ReplicaSet" &&
							objectType.GetKind() != "DeploymentConfig" &&
							objectType.GetKind() != "Service" {
							continue
						}

						// Skip object if we are already collecting its owner
						if _, exists := resourceMap[owner.UID]; exists {
							collect = false
							break
						}
					}
				}
				// If the owner isn't being collected delete the owner reference
				metadata.SetOwnerReferences(nil)
			}
		}
		if collect {
			updatedObjects = append(updatedObjects, o)
		}
	}

	return updatedObjects, pvcObjectsWithOwnerRef, nil

}

func (r *ResourceCollector) prepareResourcesForCollection(
	objects []runtime.Unstructured,
	namespaces []string,
	opts Options,
	crdList *stork_api.ApplicationRegistrationList,
) error {
	for _, o := range objects {
		metadata, err := meta.Accessor(o)
		if err != nil {
			return err
		}

		switch o.GetObjectKind().GroupVersionKind().Kind {
		case "PersistentVolume":
			err := r.preparePVResourceForCollection(o)
			if err != nil {
				return fmt.Errorf("error preparing PV resource %v: %v", metadata.GetName(), err)
			}
		case "Service":
			if !opts.SkipServices {
				err := r.prepareServiceResourceForCollection(o)
				if err != nil {
					return fmt.Errorf("error preparing Service resource %v/%v: %v", metadata.GetNamespace(), metadata.GetName(), err)
				}
			}
		case "ClusterRoleBinding":
			err := r.prepareClusterRoleBindingForCollection(o, namespaces)
			if err != nil {
				return fmt.Errorf("error preparing ClusterRoleBindings resource %v: %v", metadata.GetName(), err)
			}
		case "Job":
			err := r.prepareJobForCollection(o, namespaces)
			if err != nil {
				return fmt.Errorf("error preparing job resource %v: %v", metadata.GetName(), err)
			}

		case "NetworkPolicy":
			err := r.prepareRancherNetworkPolicy(o, opts)
			if err != nil {
				return fmt.Errorf("error preparing NetworkPolicy resource %v: %v", metadata.GetName(), err)
			}

		case "VirtualMachine":
			err := r.prepareVirtualMachineForCollection(o, namespaces)
			if err != nil {
				return fmt.Errorf("error preparing VirtualMachine resource %v: %v", metadata.GetName(), err)
			}
		}

		content := o.UnstructuredContent()
		if crdList != nil {
			resourceKind := o.GetObjectKind().GroupVersionKind()
			for _, crd := range crdList.Items {
				for _, kind := range crd.Resources {
					if kind.Kind == resourceKind.Kind && kind.Group == resourceKind.Group &&
						kind.Version == resourceKind.Version {
						// remove status from crd
						if !kind.KeepStatus {
							delete(content, "status")
						}

					}
				}
			}
		}
		// remove metadata annotations
		metadataMap := content["metadata"].(map[string]interface{})
		// Remove all metadata except some well-known ones
		for key := range metadataMap {
			switch key {
			case "name", "namespace", "labels", "annotations":
			default:
				delete(metadataMap, key)
			}
		}

	}
	return nil
}

// includeObject determines whether to include an object or not
// based on the object kind
func (r *ResourceCollector) includeObject(
	object runtime.Unstructured,
	includeObjects map[stork_api.ObjectInfo]bool,
	namespace string,
) (bool, error) {
	if len(includeObjects) == 0 {
		return true, nil
	}

	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return false, err
	}

	metadata, err := meta.Accessor(object)
	if err != nil {
		return false, err
	}

	// Even if PV isn't specified need to check if the corresponding PVC is, so
	// skip the check here
	if objectType.GetKind() != "PersistentVolume" {
		info := stork_api.ObjectInfo{
			GroupVersionKind: metav1.GroupVersionKind{
				Group:   object.GetObjectKind().GroupVersionKind().Group,
				Version: object.GetObjectKind().GroupVersionKind().Version,
				Kind:    object.GetObjectKind().GroupVersionKind().Kind,
			},
			Name:      metadata.GetName(),
			Namespace: metadata.GetNamespace(),
		}
		if info.Group == "" {
			info.Group = "core"
		}
		if namespace != "" {
			info.Namespace = namespace
		}
		if val, present := includeObjects[info]; !present || !val {
			return false, nil
		}
	}
	return true, nil
}

// PrepareResourceForApply prepares the resource for apply including update
// namespace and any PV name updates. Should be called before DeleteResources
// and ApplyResource
func (r *ResourceCollector) PrepareResourceForApply(
	object runtime.Unstructured,
	allObjects []runtime.Unstructured,
	includeObjects map[stork_api.ObjectInfo]bool,
	namespaceMappings map[string]string,
	storageClassMappings map[string]string,
	pvNameMappings map[string]string,
	optionalResourceTypes []string,
	vInfo []*stork_api.ApplicationRestoreVolumeInfo,
	opts *Options,
) (bool, error) {
	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return false, err
	}

	metadata, err := meta.Accessor(object)
	if err != nil {
		return false, err
	}

	if include, err := r.includeObject(object, includeObjects, ""); err != nil {
		return true, err
	} else if !include {
		return true, nil
	}

	if metadata.GetNamespace() != "" {
		var val string
		var present bool
		// Skip the object if it isn't in the namespace mapping
		if val, present = namespaceMappings[metadata.GetNamespace()]; !present {
			return true, nil
		}
		// Update the namespace of the object, will be no-op for clustered resources
		metadata.SetNamespace(val)
	}
	switch objectType.GetKind() {
	case "Job":
		if slice.ContainsString(optionalResourceTypes, "job", strings.ToLower) ||
			slice.ContainsString(optionalResourceTypes, "jobs", strings.ToLower) {
			return false, nil
		}
		return true, nil
	case "PersistentVolume":
		return r.preparePVResourceForApply(object, pvNameMappings, vInfo, storageClassMappings, namespaceMappings, *opts)
	case "PersistentVolumeClaim":
		return r.preparePVCResourceForApply(object, allObjects, pvNameMappings, storageClassMappings, vInfo, *opts)
	case "ClusterRoleBinding":
		return false, r.prepareClusterRoleBindingForApply(object, namespaceMappings)
	case "RoleBinding":
		return false, r.prepareRoleBindingForApply(object, namespaceMappings)
	case "ValidatingWebhookConfiguration":
		return false, r.prepareValidatingWebHookForApply(object, namespaceMappings)
	case "MutatingWebhookConfiguration":
		return false, r.prepareMutatingWebHookForApply(object, namespaceMappings)
	case "Secret":
		return false, r.prepareSecretForApply(object)
	case "NetworkPolicy":
		return false, r.prepareRancherNetworkPolicy(object, *opts)
	case "Deployment", "StatefulSet", "DeploymentConfig", "IBPPeer", "IBPCA", "IBPConsole", "IBPOrderer", "ReplicaSet":
		return false, r.prepareRancherApplicationResource(object, *opts)

	}
	return false, nil
}

func (r *ResourceCollector) mergeSupportedForResource(
	object runtime.Unstructured,
) bool {
	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return false
	}
	switch objectType.GetKind() {
	case "ClusterRoleBinding",
		"ServiceAccount":
		return true
	}
	return false
}

func (r *ResourceCollector) mergeAndUpdateResource(
	object runtime.Unstructured,
) error {
	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return err
	}

	switch objectType.GetKind() {
	case "ClusterRoleBinding":
		return r.mergeAndUpdateClusterRoleBinding(object)
	case "ServiceAccount":
		return r.mergeAndUpdateServiceAccount(object)
	}
	return nil
}

// ApplyResource applies a given resource using the provided client interface
func (r *ResourceCollector) ApplyResource(
	dynamicInterface dynamic.Interface,
	object runtime.Unstructured,
) error {
	dynamicClient, err := r.getDynamicClient(dynamicInterface, object)
	if err != nil {
		return err
	}
	_, err = dynamicClient.Create(context.TODO(), object.(*unstructured.Unstructured), metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) || strings.Contains(err.Error(), portallocator.ErrAllocated.Error()) {
			if r.mergeSupportedForResource(object) {
				return r.mergeAndUpdateResource(object)
			} else if strings.Contains(err.Error(), portallocator.ErrAllocated.Error()) {
				err = r.updateService(object)
				if err != nil {
					return err
				}
			} else {
				return err
			}
			_, err = dynamicClient.Create(context.TODO(), object.(*unstructured.Unstructured), metav1.CreateOptions{})
			return err
		}
	}

	return err
}

// DeleteResources deletes given resources using the provided client interface
func (r *ResourceCollector) DeleteResources(
	dynamicInterface dynamic.Interface,
	objects []runtime.Unstructured,
) error {
	// First delete all the objects
	deleteStart := metav1.Now()
	for _, object := range objects {
		// Don't delete objects that support merging
		if r.mergeSupportedForResource(object) {
			continue
		}

		metadata, err := meta.Accessor(object)
		if err != nil {
			return err
		}

		dynamicClient, err := r.getDynamicClient(dynamicInterface, object)
		if err != nil {
			return err
		}

		// Delete the resource if it already exists on the destination
		// cluster and try creating again
		err = dynamicClient.Delete(context.TODO(), metadata.GetName(), metav1.DeleteOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return err
		}
	}

	// Then wait for them to actually be deleted
	for _, object := range objects {
		// Objects that support merging aren't deleted
		if r.mergeSupportedForResource(object) {
			continue
		}

		metadata, err := meta.Accessor(object)
		if err != nil {
			return err
		}

		dynamicClient, err := r.getDynamicClient(dynamicInterface, object)
		if err != nil {
			return err
		}

		// Wait for up to 2 minutes for the object to be deleted
		for i := 0; i < deletedMaxRetries; i++ {
			obj, err := dynamicClient.Get(context.TODO(), metadata.GetName(), metav1.GetOptions{})
			if err != nil && apierrors.IsNotFound(err) {
				break
			}
			createTime := obj.GetCreationTimestamp()
			if deleteStart.Before(&createTime) {
				logrus.Warnf("Object[%v] got re-created after deletion. So, Ignore wait. deleteStart time:[%v], create time:[%v]",
					obj.GetName(), deleteStart, createTime)
				break
			}
			logrus.Warnf("Object %v still present, retrying in %v", metadata.GetName(), deletedRetryInterval)
			time.Sleep(deletedRetryInterval)
		}
	}
	return nil
}

func (r *ResourceCollector) getDynamicClient(
	dynamicInterface dynamic.Interface,
	object runtime.Unstructured,
) (dynamic.ResourceInterface, error) {
	metadata, err := meta.Accessor(object)
	if err != nil {
		return nil, err
	}
	objectType, err := meta.TypeAccessor(object)
	if err != nil {
		return nil, err
	}

	// The default ruleset doesn't pluralize quotas correctly, so add that
	ruleset := inflect.NewDefaultRuleset()
	ruleset.AddPlural("quota", "quotas")
	ruleset.AddPlural("prometheus", "prometheuses")
	ruleset.AddPlural("mongodbcommunity", "mongodbcommunity")
	ruleset.AddPlural("mongodbopsmanager", "opsmanagers")
	ruleset.AddPlural("mongodb", "mongodb")
	resource := &metav1.APIResource{
		Name:       ruleset.Pluralize(strings.ToLower(objectType.GetKind())),
		Namespaced: len(metadata.GetNamespace()) > 0,
	}

	destNamespace := ""
	if resource.Namespaced {
		destNamespace = metadata.GetNamespace()
	}
	return dynamicInterface.Resource(
		object.GetObjectKind().GroupVersionKind().GroupVersion().WithResource(resource.Name)).Namespace(destNamespace), nil
}

func (r *ResourceCollector) prepareRancherApplicationResource(
	object runtime.Unstructured,
	opts Options,
) error {
	content := object.UnstructuredContent()
	if len(opts.RancherProjectMappings) > 0 {
		podSpecField, found, err := unstructured.NestedFieldCopy(content, "spec", "template", "spec")
		if err != nil {
			logrus.Warnf("Unable to parse object %v while handling"+
				" rancher project mappings", object.GetObjectKind().GroupVersionKind().Kind)
		}
		var podSpec v1.PodSpec
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(podSpecField.(map[string]interface{}), &podSpec); err != nil {
			return fmt.Errorf("error converting to podSpec: %v", err)
		}
		if found {
			podSpecPtr := PreparePodSpecNamespaceSelector(
				&podSpec,
				opts.RancherProjectMappings,
			)
			o, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&podSpecPtr)
			if err != nil {
				return err
			}
			if err := unstructured.SetNestedField(content, o, "spec", "template", "spec"); err != nil {
				logrus.Warnf("Unable to set namespace selector for object %v while handling"+
					" rancher project mappings", object.GetObjectKind().GroupVersionKind().Kind)
			}
		}
	}
	return nil
}
