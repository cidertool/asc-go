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
	"bytes"
	"context"
	"io"
)

// DownloadFinanceReportsQuery are query options for DownloadFinanceReports
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_finance_reports
type DownloadFinanceReportsQuery struct {
	FilterRegionCode   []string `url:"filter[regionCode]"`
	FilterReportDate   []string `url:"filter[reportDate]"`
	FilterReportType   []string `url:"filter[reportType]"`
	FilterVendorNumber []string `url:"filter[vendorNumber]"`
}

// DownloadSalesAndTrendsReportsQuery are query options for DownloadSalesAndTrendsReports
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_sales_and_trends_reports
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
func (s *ReportingService) DownloadFinanceReports(ctx context.Context, params *DownloadFinanceReportsQuery) (io.Reader, *Response, error) {
	buffer := new(bytes.Buffer)
	resp, err := s.client.get(ctx, "financeReports", params, buffer, withAccept("application/a-gzip"))

	return buffer, resp, err
}

// DownloadSalesAndTrendsReports downloads sales and trends reports filtered by your specified criteria.
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_sales_and_trends_reports
func (s *ReportingService) DownloadSalesAndTrendsReports(ctx context.Context, params *DownloadSalesAndTrendsReportsQuery) (io.Reader, *Response, error) {
	buffer := new(bytes.Buffer)
	resp, err := s.client.get(ctx, "salesReports", params, buffer, withAccept("application/a-gzip"))

	return buffer, resp, err
}
