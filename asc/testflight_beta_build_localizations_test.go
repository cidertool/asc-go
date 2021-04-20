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
)

func TestListBetaBuildLocalizations(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaBuildLocalizations(ctx, &ListBetaBuildLocalizationsQuery{})
	})
}

func TestGetBetaBuildLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaBuildLocalization(ctx, "10", &GetBetaBuildLocalizationQuery{})
	})
}

func TestGetBuildForBetaBuildLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildForBetaBuildLocalization(ctx, "10", &GetBuildForBetaBuildLocalizationQuery{})
	})
}

func TestListBetaBuildLocalizationsForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaBuildLocalizationsForBuild(ctx, "10", &ListBetaBuildLocalizationsForBuildQuery{})
	})
}

func TestCreateBetaBuildLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaBuildLocalization(ctx, "", nil, "")
	})
}

func TestUpdateBetaBuildLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaBuildLocalization(ctx, "10", String(""))
	})
}

func TestDeleteBetaBuildLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaBuildLocalization(ctx, "10")
	})
}
