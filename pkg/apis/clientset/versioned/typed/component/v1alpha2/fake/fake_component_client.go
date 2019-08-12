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
	v1alpha2 "github.com/halkyonio/operator/pkg/apis/clientset/versioned/typed/component/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDevexpV1alpha2 struct {
	*testing.Fake
}

func (c *FakeDevexpV1alpha2) Capabilities(namespace string) v1alpha2.CapabilityInterface {
	return &FakeCapabilities{c, namespace}
}

func (c *FakeDevexpV1alpha2) Components(namespace string) v1alpha2.ComponentInterface {
	return &FakeComponents{c, namespace}
}

func (c *FakeDevexpV1alpha2) Links(namespace string) v1alpha2.LinkInterface {
	return &FakeLinks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDevexpV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
