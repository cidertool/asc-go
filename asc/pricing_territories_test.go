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

func TestListTerritories(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritories(ctx, &ListTerritoriesQuery{})
	})
}

func TestListTerritoriesForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritoriesForApp(ctx, "10", &ListTerritoriesQuery{})
	})
}

func TestListTerritoriesForEULA(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritoriesForEULA(ctx, "10", &ListTerritoriesQuery{})
	})
}

func TestGetTerritoryForAppPrice(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &TerritoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetTerritoryForAppPrice(ctx, "10", &ListTerritoriesQuery{})
	})
}
