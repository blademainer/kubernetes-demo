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
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePods implements PodInterface
type FakePods struct {
	Fake *FakeDemoV1alpha
	ns   string
}

var podsResource = schema.GroupVersionResource{Group: "demo.xiongyingqi.com", Version: "v1alpha", Resource: "pods"}

var podsKind = schema.GroupVersionKind{Group: "demo.xiongyingqi.com", Version: "v1alpha", Kind: "Pod"}

// Get takes name of the pod, and returns the corresponding pod object, and an error if there is any.
func (c *FakePods) Get(name string, options v1.GetOptions) (result *v1alpha.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podsResource, c.ns, name), &v1alpha.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.Pod), err
}

// List takes label and field selectors, and returns the list of Pods that match those selectors.
func (c *FakePods) List(opts v1.ListOptions) (result *v1alpha.PodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podsResource, podsKind, c.ns, opts), &v1alpha.PodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.PodList{ListMeta: obj.(*v1alpha.PodList).ListMeta}
	for _, item := range obj.(*v1alpha.PodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pods.
func (c *FakePods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podsResource, c.ns, opts))

}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *FakePods) Create(pod *v1alpha.Pod) (result *v1alpha.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podsResource, c.ns, pod), &v1alpha.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.Pod), err
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *FakePods) Update(pod *v1alpha.Pod) (result *v1alpha.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podsResource, c.ns, pod), &v1alpha.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.Pod), err
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *FakePods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podsResource, c.ns, name), &v1alpha.Pod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.PodList{})
	return err
}

// Patch applies the patch and returns the patched pod.
func (c *FakePods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podsResource, c.ns, name, pt, data, subresources...), &v1alpha.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.Pod), err
}

// GetEphemeralContainers takes name of the pod, and returns the corresponding ephemeralContainers object, and an error if there is any.
func (c *FakePods) GetEphemeralContainers(podName string, options v1.GetOptions) (result *v1alpha.EphemeralContainers, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(podsResource, c.ns, "ephemeralcontainers", podName), &v1alpha.EphemeralContainers{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EphemeralContainers), err
}

// UpdateEphemeralContainers takes the representation of a ephemeralContainers and updates it. Returns the server's representation of the ephemeralContainers, and an error, if there is any.
func (c *FakePods) UpdateEphemeralContainers(podName string, ephemeralContainers *v1alpha.EphemeralContainers) (result *v1alpha.EphemeralContainers, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podsResource, "ephemeralcontainers", c.ns, ephemeralContainers), &v1alpha.EphemeralContainers{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EphemeralContainers), err
}
