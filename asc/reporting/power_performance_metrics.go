package reporting

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// DiagnosticLog defines model for DiagnosticLog.
type DiagnosticLog struct {
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// DiagnosticLogsResponse defines model for DiagnosticLogsResponse.
type DiagnosticLogsResponse struct {
	Data  []DiagnosticLog           `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// DiagnosticSignature defines model for DiagnosticSignature.
type DiagnosticSignature struct {
	Attributes *struct {
		DiagnosticType *string  `json:"diagnosticType,omitempty"`
		Signature      *string  `json:"signature,omitempty"`
		Weight         *float32 `json:"weight,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// DiagnosticSignaturesResponse defines model for DiagnosticSignaturesResponse.
type DiagnosticSignaturesResponse struct {
	Data     []DiagnosticSignature     `json:"data"`
	Included *[]DiagnosticLog          `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// PerfPowerMetric defines model for PerfPowerMetric.
type PerfPowerMetric struct {
	Attributes *struct {
		DeviceType *string `json:"deviceType,omitempty"`
		MetricType *string `json:"metricType,omitempty"`
		Platform   *string `json:"platform,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// PerfPowerMetricsResponse defines model for PerfPowerMetricsResponse.
type PerfPowerMetricsResponse struct {
	Data  []PerfPowerMetric         `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type GetPerfPowerMetricsQuery struct {
	Filter *struct {
		DeviceType *[]string `url:"deviceType,omitempty"`
		MetricType *[]string `url:"metricType,omitempty"`
		Platform   *[]string `url:"platform,omitempty"`
	} `url:"filter,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListDiagnosticsSignaturesQuery struct {
	Fields *struct {
		DiagnosticSignatures *[]string `url:"diagnosticSignatures,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		DiagnosticType *[]string `url:"diagnosticType,omitempty"`
	} `url:"filter,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type GetLogsForDiagnosticSignatureQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// GetPerfPowerMetricsForApp gets the performance and power metrics data for the most recent versions of an app.
func (s *Service) GetPerfPowerMetricsForApp(id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPerfPowerMetricsForBuild gets the performance and power metrics data for a specific build.
func (s *Service) GetPerfPowerMetricsForBuild(id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListDiagnosticSignaturesForBuild lists the aggregate backtrace signatures captured for a specific build.
func (s *Service) ListDiagnosticSignaturesForBuild(id string, params *ListDiagnosticsSignaturesQuery) (*DiagnosticSignaturesResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/diagnosticSignatures", id)
	res := new(DiagnosticSignaturesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetLogsForDiagnosticSignature gets the anonymized backtrace logs associated with a specific diagnostic signature.
func (s *Service) GetLogsForDiagnosticSignature(id string, params *GetLogsForDiagnosticSignatureQuery) (*DiagnosticLogsResponse, *common.Response, error) {
	url := fmt.Sprintf("diagnosticSignatures/%s/logs", id)
	res := new(DiagnosticLogsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
