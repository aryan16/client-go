// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "istio.io/client-go/pkg/apis/telemetry/v1"
	telemetryv1 "istio.io/client-go/pkg/applyconfiguration/telemetry/v1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TelemetriesGetter has a method to return a TelemetryInterface.
// A group's client should implement this interface.
type TelemetriesGetter interface {
	Telemetries(namespace string) TelemetryInterface
}

// TelemetryInterface has methods to work with Telemetry resources.
type TelemetryInterface interface {
	Create(ctx context.Context, telemetry *v1.Telemetry, opts metav1.CreateOptions) (*v1.Telemetry, error)
	Update(ctx context.Context, telemetry *v1.Telemetry, opts metav1.UpdateOptions) (*v1.Telemetry, error)
	UpdateStatus(ctx context.Context, telemetry *v1.Telemetry, opts metav1.UpdateOptions) (*v1.Telemetry, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Telemetry, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TelemetryList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Telemetry, err error)
	Apply(ctx context.Context, telemetry *telemetryv1.TelemetryApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Telemetry, err error)
	ApplyStatus(ctx context.Context, telemetry *telemetryv1.TelemetryApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Telemetry, err error)
	TelemetryExpansion
}

// telemetries implements TelemetryInterface
type telemetries struct {
	client rest.Interface
	ns     string
}

// newTelemetries returns a Telemetries
func newTelemetries(c *TelemetryV1Client, namespace string) *telemetries {
	return &telemetries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the telemetry, and returns the corresponding telemetry object, and an error if there is any.
func (c *telemetries) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Telemetry, err error) {
	result = &v1.Telemetry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("telemetries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Telemetries that match those selectors.
func (c *telemetries) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TelemetryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TelemetryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("telemetries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested telemetries.
func (c *telemetries) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("telemetries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a telemetry and creates it.  Returns the server's representation of the telemetry, and an error, if there is any.
func (c *telemetries) Create(ctx context.Context, telemetry *v1.Telemetry, opts metav1.CreateOptions) (result *v1.Telemetry, err error) {
	result = &v1.Telemetry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("telemetries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(telemetry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a telemetry and updates it. Returns the server's representation of the telemetry, and an error, if there is any.
func (c *telemetries) Update(ctx context.Context, telemetry *v1.Telemetry, opts metav1.UpdateOptions) (result *v1.Telemetry, err error) {
	result = &v1.Telemetry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("telemetries").
		Name(telemetry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(telemetry).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *telemetries) UpdateStatus(ctx context.Context, telemetry *v1.Telemetry, opts metav1.UpdateOptions) (result *v1.Telemetry, err error) {
	result = &v1.Telemetry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("telemetries").
		Name(telemetry.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(telemetry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the telemetry and deletes it. Returns an error if one occurs.
func (c *telemetries) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("telemetries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *telemetries) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("telemetries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched telemetry.
func (c *telemetries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Telemetry, err error) {
	result = &v1.Telemetry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("telemetries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied telemetry.
func (c *telemetries) Apply(ctx context.Context, telemetry *telemetryv1.TelemetryApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Telemetry, err error) {
	if telemetry == nil {
		return nil, fmt.Errorf("telemetry provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(telemetry)
	if err != nil {
		return nil, err
	}
	name := telemetry.Name
	if name == nil {
		return nil, fmt.Errorf("telemetry.Name must be provided to Apply")
	}
	result = &v1.Telemetry{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("telemetries").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *telemetries) ApplyStatus(ctx context.Context, telemetry *telemetryv1.TelemetryApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Telemetry, err error) {
	if telemetry == nil {
		return nil, fmt.Errorf("telemetry provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(telemetry)
	if err != nil {
		return nil, err
	}

	name := telemetry.Name
	if name == nil {
		return nil, fmt.Errorf("telemetry.Name must be provided to Apply")
	}

	result = &v1.Telemetry{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("telemetries").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
