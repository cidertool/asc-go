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

// ExternalBetaState defines model for ExternalBetaState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/externalbetastate
type ExternalBetaState string

const (
	// ExternalBetaStateApproved is an external beta state for Approved.
	ExternalBetaStateApproved ExternalBetaState = "BETA_APPROVED"
	// ExternalBetaStateRejected is an external beta state for Rejected.
	ExternalBetaStateRejected ExternalBetaState = "BETA_REJECTED"
	// ExternalBetaStateExpired is an external beta state for Expired.
	ExternalBetaStateExpired ExternalBetaState = "EXPIRED"
	// ExternalBetaStateInReview is an external beta state for InReview.
	ExternalBetaStateInReview ExternalBetaState = "IN_BETA_REVIEW"
	// ExternalBetaStateInTesting is an external beta state for InTesting.
	ExternalBetaStateInTesting ExternalBetaState = "IN_BETA_TESTING"
	// ExternalBetaStateInExportComplianceReview is an external beta state for InExportComplianceReview.
	ExternalBetaStateInExportComplianceReview ExternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	// ExternalBetaStateMissingExportCompliance is an external beta state for MissingExportCompliance.
	ExternalBetaStateMissingExportCompliance ExternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	// ExternalBetaStateProcessing is an external beta state for Processing.
	ExternalBetaStateProcessing ExternalBetaState = "PROCESSING"
	// ExternalBetaStateProcessingException is an external beta state for ProcessingException.
	ExternalBetaStateProcessingException ExternalBetaState = "PROCESSING_EXCEPTION"
	// ExternalBetaStateReadyForBetaSubmission is an external beta state for ReadyForBetaSubmission.
	ExternalBetaStateReadyForBetaSubmission ExternalBetaState = "READY_FOR_BETA_SUBMISSION"
	// ExternalBetaStateReadyForBetaTesting is an external beta state for ReadyForBetaTesting.
	ExternalBetaStateReadyForBetaTesting ExternalBetaState = "READY_FOR_BETA_TESTING"
	// ExternalBetaStateWaitingForBetaReview is an external beta state for WaitingForBetaReview.
	ExternalBetaStateWaitingForBetaReview ExternalBetaState = "WAITING_FOR_BETA_REVIEW"
)

// InternalBetaState defines model for InternalBetaState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/internalbetastate
type InternalBetaState string

const (
	// InternalBetaStateExpired is an internal beta state for Expired.
	InternalBetaStateExpired InternalBetaState = "EXPIRED"
	// InternalBetaStateInTesting is an internal beta state for InTesting.
	InternalBetaStateInTesting InternalBetaState = "IN_BETA_TESTING"
	// InternalBetaStateInExportComplianceReview is an internal beta state for InExportComplianceReview.
	InternalBetaStateInExportComplianceReview InternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	// InternalBetaStateMissingExportCompliance is an internal beta state for MissingExportCompliance.
	InternalBetaStateMissingExportCompliance InternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	// InternalBetaStateProcessing is an internal beta state for Processing.
	InternalBetaStateProcessing InternalBetaState = "PROCESSING"
	// InternalBetaStateProcessingException is an internal beta state for ProcessingException.
	InternalBetaStateProcessingException InternalBetaState = "PROCESSING_EXCEPTION"
	// InternalBetaStateReadyForBetaTesting is an internal beta state for ReadyForBetaTesting.
	InternalBetaStateReadyForBetaTesting InternalBetaState = "READY_FOR_BETA_TESTING"
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
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailupdaterequest/data
type buildBetaDetailUpdateRequest struct {
	Attributes *buildBetaDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// BuildBetaDetailUpdateRequestAttributes are attributes for BuildBetaDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetadetailupdaterequest/data/attributes
type buildBetaDetailUpdateRequestAttributes struct {
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
func (s *TestflightService) UpdateBuildBetaDetail(ctx context.Context, id string, autoNotifyEnabled *bool) (*BuildBetaDetailResponse, *Response, error) {
	req := buildBetaDetailUpdateRequest{
		ID:   id,
		Type: "buildBetaDetails",
	}

	if autoNotifyEnabled != nil {
		req.Attributes = &buildBetaDetailUpdateRequestAttributes{
			AutoNotifyEnabled: autoNotifyEnabled,
		}
	}

	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
