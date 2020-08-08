package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPreOrder(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPreOrderResponse{}
	got, resp, err := client.Publishing.GetPreOrder(context.Background(), "10", &GetPreOrderQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetPreOrderForApp(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPreOrderResponse{}
	got, resp, err := client.Publishing.GetPreOrderForApp(context.Background(), "10", &GetPreOrderForAppQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestCreatePreOrder(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPreOrderResponse{}
	got, resp, err := client.Publishing.CreatePreOrder(context.Background(), &AppPreOrderCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdatePreOrder(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppPreOrderResponse{}
	got, resp, err := client.Publishing.UpdatePreOrder(context.Background(), "10", &AppPreOrderUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestDeletePreOrder(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Publishing.DeletePreOrder(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
