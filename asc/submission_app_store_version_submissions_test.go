package asc

import (
	"context"
	"testing"
)

func TestCreateSubmission(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CreateSubmission(ctx, &AppStoreVersionSubmissionCreateRequest{})
	})
}

func TestDeleteSubmission(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Submission.DeleteSubmission(ctx, "10")
	})
}

func TestGetAppStoreVersionSubmissionForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionSubmissionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetAppStoreVersionSubmissionForAppStoreVersion(ctx, "10", &GetAppStoreVersionSubmissionForAppStoreVersionQuery{})
	})
}
