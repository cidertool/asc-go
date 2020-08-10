package asc

import (
	"context"
	"testing"
)

func TestGetAppPreview(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppPreview(ctx, "10", &GetAppPreviewQuery{})
	})
}

func TestCreateAppPreview(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppPreview(ctx, &AppPreviewCreateRequest{})
	})
}

func TestCreateAppPreviewNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppPreview(ctx, nil)
	})
}

func TestCommitAppPreview(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CommitAppPreview(ctx, "10", &AppPreviewUpdateRequest{})
	})
}

func TestCommitAppPreviewNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CommitAppPreview(ctx, "10", nil)
	})
}

func TestDeleteAppPreview(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppPreview(ctx, "10")
	})
}
