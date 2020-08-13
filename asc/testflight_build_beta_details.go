package asc

import (
	"context"
	"fmt"
)

// ExternalBetaState defines model for ExternalBetaState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/externalbetastate
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/internalbetastate
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetail
type BuildBetaDetail struct {
	Attributes    *BuildBetaDetailAttributes    `json:"attributes,omitempty"`
	ID            string                        `json:"id"`
	Links         ResourceLinks                 `json:"links"`
	Relationships *BuildBetaDetailRelationships `json:"relationships,omitempty"`
	Type          string                        `json:"type"`
}

// BuildBetaDetailAttributes defines model for BuildBetaDetail.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetail/attributes
type BuildBetaDetailAttributes struct {
	AutoNotifyEnabled  *bool              `json:"autoNotifyEnabled,omitempty"`
	ExternalBuildState *ExternalBetaState `json:"externalBuildState,omitempty"`
	InternalBuildState *InternalBetaState `json:"internalBuildState,omitempty"`
}

// BuildBetaDetailRelationships defines model for BuildBetaDetail.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetail/relationships
type BuildBetaDetailRelationships struct {
	Build *Relationship `json:"build,omitempty"`
}

// BuildBetaDetailUpdateRequest defines model for BuildBetaDetailUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailupdaterequest
type buildBetaDetailUpdateRequest struct {
	Attributes *BuildBetaDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// BuildBetaDetailUpdateRequestAttributes are attributes for BuildBetaDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailupdaterequest/data/attributes
type BuildBetaDetailUpdateRequestAttributes struct {
	AutoNotifyEnabled *bool `json:"autoNotifyEnabled,omitempty"`
}

// BuildBetaDetailResponse defines model for BuildBetaDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailresponse
type BuildBetaDetailResponse struct {
	Data     BuildBetaDetail `json:"data"`
	Included []Build         `json:"included,omitempty"`
	Links    DocumentLinks   `json:"links"`
}

// BuildBetaDetailsResponse defines model for BuildBetaDetailsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailsresponse
type BuildBetaDetailsResponse struct {
	Data     []BuildBetaDetail  `json:"data"`
	Included []Build            `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListBuildBetaDetailsQuery defines model for ListBuildBetaDetails
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_build_beta_details
type ListBuildBetaDetailsQuery struct {
	FieldsBuilds           []string `url:"fields[builds],omitempty"`
	FieldsBuildBetaDetails []string `url:"fields[buildBetaDetails],omitempty"`
	FilterID               []string `url:"filter[id],omitempty"`
	FilterBuild            []string `url:"filter[build],omitempty"`
	Include                []string `url:"include,omitempty"`
	Limit                  int      `url:"limit,omitempty"`
	Cursor                 string   `url:"cursor,omitempty"`
}

// GetBuildBetaDetailsQuery defines model for GetBuildBetaDetails
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_build_beta_detail_information
type GetBuildBetaDetailsQuery struct {
	FieldsBuilds           []string `url:"fields[builds],omitempty"`
	FieldsBuildBetaDetails []string `url:"fields[buildBetaDetails],omitempty"`
	Include                []string `url:"include,omitempty"`
}

// GetBuildForBuildBetaDetailQuery defines model for GetBuildForBuildBetaDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_build_beta_detail
type GetBuildForBuildBetaDetailQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
}

// GetBuildBetaDetailForBuildQuery defines model for GetBuildBetaDetailForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_beta_details_information_of_a_build
type GetBuildBetaDetailForBuildQuery struct {
	FieldsBuildBetaDetails []string `url:"fields[buildBetaDetails],omitempty"`
}

// ListBuildBetaDetails finds and lists build beta details for all builds.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_build_beta_details
func (s *TestflightService) ListBuildBetaDetails(ctx context.Context, params *ListBuildBetaDetailsQuery) (*BuildBetaDetailsResponse, *Response, error) {
	res := new(BuildBetaDetailsResponse)
	resp, err := s.client.get(ctx, "buildBetaDetails", params, res)
	return res, resp, err
}

// GetBuildBetaDetail gets a specific build beta details resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_build_beta_detail_information
func (s *TestflightService) GetBuildBetaDetail(ctx context.Context, id string, params *GetBuildBetaDetailsQuery) (*BuildBetaDetailResponse, *Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBuildForBuildBetaDetail gets the build information for a specific build beta details resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_build_beta_detail
func (s *TestflightService) GetBuildForBuildBetaDetail(ctx context.Context, id string, params *GetBuildForBuildBetaDetailQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBuildBetaDetailForBuild gets the beta test details for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_beta_details_information_of_a_build
func (s *TestflightService) GetBuildBetaDetailForBuild(ctx context.Context, id string, params *GetBuildBetaDetailForBuildQuery) (*BuildBetaDetailResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/buildBetaDetail", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateBuildBetaDetail updates beta test details for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_build_beta_detail
func (s *TestflightService) UpdateBuildBetaDetail(ctx context.Context, id string, attributes *BuildBetaDetailUpdateRequestAttributes) (*BuildBetaDetailResponse, *Response, error) {
	req := buildBetaDetailUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "buildBetaDetails",
	}
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.patch(ctx, url, req, res)
	return res, resp, err
}
