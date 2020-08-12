package asc

import (
	"context"
	"testing"
)

func TestListBetaAppLocalizations(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppLocalizations(ctx, &ListBetaAppLocalizationsQuery{})
	})
}

func TestGetBetaAppLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppLocalization(ctx, "10", &GetBetaAppLocalizationQuery{})
	})
}

func TestGetAppForBetaAppLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaAppLocalization(ctx, "10", &GetAppForBetaAppLocalizationQuery{})
	})
}

func TestListBetaAppLocalizationsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppLocalizationsForApp(ctx, "10", &ListBetaAppLocalizationsForAppQuery{})
	})
}

func TestCreateBetaAppLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaAppLocalization(ctx, BetaAppLocalizationCreateRequest{})
	})
}

func TestUpdateBetaAppLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaAppLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaAppLocalization(ctx, "10", BetaAppLocalizationUpdateRequest{})
	})
}

func TestDeleteBetaAppLocalization(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaAppLocalization(ctx, "10")
	})
}
