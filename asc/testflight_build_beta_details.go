package asc

import (
	"fmt"
	"net/http"
)

// ExternalBetaState defines model for ExternalBetaState.
type ExternalBetaState string

// List of ExternalBetaState
const (
	ExternalBetaStateApproved                 ExternalBetaState = "BETA_APPROVED"
	ExternalBetaStateRejected                 ExternalBetaState = "BETA_REJECTED"
	ExternalBetaStateExpired                  ExternalBetaState = "EXPIRED"
	ExternalBetaStateInReview                 ExternalBetaState = "IN_BETA_REVIEW"
	ExternalBetaStateInTesting                ExternalBetaState = "IN_BETA_TESTING"
	ExternalBetaStateInExportComplianceReview ExternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	ExternalBetaStateMissingExportCompliance  ExternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	ExternalBetaStateProcessing               ExternalBetaState = "PROCESSING"
	ExternalBetaStateProcessingException      ExternalBetaState = "PROCESSING_EXCEPTION"
	ExternalBetaStateReadyForBetaSubmission   ExternalBetaState = "READY_FOR_BETA_SUBMISSION"
	ExternalBetaStateReadyForBetaTesting      ExternalBetaState = "READY_FOR_BETA_TESTING"
	ExternalBetaStateWaitingForBetaReview     ExternalBetaState = "WAITING_FOR_BETA_REVIEW"
)

// InternalBetaState defines model for InternalBetaState.
type InternalBetaState string

// List of InternalBetaState
const (
	InternalBetaStateExpired                  InternalBetaState = "EXPIRED"
	InternalBetaStateInTesting                InternalBetaState = "IN_BETA_TESTING"
	InternalBetaStateInExportComplianceReview InternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	InternalBetaStateMissingExportCompliance  InternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	InternalBetaStateProcessing               InternalBetaState = "PROCESSING"
	InternalBetaStateProcessingException      InternalBetaState = "PROCESSING_EXCEPTION"
	InternalBetaStateReadyForBetaTesting      InternalBetaState = "READY_FOR_BETA_TESTING"
)

// BuildBetaDetail defines model for BuildBetaDetail.
type BuildBetaDetail struct {
	Attributes *struct {
		AutoNotifyEnabled  *bool              `json:"autoNotifyEnabled,omitempty"`
		ExternalBuildState *ExternalBetaState `json:"externalBuildState,omitempty"`
		InternalBuildState *InternalBetaState `json:"internalBuildState,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildBetaDetailUpdateRequest defines model for BuildBetaDetailUpdateRequest.
type BuildBetaDetailUpdateRequest struct {
	Attributes *BuildBetaDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// BuildBetaDetailUpdateRequestAttributes are attributes for BuildBetaDetailUpdateRequest
type BuildBetaDetailUpdateRequestAttributes struct {
	AutoNotifyEnabled *bool `json:"autoNotifyEnabled,omitempty"`
}

// BuildBetaDetailResponse defines model for BuildBetaDetailResponse.
type BuildBetaDetailResponse struct {
	Data     BuildBetaDetail `json:"data"`
	Included *[]Build        `json:"included,omitempty"`
	Links    DocumentLinks   `json:"links"`
}

// BuildBetaDetailsResponse defines model for BuildBetaDetailsResponse.
type BuildBetaDetailsResponse struct {
	Data     []BuildBetaDetail  `json:"data"`
	Included *[]Build           `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListBuildBetaDetailsQuery defines model for ListBuildBetaDetails
type ListBuildBetaDetailsQuery struct {
	FieldsBuilds           *[]string `url:"fields[builds],omitempty"`
	FieldsBuildBetaDetails *[]string `url:"fields[buildBetaDetails],omitempty"`
	FilterID               *[]string `url:"filter[id],omitempty"`
	FilterBuild            *[]string `url:"filter[build],omitempty"`
	Include                *[]string `url:"include,omitempty"`
	Limit                  *int      `url:"limit,omitempty"`
	Cursor                 *string   `url:"cursor,omitempty"`
}

// GetBuildBetaDetailsQuery defines model for GetBuildBetaDetails
type GetBuildBetaDetailsQuery struct {
	FieldsBuilds           *[]string `url:"fields[builds],omitempty"`
	FieldsBuildBetaDetails *[]string `url:"fields[buildBetaDetails],omitempty"`
	Include                *[]string `url:"include,omitempty"`
}

// GetBuildForBuildBetaDetailQuery defines model for GetBuildForBuildBetaDetail
type GetBuildForBuildBetaDetailQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

// GetBuildBetaDetailForBuildQuery defines model for GetBuildBetaDetailForBuild
type GetBuildBetaDetailForBuildQuery struct {
	FieldsBuildBetaDetails *[]string `url:"fields[buildBetaDetails],omitempty"`
}

// ListBuildBetaDetails finds and lists build beta details for all builds.
func (s *TestflightService) ListBuildBetaDetails(params *ListBuildBetaDetailsQuery) (*BuildBetaDetailsResponse, *http.Response, error) {
	res := new(BuildBetaDetailsResponse)
	resp, err := s.client.get("buildBetaDetails", params, res)
	return res, resp, err
}

// GetBuildBetaDetail gets a specific build beta details resource.
func (s *TestflightService) GetBuildBetaDetail(id string, params *GetBuildBetaDetailsQuery) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBuildForBuildBetaDetail gets the build information for a specific build beta details resource.
func (s *TestflightService) GetBuildForBuildBetaDetail(id string, params *GetBuildForBuildBetaDetailQuery) (*BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBuildBetaDetailForBuild gets the beta test details for a specific build.
func (s *TestflightService) GetBuildBetaDetailForBuild(id string, params *GetBuildBetaDetailForBuildQuery) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/buildBetaDetail", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// UpdateBuildBetaDetail updates beta test details for a specific build.
func (s *TestflightService) UpdateBuildBetaDetail(id string, body *BuildBetaDetailUpdateRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}
