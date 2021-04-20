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

// EndUserLicenseAgreement defines model for EndUserLicenseAgreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreement
type EndUserLicenseAgreement struct {
	Attributes    *EndUserLicenseAgreementAttributes    `json:"attributes,omitempty"`
	ID            string                                `json:"id"`
	Links         ResourceLinks                         `json:"links"`
	Relationships *EndUserLicenseAgreementRelationships `json:"relationships,omitempty"`
	Type          string                                `json:"type"`
}

// EndUserLicenseAgreementAttributes defines model for EndUserLicenseAgreement.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreement/attributes
type EndUserLicenseAgreementAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// EndUserLicenseAgreementRelationships defines model for EndUserLicenseAgreement.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreement/relationships
type EndUserLicenseAgreementRelationships struct {
	App         *Relationship      `json:"app,omitempty"`
	Territories *PagedRelationship `json:"territories,omitempty"`
}

// endUserLicenseAgreementCreateRequest defines model for EndUserLicenseAgreementCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementcreaterequest/data
type endUserLicenseAgreementCreateRequest struct {
	Attributes    endUserLicenseAgreementCreateRequestAttributes    `json:"attributes"`
	Relationships endUserLicenseAgreementCreateRequestRelationships `json:"relationships"`
	Type          string                                            `json:"type"`
}

// EndUserLicenseAgreementCreateRequestAttributes are attributes for EndUserLicenseAgreementCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementcreaterequest/data/attributes
type endUserLicenseAgreementCreateRequestAttributes struct {
	AgreementText string `json:"agreementText"`
}

// EndUserLicenseAgreementCreateRequestRelationships are relationships for EndUserLicenseAgreementCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementcreaterequest/data/relationships
type endUserLicenseAgreementCreateRequestRelationships struct {
	App         relationshipDeclaration      `json:"app"`
	Territories pagedRelationshipDeclaration `json:"territories"`
}

// endUserLicenseAgreementUpdateRequest defines model for EndUserLicenseAgreementUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementupdaterequest/data
type endUserLicenseAgreementUpdateRequest struct {
	Attributes    *endUserLicenseAgreementUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                             `json:"id"`
	Relationships *endUserLicenseAgreementUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                             `json:"type"`
}

// EndUserLicenseAgreementUpdateRequestAttributes are attributes for EndUserLicenseAgreementUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementupdaterequest/data/attributes
type endUserLicenseAgreementUpdateRequestAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// endUserLicenseAgreementUpdateRequestRelationships are relationships for EndUserLicenseAgreementUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementupdaterequest/data/relationships
type endUserLicenseAgreementUpdateRequestRelationships struct {
	Territories *pagedRelationshipDeclaration `json:"territories,omitempty"`
}

// EndUserLicenseAgreementResponse defines model for EndUserLicenseAgreementResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enduserlicenseagreementresponse
type EndUserLicenseAgreementResponse struct {
	Data     EndUserLicenseAgreement `json:"data"`
	Included []Territory             `json:"included,omitempty"`
	Links    DocumentLinks           `json:"links"`
}

// GetEULAQuery are query options for GetEULA
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_end_user_license_agreement_information
type GetEULAQuery struct {
	FieldsEndUserLicenseAgreements []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsTerritories              []string `url:"fields[territories],omitempty"`
	Include                        []string `url:"include,omitempty"`
	LimitTerritories               int      `url:"limit[territories],omitempty"`
}

// GetEULAForAppQuery are query options for GetEULAForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_end_user_license_agreement_information_of_an_app
type GetEULAForAppQuery struct {
	FieldsEndUserLicenseAgreements []string `url:"fields[endUserLicenseAgreements],omitempty"`
}

// CreateEULA adds a custom end user license agreement (EULA) to an app and configure the territories to which it applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_end_user_license_agreement
func (s *AppsService) CreateEULA(ctx context.Context, agreementText string, appID string, territoryIDs []string) (*EndUserLicenseAgreementResponse, *Response, error) {
	req := endUserLicenseAgreementCreateRequest{
		Attributes: endUserLicenseAgreementCreateRequestAttributes{
			AgreementText: agreementText,
		},
		Relationships: endUserLicenseAgreementCreateRequestRelationships{
			App:         *newRelationshipDeclaration(&appID, "apps"),
			Territories: newPagedRelationshipDeclaration(territoryIDs, "territories"),
		},
		Type: "endUserLicenseAgreements",
	}
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.post(ctx, "endUserLicenseAgreements", newRequestBody(req), res)

	return res, resp, err
}

// UpdateEULA updates the text or territories for your custom end user license agreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_end_user_license_agreement
func (s *AppsService) UpdateEULA(ctx context.Context, id string, agreementText *string, territoryIDs []string) (*EndUserLicenseAgreementResponse, *Response, error) {
	req := endUserLicenseAgreementUpdateRequest{
		ID:   id,
		Type: "endUserLicenseAgreements",
	}
	if agreementText != nil {
		req.Attributes = &endUserLicenseAgreementUpdateRequestAttributes{
			AgreementText: agreementText,
		}
	}

	if len(territoryIDs) > 0 {
		relationships := newPagedRelationshipDeclaration(territoryIDs, "territories")
		req.Relationships = &endUserLicenseAgreementUpdateRequestRelationships{
			Territories: &relationships,
		}
	}

	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteEULA deletes the custom end user license agreement that is associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_end_user_license_agreement
func (s *AppsService) DeleteEULA(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)

	return s.client.delete(ctx, url, nil)
}

// GetEULA gets the custom end user license agreement associated with an app, and the territories it applies to.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_end_user_license_agreement_information
func (s *AppsService) GetEULA(ctx context.Context, id string, params *GetEULAQuery) (*EndUserLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetEULAForApp gets the custom end user license agreement (EULA) for a specific app and the territories where the agreement applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_end_user_license_agreement_information_of_an_app
func (s *AppsService) GetEULAForApp(ctx context.Context, id string, params *GetEULAForAppQuery) (*EndUserLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/endUserLicenseAgreement", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
