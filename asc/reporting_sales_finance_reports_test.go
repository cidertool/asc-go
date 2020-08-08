package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFinanceReports(t *testing.T) {
	client, server := newServer("ahhhhhh")
	defer server.Close()

	gzip, resp, err := client.Reporting.DownloadFinanceReports(context.Background(), &DownloadFinanceReportsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, gzip)
}

func TestDownloadSalesAndTrendsReports(t *testing.T) {
	client, server := newServer("ahhhhhh")
	defer server.Close()

	gzip, resp, err := client.Reporting.DownloadSalesAndTrendsReports(context.Background(), &DownloadSalesAndTrendsReportsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, gzip)
}
