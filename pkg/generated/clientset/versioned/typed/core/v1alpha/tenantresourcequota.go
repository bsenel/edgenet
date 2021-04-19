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

package v1alpha

import (
	"context"
	"time"

	v1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/core/v1alpha"
	scheme "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TenantResourceQuotasGetter has a method to return a TenantResourceQuotaInterface.
// A group's client should implement this interface.
type TenantResourceQuotasGetter interface {
	TenantResourceQuotas() TenantResourceQuotaInterface
}

// TenantResourceQuotaInterface has methods to work with TenantResourceQuota resources.
type TenantResourceQuotaInterface interface {
	Create(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.CreateOptions) (*v1alpha.TenantResourceQuota, error)
	Update(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.UpdateOptions) (*v1alpha.TenantResourceQuota, error)
	UpdateStatus(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.UpdateOptions) (*v1alpha.TenantResourceQuota, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha.TenantResourceQuota, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha.TenantResourceQuotaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.TenantResourceQuota, err error)
	TenantResourceQuotaExpansion
}

// tenantResourceQuotas implements TenantResourceQuotaInterface
type tenantResourceQuotas struct {
	client rest.Interface
}

// newTenantResourceQuotas returns a TenantResourceQuotas
func newTenantResourceQuotas(c *CoreV1alphaClient) *tenantResourceQuotas {
	return &tenantResourceQuotas{
		client: c.RESTClient(),
	}
}

// Get takes name of the tenantResourceQuota, and returns the corresponding tenantResourceQuota object, and an error if there is any.
func (c *tenantResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.TenantResourceQuota, err error) {
	result = &v1alpha.TenantResourceQuota{}
	err = c.client.Get().
		Resource("tenantresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TenantResourceQuotas that match those selectors.
func (c *tenantResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.TenantResourceQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha.TenantResourceQuotaList{}
	err = c.client.Get().
		Resource("tenantresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tenantResourceQuotas.
func (c *tenantResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tenantresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tenantResourceQuota and creates it.  Returns the server's representation of the tenantResourceQuota, and an error, if there is any.
func (c *tenantResourceQuotas) Create(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.CreateOptions) (result *v1alpha.TenantResourceQuota, err error) {
	result = &v1alpha.TenantResourceQuota{}
	err = c.client.Post().
		Resource("tenantresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tenantResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tenantResourceQuota and updates it. Returns the server's representation of the tenantResourceQuota, and an error, if there is any.
func (c *tenantResourceQuotas) Update(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.UpdateOptions) (result *v1alpha.TenantResourceQuota, err error) {
	result = &v1alpha.TenantResourceQuota{}
	err = c.client.Put().
		Resource("tenantresourcequotas").
		Name(tenantResourceQuota.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tenantResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tenantResourceQuotas) UpdateStatus(ctx context.Context, tenantResourceQuota *v1alpha.TenantResourceQuota, opts v1.UpdateOptions) (result *v1alpha.TenantResourceQuota, err error) {
	result = &v1alpha.TenantResourceQuota{}
	err = c.client.Put().
		Resource("tenantresourcequotas").
		Name(tenantResourceQuota.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tenantResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tenantResourceQuota and deletes it. Returns an error if one occurs.
func (c *tenantResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tenantresourcequotas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tenantResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("tenantresourcequotas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tenantResourceQuota.
func (c *tenantResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.TenantResourceQuota, err error) {
	result = &v1alpha.TenantResourceQuota{}
	err = c.client.Patch(pt).
		Resource("tenantresourcequotas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
