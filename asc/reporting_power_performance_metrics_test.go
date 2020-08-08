package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPerfPowerMetricsForApp(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &PerfPowerMetricsResponse{}
	got, resp, err := client.Reporting.GetPerfPowerMetricsForApp(context.Background(), "10", &GetPerfPowerMetricsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetPerfPowerMetricsForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &PerfPowerMetricsResponse{}
	got, resp, err := client.Reporting.GetPerfPowerMetricsForBuild(context.Background(), "10", &GetPerfPowerMetricsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListDiagnosticSignaturesForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DiagnosticSignaturesResponse{}
	got, resp, err := client.Reporting.ListDiagnosticSignaturesForBuild(context.Background(), "10", &ListDiagnosticsSignaturesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetLogsForDiagnosticSignature(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DiagnosticLogsResponse{}
	got, resp, err := client.Reporting.GetLogsForDiagnosticSignature(context.Background(), "10", &GetLogsForDiagnosticSignatureQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
