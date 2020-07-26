package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal/services"
)

// BetaReviewState defines model for BetaReviewState.
type BetaReviewState string

// List of BetaReviewState
const (
	Approved         BetaReviewState = "APPROVED"
	InReview         BetaReviewState = "IN_REVIEW"
	Rejected         BetaReviewState = "REJECTED"
	WaitingForReview BetaReviewState = "WAITING_FOR_REVIEW"
)

// BetaAppReviewSubmission defines model for BetaAppReviewSubmission.
type BetaAppReviewSubmission struct {
	Attributes *struct {
		BetaReviewState *BetaReviewState `json:"betaReviewState,omitempty"`
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

// BetaAppReviewSubmissionCreateRequest defines model for BetaAppReviewSubmissionCreateRequest.
type BetaAppReviewSubmissionCreateRequest struct {
	Relationships BetaAppReviewSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                            `json:"type"`
}

// BetaAppReviewSubmissionCreateRequestRelationships are relationships for BetaAppReviewSubmissionCreateRequest
type BetaAppReviewSubmissionCreateRequestRelationships struct {
	Build struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"build"`
}

// BetaAppReviewSubmissionResponse defines model for BetaAppReviewSubmissionResponse.
type BetaAppReviewSubmissionResponse struct {
	Data     BetaAppReviewSubmission `json:"data"`
	Included *[]builds.Build         `json:"included,omitempty"`
	Links    services.DocumentLinks  `json:"links"`
}

// BetaAppReviewSubmissionsResponse defines model for BetaAppReviewSubmissionsResponse.
type BetaAppReviewSubmissionsResponse struct {
	Data     []BetaAppReviewSubmission   `json:"data"`
	Included *[]builds.Build             `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
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
func (s *Service) CreateBetaAppReviewSubmission(body *BetaAppReviewSubmissionCreateRequest) (*BetaAppReviewSubmissionResponse, *services.Response, error) {
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.Post("betaAppReviewSubmissions", body, res)
	return res, resp, err
}

// ListBetaAppReviewSubmissions finds and lists beta app review submissions for all builds.
func (s *Service) ListBetaAppReviewSubmissions(params *ListBetaAppReviewSubmissionsQuery) (*BetaAppReviewSubmissionsResponse, *services.Response, error) {
	res := new(BetaAppReviewSubmissionsResponse)
	resp, err := s.GetWithQuery("betaAppReviewSubmissions", params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmission gets a specific beta app review submission.
func (s *Service) GetBetaAppReviewSubmission(id string, params *GetBetaAppReviewSubmissionQuery) (*BetaAppReviewSubmissionResponse, *services.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBetaAppReviewSubmission gets the build information for a specific beta app review submission.
func (s *Service) GetBuildForBetaAppReviewSubmission(id string, params *GetBuildForBetaAppReviewSubmissionQuery) (*builds.BuildResponse, *services.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmissionForBuild gets the beta app review submission status for a specific build.
func (s *Service) GetBetaAppReviewSubmissionForBuild(id string, params *GetBetaAppReviewSubmissionForBuildQuery) (*BetaAppReviewSubmissionResponse, *services.Response, error) {
	url := fmt.Sprintf("builds/%s/betaAppReviewSubmission", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
