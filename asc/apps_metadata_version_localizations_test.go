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

func TestListLocalizationsForAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListLocalizationsForAppStoreVersion(ctx, "10", &ListLocalizationsForAppStoreVersionQuery{})
	})
}

func TestGetAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppStoreVersionLocalization(ctx, "10", &GetAppStoreVersionLocalizationQuery{})
	})
}

func TestGetAppStoreVersionLocalizationIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[{"type":"appScreenshotSets"},{"type":"appPreviewSets"}]}`, func(ctx context.Context, client *Client) {
		localization, _, err := client.Apps.GetAppStoreVersionLocalization(ctx, "10", &GetAppStoreVersionLocalizationQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, localization.Included)

		assert.NotNil(t, localization.Included[0].AppScreenshotSet())
		assert.NotNil(t, localization.Included[1].AppPreviewSet())

		assert.Nil(t, localization.Included[0].AppPreviewSet())
		assert.Nil(t, localization.Included[1].AppScreenshotSet())
	})
}

func TestCreateAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppStoreVersionLocalization(ctx, AppStoreVersionLocalizationCreateRequestAttributes{}, "")
	})
}

func TestUpdateAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppStoreVersionLocalization(ctx, "10", &AppStoreVersionLocalizationUpdateRequestAttributes{})
	})
}

func TestDeleteAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppStoreVersionLocalization(ctx, "10")
	})
}

func TestListAppScreenshotSetsForAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppScreenshotSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotSetsForAppStoreVersionLocalization(ctx, "10", &ListAppScreenshotSetsForAppStoreVersionLocalizationQuery{})
	})
}

func TestListAppPreviewSetsForAppStoreVersionLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewSetsForAppStoreVersionLocalization(ctx, "10", &ListAppPreviewSetsForAppStoreVersionLocalizationQuery{})
	})
}
