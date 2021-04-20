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

func TestListAppInfoLocalizationsForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppInfoLocalizationsForAppInfo(ctx, "10", &ListAppInfoLocalizationsForAppInfoQuery{})
	})
}

func TestGetAppInfoLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppInfoLocalization(ctx, "10", &GetAppInfoLocalizationQuery{})
	})
}

func TestCreateAppInfoLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppInfoLocalization(ctx, AppInfoLocalizationCreateRequestAttributes{}, "")
	})
}

func TestUpdateAppInfoLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppInfoLocalization(ctx, "10", &AppInfoLocalizationUpdateRequestAttributes{})
	})
}

func TestDeleteAppInfoLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppInfoLocalization(ctx, "10")
	})
}
