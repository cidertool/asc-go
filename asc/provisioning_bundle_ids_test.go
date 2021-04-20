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

func TestCreateBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateBundleID(ctx, BundleIDCreateRequestAttributes{})
	})
}

func TestUpdateBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.UpdateBundleID(ctx, "10", String(""))
	})
}

func TestDeleteBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.DeleteBundleID(ctx, "10")
	})
}

func TestListBundleIDs(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListBundleIDs(ctx, &ListBundleIDsQuery{})
	})
}

func TestGetBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetBundleID(ctx, "10", &GetBundleIDQuery{})
	})
}

func TestGetBundleIDIncludeds(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetAppForBundleID(ctx, "10", &GetAppForBundleIDQuery{})
	})
}

func TestListProfilesForBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ProfilesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListProfilesForBundleID(ctx, "10", &ListProfilesForBundleIDQuery{})
	})
}

func TestListCapabilitiesForBundleID(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BundleIDCapabilitiesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCapabilitiesForBundleID(ctx, "10", &ListCapabilitiesForBundleIDQuery{})
	})
}
