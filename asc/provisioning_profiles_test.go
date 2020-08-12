package asc

import (
	"context"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &ProfileResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateProfile(ctx, &ProfileCreateRequest{})
	})
}

func TestDeleteProfile(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.DeleteProfile(ctx, "10")
	})
}

func TestListProfiles(t *testing.T) {
	testEndpointWithResponse(t, "{}", &ProfilesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListProfiles(ctx, &ListProfilesQuery{})
	})
}

func TestGetProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &ProfileResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetProfile(ctx, "10", &GetProfileQuery{})
	})
}

func TestGetBundleIDForProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetBundleIDForProfile(ctx, "10", &GetBundleIDForProfileQuery{})
	})
}

func TestListCertificatesInProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &CertificatesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCertificatesInProfile(ctx, "10", &ListCertificatesForProfileQuery{})
	})
}

func TestListDevicesInProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DevicesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListDevicesInProfile(ctx, "10", &ListDevicesInProfileQuery{})
	})
}
