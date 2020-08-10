package asc

import (
	"context"
	"testing"
)

func TestGetAppPreviewSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppPreviewSet(ctx, "10", &GetAppPreviewSetQuery{})
	})
}

func TestCreateAppPreviewSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppPreviewSet(ctx, &AppPreviewSetCreateRequest{})
	})
}

func TestDeleteAppPreviewSet(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppPreviewSet(ctx, "10")
	})
}

func TestListAppPreviewsForSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewsForSet(ctx, "10", &ListAppPreviewsForSetQuery{})
	})
}

func TestListAppPreviewIDsForSet(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewSetAppPreviewsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewIDsForSet(ctx, "10", &ListAppPreviewIDsForSetQuery{})
	})
}

func TestReplaceAppPreviewsForSet(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.ReplaceAppPreviewsForSet(ctx, "10", &AppPreviewSetAppPreviewsLinkagesRequest{})
	})
}
