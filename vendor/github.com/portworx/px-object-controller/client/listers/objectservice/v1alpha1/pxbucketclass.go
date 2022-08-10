// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/portworx/px-object-controller/client/apis/objectservice/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PXBucketClassLister helps list PXBucketClasses.
// All objects returned here must be treated as read-only.
type PXBucketClassLister interface {
	// List lists all PXBucketClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PXBucketClass, err error)
	// Get retrieves the PXBucketClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PXBucketClass, error)
	PXBucketClassListerExpansion
}

// pXBucketClassLister implements the PXBucketClassLister interface.
type pXBucketClassLister struct {
	indexer cache.Indexer
}

// NewPXBucketClassLister returns a new PXBucketClassLister.
func NewPXBucketClassLister(indexer cache.Indexer) PXBucketClassLister {
	return &pXBucketClassLister{indexer: indexer}
}

// List lists all PXBucketClasses in the indexer.
func (s *pXBucketClassLister) List(selector labels.Selector) (ret []*v1alpha1.PXBucketClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PXBucketClass))
	})
	return ret, err
}

// Get retrieves the PXBucketClass from the index for a given name.
func (s *pXBucketClassLister) Get(name string) (*v1alpha1.PXBucketClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pxbucketclass"), name)
	}
	return obj.(*v1alpha1.PXBucketClass), nil
}