package asc

import (
	"context"
	"testing"
)

func TestCreateReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CreateReviewDetail(ctx, &AppStoreReviewDetailCreateRequest{})
	})
}

func TestCreateReviewDetailApplyRequestTypes(t *testing.T) {

}

func TestGetReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetReviewDetail(ctx, "10", &GetReviewDetailQuery{})
	})
}

func TestGetReviewDetailsForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetReviewDetailsForAppStoreVersion(ctx, "10", &GetAppStoreReviewDetailsForAppStoreVersionQuery{})
	})
}

func TestUpdateReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.UpdateReviewDetail(ctx, "10", &AppStoreReviewDetailUpdateRequest{})
	})
}

func TestUpdateReviewDetailApplyRequestTypes(t *testing.T) {

}
