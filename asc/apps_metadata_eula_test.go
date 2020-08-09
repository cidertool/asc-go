package asc

import (
	"context"
	"testing"
)

func TestCreateEULA(t *testing.T) {
	testEndpointWithResponse(t, "{}", &EndUserLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateEULA(ctx, &EndUserLicenseAgreementCreateRequest{})
	})
}

func TestUpdateEULA(t *testing.T) {
	testEndpointWithResponse(t, "{}", &EndUserLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateEULA(ctx, "10", &EndUserLicenseAgreementUpdateRequest{})
	})
}

func TestDeleteEULA(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteEULA(ctx, "10")
	})
}

func TestGetEULA(t *testing.T) {
	testEndpointWithResponse(t, "{}", &EndUserLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetEULA(ctx, "10", &GetEULAQuery{})
	})
}

func TestGetEULAForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &EndUserLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetEULAForApp(ctx, "10", &GetEULAForAppQuery{})
	})
}
