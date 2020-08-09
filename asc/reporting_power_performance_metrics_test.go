package asc

import (
	"context"
	"testing"
)

func TestGetPerfPowerMetricsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PerfPowerMetricsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetPerfPowerMetricsForApp(ctx, "10", &GetPerfPowerMetricsQuery{})
	})
}

func TestGetPerfPowerMetricsForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &PerfPowerMetricsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetPerfPowerMetricsForBuild(ctx, "10", &GetPerfPowerMetricsQuery{})
	})
}

func TestListDiagnosticSignaturesForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DiagnosticSignaturesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.ListDiagnosticSignaturesForBuild(ctx, "10", &ListDiagnosticsSignaturesQuery{})
	})
}

func TestGetLogsForDiagnosticSignature(t *testing.T) {
	testEndpointWithResponse(t, "{}", &DiagnosticLogsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetLogsForDiagnosticSignature(ctx, "10", &GetLogsForDiagnosticSignatureQuery{})
	})
}
