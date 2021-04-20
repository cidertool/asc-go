/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ProfileResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateProfile(ctx, "", "", "", []string{"10"}, []string{"10"})
	})
}

func TestDeleteProfile(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.DeleteProfile(ctx, "10")
	})
}

func TestListProfiles(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ProfilesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListProfiles(ctx, &ListProfilesQuery{})
	})
}

func TestGetProfile(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ProfileResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetProfile(ctx, "10", &GetProfileQuery{})
	})
}

func TestGetProfileIncludeds(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetBundleIDForProfile(ctx, "10", &GetBundleIDForProfileQuery{})
	})
}

func TestListCertificatesInProfile(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CertificatesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCertificatesInProfile(ctx, "10", &ListCertificatesForProfileQuery{})
	})
}

func TestListDevicesInProfile(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &DevicesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListDevicesInProfile(ctx, "10", &ListDevicesInProfileQuery{})
	})
}
