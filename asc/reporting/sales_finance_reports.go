package reporting

import (
	"bytes"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

type DownloadFinanceReportsQuery struct {
	Filter struct {
		RegionCode   []string `url:"regionCode"`
		ReportDate   []string `url:"reportDate"`
		ReportType   []string `url:"reportType"`
		VendorNumber []string `url:"vendorNumber"`
	} `url:"filter"`
}

type DownloadSalesAndTrendsReportsQuery struct {
	Filter struct {
		Frequency     []string  `url:"frequency"`
		ReportDate    *[]string `url:"reportDate,omitempty"`
		ReportSubType []string  `url:"reportSubType"`
		ReportType    []string  `url:"reportType"`
		VendorNumber  []string  `url:"vendorNumber"`
		Version       *[]string `url:"version,omitempty"`
	} `url:"filter"`
}

// DownloadFinanceReports downloads finance reports filtered by your specified criteria.
func (s *Service) DownloadFinanceReports(params *DownloadFinanceReportsQuery) (*bytes.Buffer, *common.Response, error) {
	res := new(bytes.Buffer)
	resp, err := s.GetWithQuery("financeReports", params, res)
	return res, resp, err
}

// DownloadSalesAndTrendsReports downloads sales and trends reports filtered by your specified criteria.
func (s *Service) DownloadSalesAndTrendsReports(params *DownloadSalesAndTrendsReportsQuery) (*bytes.Buffer, *common.Response, error) {
	res := new(bytes.Buffer)
	resp, err := s.GetWithQuery("salesReports", params, res)
	return res, resp, err
}
