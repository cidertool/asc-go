package asc

import (
	"context"
	"testing"
)

func TestGetAppScreenshotSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppScreenshotSet(ctx, "10", &GetAppScreenshotSetQuery{})
	})
}

func TestCreateAppScreenshotSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppScreenshotSet(ctx, &AppScreenshotSetCreateRequest{})
	})
}

func TestDeleteAppScreenshotSet(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppScreenshotSet(ctx, "10")
	})
}

func TestListAppScreenshotsForSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotsForSet(ctx, "10", &ListAppScreenshotsForSetQuery{})
	})
}

func TestListAppScreenshotIDsForSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotSetAppScreenshotsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotIDsForSet(ctx, "10", &ListAppScreenshotIDsForSetQuery{})
	})
}

func TestReplaceAppScreenshotsForSet(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.ReplaceAppScreenshotsForSet(ctx, "10", &[]RelationshipsData{})
	})
}
