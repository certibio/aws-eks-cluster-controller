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

package v1alpha1

import (
	v1alpha1 "github.com/awslabs/aws-eks-cluster-controller/pkg/apis/cluster/v1alpha1"
	scheme "github.com/awslabs/aws-eks-cluster-controller/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EKSsGetter has a method to return a EKSInterface.
// A group's client should implement this interface.
type EKSsGetter interface {
	EKSs(namespace string) EKSInterface
}

// EKSInterface has methods to work with EKS resources.
type EKSInterface interface {
	Create(*v1alpha1.EKS) (*v1alpha1.EKS, error)
	Update(*v1alpha1.EKS) (*v1alpha1.EKS, error)
	UpdateStatus(*v1alpha1.EKS) (*v1alpha1.EKS, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EKS, error)
	List(opts v1.ListOptions) (*v1alpha1.EKSList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EKS, err error)
	EKSExpansion
}

// eKSs implements EKSInterface
type eKSs struct {
	client rest.Interface
	ns     string
}

// newEKSs returns a EKSs
func newEKSs(c *ClusterV1alpha1Client, namespace string) *eKSs {
	return &eKSs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eKS, and returns the corresponding eKS object, and an error if there is any.
func (c *eKSs) Get(name string, options v1.GetOptions) (result *v1alpha1.EKS, err error) {
	result = &v1alpha1.EKS{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ekss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EKSs that match those selectors.
func (c *eKSs) List(opts v1.ListOptions) (result *v1alpha1.EKSList, err error) {
	result = &v1alpha1.EKSList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ekss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eKSs.
func (c *eKSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ekss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eKS and creates it.  Returns the server's representation of the eKS, and an error, if there is any.
func (c *eKSs) Create(eKS *v1alpha1.EKS) (result *v1alpha1.EKS, err error) {
	result = &v1alpha1.EKS{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ekss").
		Body(eKS).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eKS and updates it. Returns the server's representation of the eKS, and an error, if there is any.
func (c *eKSs) Update(eKS *v1alpha1.EKS) (result *v1alpha1.EKS, err error) {
	result = &v1alpha1.EKS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ekss").
		Name(eKS.Name).
		Body(eKS).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eKSs) UpdateStatus(eKS *v1alpha1.EKS) (result *v1alpha1.EKS, err error) {
	result = &v1alpha1.EKS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ekss").
		Name(eKS.Name).
		SubResource("status").
		Body(eKS).
		Do().
		Into(result)
	return
}

// Delete takes name of the eKS and deletes it. Returns an error if one occurs.
func (c *eKSs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ekss").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eKSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ekss").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eKS.
func (c *eKSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EKS, err error) {
	result = &v1alpha1.EKS{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ekss").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
