package asc

import (
	"context"
	"testing"
)

func TestGetAppScreenshot(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppScreenshot(ctx, "10", &GetAppScreenshotQuery{})
	})
}

func TestCreateAppScreenshot(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppScreenshot(ctx, "", 0, "")
	})
}

func TestCommitAppScreenshot(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CommitAppScreenshot(ctx, "10", AppScreenshotUpdateRequest{})
	})
}

func TestDeleteAppScreenshot(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppScreenshot(ctx, "10")
	})
}
