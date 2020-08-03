package asc

import (
	"fmt"
	"net/http"
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
	FieldsAppStoreVersions           *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions *[]string `url:"fields[appStoreVersionSubmissions],omitempty"`
	Include                          *[]string `url:"include,omitempty"`
}

// CreateSubmission submits an App Store version to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_submission
func (s *SubmissionService) CreateSubmission(body *AppStoreVersionSubmissionCreateRequest) (*AppStoreVersionSubmissionResponse, *http.Response, error) {
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.Post("appStoreVersionSubmissions", body, res)
	return res, resp, err
}

// DeleteSubmission removes a version from App Store review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_submission
func (s *SubmissionService) DeleteSubmission(id string) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.client.Delete(url, nil)
}

// GetAppStoreVersionSubmissionForAppStoreVersion reads the App Store Version Submission Information of an App Store Version
func (s *SubmissionService) GetAppStoreVersionSubmissionForAppStoreVersion(id string, params *GetAppStoreVersionSubmissionForAppStoreVersionQuery) (*AppStoreVersionSubmissionResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionSubmission", id)
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}
