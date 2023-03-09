package cache

import (
	"context"
	"fmt"
	"sync"

	storkv1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/client-go/rest"
	controllercache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// SharedInformerCache  is an eventually consistent cache. The cache interface
// provides APIs to fetch specific k8s objects from the cache. Only a subset
// of k8s objects are currently managed by this cache.
// DO NOT USE it when you need the latest and accurate copy of a CR.
type SharedInformerCache interface {
	// GetStorageClass returns the storage class if present in the cache.
	GetStorageClass(storageClassName string) (*storagev1.StorageClass, error)

	// GetStorageClassForPVC returns the storage class for the provided PVC if present in cache
	GetStorageClassForPVC(pvc *corev1.PersistentVolumeClaim) (*storagev1.StorageClass, error)

	// ListStorageClasses returns a list of storage classes from the cache
	ListStorageClasses() (*storagev1.StorageClassList, error)

	// GetApplicationRegistration returns the ApplicationRegistration CR from the cache
	GetApplicationRegistration(name string) (*storkv1alpha1.ApplicationRegistration, error)

	// ListApplicationRegistrations lists the application registration CRs from the cache
	ListApplicationRegistrations() (*storkv1alpha1.ApplicationRegistrationList, error)

	// ListPods lists the all the Pods from the cache
	ListTransformedPods() (*corev1.PodList, error)
}

type cache struct {
	controllerCache controllercache.Cache
}

var (
	cacheLock           sync.Mutex
	sharedInformerCache *cache
)

func CreateSharedInformerCache(mgr manager.Manager) error {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	if sharedInformerCache != nil {
		return fmt.Errorf("shared informer cache already initialized")
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// Define a transform map to modify the informer's objects before they're added to the cache.
	transformMap := controllercache.TransformByObject{
		&corev1.Pod{}: func(obj interface{}) (interface{}, error) {
			podResource, ok := obj.(*corev1.Pod)
			if !ok {
				return nil, fmt.Errorf("unexpected object type: %T", obj)
			}
			currPod := corev1.Pod{}
			currPod.Name = podResource.Name
			currPod.Namespace = podResource.Namespace
			currPod.Namespace = podResource.Namespace

			currPod.Spec.Volumes = podResource.Spec.Volumes
			currPod.Spec.Containers = podResource.Spec.Containers
			currPod.Spec.NodeName = podResource.Spec.NodeName

			currPod.Status.Conditions = podResource.Status.Conditions
			currPod.Status.Reason = podResource.Status.Reason
			currPod.Status.Phase = podResource.Status.Phase
			currPod.Status.HostIP = podResource.Status.HostIP
			return &currPod, nil
		},
	}

	sharedInformerCache = &cache{}
	sharedInformerCache.controllerCache, err = controllercache.New(config, controllercache.Options{
		TransformByObject: transformMap,
	})
	if err != nil {
		return err
	}
	go sharedInformerCache.controllerCache.Start(context.Background())

	synced := sharedInformerCache.controllerCache.WaitForCacheSync(context.Background())
	if !synced {
		return fmt.Errorf("error syncing the shared informer cache")
	}
	return nil
}

func Instance() SharedInformerCache {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	return sharedInformerCache
}

// GetStorageClass returns the storage class if present in the cache.
func (c *cache) GetStorageClass(storageClassName string) (*storagev1.StorageClass, error) {
	sc := &storagev1.StorageClass{}
	if err := c.controllerCache.Get(context.Background(), client.ObjectKey{Name: storageClassName}, sc); err != nil {
		return nil, err
	}
	return sc, nil
}

// ListStorageClasses returns a list of storage classes from the cache
func (c *cache) ListStorageClasses() (*storagev1.StorageClassList, error) {
	scList := &storagev1.StorageClassList{}
	if err := c.controllerCache.List(context.Background(), scList); err != nil {
		return nil, err
	}
	return scList, nil
}

// GetStorageClassForPVC returns the storage class for the provided PVC if present in cache
func (c *cache) GetStorageClassForPVC(pvc *corev1.PersistentVolumeClaim) (*storagev1.StorageClass, error) {
	var scName string
	if pvc.Spec.StorageClassName != nil && len(*pvc.Spec.StorageClassName) > 0 {
		scName = *pvc.Spec.StorageClassName
	} else {
		scName = pvc.Annotations[corev1.BetaStorageClassAnnotation]
	}

	if len(scName) == 0 {
		return nil, fmt.Errorf("PVC: %s does not have a storage class", pvc.Name)
	}
	return c.GetStorageClass(scName)
}

// GetApplicationRegistration returns the ApplicationRegistration CR from the cache
func (c *cache) GetApplicationRegistration(name string) (*storkv1alpha1.ApplicationRegistration, error) {
	appReg := &storkv1alpha1.ApplicationRegistration{}
	if err := c.controllerCache.Get(context.Background(), client.ObjectKey{Name: name}, appReg); err != nil {
		return nil, err
	}
	return appReg, nil
}

// ListApplicationRegistrations lists the application registration CRs from the cache
func (c *cache) ListApplicationRegistrations() (*storkv1alpha1.ApplicationRegistrationList, error) {
	appRegList := &storkv1alpha1.ApplicationRegistrationList{}
	if err := c.controllerCache.List(context.Background(), appRegList); err != nil {
		return nil, err
	}
	return appRegList, nil
}

// ListPods lists the all the Pods from the cache
func (c *cache) ListTransformedPods() (*corev1.PodList, error) {
	podList := &corev1.PodList{}
	if err := c.controllerCache.List(context.Background(), podList); err != nil {
		return nil, err
	}
	return podList, nil
}
