// Copyright Red Hat

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/identitatem/idp-client-api/api/identitatem/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOAuths implements ClusterOAuthInterface
type FakeClusterOAuths struct {
	Fake *FakeIdentityconfigV1alpha1
	ns   string
}

var clusteroauthsResource = schema.GroupVersionResource{Group: "identityconfig.identitatem.io", Version: "v1alpha1", Resource: "clusteroauths"}

var clusteroauthsKind = schema.GroupVersionKind{Group: "identityconfig.identitatem.io", Version: "v1alpha1", Kind: "ClusterOAuth"}

// Get takes name of the clusterOAuth, and returns the corresponding clusterOAuth object, and an error if there is any.
func (c *FakeClusterOAuths) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterOAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteroauthsResource, c.ns, name), &v1alpha1.ClusterOAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOAuth), err
}

// List takes label and field selectors, and returns the list of ClusterOAuths that match those selectors.
func (c *FakeClusterOAuths) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterOAuthList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteroauthsResource, clusteroauthsKind, c.ns, opts), &v1alpha1.ClusterOAuthList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterOAuthList{ListMeta: obj.(*v1alpha1.ClusterOAuthList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterOAuthList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOAuths.
func (c *FakeClusterOAuths) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteroauthsResource, c.ns, opts))

}

// Create takes the representation of a clusterOAuth and creates it.  Returns the server's representation of the clusterOAuth, and an error, if there is any.
func (c *FakeClusterOAuths) Create(ctx context.Context, clusterOAuth *v1alpha1.ClusterOAuth, opts v1.CreateOptions) (result *v1alpha1.ClusterOAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteroauthsResource, c.ns, clusterOAuth), &v1alpha1.ClusterOAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOAuth), err
}

// Update takes the representation of a clusterOAuth and updates it. Returns the server's representation of the clusterOAuth, and an error, if there is any.
func (c *FakeClusterOAuths) Update(ctx context.Context, clusterOAuth *v1alpha1.ClusterOAuth, opts v1.UpdateOptions) (result *v1alpha1.ClusterOAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteroauthsResource, c.ns, clusterOAuth), &v1alpha1.ClusterOAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOAuth), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOAuths) UpdateStatus(ctx context.Context, clusterOAuth *v1alpha1.ClusterOAuth, opts v1.UpdateOptions) (*v1alpha1.ClusterOAuth, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusteroauthsResource, "status", c.ns, clusterOAuth), &v1alpha1.ClusterOAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOAuth), err
}

// Delete takes name of the clusterOAuth and deletes it. Returns an error if one occurs.
func (c *FakeClusterOAuths) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusteroauthsResource, c.ns, name), &v1alpha1.ClusterOAuth{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOAuths) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteroauthsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterOAuthList{})
	return err
}

// Patch applies the patch and returns the patched clusterOAuth.
func (c *FakeClusterOAuths) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterOAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteroauthsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterOAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOAuth), err
}