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

func TestListAppStoreVersionsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppStoreVersionsForApp(ctx, "10", &ListAppStoreVersionsQuery{})
	})
}

func TestGetAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppStoreVersion(ctx, "10", &GetAppStoreVersionQuery{})
	})
}

func TestGetAppStoreVersionIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[
		{"type":"ageRatingDeclarations"},{"type":"appStoreVersionLocalizations"},
		{"type":"builds"},{"type":"appStoreVersionPhasedReleases"},
		{"type":"routingAppCoverages"},{"type":"appStoreReviewDetails"},
		{"type":"appStoreVersionSubmissions"},{"type":"idfaDeclarations"}
		]}`, func(ctx context.Context, client *Client) {
		version, _, err := client.Apps.GetAppStoreVersion(ctx, "10", &GetAppStoreVersionQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, version.Included)

		assert.NotNil(t, version.Included[0].AgeRatingDeclaration())
		assert.NotNil(t, version.Included[1].AppStoreVersionLocalization())
		assert.NotNil(t, version.Included[2].Build())
		assert.NotNil(t, version.Included[3].AppStoreVersionPhasedRelease())
		assert.NotNil(t, version.Included[4].RoutingAppCoverage())
		assert.NotNil(t, version.Included[5].AppStoreReviewDetail())
		assert.NotNil(t, version.Included[6].AppStoreVersionSubmission())
		assert.NotNil(t, version.Included[7].IDFADeclaration())

		assert.Nil(t, version.Included[0].AppStoreVersionLocalization())
		assert.Nil(t, version.Included[0].Build())
		assert.Nil(t, version.Included[0].AppStoreVersionPhasedRelease())
		assert.Nil(t, version.Included[0].RoutingAppCoverage())
		assert.Nil(t, version.Included[0].AppStoreReviewDetail())
		assert.Nil(t, version.Included[0].AppStoreVersionSubmission())
		assert.Nil(t, version.Included[0].IDFADeclaration())
		assert.Nil(t, version.Included[1].AgeRatingDeclaration())
	})
}

func TestCreateAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppStoreVersion(ctx, AppStoreVersionCreateRequestAttributes{}, "10", String("10"))
	})
}

func TestUpdateAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppStoreVersion(ctx, "10", &AppStoreVersionUpdateRequestAttributes{}, String("10"))
	})
}

func TestDeleteAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppStoreVersion(ctx, "10")
	})
}

func TestGetBuildIDForAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionBuildLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetBuildIDForAppStoreVersion(ctx, "10")
	})
}

func TestUpdateBuildForAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionBuildLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateBuildForAppStoreVersion(ctx, "10", String("10"))
	})
}
