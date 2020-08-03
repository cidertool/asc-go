package asc

import (
	"fmt"
	"net/http"
)

// BetaReviewState defines model for BetaReviewState.
type BetaReviewState string

// List of BetaReviewState
const (
	BetaReviewStateApproved         BetaReviewState = "APPROVED"
	BetaReviewStateInReview         BetaReviewState = "IN_REVIEW"
	BetaReviewStateRejected         BetaReviewState = "REJECTED"
	BetaReviewStateWaitingForReview BetaReviewState = "WAITING_FOR_REVIEW"
)

// BetaAppReviewSubmission defines model for BetaAppReviewSubmission.
type BetaAppReviewSubmission struct {
	Attributes *struct {
		BetaReviewState *BetaReviewState `json:"betaReviewState,omitempty"`
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

// BetaAppReviewSubmissionCreateRequest defines model for BetaAppReviewSubmissionCreateRequest.
type BetaAppReviewSubmissionCreateRequest struct {
	Relationships BetaAppReviewSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                            `json:"type"`
}

// BetaAppReviewSubmissionCreateRequestRelationships are relationships for BetaAppReviewSubmissionCreateRequest
type BetaAppReviewSubmissionCreateRequestRelationships struct {
	Build struct {
		Data RelationshipsData `json:"data"`
	} `json:"build"`
}

// BetaAppReviewSubmissionResponse defines model for BetaAppReviewSubmissionResponse.
type BetaAppReviewSubmissionResponse struct {
	Data     BetaAppReviewSubmission `json:"data"`
	Included *[]Build                `json:"included,omitempty"`
	Links    DocumentLinks           `json:"links"`
}

// BetaAppReviewSubmissionsResponse defines model for BetaAppReviewSubmissionsResponse.
type BetaAppReviewSubmissionsResponse struct {
	Data     []BetaAppReviewSubmission `json:"data"`
	Included *[]Build                  `json:"included,omitempty"`
	Links    PagedDocumentLinks        `json:"links"`
	Meta     *PagingInformation        `json:"meta,omitempty"`
}

// ListBetaAppReviewSubmissionsQuery defines model for ListBetaAppReviewSubmissions
type ListBetaAppReviewSubmissionsQuery struct {
	FieldsBuilds                   *[]string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FilterBuild                    *[]string `url:"filter[build],omitempty"`
	FilterBetaReviewState          *[]string `url:"filter[betaReviewState],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
	Limit                          *int      `url:"limit,omitempty"`
	Cursor                         *string   `url:"cursor,omitempty"`
}

// GetBetaAppReviewSubmissionQuery defines model for GetBetaAppReviewSubmission
type GetBetaAppReviewSubmissionQuery struct {
	FieldsBuilds                   *[]string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
}

// GetBuildForBetaAppReviewSubmissionQuery defines model for GetBuildForBetaAppReviewSubmission
type GetBuildForBetaAppReviewSubmissionQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

// GetBetaAppReviewSubmissionForBuildQuery defines model for GetBetaAppReviewSubmissionForBuild
type GetBetaAppReviewSubmissionForBuildQuery struct {
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
}

// CreateBetaAppReviewSubmission submits an app for beta app review to allow external testing.
func (s *TestflightService) CreateBetaAppReviewSubmission(body *BetaAppReviewSubmissionCreateRequest) (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.post("betaAppReviewSubmissions", body, res)
	return res, resp, err
}

// ListBetaAppReviewSubmissions finds and lists beta app review submissions for all builds.
func (s *TestflightService) ListBetaAppReviewSubmissions(params *ListBetaAppReviewSubmissionsQuery) (*BetaAppReviewSubmissionsResponse, *http.Response, error) {
	res := new(BetaAppReviewSubmissionsResponse)
	resp, err := s.client.get("betaAppReviewSubmissions", params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmission gets a specific beta app review submission.
func (s *TestflightService) GetBetaAppReviewSubmission(id string, params *GetBetaAppReviewSubmissionQuery) (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBuildForBetaAppReviewSubmission gets the build information for a specific beta app review submission.
func (s *TestflightService) GetBuildForBetaAppReviewSubmission(id string, params *GetBuildForBetaAppReviewSubmissionQuery) (*BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmissionForBuild gets the beta app review submission status for a specific build.
func (s *TestflightService) GetBetaAppReviewSubmissionForBuild(id string, params *GetBetaAppReviewSubmissionForBuildQuery) (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/betaAppReviewSubmission", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
