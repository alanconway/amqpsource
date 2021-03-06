/*
Copyright 2018 The Knative Authors

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

// Someday...  // Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing-sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/knative/eventing-sources/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AmqpSourcesGetter has a method to return a AmqpSourceInterface.
// A group's client should implement this interface.
type AmqpSourcesGetter interface {
	AmqpSources(namespace string) AmqpSourceInterface
}

// AmqpSourceInterface has methods to work with AmqpSource resources.
type AmqpSourceInterface interface {
	Create(*v1alpha1.AmqpSource) (*v1alpha1.AmqpSource, error)
	Update(*v1alpha1.AmqpSource) (*v1alpha1.AmqpSource, error)
	UpdateStatus(*v1alpha1.AmqpSource) (*v1alpha1.AmqpSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AmqpSource, error)
	List(opts v1.ListOptions) (*v1alpha1.AmqpSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AmqpSource, err error)
	AmqpSourceExpansion
}

// amqpSources implements AmqpSourceInterface
type amqpSources struct {
	client rest.Interface
	ns     string
}

// newAmqpSources returns a AmqpSources
func newAmqpSources(c *SourcesV1alpha1Client, namespace string) *amqpSources {
	return &amqpSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the amqpSource, and returns the corresponding amqpSource object, and an error if there is any.
func (c *amqpSources) Get(name string, options v1.GetOptions) (result *v1alpha1.AmqpSource, err error) {
	result = &v1alpha1.AmqpSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("amqpsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AmqpSources that match those selectors.
func (c *amqpSources) List(opts v1.ListOptions) (result *v1alpha1.AmqpSourceList, err error) {
	result = &v1alpha1.AmqpSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("amqpsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested amqpSources.
func (c *amqpSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("amqpsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a amqpSource and creates it.  Returns the server's representation of the amqpSource, and an error, if there is any.
func (c *amqpSources) Create(amqpSource *v1alpha1.AmqpSource) (result *v1alpha1.AmqpSource, err error) {
	result = &v1alpha1.AmqpSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("amqpsources").
		Body(amqpSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a amqpSource and updates it. Returns the server's representation of the amqpSource, and an error, if there is any.
func (c *amqpSources) Update(amqpSource *v1alpha1.AmqpSource) (result *v1alpha1.AmqpSource, err error) {
	result = &v1alpha1.AmqpSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("amqpsources").
		Name(amqpSource.Name).
		Body(amqpSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *amqpSources) UpdateStatus(amqpSource *v1alpha1.AmqpSource) (result *v1alpha1.AmqpSource, err error) {
	result = &v1alpha1.AmqpSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("amqpsources").
		Name(amqpSource.Name).
		SubResource("status").
		Body(amqpSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the amqpSource and deletes it. Returns an error if one occurs.
func (c *amqpSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("amqpsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *amqpSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("amqpsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched amqpSource.
func (c *amqpSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AmqpSource, err error) {
	result = &v1alpha1.AmqpSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("amqpsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
