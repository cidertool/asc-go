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

func TestListAppPriceTiers(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPriceTiersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListAppPriceTiers(ctx, &ListAppPriceTiersQuery{})
	})
}

func TestGetAppPriceTier(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPriceTierResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetAppPriceTier(ctx, "10", &GetAppPriceTierQuery{})
	})
}

func TestListPricePointsForAppPriceTier(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPricePointsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListPricePointsForAppPriceTier(ctx, "10", &ListPricePointsForAppPriceTierQuery{})
	})
}

func TestListAppPricePoints(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPricePointsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListAppPricePoints(ctx, &ListAppPricePointsQuery{})
	})
}

func TestGetTerritoryForAppPricePoint(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &TerritoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetTerritoryForAppPricePoint(ctx, "10", &GetTerritoryForAppPricePointQuery{})
	})
}

func TestGetAppPricePoint(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPricePointResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetAppPricePoint(ctx, "10", &GetAppPricePointQuery{})
	})
}
