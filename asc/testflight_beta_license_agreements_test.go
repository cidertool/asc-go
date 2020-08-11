package asc

import (
	"context"
	"testing"
)

func TestListBetaLicenseAgreements(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaLicenseAgreements(ctx, &ListBetaLicenseAgreementsQuery{})
	})
}

func TestGetBetaLicenseAgreement(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaLicenseAgreement(ctx, "10", &GetBetaLicenseAgreementQuery{})
	})
}

func TestGetAppForBetaLicenseAgreement(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaLicenseAgreement(ctx, "10", &GetAppForBetaLicenseAgreementQuery{})
	})
}

func TestGetBetaLicenseAgreementForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaLicenseAgreementForApp(ctx, "10", &GetBetaLicenseAgreementForAppQuery{})
	})
}

func TestUpdateBetaLicenseAgreement(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaLicenseAgreement(ctx, "10", &BetaLicenseAgreementUpdateRequest{})
	})
}

func TestUpdateBetaLicenseAgreementApplyRequestTypes(t *testing.T) {

}
