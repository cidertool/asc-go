package asc

import (
	"context"
	"testing"
)

func TestListBetaBuildLocalizations(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaBuildLocalizations(ctx, &ListBetaBuildLocalizationsQuery{})
	})
}

func TestGetBetaBuildLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaBuildLocalization(ctx, "10", &GetBetaBuildLocalizationQuery{})
	})
}

func TestGetBuildForBetaBuildLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBuildForBetaBuildLocalization(ctx, "10", &GetBuildForBetaBuildLocalizationQuery{})
	})
}

func TestListBetaBuildLocalizationsForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaBuildLocalizationsForBuild(ctx, "10", &ListBetaBuildLocalizationsForBuildQuery{})
	})
}

func TestCreateBetaBuildLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaBuildLocalization(ctx, "", nil, "")
	})
}

func TestUpdateBetaBuildLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaBuildLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaBuildLocalization(ctx, "10", BetaBuildLocalizationUpdateRequest{})
	})
}

func TestDeleteBetaBuildLocalization(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaBuildLocalization(ctx, "10")
	})
}
