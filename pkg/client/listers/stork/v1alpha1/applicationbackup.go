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

// ApplicationBackupLister helps list ApplicationBackups.
// All objects returned here must be treated as read-only.
type ApplicationBackupLister interface {
	// List lists all ApplicationBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationBackup, err error)
	// ApplicationBackups returns an object that can list and get ApplicationBackups.
	ApplicationBackups(namespace string) ApplicationBackupNamespaceLister
	ApplicationBackupListerExpansion
}

// applicationBackupLister implements the ApplicationBackupLister interface.
type applicationBackupLister struct {
	indexer cache.Indexer
}

// NewApplicationBackupLister returns a new ApplicationBackupLister.
func NewApplicationBackupLister(indexer cache.Indexer) ApplicationBackupLister {
	return &applicationBackupLister{indexer: indexer}
}

// List lists all ApplicationBackups in the indexer.
func (s *applicationBackupLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationBackup))
	})
	return ret, err
}

// ApplicationBackups returns an object that can list and get ApplicationBackups.
func (s *applicationBackupLister) ApplicationBackups(namespace string) ApplicationBackupNamespaceLister {
	return applicationBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationBackupNamespaceLister helps list and get ApplicationBackups.
// All objects returned here must be treated as read-only.
type ApplicationBackupNamespaceLister interface {
	// List lists all ApplicationBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationBackup, err error)
	// Get retrieves the ApplicationBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApplicationBackup, error)
	ApplicationBackupNamespaceListerExpansion
}

// applicationBackupNamespaceLister implements the ApplicationBackupNamespaceLister
// interface.
type applicationBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationBackups in the indexer for a given namespace.
func (s applicationBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationBackup))
	})
	return ret, err
}

// Get retrieves the ApplicationBackup from the indexer for a given namespace and name.
func (s applicationBackupNamespaceLister) Get(name string) (*v1alpha1.ApplicationBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationbackup"), name)
	}
	return obj.(*v1alpha1.ApplicationBackup), nil
}
