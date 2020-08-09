package asc

import (
	"context"
	"testing"
)

func TestListTerritories(t *testing.T) {
	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritories(ctx, &ListTerritoriesQuery{})
	})
}

func TestListTerritoriesForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritoriesForApp(ctx, "10", &ListTerritoriesQuery{})
	})
}

func TestListTerritoriesForEULA(t *testing.T) {
	testEndpointWithResponse(t, "{}", &TerritoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListTerritoriesForEULA(ctx, "10", &ListTerritoriesQuery{})
	})
}

func TestGetTerritoryForAppPrice(t *testing.T) {
	testEndpointWithResponse(t, "{}", &TerritoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetTerritoryForAppPrice(ctx, "10", &ListTerritoriesQuery{})
	})
}
