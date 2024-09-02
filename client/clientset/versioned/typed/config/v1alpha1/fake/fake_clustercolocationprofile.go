/*
Copyright 2022 The Koordinator Authors.

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

	v1alpha1 "github.com/koordinator-sh/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterColocationProfiles implements ClusterColocationProfileInterface
type FakeClusterColocationProfiles struct {
	Fake *FakeConfigV1alpha1
}

var clustercolocationprofilesResource = v1alpha1.SchemeGroupVersion.WithResource("clustercolocationprofiles")

var clustercolocationprofilesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterColocationProfile")

// Get takes name of the clusterColocationProfile, and returns the corresponding clusterColocationProfile object, and an error if there is any.
func (c *FakeClusterColocationProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterColocationProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustercolocationprofilesResource, name), &v1alpha1.ClusterColocationProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterColocationProfile), err
}

// List takes label and field selectors, and returns the list of ClusterColocationProfiles that match those selectors.
func (c *FakeClusterColocationProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterColocationProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustercolocationprofilesResource, clustercolocationprofilesKind, opts), &v1alpha1.ClusterColocationProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterColocationProfileList{ListMeta: obj.(*v1alpha1.ClusterColocationProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterColocationProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterColocationProfiles.
func (c *FakeClusterColocationProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustercolocationprofilesResource, opts))
}

// Create takes the representation of a clusterColocationProfile and creates it.  Returns the server's representation of the clusterColocationProfile, and an error, if there is any.
func (c *FakeClusterColocationProfiles) Create(ctx context.Context, clusterColocationProfile *v1alpha1.ClusterColocationProfile, opts v1.CreateOptions) (result *v1alpha1.ClusterColocationProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustercolocationprofilesResource, clusterColocationProfile), &v1alpha1.ClusterColocationProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterColocationProfile), err
}

// Update takes the representation of a clusterColocationProfile and updates it. Returns the server's representation of the clusterColocationProfile, and an error, if there is any.
func (c *FakeClusterColocationProfiles) Update(ctx context.Context, clusterColocationProfile *v1alpha1.ClusterColocationProfile, opts v1.UpdateOptions) (result *v1alpha1.ClusterColocationProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustercolocationprofilesResource, clusterColocationProfile), &v1alpha1.ClusterColocationProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterColocationProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterColocationProfiles) UpdateStatus(ctx context.Context, clusterColocationProfile *v1alpha1.ClusterColocationProfile, opts v1.UpdateOptions) (*v1alpha1.ClusterColocationProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustercolocationprofilesResource, "status", clusterColocationProfile), &v1alpha1.ClusterColocationProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterColocationProfile), err
}

// Delete takes name of the clusterColocationProfile and deletes it. Returns an error if one occurs.
func (c *FakeClusterColocationProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustercolocationprofilesResource, name, opts), &v1alpha1.ClusterColocationProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterColocationProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustercolocationprofilesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterColocationProfileList{})
	return err
}

// Patch applies the patch and returns the patched clusterColocationProfile.
func (c *FakeClusterColocationProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterColocationProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustercolocationprofilesResource, name, pt, data, subresources...), &v1alpha1.ClusterColocationProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterColocationProfile), err
}
