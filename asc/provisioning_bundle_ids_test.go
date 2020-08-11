package asc

import (
	"context"
	"testing"
)

func TestCreateBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateBundleID(ctx, &BundleIDCreateRequest{})
	})
}

func TestCreateBundleIDApplyRequestTypes(t *testing.T) {

}

func TestUpdateBundleID(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.UpdateBundleID(ctx, "10", &BundleIDUpdateRequest{})
	})
}

func TestUpdateBundleIDApplyRequestTypes(t *testing.T) {

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
