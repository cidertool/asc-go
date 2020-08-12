package asc

import (
	"context"
	"testing"
)

func TestCreateIDFADeclaration(t *testing.T) {
	testEndpointWithResponse(t, "{}", &IDFADeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CreateIDFADeclaration(ctx, &IDFADeclarationCreateRequest{})
	})
}

func TestUpdateIDFADeclaration(t *testing.T) {
	testEndpointWithResponse(t, "{}", &IDFADeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.UpdateIDFADeclaration(ctx, "10", &IDFADeclarationUpdateRequest{})
	})
}

func TestDeleteIDFADeclaration(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Submission.DeleteIDFADeclaration(ctx, "10")
	})
}

func TestGetIDFADeclarationForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &IDFADeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetIDFADeclarationForAppStoreVersion(ctx, "10", &GetIDFADeclarationForAppStoreVersionQuery{})
	})
}
