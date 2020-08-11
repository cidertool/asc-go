package asc

import (
	"context"
	"testing"
)

func TestListAppEncryptionDeclarations(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListAppEncryptionDeclarations(ctx, &ListAppEncryptionDeclarationsQuery{})
	})
}

func TestGetAppEncryptionDeclaration(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclaration(ctx, "10", &GetAppEncryptionDeclarationQuery{})
	})
}

func TestGetAppForAppEncryptionDeclaration(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppForAppEncryptionDeclaration(ctx, "10", &GetAppForEncryptionDeclarationQuery{})
	})
}

func TestAssignBuildsToAppEncryptionDeclaration(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.AssignBuildsToAppEncryptionDeclaration(ctx, "10", AppEncryptionDeclarationBuildsLinkagesRequest{})
	})
}

func TestAssignBuildsToAppEncryptionDeclarationApplyRequestTypes(t *testing.T) {

}
