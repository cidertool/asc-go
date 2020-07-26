package testflight

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal/services"
)

// ExternalBetaState defines model for ExternalBetaState.
type ExternalBetaState string

// List of ExternalBetaState
const (
	ExternalBetaApproved                 ExternalBetaState = "BETA_APPROVED"
	ExternalBetaRejected                 ExternalBetaState = "BETA_REJECTED"
	ExternalBetaExpired                  ExternalBetaState = "EXPIRED"
	ExternalBetaInReview                 ExternalBetaState = "IN_BETA_REVIEW"
	ExternalBetaInTesting                ExternalBetaState = "IN_BETA_TESTING"
	ExternalBetaInExportComplianceReview ExternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	ExternalBetaMissingExportCompliance  ExternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	ExternalBetaProcessing               ExternalBetaState = "PROCESSING"
	ExternalBetaProcessingException      ExternalBetaState = "PROCESSING_EXCEPTION"
	ExternalBetaReadyForBetaSubmission   ExternalBetaState = "READY_FOR_BETA_SUBMISSION"
	ExternalBetaReadyForBetaTesting      ExternalBetaState = "READY_FOR_BETA_TESTING"
	ExternalBetaWaitingForBetaReview     ExternalBetaState = "WAITING_FOR_BETA_REVIEW"
)

// InternalBetaState defines model for InternalBetaState.
type InternalBetaState string

// List of InternalBetaState
const (
	InternalBetaExpired                  InternalBetaState = "EXPIRED"
	InternalBetaInTesting                InternalBetaState = "IN_BETA_TESTING"
	InternalBetaInExportComplianceReview InternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	InternalBetaMissingExportCompliance  InternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	InternalBetaProcessing               InternalBetaState = "PROCESSING"
	InternalBetaProcessingException      InternalBetaState = "PROCESSING_EXCEPTION"
	InternalBetaReadyForBetaTesting      InternalBetaState = "READY_FOR_BETA_TESTING"
)

// BuildBetaDetail defines model for BuildBetaDetail.
type BuildBetaDetail struct {
	Attributes *struct {
		AutoNotifyEnabled  *bool              `json:"autoNotifyEnabled,omitempty"`
		ExternalBuildState *ExternalBetaState `json:"externalBuildState,omitempty"`
		InternalBuildState *InternalBetaState `json:"internalBuildState,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
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
	Data     BuildBetaDetail        `json:"data"`
	Included *[]builds.Build        `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// BuildBetaDetailsResponse defines model for BuildBetaDetailsResponse.
type BuildBetaDetailsResponse struct {
	Data     []BuildBetaDetail           `json:"data"`
	Included *[]builds.Build             `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
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
func (s *Service) ListBuildBetaDetails(params *ListBuildBetaDetailsQuery) (*BuildBetaDetailsResponse, *http.Response, error) {
	res := new(BuildBetaDetailsResponse)
	resp, err := s.GetWithQuery("buildBetaDetails", params, res)
	return res, resp, err
}

// GetBuildBetaDetail gets a specific build beta details resource.
func (s *Service) GetBuildBetaDetail(id string, params *GetBuildBetaDetailsQuery) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBuildBetaDetail gets the build information for a specific build beta details resource.
func (s *Service) GetBuildForBuildBetaDetail(id string, params *GetBuildForBuildBetaDetailQuery) (*builds.BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildBetaDetailForBuild gets the beta test details for a specific build.
func (s *Service) GetBuildBetaDetailForBuild(id string, params *GetBuildBetaDetailForBuildQuery) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/buildBetaDetail", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBuildBetaDetail updates beta test details for a specific build.
func (s *Service) UpdateBuildBetaDetail(id string, body *BuildBetaDetailUpdateRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
