/*
Copyright 2019 blademainer.

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
	v1alpha "github.com/blademainer/kubernetes-demo/pkg/apis/demo/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodSpecs implements PodSpecInterface
type FakePodSpecs struct {
	Fake *FakeDemoV1alpha
	ns   string
}

var podspecsResource = schema.GroupVersionResource{Group: "demo.xiongyingqi.com", Version: "v1alpha", Resource: "podspecs"}

var podspecsKind = schema.GroupVersionKind{Group: "demo.xiongyingqi.com", Version: "v1alpha", Kind: "PodSpec"}

// Get takes name of the podSpec, and returns the corresponding podSpec object, and an error if there is any.
func (c *FakePodSpecs) Get(name string, options v1.GetOptions) (result *v1alpha.PodSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podspecsResource, c.ns, name), &v1alpha.PodSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpec), err
}

// List takes label and field selectors, and returns the list of PodSpecs that match those selectors.
func (c *FakePodSpecs) List(opts v1.ListOptions) (result *v1alpha.PodSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podspecsResource, podspecsKind, c.ns, opts), &v1alpha.PodSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpecList), err
}

// Watch returns a watch.Interface that watches the requested podSpecs.
func (c *FakePodSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podspecsResource, c.ns, opts))

}

// Create takes the representation of a podSpec and creates it.  Returns the server's representation of the podSpec, and an error, if there is any.
func (c *FakePodSpecs) Create(podSpec *v1alpha.PodSpec) (result *v1alpha.PodSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podspecsResource, c.ns, podSpec), &v1alpha.PodSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpec), err
}

// Update takes the representation of a podSpec and updates it. Returns the server's representation of the podSpec, and an error, if there is any.
func (c *FakePodSpecs) Update(podSpec *v1alpha.PodSpec) (result *v1alpha.PodSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podspecsResource, c.ns, podSpec), &v1alpha.PodSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpec), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodSpecs) UpdateStatus(podSpec *v1alpha.PodSpec) (*v1alpha.PodSpec, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podspecsResource, "status", c.ns, podSpec), &v1alpha.PodSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpec), err
}

// Delete takes name of the podSpec and deletes it. Returns an error if one occurs.
func (c *FakePodSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podspecsResource, c.ns, name), &v1alpha.PodSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podspecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.PodSpecList{})
	return err
}

// Patch applies the patch and returns the patched podSpec.
func (c *FakePodSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.PodSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podspecsResource, c.ns, name, pt, data, subresources...), &v1alpha.PodSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.PodSpec), err
}
