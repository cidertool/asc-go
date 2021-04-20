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

// PhasedReleaseState defines model for PhasedReleaseState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/phasedreleasestate
type PhasedReleaseState string

const (
	// PhasedReleaseStateInactive is a representation of the INACTIVE state.
	PhasedReleaseStateInactive PhasedReleaseState = "INACTIVE"
	// PhasedReleaseStateActive is a representation of the ACTIVE state.
	PhasedReleaseStateActive PhasedReleaseState = "ACTIVE"
	// PhasedReleaseStatePaused is a representation of the PAUSED state.
	PhasedReleaseStatePaused PhasedReleaseState = "PAUSED"
	// PhasedReleaseStateComplete is a representation of the COMPLETE state.
	PhasedReleaseStateComplete PhasedReleaseState = "COMPLETE"
)

// AppStoreVersionPhasedRelease defines model for AppStoreVersionPhasedRelease.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedrelease
type AppStoreVersionPhasedRelease struct {
	Attributes *AppStoreVersionPhasedReleaseAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Links      ResourceLinks                           `json:"links"`
	Type       string                                  `json:"type"`
}

// AppStoreVersionPhasedReleaseAttributes defines model for AppStoreVersionPhasedRelease.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedrelease/attributes
type AppStoreVersionPhasedReleaseAttributes struct {
	CurrentDayNumber   *int                `json:"currentDayNumber,omitempty"`
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
	StartDate          *DateTime           `json:"startDate,omitempty"`
	TotalPauseDuration *int                `json:"totalPauseDuration,omitempty"`
}

// appStoreVersionPhasedReleaseCreateRequest defines model for appStoreVersionPhasedReleaseCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest/data
type appStoreVersionPhasedReleaseCreateRequest struct {
	Attributes    *appStoreVersionPhasedReleaseCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships appStoreVersionPhasedReleaseCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// AppStoreVersionPhasedReleaseCreateRequestAttributes are attributes for AppStoreVersionPhasedReleaseCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest/data/attributes
type appStoreVersionPhasedReleaseCreateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// AppStoreVersionPhasedReleaseCreateRequestRelationships are relationships for AppStoreVersionPhasedReleaseCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest/data/relationships
type appStoreVersionPhasedReleaseCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreVersionPhasedReleaseResponse defines model for AppStoreVersionPhasedReleaseResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseresponse
type AppStoreVersionPhasedReleaseResponse struct {
	Data  AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks                `json:"links"`
}

// AppStoreVersionPhasedReleaseUpdateRequest defines model for AppStoreVersionPhasedReleaseUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseupdaterequest/data
type appStoreVersionPhasedReleaseUpdateRequest struct {
	Attributes *appStoreVersionPhasedReleaseUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                               `json:"id"`
	Type       string                                               `json:"type"`
}

// AppStoreVersionPhasedReleaseUpdateRequestAttributes are attributes for AppStoreVersionPhasedReleaseUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseupdaterequest/data/attributes
type appStoreVersionPhasedReleaseUpdateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery are query options for GetAppStoreVersionPhasedReleaseForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_phased_release_information_of_an_app_store_version
type GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery struct {
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
}

// CreatePhasedRelease enables phased release for an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_phased_release
func (s *PublishingService) CreatePhasedRelease(ctx context.Context, phasedReleaseState *PhasedReleaseState, appStoreVersionID string) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	req := appStoreVersionPhasedReleaseCreateRequest{
		Relationships: appStoreVersionPhasedReleaseCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreVersionPhasedReleases",
	}

	if phasedReleaseState != nil {
		req.Attributes = &appStoreVersionPhasedReleaseCreateRequestAttributes{
			PhasedReleaseState: phasedReleaseState,
		}
	}

	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.post(ctx, "appStoreVersionPhasedReleases", newRequestBody(req), res)

	return res, resp, err
}

// UpdatePhasedRelease pauses or resumes a phased release, or immediately release an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_version_phased_release
func (s *PublishingService) UpdatePhasedRelease(ctx context.Context, id string, state *PhasedReleaseState) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	req := appStoreVersionPhasedReleaseUpdateRequest{
		ID:   id,
		Type: "appStoreVersionPhasedReleases",
	}

	if state != nil {
		req.Attributes = &appStoreVersionPhasedReleaseUpdateRequestAttributes{
			PhasedReleaseState: state,
		}
	}

	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeletePhasedRelease cancels a planned phased release that has not been started.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_phased_release
func (s *PublishingService) DeletePhasedRelease(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)

	return s.client.delete(ctx, url, nil)
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersion reads the phased release status and configuration for a version with phased release enabled.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_phased_release_information_of_an_app_store_version
func (s *PublishingService) GetAppStoreVersionPhasedReleaseForAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionPhasedRelease", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
