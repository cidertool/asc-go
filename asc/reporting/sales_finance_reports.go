package reporting

import (
	"bytes"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

type DownloadFinanceReportsQuery struct {
	FilterRegionCode   []string `url:"filter[regionCode]"`
	FilterReportDate   []string `url:"filter[reportDate]"`
	FilterReportType   []string `url:"filter[reportType]"`
	FilterVendorNumber []string `url:"filter[vendorNumber]"`
}

type DownloadSalesAndTrendsReportsQuery struct {
	FilterFrequency     []string  `url:"filter[frequency]"`
	FilterReportDate    *[]string `url:"filter[reportDate],omitempty"`
	FilterReportSubType []string  `url:"filter[reportSubType]"`
	FilterReportType    []string  `url:"filter[reportType]"`
	FilterVendorNumber  []string  `url:"filter[vendorNumber]"`
	FilterVersion       *[]string `url:"filter[version],omitempty"`
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
