package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	testEndpointWithResponse(t, "{}", &ProfileResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateProfile(ctx, "", "", "", []string{"10"}, []string{"10"})
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

func TestGetProfileIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"bundleIds"},{"type":"certificates"},{"type":"devices"}]}`, func(ctx context.Context, client *Client) {
		profile, _, err := client.Provisioning.GetProfile(ctx, "10", &GetProfileQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, profile.Included)

		assert.NotNil(t, profile.Included[0].BundleID())
		assert.NotNil(t, profile.Included[1].Certificate())
		assert.NotNil(t, profile.Included[2].Device())

		assert.Nil(t, profile.Included[0].Certificate())
		assert.Nil(t, profile.Included[0].Device())
		assert.Nil(t, profile.Included[1].BundleID())
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
