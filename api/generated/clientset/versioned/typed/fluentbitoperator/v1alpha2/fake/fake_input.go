/*

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubesphere.io/fluentbit-operator/api/fluentbitoperator/v1alpha2"
)

// FakeInputs implements InputInterface
type FakeInputs struct {
	Fake *FakeFluentbitoperatorV1alpha2
	ns   string
}

var inputsResource = schema.GroupVersionResource{Group: "fluentbitoperator", Version: "v1alpha2", Resource: "inputs"}

var inputsKind = schema.GroupVersionKind{Group: "fluentbitoperator", Version: "v1alpha2", Kind: "Input"}

// Get takes name of the input, and returns the corresponding input object, and an error if there is any.
func (c *FakeInputs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Input, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inputsResource, c.ns, name), &v1alpha2.Input{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Input), err
}

// List takes label and field selectors, and returns the list of Inputs that match those selectors.
func (c *FakeInputs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.InputList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inputsResource, inputsKind, c.ns, opts), &v1alpha2.InputList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.InputList{ListMeta: obj.(*v1alpha2.InputList).ListMeta}
	for _, item := range obj.(*v1alpha2.InputList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inputs.
func (c *FakeInputs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inputsResource, c.ns, opts))

}

// Create takes the representation of a input and creates it.  Returns the server's representation of the input, and an error, if there is any.
func (c *FakeInputs) Create(ctx context.Context, input *v1alpha2.Input, opts v1.CreateOptions) (result *v1alpha2.Input, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inputsResource, c.ns, input), &v1alpha2.Input{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Input), err
}

// Update takes the representation of a input and updates it. Returns the server's representation of the input, and an error, if there is any.
func (c *FakeInputs) Update(ctx context.Context, input *v1alpha2.Input, opts v1.UpdateOptions) (result *v1alpha2.Input, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inputsResource, c.ns, input), &v1alpha2.Input{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Input), err
}

// Delete takes name of the input and deletes it. Returns an error if one occurs.
func (c *FakeInputs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inputsResource, c.ns, name), &v1alpha2.Input{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInputs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inputsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.InputList{})
	return err
}

// Patch applies the patch and returns the patched input.
func (c *FakeInputs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Input, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inputsResource, c.ns, name, pt, data, subresources...), &v1alpha2.Input{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Input), err
}