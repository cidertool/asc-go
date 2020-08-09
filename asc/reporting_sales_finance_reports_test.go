package asc

import (
	"context"
	"testing"
)

func TestDownloadFinanceReports(t *testing.T) {
	testEndpointWithResponse(t, "ahhhhhhh", nil, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.DownloadFinanceReports(ctx, &DownloadFinanceReportsQuery{})
	})
}

func TestDownloadSalesAndTrendsReports(t *testing.T) {
	testEndpointWithResponse(t, "ahhhhhhh", nil, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.DownloadSalesAndTrendsReports(ctx, &DownloadSalesAndTrendsReportsQuery{})
	})
}
