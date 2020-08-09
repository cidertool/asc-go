package asc

import (
	"context"
	"testing"
)

func TestListLocalizationsForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListLocalizationsForAppStoreVersion(ctx, "10", &ListLocalizationsForAppStoreVersionQuery{})
	})
}

func TestGetAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppStoreVersionLocalization(ctx, "10", &GetAppStoreVersionLocalizationQuery{})
	})
}

func TestCreateAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppStoreVersionLocalization(ctx, &AppStoreVersionLocalizationCreateRequest{})
	})
}

func TestUpdateAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppStoreVersionLocalization(ctx, "10", &AppStoreVersionLocalizationUpdateRequest{})
	})
}

func TestDeleteAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppStoreVersionLocalization(ctx, "10")
	})
}

func TestListAppScreenshotSetsForAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotSetsForAppStoreVersionLocalization(ctx, "10", &ListAppScreenshotSetsForAppStoreVersionLocalizationQuery{})
	})
}

func TestListAppPreviewSetsForAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewSetsForAppStoreVersionLocalization(ctx, "10", &ListAppPreviewSetsForAppStoreVersionLocalizationQuery{})
	})
}
