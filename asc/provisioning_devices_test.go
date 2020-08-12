package asc

import (
	"context"
	"testing"
)

func TestCreateDevice(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DeviceResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateDevice(ctx, DeviceCreateRequest{})
	})
}

func TestListDevices(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DevicesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListDevices(ctx, &ListDevicesQuery{})
	})
}

func TestGetDevice(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DeviceResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetDevice(ctx, "10", &GetDeviceQuery{})
	})
}

func TestUpdateDevice(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DeviceResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.UpdateDevice(ctx, "10", DeviceUpdateRequest{})
	})
}
