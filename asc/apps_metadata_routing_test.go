package asc

import (
	"context"
	"testing"
)

func TestGetRoutingAppCoverageForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &RoutingAppCoverageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetRoutingAppCoverageForAppStoreVersion(ctx, "10", &GetRoutingAppCoverageForVersionQuery{})
	})
}

func TestGetRoutingAppCoverage(t *testing.T) {
	testEndpointWithResponse(t, "{}", &RoutingAppCoverageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetRoutingAppCoverage(ctx, "10", &GetRoutingAppCoverageQuery{})
	})
}

func TestCreateRoutingAppCoverage(t *testing.T) {
	testEndpointWithResponse(t, "{}", &RoutingAppCoverageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateRoutingAppCoverage(ctx, "", 0, "")
	})
}

func TestCommitRoutingAppCoverage(t *testing.T) {
	testEndpointWithResponse(t, "{}", &RoutingAppCoverageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CommitRoutingAppCoverage(ctx, "10", Bool(true), String("10"))
	})
}

func TestDeleteRoutingAppCoverage(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteRoutingAppCoverage(ctx, "10")
	})
}
