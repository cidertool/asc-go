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

func TestCreateBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaGroup(ctx, BetaGroupCreateRequestAttributes{}, "", []string{"10"}, []string{"10"})
	})
}

func TestUpdateBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaGroup(ctx, "10", &BetaGroupUpdateRequestAttributes{})
	})
}

func TestDeleteBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaGroup(ctx, "10")
	})
}

func TestListBetaGroups(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroups(ctx, &ListBetaGroupsQuery{})
	})
}

func TestGetBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaGroup(ctx, "10", &GetBetaGroupQuery{})
	})
}

func TestGetBetaGroupIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[{"type":"apps"},{"type":"builds"},{"type":"betaTesters"}]}`, func(ctx context.Context, client *Client) {
		group, _, err := client.TestFlight.GetBetaGroup(ctx, "10", &GetBetaGroupQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, group.Included)

		assert.NotNil(t, group.Included[0].App())
		assert.NotNil(t, group.Included[1].Build())
		assert.NotNil(t, group.Included[2].BetaTester())

		assert.Nil(t, group.Included[0].Build())
		assert.Nil(t, group.Included[0].BetaTester())
	})
}

func TestGetAppForBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaGroup(ctx, "10", &GetAppForBetaGroupQuery{})
	})
}

func TestListBetaGroupsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroupsForApp(ctx, "10", &ListBetaGroupsForAppQuery{})
	})
}

func TestAddBetaTestersToBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AddBetaTestersToBetaGroup(ctx, "10", []string{"10"})
	})
}

func TestRemoveBetaTestersFromBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveBetaTestersFromBetaGroup(ctx, "10", []string{"10"})
	})
}

func TestAddBuildsToBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AddBuildsToBetaGroup(ctx, "10", []string{"10"})
	})
}

func TestRemoveBuildsFromBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveBuildsFromBetaGroup(ctx, "10", []string{"10"})
	})
}

func TestListBuildsForBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildsForBetaGroup(ctx, "10", &ListBuildsForBetaGroupQuery{})
	})
}

func TestListBuildIDsForBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupBuildsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildIDsForBetaGroup(ctx, "10", &ListBuildIDsForBetaGroupQuery{})
	})
}

func TestListBetaTestersForBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaTestersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaTestersForBetaGroup(ctx, "10", &ListBetaTestersForBetaGroupQuery{})
	})
}

func TestListBetaTesterIDsForBetaGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaGroupBetaTestersLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaTesterIDsForBetaGroup(ctx, "10", &ListBetaTesterIDsForBetaGroupQuery{})
	})
}
