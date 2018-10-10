/*
Copyright The Flagger Authors.

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
	v1beta1 "github.com/stefanprodan/flagger/pkg/apis/flagger/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCanaryDeployments implements CanaryDeploymentInterface
type FakeCanaryDeployments struct {
	Fake *FakeFlaggerV1beta1
	ns   string
}

var canarydeploymentsResource = schema.GroupVersionResource{Group: "flagger.app", Version: "v1beta1", Resource: "canarydeployments"}

var canarydeploymentsKind = schema.GroupVersionKind{Group: "flagger.app", Version: "v1beta1", Kind: "CanaryDeployment"}

// Get takes name of the canaryDeployment, and returns the corresponding canaryDeployment object, and an error if there is any.
func (c *FakeCanaryDeployments) Get(name string, options v1.GetOptions) (result *v1beta1.CanaryDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(canarydeploymentsResource, c.ns, name), &v1beta1.CanaryDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CanaryDeployment), err
}

// List takes label and field selectors, and returns the list of CanaryDeployments that match those selectors.
func (c *FakeCanaryDeployments) List(opts v1.ListOptions) (result *v1beta1.CanaryDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(canarydeploymentsResource, canarydeploymentsKind, c.ns, opts), &v1beta1.CanaryDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CanaryDeploymentList{ListMeta: obj.(*v1beta1.CanaryDeploymentList).ListMeta}
	for _, item := range obj.(*v1beta1.CanaryDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested canaryDeployments.
func (c *FakeCanaryDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(canarydeploymentsResource, c.ns, opts))

}

// Create takes the representation of a canaryDeployment and creates it.  Returns the server's representation of the canaryDeployment, and an error, if there is any.
func (c *FakeCanaryDeployments) Create(canaryDeployment *v1beta1.CanaryDeployment) (result *v1beta1.CanaryDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(canarydeploymentsResource, c.ns, canaryDeployment), &v1beta1.CanaryDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CanaryDeployment), err
}

// Update takes the representation of a canaryDeployment and updates it. Returns the server's representation of the canaryDeployment, and an error, if there is any.
func (c *FakeCanaryDeployments) Update(canaryDeployment *v1beta1.CanaryDeployment) (result *v1beta1.CanaryDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(canarydeploymentsResource, c.ns, canaryDeployment), &v1beta1.CanaryDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CanaryDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCanaryDeployments) UpdateStatus(canaryDeployment *v1beta1.CanaryDeployment) (*v1beta1.CanaryDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(canarydeploymentsResource, "status", c.ns, canaryDeployment), &v1beta1.CanaryDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CanaryDeployment), err
}

// Delete takes name of the canaryDeployment and deletes it. Returns an error if one occurs.
func (c *FakeCanaryDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(canarydeploymentsResource, c.ns, name), &v1beta1.CanaryDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCanaryDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(canarydeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CanaryDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched canaryDeployment.
func (c *FakeCanaryDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CanaryDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(canarydeploymentsResource, c.ns, name, data, subresources...), &v1beta1.CanaryDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CanaryDeployment), err
}