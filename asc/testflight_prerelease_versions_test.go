package asc

import (
	"context"
	"testing"
)

func TestListPrereleaseVersions(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PrereleaseVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListPrereleaseVersions(ctx, &ListPrereleaseVersionsQuery{})
	})
}

func TestGetPrereleaseVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PrereleaseVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetPrereleaseVersion(ctx, "10", &GetPrereleaseVersionQuery{})
	})
}

func TestGetAppForPrereleaseVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForPrereleaseVersion(ctx, "10", &GetAppForPrereleaseVersionQuery{})
	})
}

func TestListPrereleaseVersionsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PrereleaseVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListPrereleaseVersionsForApp(ctx, "10", &ListPrereleaseVersionsForAppQuery{})
	})
}

func TestListBuildsForPrereleaseVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildsForPrereleaseVersion(ctx, "10", &ListBuildsForPrereleaseVersionQuery{})
	})
}

func TestGetPrereleaseVersionForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PrereleaseVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetPrereleaseVersionForBuild(ctx, "10", &GetPrereleaseVersionForBuildQuery{})
	})
}
