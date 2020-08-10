package asc

import (
	"context"
	"testing"
)

func TestCreateAvailableBuildNotification(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaNotificationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateAvailableBuildNotification(ctx, &BuildBetaNotificationCreateRequest{})
	})
}

func TestCreateAvailableBuildNotificationNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaNotificationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateAvailableBuildNotification(ctx, nil)
	})
}
