package asc

import (
	"context"
	"testing"
)

func TestListPricesForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPricesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.ListPricesForApp(ctx, "10", &ListPricesQuery{})
	})
}

func TestGetPrice(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPriceResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Pricing.GetPrice(ctx, "10", &GetPriceQuery{})
	})
}
