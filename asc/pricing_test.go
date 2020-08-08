package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPricesForApp(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPricesResponse{}
	got, resp, err := client.Pricing.ListPricesForApp(context.Background(), "10", &ListPricesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetPrice(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPriceResponse{}
	got, resp, err := client.Pricing.GetPrice(context.Background(), "10", &GetPriceQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
