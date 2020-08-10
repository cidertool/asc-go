package asc

import (
	"context"
	"testing"
)

func TestGetPreOrder(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.GetPreOrder(ctx, "10", &GetPreOrderQuery{})
	})
}

func TestGetPreOrderForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.GetPreOrderForApp(ctx, "10", &GetPreOrderForAppQuery{})
	})
}

func TestCreatePreOrder(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.CreatePreOrder(ctx, &AppPreOrderCreateRequest{})
	})
}

func TestCreatePreOrderNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.CreatePreOrder(ctx, nil)
	})
}

func TestUpdatePreOrder(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.UpdatePreOrder(ctx, "10", &AppPreOrderUpdateRequest{})
	})
}

func TestUpdatePreOrderNoRequest(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreOrderResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Publishing.UpdatePreOrder(ctx, "10", nil)
	})
}

func TestDeletePreOrder(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Publishing.DeletePreOrder(ctx, "10")
	})
}
