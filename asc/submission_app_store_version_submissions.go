package asc

import (
	"context"
	"fmt"
)

// AppStoreVersionSubmission defines model for AppStoreVersionSubmission.
type AppStoreVersionSubmission struct {
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionSubmissionCreateRequest defines model for AppStoreVersionSubmissionCreateRequest.
type AppStoreVersionSubmissionCreateRequest struct {
	Relationships AppStoreVersionSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                              `json:"type"`
}

// AppStoreVersionSubmissionCreateRequestRelationships are attributes for AppStoreVersionSubmissionCreateRequest
type AppStoreVersionSubmissionCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// AppStoreVersionSubmissionResponse defines model for AppStoreVersionSubmissionResponse.
type AppStoreVersionSubmissionResponse struct {
	Data  AppStoreVersionSubmission `json:"data"`
	Links DocumentLinks             `json:"links"`
}

// GetAppStoreVersionSubmissionForAppStoreVersionQuery are query options for GetAppStoreVersionSubmissionForAppStoreVersion
type GetAppStoreVersionSubmissionForAppStoreVersionQuery struct {
	FieldsAppStoreVersions           []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	Include                          []string `url:"include,omitempty"`
}

// CreateSubmission submits an App Store version to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_submission
func (s *SubmissionService) CreateSubmission(ctx context.Context, body *AppStoreVersionSubmissionCreateRequest) (*AppStoreVersionSubmissionResponse, *Response, error) {
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.post(ctx, "appStoreVersionSubmissions", body, res)
	return res, resp, err
}

// DeleteSubmission removes a version from App Store review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_submission
func (s *SubmissionService) DeleteSubmission(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.client.delete(ctx, url, nil)
}

// GetAppStoreVersionSubmissionForAppStoreVersion reads the App Store Version Submission Information of an App Store Version
func (s *SubmissionService) GetAppStoreVersionSubmissionForAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionSubmissionForAppStoreVersionQuery) (*AppStoreVersionSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionSubmission", id)
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
