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

func TestGetPerfPowerMetricsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PerfPowerMetricsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetPerfPowerMetricsForApp(ctx, "10", &GetPerfPowerMetricsQuery{})
	})
}

func TestGetPerfPowerMetricsForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &PerfPowerMetricsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetPerfPowerMetricsForBuild(ctx, "10", &GetPerfPowerMetricsQuery{})
	})
}

func TestListDiagnosticSignaturesForBuild(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &DiagnosticSignaturesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.ListDiagnosticSignaturesForBuild(ctx, "10", &ListDiagnosticsSignaturesQuery{})
	})
}

func TestGetLogsForDiagnosticSignature(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &DiagnosticLogsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetLogsForDiagnosticSignature(ctx, "10", &GetLogsForDiagnosticSignatureQuery{})
	})
}
