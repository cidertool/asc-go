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

// IDFADeclaration defines model for IDFADeclaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclaration
type IDFADeclaration struct {
	Attributes    *IDFADeclarationAttributes    `json:"attributes,omitempty"`
	ID            string                        `json:"id"`
	Links         ResourceLinks                 `json:"links"`
	Relationships *IDFADeclarationRelationships `json:"relationships,omitempty"`
	Type          string                        `json:"type"`
}

// IDFADeclarationAttributes defines model for IDFADeclaration.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclaration/attributes
type IDFADeclarationAttributes struct {
	AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
	AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
	HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
	ServesAds                             *bool `json:"servesAds,omitempty"`
}

// IDFADeclarationRelationships defines model for IDFADeclaration.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclaration/relationships
type IDFADeclarationRelationships struct {
	AppStoreVersion *Relationship `json:"appStoreVersion,omitempty"`
}

// idfaDeclarationCreateRequest defines model for idfaDeclarationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationcreaterequest/data
type idfaDeclarationCreateRequest struct {
	Attributes    IDFADeclarationCreateRequestAttributes    `json:"attributes"`
	Relationships idfaDeclarationCreateRequestRelationships `json:"relationships"`
	Type          string                                    `json:"type"`
}

// IDFADeclarationCreateRequestAttributes are attributes for IDFADeclarationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationcreaterequest/data/attributes
type IDFADeclarationCreateRequestAttributes struct {
	AttributesActionWithPreviousAd        bool `json:"attributesActionWithPreviousAd"`
	AttributesAppInstallationToPreviousAd bool `json:"attributesAppInstallationToPreviousAd"`
	HonorsLimitedAdTracking               bool `json:"honorsLimitedAdTracking"`
	ServesAds                             bool `json:"servesAds"`
}

// idfaDeclarationCreateRequestRelationships are relationships for IDFADeclarationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationcreaterequest/data/relationships
type idfaDeclarationCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
}

// IDFADeclarationUpdateRequest defines model for IDFADeclarationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationupdaterequest/data
type idfaDeclarationUpdateRequest struct {
	Attributes *IDFADeclarationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// IDFADeclarationUpdateRequestAttributes are attributes for IDFADeclarationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationupdaterequest/data/attributes
type IDFADeclarationUpdateRequestAttributes struct {
	AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
	AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
	HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
	ServesAds                             *bool `json:"servesAds,omitempty"`
}

// IDFADeclarationResponse defines model for IDFADeclarationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/idfadeclarationresponse
type IDFADeclarationResponse struct {
	Data  IDFADeclaration `json:"data"`
	Links DocumentLinks   `json:"links"`
}

// GetIDFADeclarationForAppStoreVersionQuery are query options for GetIDFADeclarationForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_idfa_declaration_information_of_an_app_store_version
type GetIDFADeclarationForAppStoreVersionQuery struct {
	FieldsIDFADeclarations []string `url:"fields[idfaDeclarations],omitempty"`
}

// CreateIDFADeclaration declares the IDFA usage for an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_idfa_declaration
func (s *SubmissionService) CreateIDFADeclaration(ctx context.Context, attributes IDFADeclarationCreateRequestAttributes, appStoreVersionID string) (*IDFADeclarationResponse, *Response, error) {
	req := idfaDeclarationCreateRequest{
		Attributes: attributes,
		Relationships: idfaDeclarationCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "idfaDeclarations",
	}
	res := new(IDFADeclarationResponse)
	resp, err := s.client.post(ctx, "idfaDeclarations", newRequestBody(req), res)

	return res, resp, err
}

// UpdateIDFADeclaration updates your declared IDFA usage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_idfa_declaration
func (s *SubmissionService) UpdateIDFADeclaration(ctx context.Context, id string, attributes *IDFADeclarationUpdateRequestAttributes) (*IDFADeclarationResponse, *Response, error) {
	req := idfaDeclarationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "idfaDeclarations",
	}
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteIDFADeclaration deletes the IDFA declaration that is associated with a version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_idfa_declaration
func (s *SubmissionService) DeleteIDFADeclaration(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)

	return s.client.delete(ctx, url, nil)
}

// GetIDFADeclarationForAppStoreVersion reads your declared Advertising Identifier (IDFA) usage responses.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_idfa_declaration_information_of_an_app_store_version
func (s *SubmissionService) GetIDFADeclarationForAppStoreVersion(ctx context.Context, id string, params *GetIDFADeclarationForAppStoreVersionQuery) (*IDFADeclarationResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/idfaDeclaration", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
