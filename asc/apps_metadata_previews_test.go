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

func TestCreateAppPreviewApplyRequestTypes(t *testing.T) {

}

func TestCommitAppPreview(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CommitAppPreview(ctx, "10", &AppPreviewUpdateRequest{})
	})
}

func TestCommitAppPreviewApplyRequestTypes(t *testing.T) {

}

func TestDeleteAppPreview(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppPreview(ctx, "10")
	})
}
