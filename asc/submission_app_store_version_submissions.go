package asc

import (
	"context"
	"fmt"
)

// AppStoreVersionSubmission defines model for AppStoreVersionSubmission.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmission
type AppStoreVersionSubmission struct {
	ID            string                                  `json:"id"`
	Links         ResourceLinks                           `json:"links"`
	Relationships *AppStoreVersionSubmissionRelationships `json:"relationships,omitempty"`
	Type          string                                  `json:"type"`
}

// AppStoreVersionSubmissionRelationships defines model for AppStoreVersionSubmission.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmission/relationships
type AppStoreVersionSubmissionRelationships struct {
	AppStoreVersion *Relationship `json:"appStoreVersion,omitempty"`
}

// AppStoreVersionSubmissionCreateRequest defines model for AppStoreVersionSubmissionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest
type AppStoreVersionSubmissionCreateRequest struct {
	Relationships AppStoreVersionSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                              `json:"type"`
}

// AppStoreVersionSubmissionCreateRequestRelationships are attributes for AppStoreVersionSubmissionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest/data/relationships
type AppStoreVersionSubmissionCreateRequestRelationships struct {
	AppStoreVersion RelationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreVersionSubmissionResponse defines model for AppStoreVersionSubmissionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissionresponse
type AppStoreVersionSubmissionResponse struct {
	Data  AppStoreVersionSubmission `json:"data"`
	Links DocumentLinks             `json:"links"`
}

// GetAppStoreVersionSubmissionForAppStoreVersionQuery are query options for GetAppStoreVersionSubmissionForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_submission_information_of_an_app_store_version
type GetAppStoreVersionSubmissionForAppStoreVersionQuery struct {
	FieldsAppStoreVersions           []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	Include                          []string `url:"include,omitempty"`
}

func (r *AppStoreVersionSubmissionCreateRequest) applyTypes() {
	if r == nil {
		return
	}
	if r.Type == "" {
		r.Type = "appStoreVersionSubmissions"
	}
	r.Relationships.AppStoreVersion.applyType("appStoreVersions")
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_submission_information_of_an_app_store_version
func (s *SubmissionService) GetAppStoreVersionSubmissionForAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionSubmissionForAppStoreVersionQuery) (*AppStoreVersionSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionSubmission", id)
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
