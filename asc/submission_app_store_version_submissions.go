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

// appStoreVersionSubmissionCreateRequest defines model for appStoreVersionSubmissionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest/data
type appStoreVersionSubmissionCreateRequest struct {
	Relationships appStoreVersionSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                              `json:"type"`
}

// appStoreVersionSubmissionCreateRequestRelationships are attributes for AppStoreVersionSubmissionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest/data/relationships
type appStoreVersionSubmissionCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
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

// CreateSubmission submits an App Store version to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_submission
func (s *SubmissionService) CreateSubmission(ctx context.Context, appStoreVersionID string) (*AppStoreVersionSubmissionResponse, *Response, error) {
	req := appStoreVersionSubmissionCreateRequest{
		Relationships: appStoreVersionSubmissionCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreVersionSubmissions",
	}
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.client.post(ctx, "appStoreVersionSubmissions", newRequestBody(req), res)

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
