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
)

// appStoreVersionReleaseRequestCreateRequest defines model for appStoreVersionReleaseRequestCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionreleaserequestcreaterequest/data
type appStoreVersionReleaseRequestCreateRequest struct {
	Relationships appStoreVersionReleaseRequestCreateRequestRelationships `json:"relationships"`
	Type          string                                                  `json:"type"`
}

// appStoreVersionReleaseRequestCreateRequestRelationships are relationships for appStoreVersionReleaseRequestCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionreleaserequestcreaterequest/data/relationships
type appStoreVersionReleaseRequestCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreVersionReleaseRequestResponse defines model for AppStoreVersionReleaseRequestResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionreleaserequestresponse
type AppStoreVersionReleaseRequestResponse struct {
	Data  AppStoreVersionReleaseRequest `json:"data"`
	Links DocumentLinks                 `json:"links"`
}

// AppStoreVersionReleaseRequest defines model for AppStoreVersionReleaseRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionreleaserequest
type AppStoreVersionReleaseRequest struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// ManuallyReleaseApprovedVersion release an approved version of your app to the App Store.
//
// https://developer.apple.com/documentation/appstoreconnectapi/manually_release_an_app_store_approved_version_of_your_app
func (s *PublishingService) ManuallyReleaseApprovedVersion(ctx context.Context, appStoreVersionID string) (*AppStoreVersionReleaseRequestResponse, *Response, error) {
	req := appStoreVersionReleaseRequestCreateRequest{
		Relationships: appStoreVersionReleaseRequestCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreVersionReleaseRequests",
	}

	res := new(AppStoreVersionReleaseRequestResponse)
	resp, err := s.client.post(ctx, "appStoreVersionReleaseRequests", newRequestBody(req), res)

	return res, resp, err
}
