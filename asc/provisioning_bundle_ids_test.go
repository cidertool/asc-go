package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateBundleID(ctx, BundleIDCreateRequestAttributes{})
	})
}

func TestUpdateBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.UpdateBundleID(ctx, "10", String(""))
	})
}

func TestDeleteBundleID(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.DeleteBundleID(ctx, "10")
	})
}

func TestListBundleIDs(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListBundleIDs(ctx, &ListBundleIDsQuery{})
	})
}

func TestGetBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetBundleID(ctx, "10", &GetBundleIDQuery{})
	})
}

func TestGetBundleIDIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"profiles"},{"type":"bundleIdCapabilities"},{"type":"apps"}]}`, func(ctx context.Context, client *Client) {
		bundleID, _, err := client.Provisioning.GetBundleID(ctx, "10", &GetBundleIDQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, bundleID.Included)

		assert.NotNil(t, bundleID.Included[0].Profile())
		assert.NotNil(t, bundleID.Included[1].BundleIDCapability())
		assert.NotNil(t, bundleID.Included[2].App())

		assert.Nil(t, bundleID.Included[0].BundleIDCapability())
		assert.Nil(t, bundleID.Included[0].App())
		assert.Nil(t, bundleID.Included[1].Profile())
	})
}

func TestGetAppForBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetAppForBundleID(ctx, "10", &GetAppForBundleIDQuery{})
	})
}

func TestListProfilesForBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &ProfilesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListProfilesForBundleID(ctx, "10", &ListProfilesForBundleIDQuery{})
	})
}

func TestListCapabilitiesForBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDCapabilitiesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCapabilitiesForBundleID(ctx, "10", &ListCapabilitiesForBundleIDQuery{})
	})
}
