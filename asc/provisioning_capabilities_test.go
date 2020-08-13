package asc

import (
	"context"
	"testing"
)

func TestEnableCapability(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDCapabilityResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.EnableCapability(ctx, CapabilityTypeAccessWifiInformation, []CapabilitySetting{}, "")
	})
}

func TestDisableCapability(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.DisableCapability(ctx, "10")
	})
}

func TestUpdateCapability(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BundleIDCapabilityResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		capability := CapabilityTypeAppGroups
		return client.Provisioning.UpdateCapability(ctx, "10", &capability, []CapabilitySetting{})
	})
}
