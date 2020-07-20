package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
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
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppReviewSubmissionCreateRequest defines model for BetaAppReviewSubmissionCreateRequest.
type BetaAppReviewSubmissionCreateRequest struct {
	Data struct {
		Relationships struct {
			Build struct {
				Data common.RelationshipsData `json:"data"`
			} `json:"build"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaAppReviewSubmissionResponse defines model for BetaAppReviewSubmissionResponse.
type BetaAppReviewSubmissionResponse struct {
	Data     BetaAppReviewSubmission `json:"data"`
	Included *[]builds.Build         `json:"included,omitempty"`
	Links    common.DocumentLinks    `json:"links"`
}

// BetaAppReviewSubmissionsResponse defines model for BetaAppReviewSubmissionsResponse.
type BetaAppReviewSubmissionsResponse struct {
	Data     []BetaAppReviewSubmission `json:"data"`
	Included *[]builds.Build           `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaAppReviewSubmissionsQuery struct {
	FieldsBuilds                   *[]string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FilterBuild                    *[]string `url:"filter[build],omitempty"`
	FilterBetaReviewState          *[]string `url:"filter[betaReviewState],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
	Limit                          *int      `url:"limit,omitempty"`
	Cursor                         *string   `url:"cursor,omitempty"`
}

type GetBetaAppReviewSubmissionQuery struct {
	FieldsBuilds                   *[]string `url:"fields[builds],omitempty"`
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
}

type GetBuildForBetaAppReviewSubmissionQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

type GetBetaAppReviewSubmissionForBuildQuery struct {
	FieldsBetaAppReviewSubmissions *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
}

// CreateBetaAppReviewSubmission submits an app for beta app review to allow external testing.
func (s *Service) CreateBetaAppReviewSubmission(body *BetaAppReviewSubmissionCreateRequest) (*BetaAppReviewSubmissionResponse, *common.Response, error) {
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.Post("betaAppReviewSubmissions", body, res)
	return res, resp, err
}

// ListBetaAppReviewSubmissions finds and lists beta app review submissions for all builds.
func (s *Service) ListBetaAppReviewSubmissions(params *ListBetaAppReviewSubmissionsQuery) (*BetaAppReviewSubmissionsResponse, *common.Response, error) {
	res := new(BetaAppReviewSubmissionsResponse)
	resp, err := s.GetWithQuery("betaAppReviewSubmissions", params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmission gets a specific beta app review submission.
func (s *Service) GetBetaAppReviewSubmission(id string, params *GetBetaAppReviewSubmissionQuery) (*BetaAppReviewSubmissionResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBetaAppReviewSubmission gets the build information for a specific beta app review submission.
func (s *Service) GetBuildForBetaAppReviewSubmission(id string, params *GetBuildForBetaAppReviewSubmissionQuery) (*builds.BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppReviewSubmissions/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewSubmissionForBuild gets the beta app review submission status for a specific build.
func (s *Service) GetBetaAppReviewSubmissionForBuild(id string, params *GetBetaAppReviewSubmissionForBuildQuery) (*BetaAppReviewSubmissionResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/betaAppReviewSubmission", id)
	res := new(BetaAppReviewSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
