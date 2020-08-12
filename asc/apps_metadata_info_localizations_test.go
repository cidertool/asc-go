package asc

import (
	"context"
	"testing"
)

func TestListAppInfoLocalizationsForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppInfoLocalizationsForAppInfo(ctx, "10", &ListAppInfoLocalizationsForAppInfoQuery{})
	})
}

func TestGetAppInfoLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppInfoLocalization(ctx, "10", &GetAppInfoLocalizationQuery{})
	})
}

func TestCreateAppInfoLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppInfoLocalization(ctx, AppInfoLocalizationCreateRequest{})
	})
}

func TestUpdateAppInfoLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppInfoLocalization(ctx, "10", AppInfoLocalizationUpdateRequest{})
	})
}

func TestDeleteAppInfoLocalization(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppInfoLocalization(ctx, "10")
	})
}
