package asc

import (
	"context"
	"testing"
)

func TestListBuildBetaDetails(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaDetailsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildBetaDetails(ctx, &ListBuildBetaDetailsQuery{})
	})
}

func TestGetBuildBetaDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildBetaDetail(ctx, "10", &GetBuildBetaDetailsQuery{})
	})
}

func TestGetBuildForBuildBetaDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildForBuildBetaDetail(ctx, "10", &GetBuildForBuildBetaDetailQuery{})
	})
}

func TestGetBuildBetaDetailForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildBetaDetailForBuild(ctx, "10", &GetBuildBetaDetailForBuildQuery{})
	})
}

func TestUpdateBuildBetaDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBuildBetaDetail(ctx, "10", Bool(false))
	})
}
