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

// BetaReviewState defines model for BetaReviewState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betareviewstate
type BetaReviewState string

const (
	// BetaReviewStateApproved is a beta review satte for Approved.
	BetaReviewStateApproved BetaReviewState = "APPROVED"
	// BetaReviewStateInReview is a beta review satte for InReview.
	BetaReviewStateInReview BetaReviewState = "IN_REVIEW"
	// BetaReviewStateRejected is a beta review satte for Rejected.
	BetaReviewStateRejected BetaReviewState = "REJECTED"
	// BetaReviewStateWaitingForReview is a beta review satte for WaitingForReview.
	BetaReviewStateWaitingForReview BetaReviewState = "WAITING_FOR_REVIEW"
)

// BetaAppReviewSubmission defines model for BetaAppReviewSubmission.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmission
type BetaAppReviewSubmission struct {
	Attributes    *BetaAppReviewSubmissionAttributes    `json:"attributes,omitempty"`
	ID            string                                `json:"id"`
	Links         ResourceLinks                         `json:"links"`
	Relationships *BetaAppReviewSubmissionRelationships `json:"relationships,omitempty"`
	Type          string                                `json:"type"`
}

// BetaAppReviewSubmissionAttributes defines model for BetaAppReviewSubmission.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmission/attributes
type BetaAppReviewSubmissionAttributes struct {
	BetaReviewState *BetaReviewState `json:"betaReviewState,omitempty"`
}

// BetaAppReviewSubmissionRelationships defines model for BetaAppReviewSubmission.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmission/relationships
type BetaAppReviewSubmissionRelationships struct {
	Build *Relationship `json:"build,omitempty"`
}

// betaAppReviewSubmissionCreateRequest defines model for betaAppReviewSubmissionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmissioncreaterequest/data
type betaAppReviewSubmissionCreateRequest struct {
	Relationships betaAppReviewSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                            `json:"type"`
}

// betaAppReviewSubmissionCreateRequestRelationships are relationships for BetaAppReviewSubmissionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmissioncreaterequest/data/relationships
type betaAppReviewSubmissionCreateRequestRelationships struct {
	Build relationshipDeclaration `json:"build"`
}

// BetaAppReviewSubmissionResponse defines model for BetaAppReviewSubmissionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmissionresponse
type BetaAppReviewSubmissionResponse struct {
	Data     BetaAppReviewSubmission `json:"data"`
	Included []Build                 `json:"included,omitempty"`
	Links    DocumentLinks           `json:"links"`
}

// BetaAppReviewSubmissionsResponse defines model for BetaAppReviewSubmissionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewsubmissionsresponse
type BetaAppReviewSubmissionsResponse struct {
	Data     []BetaAppReviewSubmission `json:"data"`
	Included []Build                   `json:"included,omitempty"`
	Links    PagedDocumentLinks        `json:"links"`
	Meta     *PagingInformation        `json:"meta,omitempty"`
}

// ListBetaAppReviewSubmissionsQuery defines model for ListBetaAppReviewSubmissions
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_review_submissions
type ListBetaAppReviewSubmissionsQuery struct {
	FieldsBuilds                   []string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions []string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FilterBuild                    []string `url:"filter[build],omitempty"`
	FilterBetaReviewState          []string `url:"filter[betaReviewState],omitempty"`
	Include                        []string `url:"include,omitempty"`
	Limit                          int      `url:"limit,omitempty"`
	Cursor                         string   `url:"cursor,omitempty"`
}

// GetBetaAppReviewSubmissionQuery defines model for GetBetaAppReviewSubmission
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_review_submission_information
type GetBetaAppReviewSubmissionQuery struct {
	FieldsBuilds                   []string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions []string `url:"fields[betaAppReviewSubmissions],omitempty"`
	Include                        []string `url:"include,omitempty"`
}

// GetBuildForBetaAppReviewSubmissionQuery defines model for GetBuildForBetaAppReviewSubmission
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_beta_app_review_submission
type GetBuildForBetaAppReviewSubmissionQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
}

// GetBetaAppReviewSubmissionForBuildQuery defines model for GetBetaAppReviewSubmissionForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_app_review_submission_of_a_build
type GetBetaAppReviewSubmissionForBuildQuery struct {
	FieldsBetaAppReviewSubmissions []string `url:"fields[betaAppReviewSubmissions],omitempty"`
}

// CreateBetaAppReviewSubmission submits an app for beta app review to allow external testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/submit_an_app_for_beta_review
func (s *TestflightService) CreateBetaAppReviewSubmission(ctx context.Context, buildID string) (*BetaAppReviewSubmissionResponse, *Response, error) {
	req := betaAppReviewSubmissionCreateRequest{
		Relationships: betaAppReviewSubmissionCreateRequestRelationships{
			Build: relationshipDeclaration{
				Data: RelationshipData{
					ID:   buildID,
					Type: "builds",
				},
			},
		},
		Type: "betaAppReviewSubmissions",
	}
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.post(ctx, "betaAppReviewSubmissions", newRequestBody(req), res)

	return res, resp, err
}

// ListBetaAppReviewSubmissions finds and lists beta app review submissions for all builds.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_review_submissions
func (s *TestflightService) ListBetaAppReviewSubmissions(ctx context.Context, params *ListBetaAppReviewSubmissionsQuery) (*BetaAppReviewSubmissionsResponse, *Response, error) {
	res := new(BetaAppReviewSubmissionsResponse)
	resp, err := s.client.get(ctx, "betaAppReviewSubmissions", params, res)

	return res, resp, err
}

// GetBetaAppReviewSubmission gets a specific beta app review submission.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_review_submission_information
func (s *TestflightService) GetBetaAppReviewSubmission(ctx context.Context, id string, params *GetBetaAppReviewSubmissionQuery) (*BetaAppReviewSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetBuildForBetaAppReviewSubmission gets the build information for a specific beta app review submission.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_beta_app_review_submission
func (s *TestflightService) GetBuildForBetaAppReviewSubmission(ctx context.Context, id string, params *GetBuildForBetaAppReviewSubmissionQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetBetaAppReviewSubmissionForBuild gets the beta app review submission status for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_app_review_submission_of_a_build
func (s *TestflightService) GetBetaAppReviewSubmissionForBuild(ctx context.Context, id string, params *GetBetaAppReviewSubmissionForBuildQuery) (*BetaAppReviewSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/betaAppReviewSubmission", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
