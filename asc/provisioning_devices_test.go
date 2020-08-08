package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DeviceResponse{}
	got, resp, err := client.Provisioning.CreateDevice(context.Background(), &DeviceCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListDevices(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DevicesResponse{}
	got, resp, err := client.Provisioning.ListDevices(context.Background(), &ListDevicesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetDevice(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DeviceResponse{}
	got, resp, err := client.Provisioning.GetDevice(context.Background(), "10", &GetDeviceQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdateDevice(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DeviceResponse{}
	got, resp, err := client.Provisioning.UpdateDevice(context.Background(), "10", &DeviceUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
