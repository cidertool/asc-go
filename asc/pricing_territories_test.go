package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTerritories(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &TerritoriesResponse{}
	got, resp, err := client.Pricing.ListTerritories(context.Background(), &ListTerritoriesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListTerritoriesForApp(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &TerritoriesResponse{}
	got, resp, err := client.Pricing.ListTerritoriesForApp(context.Background(), "10", &ListTerritoriesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListTerritoriesForEULA(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &TerritoriesResponse{}
	got, resp, err := client.Pricing.ListTerritoriesForEULA(context.Background(), "10", &ListTerritoriesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetTerritoryForAppPrice(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &TerritoryResponse{}
	got, resp, err := client.Pricing.GetTerritoryForAppPrice(context.Background(), "10", &ListTerritoriesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
