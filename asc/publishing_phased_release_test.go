package asc

import (
	"context"
	"testing"
)

func TestCreatePhasedRelease(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionPhasedReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.CreatePhasedRelease(ctx, &AppStoreVersionPhasedReleaseCreateRequest{})
	})
}

func TestCreatePhasedReleaseApplyRequestTypes(t *testing.T) {

}

func TestUpdatePhasedRelease(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionPhasedReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.UpdatePhasedRelease(ctx, "10", &AppStoreVersionPhasedReleaseUpdateRequest{})
	})
}

func TestUpdatePhasedReleaseApplyRequestTypes(t *testing.T) {

}

func TestDeletePhasedRelease(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Publishing.DeletePhasedRelease(ctx, "10")
	})
}

func TestGetAppStoreVersionPhasedReleaseForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionPhasedReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.GetAppStoreVersionPhasedReleaseForAppStoreVersion(ctx, "10", &GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery{})
	})
}
