package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableCapability(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDCapabilityResponse{}
	got, resp, err := client.Provisioning.EnableCapability(context.Background(), &BundleIDCapabilityCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestDisableCapability(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Provisioning.DisableCapability(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUpdateCapability(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDCapabilityResponse{}
	got, resp, err := client.Provisioning.UpdateCapability(context.Background(), "10", &BundleIDCapabilityUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
