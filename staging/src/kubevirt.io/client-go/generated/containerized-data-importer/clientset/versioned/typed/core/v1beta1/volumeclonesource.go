/*
Copyright 2023 The KubeVirt Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned/scheme"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// VolumeCloneSourcesGetter has a method to return a VolumeCloneSourceInterface.
// A group's client should implement this interface.
type VolumeCloneSourcesGetter interface {
	VolumeCloneSources(namespace string) VolumeCloneSourceInterface
}

// VolumeCloneSourceInterface has methods to work with VolumeCloneSource resources.
type VolumeCloneSourceInterface interface {
	Create(ctx context.Context, volumeCloneSource *v1beta1.VolumeCloneSource, opts v1.CreateOptions) (*v1beta1.VolumeCloneSource, error)
	Update(ctx context.Context, volumeCloneSource *v1beta1.VolumeCloneSource, opts v1.UpdateOptions) (*v1beta1.VolumeCloneSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeCloneSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeCloneSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeCloneSource, err error)
	VolumeCloneSourceExpansion
}

// volumeCloneSources implements VolumeCloneSourceInterface
type volumeCloneSources struct {
	client rest.Interface
	ns     string
}

// newVolumeCloneSources returns a VolumeCloneSources
func newVolumeCloneSources(c *CdiV1beta1Client, namespace string) *volumeCloneSources {
	return &volumeCloneSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeCloneSource, and returns the corresponding volumeCloneSource object, and an error if there is any.
func (c *volumeCloneSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeCloneSource, err error) {
	result = &v1beta1.VolumeCloneSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeclonesources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeCloneSources that match those selectors.
func (c *volumeCloneSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeCloneSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeCloneSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeclonesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeCloneSources.
func (c *volumeCloneSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumeclonesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeCloneSource and creates it.  Returns the server's representation of the volumeCloneSource, and an error, if there is any.
func (c *volumeCloneSources) Create(ctx context.Context, volumeCloneSource *v1beta1.VolumeCloneSource, opts v1.CreateOptions) (result *v1beta1.VolumeCloneSource, err error) {
	result = &v1beta1.VolumeCloneSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumeclonesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeCloneSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeCloneSource and updates it. Returns the server's representation of the volumeCloneSource, and an error, if there is any.
func (c *volumeCloneSources) Update(ctx context.Context, volumeCloneSource *v1beta1.VolumeCloneSource, opts v1.UpdateOptions) (result *v1beta1.VolumeCloneSource, err error) {
	result = &v1beta1.VolumeCloneSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumeclonesources").
		Name(volumeCloneSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeCloneSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeCloneSource and deletes it. Returns an error if one occurs.
func (c *volumeCloneSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeclonesources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeCloneSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeclonesources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeCloneSource.
func (c *volumeCloneSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeCloneSource, err error) {
	result = &v1beta1.VolumeCloneSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumeclonesources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
