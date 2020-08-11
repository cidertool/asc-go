package asc

import (
	"context"
	"testing"
)

func TestGetAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppInfo(ctx, "10", &GetAppInfoQuery{})
	})
}

func TestListAppInfosForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfosResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppInfosForApp(ctx, "10", &ListAppInfosForAppQuery{})
	})
}

func TestUpdateAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppInfo(ctx, "10", &AppInfoUpdateRequest{})
	})
}

func TestUpdateAppInfoApplyRequestTypes(t *testing.T) {

}
