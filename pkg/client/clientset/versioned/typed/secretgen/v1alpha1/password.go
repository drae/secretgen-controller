// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "carvel.dev/secretgen-controller/pkg/apis/secretgen/v1alpha1"
	scheme "carvel.dev/secretgen-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PasswordsGetter has a method to return a PasswordInterface.
// A group's client should implement this interface.
type PasswordsGetter interface {
	Passwords(namespace string) PasswordInterface
}

// PasswordInterface has methods to work with Password resources.
type PasswordInterface interface {
	Create(ctx context.Context, password *v1alpha1.Password, opts v1.CreateOptions) (*v1alpha1.Password, error)
	Update(ctx context.Context, password *v1alpha1.Password, opts v1.UpdateOptions) (*v1alpha1.Password, error)
	UpdateStatus(ctx context.Context, password *v1alpha1.Password, opts v1.UpdateOptions) (*v1alpha1.Password, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Password, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PasswordList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Password, err error)
	PasswordExpansion
}

// passwords implements PasswordInterface
type passwords struct {
	client rest.Interface
	ns     string
}

// newPasswords returns a Passwords
func newPasswords(c *SecretgenV1alpha1Client, namespace string) *passwords {
	return &passwords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the password, and returns the corresponding password object, and an error if there is any.
func (c *passwords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Password, err error) {
	result = &v1alpha1.Password{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("passwords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Passwords that match those selectors.
func (c *passwords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PasswordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PasswordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("passwords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested passwords.
func (c *passwords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("passwords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a password and creates it.  Returns the server's representation of the password, and an error, if there is any.
func (c *passwords) Create(ctx context.Context, password *v1alpha1.Password, opts v1.CreateOptions) (result *v1alpha1.Password, err error) {
	result = &v1alpha1.Password{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("passwords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(password).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a password and updates it. Returns the server's representation of the password, and an error, if there is any.
func (c *passwords) Update(ctx context.Context, password *v1alpha1.Password, opts v1.UpdateOptions) (result *v1alpha1.Password, err error) {
	result = &v1alpha1.Password{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("passwords").
		Name(password.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(password).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *passwords) UpdateStatus(ctx context.Context, password *v1alpha1.Password, opts v1.UpdateOptions) (result *v1alpha1.Password, err error) {
	result = &v1alpha1.Password{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("passwords").
		Name(password.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(password).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the password and deletes it. Returns an error if one occurs.
func (c *passwords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("passwords").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *passwords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("passwords").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched password.
func (c *passwords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Password, err error) {
	result = &v1alpha1.Password{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("passwords").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
