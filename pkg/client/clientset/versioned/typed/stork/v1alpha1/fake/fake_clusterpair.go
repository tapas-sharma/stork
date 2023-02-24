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

package fake

import (
	"context"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPairs implements ClusterPairInterface
type FakeClusterPairs struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var clusterpairsResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "clusterpairs"}

var clusterpairsKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "ClusterPair"}

// Get takes name of the clusterPair, and returns the corresponding clusterPair object, and an error if there is any.
func (c *FakeClusterPairs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterpairsResource, c.ns, name), &v1alpha1.ClusterPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPair), err
}

// List takes label and field selectors, and returns the list of ClusterPairs that match those selectors.
func (c *FakeClusterPairs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPairList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterpairsResource, clusterpairsKind, c.ns, opts), &v1alpha1.ClusterPairList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPairList{ListMeta: obj.(*v1alpha1.ClusterPairList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPairList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPairs.
func (c *FakeClusterPairs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterpairsResource, c.ns, opts))

}

// Create takes the representation of a clusterPair and creates it.  Returns the server's representation of the clusterPair, and an error, if there is any.
func (c *FakeClusterPairs) Create(ctx context.Context, clusterPair *v1alpha1.ClusterPair, opts v1.CreateOptions) (result *v1alpha1.ClusterPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterpairsResource, c.ns, clusterPair), &v1alpha1.ClusterPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPair), err
}

// Update takes the representation of a clusterPair and updates it. Returns the server's representation of the clusterPair, and an error, if there is any.
func (c *FakeClusterPairs) Update(ctx context.Context, clusterPair *v1alpha1.ClusterPair, opts v1.UpdateOptions) (result *v1alpha1.ClusterPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterpairsResource, c.ns, clusterPair), &v1alpha1.ClusterPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPair), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPairs) UpdateStatus(ctx context.Context, clusterPair *v1alpha1.ClusterPair, opts v1.UpdateOptions) (*v1alpha1.ClusterPair, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterpairsResource, "status", c.ns, clusterPair), &v1alpha1.ClusterPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPair), err
}

// Delete takes name of the clusterPair and deletes it. Returns an error if one occurs.
func (c *FakeClusterPairs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterpairsResource, c.ns, name, opts), &v1alpha1.ClusterPair{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPairs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterpairsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPairList{})
	return err
}

// Patch applies the patch and returns the patched clusterPair.
func (c *FakeClusterPairs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterpairsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPair), err
}
