package asc

import (
	"context"
	"testing"
)

func TestCreateBetaAppReviewSubmission(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaAppReviewSubmission(ctx, &BetaAppReviewSubmissionCreateRequest{})
	})
}

func TestCreateBetaAppReviewSubmissionNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaAppReviewSubmission(ctx, nil)
	})
}

func TestListBetaAppReviewSubmissions(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewSubmissionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppReviewSubmissions(ctx, &ListBetaAppReviewSubmissionsQuery{})
	})
}

func TestGetBetaAppReviewSubmission(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewSubmission(ctx, "10", &GetBetaAppReviewSubmissionQuery{})
	})
}

func TestGetBuildForBetaAppReviewSubmission(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildForBetaAppReviewSubmission(ctx, "10", &GetBuildForBetaAppReviewSubmissionQuery{})
	})
}

func TestGetBetaAppReviewSubmissionForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppReviewSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewSubmissionForBuild(ctx, "10", &GetBetaAppReviewSubmissionForBuildQuery{})
	})
}
