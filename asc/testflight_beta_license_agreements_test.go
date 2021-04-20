/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"testing"
)

func TestListBetaLicenseAgreements(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaLicenseAgreements(ctx, &ListBetaLicenseAgreementsQuery{})
	})
}

func TestGetBetaLicenseAgreement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaLicenseAgreement(ctx, "10", &GetBetaLicenseAgreementQuery{})
	})
}

func TestGetAppForBetaLicenseAgreement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaLicenseAgreement(ctx, "10", &GetAppForBetaLicenseAgreementQuery{})
	})
}

func TestGetBetaLicenseAgreementForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaLicenseAgreementForApp(ctx, "10", &GetBetaLicenseAgreementForAppQuery{})
	})
}

func TestUpdateBetaLicenseAgreement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaLicenseAgreementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaLicenseAgreement(ctx, "10", String(""))
	})
}
