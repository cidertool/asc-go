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

func TestGetAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppInfo(ctx, "10", &GetAppInfoQuery{})
	})
}

func TestGetAppInfoIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[{"type":"appInfoLocalizations"},{"type":"appCategories"}]}`, func(ctx context.Context, client *Client) {
		infos, _, err := client.Apps.GetAppInfo(ctx, "10", &GetAppInfoQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, infos.Included)

		assert.NotNil(t, infos.Included[0].AppInfoLocalization())
		assert.NotNil(t, infos.Included[1].AppCategory())

		assert.Nil(t, infos.Included[0].AppCategory())
		assert.Nil(t, infos.Included[1].AppInfoLocalization())
	})
}

func TestListAppInfosForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfosResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppInfosForApp(ctx, "10", &ListAppInfosForAppQuery{})
	})
}

func TestUpdateAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppInfo(ctx, "10", &AppInfoUpdateRequestRelationships{})
	})
}

func TestGetAgeRatingDeclarationForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AgeRatingDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAgeRatingDeclarationForAppInfo(ctx, "10", &GetAgeRatingDeclarationForAppInfoQuery{})
	})
}
