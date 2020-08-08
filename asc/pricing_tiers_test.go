package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppPriceTiers(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPriceTiersResponse{}
	got, resp, err := client.Pricing.ListAppPriceTiers(context.Background(), &ListAppPriceTiersQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppPriceTier(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPriceTierResponse{}
	got, resp, err := client.Pricing.GetAppPriceTier(context.Background(), "10", &GetAppPriceTierQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListPricePointsForAppPriceTier(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPricePointsResponse{}
	got, resp, err := client.Pricing.ListPricePointsForAppPriceTier(context.Background(), "10", &ListPricePointsForAppPriceTierQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListAppPricePoints(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPricePointsResponse{}
	got, resp, err := client.Pricing.ListAppPricePoints(context.Background(), &ListAppPricePointsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetTerritoryForAppPricePoint(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &TerritoryResponse{}
	got, resp, err := client.Pricing.GetTerritoryForAppPricePoint(context.Background(), "10", &GetTerritoryForAppPricePointQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppPricePoint(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPricePointResponse{}
	got, resp, err := client.Pricing.GetAppPricePoint(context.Background(), "10", &GetAppPricePointQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
