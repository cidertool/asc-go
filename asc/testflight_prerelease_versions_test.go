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

func TestListPrereleaseVersions(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PrereleaseVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListPrereleaseVersions(ctx, &ListPrereleaseVersionsQuery{})
	})
}

func TestGetPrereleaseVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PrereleaseVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetPrereleaseVersion(ctx, "10", &GetPrereleaseVersionQuery{})
	})
}

func TestGetPrereleaseVersionIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[{"type":"apps"},{"type":"builds"}]}`, func(ctx context.Context, client *Client) {
		version, _, err := client.TestFlight.GetPrereleaseVersion(ctx, "10", &GetPrereleaseVersionQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, version.Included)

		assert.NotNil(t, version.Included[0].App())
		assert.NotNil(t, version.Included[1].Build())

		assert.Nil(t, version.Included[0].Build())
	})
}

func TestGetAppForPrereleaseVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForPrereleaseVersion(ctx, "10", &GetAppForPrereleaseVersionQuery{})
	})
}

func TestListPrereleaseVersionsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PrereleaseVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListPrereleaseVersionsForApp(ctx, "10", &ListPrereleaseVersionsForAppQuery{})
	})
}

func TestListBuildsForPrereleaseVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildsForPrereleaseVersion(ctx, "10", &ListBuildsForPrereleaseVersionQuery{})
	})
}

func TestGetPrereleaseVersionForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PrereleaseVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetPrereleaseVersionForBuild(ctx, "10", &GetPrereleaseVersionForBuildQuery{})
	})
}
