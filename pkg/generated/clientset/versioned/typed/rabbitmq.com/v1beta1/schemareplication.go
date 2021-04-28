/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/rabbitmq/messaging-topology-operator/api/v1beta1"
	scheme "github.com/rabbitmq/messaging-topology-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchemaReplicationsGetter has a method to return a SchemaReplicationInterface.
// A group's client should implement this interface.
type SchemaReplicationsGetter interface {
	SchemaReplications(namespace string) SchemaReplicationInterface
}

// SchemaReplicationInterface has methods to work with SchemaReplication resources.
type SchemaReplicationInterface interface {
	Create(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.CreateOptions) (*v1beta1.SchemaReplication, error)
	Update(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.UpdateOptions) (*v1beta1.SchemaReplication, error)
	UpdateStatus(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.UpdateOptions) (*v1beta1.SchemaReplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SchemaReplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SchemaReplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SchemaReplication, err error)
	SchemaReplicationExpansion
}

// schemaReplications implements SchemaReplicationInterface
type schemaReplications struct {
	client rest.Interface
	ns     string
}

// newSchemaReplications returns a SchemaReplications
func newSchemaReplications(c *RabbitmqV1beta1Client, namespace string) *schemaReplications {
	return &schemaReplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the schemaReplication, and returns the corresponding schemaReplication object, and an error if there is any.
func (c *schemaReplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SchemaReplication, err error) {
	result = &v1beta1.SchemaReplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schemareplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchemaReplications that match those selectors.
func (c *schemaReplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SchemaReplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SchemaReplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schemareplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schemaReplications.
func (c *schemaReplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schemareplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a schemaReplication and creates it.  Returns the server's representation of the schemaReplication, and an error, if there is any.
func (c *schemaReplications) Create(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.CreateOptions) (result *v1beta1.SchemaReplication, err error) {
	result = &v1beta1.SchemaReplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schemareplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schemaReplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a schemaReplication and updates it. Returns the server's representation of the schemaReplication, and an error, if there is any.
func (c *schemaReplications) Update(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.UpdateOptions) (result *v1beta1.SchemaReplication, err error) {
	result = &v1beta1.SchemaReplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schemareplications").
		Name(schemaReplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schemaReplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *schemaReplications) UpdateStatus(ctx context.Context, schemaReplication *v1beta1.SchemaReplication, opts v1.UpdateOptions) (result *v1beta1.SchemaReplication, err error) {
	result = &v1beta1.SchemaReplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schemareplications").
		Name(schemaReplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schemaReplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the schemaReplication and deletes it. Returns an error if one occurs.
func (c *schemaReplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schemareplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schemaReplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schemareplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched schemaReplication.
func (c *schemaReplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SchemaReplication, err error) {
	result = &v1beta1.SchemaReplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schemareplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
