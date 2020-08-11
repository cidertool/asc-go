package asc

import (
	"context"
	"testing"
)

func TestListAppStoreVersionsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppStoreVersionsForApp(ctx, "10", &ListAppStoreVersionsQuery{})
	})
}

func TestGetAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppStoreVersion(ctx, "10", &GetAppStoreVersionQuery{})
	})
}

func TestCreateAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppStoreVersion(ctx, &AppStoreVersionCreateRequest{})
	})
}

func TestCreateAppStoreVersionApplyRequestTypes(t *testing.T) {

}

func TestUpdateAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppStoreVersion(ctx, "10", &AppStoreVersionUpdateRequest{})
	})
}

func TestUpdateAppStoreVersionApplyRequestTypes(t *testing.T) {

}

func TestDeleteAppStoreVersion(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppStoreVersion(ctx, "10")
	})
}

func TestGetBuildIDForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionBuildLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetBuildIDForAppStoreVersion(ctx, "10")
	})
}

func TestUpdateBuildForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionBuildLinkageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateBuildForAppStoreVersion(ctx, "10", &AppStoreVersionBuildLinkageRequest{})
	})
}

func TestUpdateBuildForAppStoreVersionApplyRequestTypes(t *testing.T) {

}

func TestGetAgeRatingDeclarationForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AgeRatingDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAgeRatingDeclarationForAppStoreVersion(ctx, "10", &GetAgeRatingDeclarationForAppStoreVersionQuery{})
	})
}
