package asc

import (
	"context"
	"testing"
)

func TestListApps(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListApps(ctx, &ListAppsQuery{})
	})
}

func TestGetApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetApp(ctx, "10", &GetAppQuery{})
	})
}

func TestUpdateApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateApp(ctx, "10", &AppUpdateRequest{})
	})
}

func TestUpdateAppApplyRequestTypes(t *testing.T) {

}

func TestRemoveBetaTestersFromApp(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.RemoveBetaTestersFromApp(ctx, "10", AppBetaTestersLinkagesRequest{})
	})
}

func TestRemoveBetaTestersFromAppApplyRequestTypes(t *testing.T) {

}

func TestListInAppPurchasesForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &InAppPurchasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListInAppPurchasesForApp(ctx, "10", &ListInAppPurchasesQuery{})
	})
}

func TestGetInAppPurchase(t *testing.T) {
	testEndpointWithResponse(t, "{}", &InAppPurchaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetInAppPurchase(ctx, "10", &GetInAppPurchaseQuery{})
	})
}
