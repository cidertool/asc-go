package asc

import (
	"context"
	"testing"
)

func TestListBuilds(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListBuilds(ctx, &ListBuildsQuery{})
	})
}

func TestListBuildsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListBuildsForApp(ctx, "10", &ListBuildsForAppQuery{})
	})
}

func TestGetBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetBuild(ctx, "10", &GetBuildQuery{})
	})
}

func TestGetAppForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppForBuild(ctx, "10", &GetAppForBuildQuery{})
	})
}

func TestGetAppStoreVersionForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppStoreVersionForBuild(ctx, "10", &GetAppStoreVersionForBuildQuery{})
	})
}

func TestGetBuildForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetBuildForAppStoreVersion(ctx, "10", &GetBuildForAppStoreVersionQuery{})
	})
}

func TestUpdateBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.UpdateBuild(ctx, "10", &BuildUpdateRequest{})
	})
}

func TestUpdateAppEncryptionDeclarationForBuild(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.UpdateAppEncryptionDeclarationForBuild(ctx, "10", &RelationshipData{})
	})
}

func TestCreateAccessForBetaGroupsToBuild(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.CreateAccessForBetaGroupsToBuild(ctx, "10", &[]RelationshipData{})
	})
}

func TestRemoveAccessForBetaGroupsFromBuild(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.RemoveAccessForBetaGroupsFromBuild(ctx, "10", &[]RelationshipData{})
	})
}

func TestCreateAccessForIndividualTestersToBuild(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.CreateAccessForIndividualTestersToBuild(ctx, "10", &[]RelationshipData{})
	})
}

func TestRemoveAccessForIndividualTestersFromBuild(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.RemoveAccessForIndividualTestersFromBuild(ctx, "10", &[]RelationshipData{})
	})
}

func TestListResourceIDsForIndividualTestersForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildIndividualTestersLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListResourceIDsForIndividualTestersForBuild(ctx, "10", &ListResourceIDsForIndividualTestersForBuildQuery{})
	})
}

func TestGetAppEncryptionDeclarationForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclarationForBuild(ctx, "10", &GetAppEncryptionDeclarationForBuildQuery{})
	})
}

func TestGetAppEncryptionDeclarationIDForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildAppEncryptionDeclarationLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclarationIDForBuild(ctx, "10")
	})
}
