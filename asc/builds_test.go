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

func TestListBuilds(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListBuilds(ctx, &ListBuildsQuery{})
	})
}

func TestListBuildsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListBuildsForApp(ctx, "10", &ListBuildsForAppQuery{})
	})
}

func TestGetBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetBuild(ctx, "10", &GetBuildQuery{})
	})
}

func TestGetBuildIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[
		{"type":"preReleaseVersions"},{"type":"betaTesters"},{"type":"betaBuildLocalizations"},
		{"type":"appEncryptionDeclarations"},{"type":"betaAppReviewSubmissions"},{"type":"apps"},
		{"type":"buildBetaDetails"},{"type":"appStoreVersions"},{"type":"buildIcons"},
		{"type":"perfPowerMetrics"},{"type":"diagnosticSignatures"}
		]}`, func(ctx context.Context, client *Client) {
		build, _, err := client.Builds.GetBuild(ctx, "10", &GetBuildQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, build.Included)

		assert.NotNil(t, build.Included[0].PrereleaseVersion())
		assert.NotNil(t, build.Included[1].BetaTester())
		assert.NotNil(t, build.Included[2].BetaBuildLocalization())
		assert.NotNil(t, build.Included[3].AppEncryptionDeclaration())
		assert.NotNil(t, build.Included[4].BetaAppReviewSubmission())
		assert.NotNil(t, build.Included[5].App())
		assert.NotNil(t, build.Included[6].BuildBetaDetail())
		assert.NotNil(t, build.Included[7].AppStoreVersion())
		assert.NotNil(t, build.Included[8].BuildIcon())
		assert.NotNil(t, build.Included[9].PerfPowerMetric())
		assert.NotNil(t, build.Included[10].DiagnosticSignature())

		assert.Nil(t, build.Included[0].BetaTester())
		assert.Nil(t, build.Included[0].BetaBuildLocalization())
		assert.Nil(t, build.Included[0].AppEncryptionDeclaration())
		assert.Nil(t, build.Included[0].BetaAppReviewSubmission())
		assert.Nil(t, build.Included[0].App())
		assert.Nil(t, build.Included[0].BuildBetaDetail())
		assert.Nil(t, build.Included[0].AppStoreVersion())
		assert.Nil(t, build.Included[0].BuildIcon())
		assert.Nil(t, build.Included[0].PerfPowerMetric())
		assert.Nil(t, build.Included[0].DiagnosticSignature())
	})
}

func TestGetAppForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppForBuild(ctx, "10", &GetAppForBuildQuery{})
	})
}

func TestGetAppStoreVersionForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppStoreVersionForBuild(ctx, "10", &GetAppStoreVersionForBuildQuery{})
	})
}

func TestGetBuildForAppStoreVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetBuildForAppStoreVersion(ctx, "10", &GetBuildForAppStoreVersionQuery{})
	})
}

func TestUpdateBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.UpdateBuild(ctx, "10", Bool(true), nil, String("10"))
	})
}

func TestUpdateAppEncryptionDeclarationForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.UpdateAppEncryptionDeclarationForBuild(ctx, "10", String("10"))
	})
}

func TestCreateAccessForBetaGroupsToBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.CreateAccessForBetaGroupsToBuild(ctx, "10", []string{"10"})
	})
}

func TestRemoveAccessForBetaGroupsFromBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.RemoveAccessForBetaGroupsFromBuild(ctx, "10", []string{"10"})
	})
}

func TestCreateAccessForIndividualTestersToBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.CreateAccessForIndividualTestersToBuild(ctx, "10", []string{"10"})
	})
}

func TestRemoveAccessForIndividualTestersFromBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.RemoveAccessForIndividualTestersFromBuild(ctx, "10", []string{"10"})
	})
}

func TestListResourceIDsForIndividualTestersForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildIndividualTestersLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListResourceIDsForIndividualTestersForBuild(ctx, "10", &ListResourceIDsForIndividualTestersForBuildQuery{})
	})
}

func TestGetAppEncryptionDeclarationForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclarationForBuild(ctx, "10", &GetAppEncryptionDeclarationForBuildQuery{})
	})
}

func TestGetAppEncryptionDeclarationIDForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BuildAppEncryptionDeclarationLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclarationIDForBuild(ctx, "10")
	})
}
