package asc

import (
	"context"
	"testing"
)

func TestListBetaAppReviewDetails(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppReviewDetails(ctx, &ListBetaAppReviewDetailsQuery{})
	})
}

func TestGetBetaAppReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewDetail(ctx, "10", &GetBetaAppReviewDetailQuery{})
	})
}

func TestGetAppForBetaAppReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaAppReviewDetail(ctx, "10", &GetAppForBetaAppReviewDetailQuery{})
	})
}

func TestGetBetaAppReviewDetailsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewDetailsForApp(ctx, "10", &GetBetaAppReviewDetailsForAppQuery{})
	})
}

func TestUpdateBetaAppReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaAppReviewDetail(ctx, "10", BetaAppReviewDetailUpdateRequest{})
	})
}
