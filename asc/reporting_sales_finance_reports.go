package asc

import (
	"bytes"
	"io"
)

// DownloadFinanceReportsQuery are query options for DownloadFinanceReports
type DownloadFinanceReportsQuery struct {
	FilterRegionCode   []string `url:"filter[regionCode]"`
	FilterReportDate   []string `url:"filter[reportDate]"`
	FilterReportType   []string `url:"filter[reportType]"`
	FilterVendorNumber []string `url:"filter[vendorNumber]"`
}

// DownloadSalesAndTrendsReportsQuery are query options for DownloadSalesAndTrendsReports
type DownloadSalesAndTrendsReportsQuery struct {
	FilterFrequency     []string `url:"filter[frequency]"`
	FilterReportDate    []string `url:"filter[reportDate],omitempty"`
	FilterReportSubType []string `url:"filter[reportSubType]"`
	FilterReportType    []string `url:"filter[reportType]"`
	FilterVendorNumber  []string `url:"filter[vendorNumber]"`
	FilterVersion       []string `url:"filter[version],omitempty"`
}

// DownloadFinanceReports downloads finance reports filtered by your specified criteria.
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_finance_reports
func (s *ReportingService) DownloadFinanceReports(params *DownloadFinanceReportsQuery) (io.Reader, *Response, error) {
	buffer := new(bytes.Buffer)
	resp, err := s.client.get("financeReports", params, buffer, withAccept("application/a-gzip"))
	return buffer, resp, err
}

// DownloadSalesAndTrendsReports downloads sales and trends reports filtered by your specified criteria.
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_sales_and_trends_reports
func (s *ReportingService) DownloadSalesAndTrendsReports(params *DownloadSalesAndTrendsReportsQuery) (io.Reader, *Response, error) {
	buffer := new(bytes.Buffer)
	resp, err := s.client.get("salesReports", params, buffer, withAccept("application/a-gzip"))
	return buffer, resp, err
}
