package asc

import (
	"context"
	"testing"
)

func TestListAppPriceTiers(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPriceTiersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListAppPriceTiers(ctx, &ListAppPriceTiersQuery{})
	})
}

func TestGetAppPriceTier(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPriceTierResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetAppPriceTier(ctx, "10", &GetAppPriceTierQuery{})
	})
}

func TestListPricePointsForAppPriceTier(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPricePointsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListPricePointsForAppPriceTier(ctx, "10", &ListPricePointsForAppPriceTierQuery{})
	})
}

func TestListAppPricePoints(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPricePointsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListAppPricePoints(ctx, &ListAppPricePointsQuery{})
	})
}

func TestGetTerritoryForAppPricePoint(t *testing.T) {
	testEndpointWithResponse(t, "{}", &TerritoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetTerritoryForAppPricePoint(ctx, "10", &GetTerritoryForAppPricePointQuery{})
	})
}

func TestGetAppPricePoint(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPricePointResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetAppPricePoint(ctx, "10", &GetAppPricePointQuery{})
	})
}
