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
	v1alpha "github.com/blademainer/kubernetes-demo/pkg/generated/clientset/versioned/typed/demo/v1alpha"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDemoV1alpha struct {
	*testing.Fake
}

func (c *FakeDemoV1alpha) Pods(namespace string) v1alpha.PodInterface {
	return &FakePods{c, namespace}
}

func (c *FakeDemoV1alpha) PodSpecs(namespace string) v1alpha.PodSpecInterface {
	return &FakePodSpecs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDemoV1alpha) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}