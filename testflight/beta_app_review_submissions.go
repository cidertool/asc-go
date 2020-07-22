package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal"
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
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppReviewSubmissionCreateRequest defines model for BetaAppReviewSubmissionCreateRequest.
type BetaAppReviewSubmissionCreateRequest struct {
	Data struct {
		Relationships struct {
			Build struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"build"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaAppReviewSubmissionResponse defines model for BetaAppReviewSubmissionResponse.
type BetaAppReviewSubmissionResponse struct {
	Data     BetaAppReviewSubmission `json:"data"`
	Included *[]builds.Build         `json:"included,omitempty"`
	Links    internal.DocumentLinks  `json:"links"`
}

// BetaAppReviewSubmissionsResponse defines model for BetaAppReviewSubmissionsResponse.
type BetaAppReviewSubmissionsResponse struct {
	Data     []BetaAppReviewSubmission   `json:"data"`
	Included *[]builds.Build             `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
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
func (s *Service) CreateBetaAppReviewSubmission(body *BetaAppReviewSubmissionCreateRequest) (*BetaAppReviewSubmissionResponse, *internal.Response, error) {
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.Post("betaAppReviewSubmissions", body, res)
	return res, resp, err
}

// ListBetaAppReviewSubmissions finds and lists beta app review submissions for all builds.
func (s *Service) ListBetaAppReviewSubmissions(params *ListBetaAppReviewSubmissionsQuery) (*BetaAppReviewSubmissionsResponse, *internal.Response, error) {
	res := new(BetaAppReviewSubmissionsResponse)
	resp, err := s.GetWithQuery("betaAppReviewSubmissions", params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmission gets a specific beta app review submission.
func (s *Service) GetBetaAppReviewSubmission(id string, params *GetBetaAppReviewSubmissionQuery) (*BetaAppReviewSubmissionResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBetaAppReviewSubmission gets the build information for a specific beta app review submission.
func (s *Service) GetBuildForBetaAppReviewSubmission(id string, params *GetBuildForBetaAppReviewSubmissionQuery) (*builds.BuildResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmissionForBuild gets the beta app review submission status for a specific build.
func (s *Service) GetBetaAppReviewSubmissionForBuild(id string, params *GetBetaAppReviewSubmissionForBuildQuery) (*BetaAppReviewSubmissionResponse, *internal.Response, error) {
	url := fmt.Sprintf("builds/%s/betaAppReviewSubmission", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
