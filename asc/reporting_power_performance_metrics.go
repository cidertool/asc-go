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
	"fmt"
)

// DiagnosticLog defines model for DiagnosticLog.
//
// https://developer.apple.com/documentation/appstoreconnectapi/diagnosticlog
type DiagnosticLog struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// DiagnosticLogsResponse defines model for DiagnosticLogsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/diagnosticlogsresponse
type DiagnosticLogsResponse struct {
	Data  []DiagnosticLog    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// DiagnosticSignature defines model for DiagnosticSignature.
//
// https://developer.apple.com/documentation/appstoreconnectapi/diagnosticsignature
type DiagnosticSignature struct {
	Attributes *DiagnosticSignatureAttributes `json:"attributes,omitempty"`
	ID         string                         `json:"id"`
	Links      ResourceLinks                  `json:"links"`
	Type       string                         `json:"type"`
}

// DiagnosticSignatureAttributes defines model for DiagnosticSignature.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/diagnosticsignature/attributes
type DiagnosticSignatureAttributes struct {
	DiagnosticType *string  `json:"diagnosticType,omitempty"`
	Signature      *string  `json:"signature,omitempty"`
	Weight         *float32 `json:"weight,omitempty"`
}

// DiagnosticSignaturesResponse defines model for DiagnosticSignaturesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/diagnosticsignaturesresponse
type DiagnosticSignaturesResponse struct {
	Data     []DiagnosticSignature `json:"data"`
	Included []DiagnosticLog       `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// PerfPowerMetric defines model for PerfPowerMetric.
//
// https://developer.apple.com/documentation/appstoreconnectapi/perfpowermetric
type PerfPowerMetric struct {
	Attributes *PerfPowerMetricAttributes `json:"attributes,omitempty"`
	ID         string                     `json:"id"`
	Links      ResourceLinks              `json:"links"`
	Type       string                     `json:"type"`
}

// PerfPowerMetricAttributes defines model for PerfPowerMetric.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/perfpowermetric/attributes
type PerfPowerMetricAttributes struct {
	DeviceType *string `json:"deviceType,omitempty"`
	MetricType *string `json:"metricType,omitempty"`
	Platform   *string `json:"platform,omitempty"`
}

// PerfPowerMetricsResponse defines model for PerfPowerMetricsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/perfpowermetricsresponse
type PerfPowerMetricsResponse struct {
	Data  []PerfPowerMetric  `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GetPerfPowerMetricsQuery are query options for GetPerfPowerMetrics
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_power_and_performance_metrics_for_a_build
type GetPerfPowerMetricsQuery struct {
	FilterDeviceType []string `url:"filter[deviceType],omitempty"`
	FilterMetricType []string `url:"filter[metricType],omitempty"`
	FilterPlatform   []string `url:"filter[platform],omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListDiagnosticsSignaturesQuery are query options for ListDiagnosticsSignatures
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_diagnostic_signatures_for_a_build
type ListDiagnosticsSignaturesQuery struct {
	FieldsDiagnosticSignatures []string `url:"fields[diagnosticSignatures],omitempty"`
	FilterDiagnosticType       []string `url:"filter[diagnosticType],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetLogsForDiagnosticSignatureQuery are query options for GetLogsForDiagnosticSignature
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_logs_for_a_diagnostic_signature
type GetLogsForDiagnosticSignatureQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetPerfPowerMetricsForApp gets the performance and power metrics data for the most recent versions of an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_power_and_performance_metrics_for_an_app
func (s *ReportingService) GetPerfPowerMetricsForApp(ctx context.Context, id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetPerfPowerMetricsForBuild gets the performance and power metrics data for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_power_and_performance_metrics_for_a_build
func (s *ReportingService) GetPerfPowerMetricsForBuild(ctx context.Context, id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListDiagnosticSignaturesForBuild lists the aggregate backtrace signatures captured for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_diagnostic_signatures_for_a_build
func (s *ReportingService) ListDiagnosticSignaturesForBuild(ctx context.Context, id string, params *ListDiagnosticsSignaturesQuery) (*DiagnosticSignaturesResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/diagnosticSignatures", id)
	res := new(DiagnosticSignaturesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetLogsForDiagnosticSignature gets the anonymized backtrace logs associated with a specific diagnostic signature.
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_logs_for_a_diagnostic_signature
func (s *ReportingService) GetLogsForDiagnosticSignature(ctx context.Context, id string, params *GetLogsForDiagnosticSignatureQuery) (*DiagnosticLogsResponse, *Response, error) {
	url := fmt.Sprintf("diagnosticSignatures/%s/logs", id)
	res := new(DiagnosticLogsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
