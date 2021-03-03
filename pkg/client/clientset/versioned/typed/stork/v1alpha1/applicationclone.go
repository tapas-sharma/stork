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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationClonesGetter has a method to return a ApplicationCloneInterface.
// A group's client should implement this interface.
type ApplicationClonesGetter interface {
	ApplicationClones(namespace string) ApplicationCloneInterface
}

// ApplicationCloneInterface has methods to work with ApplicationClone resources.
type ApplicationCloneInterface interface {
	Create(*v1alpha1.ApplicationClone) (*v1alpha1.ApplicationClone, error)
	Update(*v1alpha1.ApplicationClone) (*v1alpha1.ApplicationClone, error)
	UpdateStatus(*v1alpha1.ApplicationClone) (*v1alpha1.ApplicationClone, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationClone, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationCloneList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationClone, err error)
	ApplicationCloneExpansion
}

// applicationClones implements ApplicationCloneInterface
type applicationClones struct {
	client rest.Interface
	ns     string
}

// newApplicationClones returns a ApplicationClones
func newApplicationClones(c *StorkV1alpha1Client, namespace string) *applicationClones {
	return &applicationClones{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationClone, and returns the corresponding applicationClone object, and an error if there is any.
func (c *applicationClones) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationClone, err error) {
	result = &v1alpha1.ApplicationClone{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationclones").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationClones that match those selectors.
func (c *applicationClones) List(opts v1.ListOptions) (result *v1alpha1.ApplicationCloneList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationCloneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationclones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationClones.
func (c *applicationClones) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationclones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a applicationClone and creates it.  Returns the server's representation of the applicationClone, and an error, if there is any.
func (c *applicationClones) Create(applicationClone *v1alpha1.ApplicationClone) (result *v1alpha1.ApplicationClone, err error) {
	result = &v1alpha1.ApplicationClone{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationclones").
		Body(applicationClone).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a applicationClone and updates it. Returns the server's representation of the applicationClone, and an error, if there is any.
func (c *applicationClones) Update(applicationClone *v1alpha1.ApplicationClone) (result *v1alpha1.ApplicationClone, err error) {
	result = &v1alpha1.ApplicationClone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationclones").
		Name(applicationClone.Name).
		Body(applicationClone).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationClones) UpdateStatus(applicationClone *v1alpha1.ApplicationClone) (result *v1alpha1.ApplicationClone, err error) {
	result = &v1alpha1.ApplicationClone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationclones").
		Name(applicationClone.Name).
		SubResource("status").
		Body(applicationClone).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the applicationClone and deletes it. Returns an error if one occurs.
func (c *applicationClones) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationclones").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationClones) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationclones").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched applicationClone.
func (c *applicationClones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationClone, err error) {
	result = &v1alpha1.ApplicationClone{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationclones").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
