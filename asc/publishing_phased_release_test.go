package asc

import (
	"context"
	"testing"
)

func TestCreatePhasedRelease(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionPhasedReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		state := PhasedReleaseStateActive
		return client.Publishing.CreatePhasedRelease(ctx, &state, "")
	})
}

func TestUpdatePhasedRelease(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionPhasedReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		state := PhasedReleaseStateActive
		return client.Publishing.UpdatePhasedRelease(ctx, "10", &state)
	})
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
