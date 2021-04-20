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

func TestListBetaAppLocalizations(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppLocalizations(ctx, &ListBetaAppLocalizationsQuery{})
	})
}

func TestGetBetaAppLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppLocalization(ctx, "10", &GetBetaAppLocalizationQuery{})
	})
}

func TestGetAppForBetaAppLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaAppLocalization(ctx, "10", &GetAppForBetaAppLocalizationQuery{})
	})
}

func TestListBetaAppLocalizationsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppLocalizationsForApp(ctx, "10", &ListBetaAppLocalizationsForAppQuery{})
	})
}

func TestCreateBetaAppLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaAppLocalization(ctx, BetaAppLocalizationCreateRequestAttributes{}, "")
	})
}

func TestUpdateBetaAppLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaAppLocalization(ctx, "10", &BetaAppLocalizationUpdateRequestAttributes{})
	})
}

func TestDeleteBetaAppLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaAppLocalization(ctx, "10")
	})
}
