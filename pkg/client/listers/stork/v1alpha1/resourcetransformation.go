/*
Copyright 2018 Openstorage.org

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceTransformationLister helps list ResourceTransformations.
// All objects returned here must be treated as read-only.
type ResourceTransformationLister interface {
	// List lists all ResourceTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceTransformation, err error)
	// ResourceTransformations returns an object that can list and get ResourceTransformations.
	ResourceTransformations(namespace string) ResourceTransformationNamespaceLister
	ResourceTransformationListerExpansion
}

// resourceTransformationLister implements the ResourceTransformationLister interface.
type resourceTransformationLister struct {
	indexer cache.Indexer
}

// NewResourceTransformationLister returns a new ResourceTransformationLister.
func NewResourceTransformationLister(indexer cache.Indexer) ResourceTransformationLister {
	return &resourceTransformationLister{indexer: indexer}
}

// List lists all ResourceTransformations in the indexer.
func (s *resourceTransformationLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceTransformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceTransformation))
	})
	return ret, err
}

// ResourceTransformations returns an object that can list and get ResourceTransformations.
func (s *resourceTransformationLister) ResourceTransformations(namespace string) ResourceTransformationNamespaceLister {
	return resourceTransformationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceTransformationNamespaceLister helps list and get ResourceTransformations.
// All objects returned here must be treated as read-only.
type ResourceTransformationNamespaceLister interface {
	// List lists all ResourceTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceTransformation, err error)
	// Get retrieves the ResourceTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceTransformation, error)
	ResourceTransformationNamespaceListerExpansion
}

// resourceTransformationNamespaceLister implements the ResourceTransformationNamespaceLister
// interface.
type resourceTransformationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceTransformations in the indexer for a given namespace.
func (s resourceTransformationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceTransformation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceTransformation))
	})
	return ret, err
}

// Get retrieves the ResourceTransformation from the indexer for a given namespace and name.
func (s resourceTransformationNamespaceLister) Get(name string) (*v1alpha1.ResourceTransformation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcetransformation"), name)
	}
	return obj.(*v1alpha1.ResourceTransformation), nil
}
