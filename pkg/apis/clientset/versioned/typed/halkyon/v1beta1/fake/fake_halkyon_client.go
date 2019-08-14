/*
Copyright The Kubernetes Authors.

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
	v1beta1 "halkyon.io/operator/pkg/apis/clientset/versioned/typed/halkyon/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHalkyonV1beta1 struct {
	*testing.Fake
}

func (c *FakeHalkyonV1beta1) Capabilities(namespace string) v1beta1.CapabilityInterface {
	return &FakeCapabilities{c, namespace}
}

func (c *FakeHalkyonV1beta1) Components(namespace string) v1beta1.ComponentInterface {
	return &FakeComponents{c, namespace}
}

func (c *FakeHalkyonV1beta1) Links(namespace string) v1beta1.LinkInterface {
	return &FakeLinks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHalkyonV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
