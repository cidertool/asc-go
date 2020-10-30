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

// AppEncryptionDeclarationState defines model for AppEncryptionDeclarationState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationstate
type AppEncryptionDeclarationState string

const (
	// AppEncryptionDeclarationStateApproved is an app encryption declaration state type for Approved.
	AppEncryptionDeclarationStateApproved AppEncryptionDeclarationState = "APPROVED"
	// AppEncryptionDeclarationStateExpired is an app encryption declaration state type for Expired.
	AppEncryptionDeclarationStateExpired AppEncryptionDeclarationState = "EXPIRED"
	// AppEncryptionDeclarationStateInvalid is an app encryption declaration state type for Invalid.
	AppEncryptionDeclarationStateInvalid AppEncryptionDeclarationState = "INVALID"
	// AppEncryptionDeclarationStateInReview is an app encryption declaration state type for InReview.
	AppEncryptionDeclarationStateInReview AppEncryptionDeclarationState = "IN_REVIEW"
	// AppEncryptionDeclarationStateRejected is an app encryption declaration state type for Rejected.
	AppEncryptionDeclarationStateRejected AppEncryptionDeclarationState = "REJECTED"
)

// AppEncryptionDeclaration defines model for AppEncryptionDeclaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclaration
type AppEncryptionDeclaration struct {
	Attributes    *AppEncryptionDeclarationAttributes    `json:"attributes,omitempty"`
	ID            string                                 `json:"id"`
	Links         ResourceLinks                          `json:"links"`
	Relationships *AppEncryptionDeclarationRelationships `json:"relationships,omitempty"`
	Type          string                                 `json:"type"`
}

// AppEncryptionDeclarationAttributes defines model for AppEncryptionDeclaration.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclaration/attributes
type AppEncryptionDeclarationAttributes struct {
	AppEncryptionDeclarationState   *AppEncryptionDeclarationState `json:"appEncryptionDeclarationState,omitempty"`
	AvailableOnFrenchStore          *bool                          `json:"availableOnFrenchStore,omitempty"`
	CodeValue                       *string                        `json:"codeValue,omitempty"`
	ContainsProprietaryCryptography *bool                          `json:"containsProprietaryCryptography,omitempty"`
	ContainsThirdPartyCryptography  *bool                          `json:"containsThirdPartyCryptography,omitempty"`
	DocumentName                    *string                        `json:"documentName,omitempty"`
	DocumentType                    *string                        `json:"documentType,omitempty"`
	DocumentURL                     *string                        `json:"documentUrl,omitempty"`
	Exempt                          *bool                          `json:"exempt,omitempty"`
	Platform                        *Platform                      `json:"platform,omitempty"`
	UploadedDate                    *DateTime                      `json:"uploadedDate,omitempty"`
	UsesEncryption                  *bool                          `json:"usesEncryption,omitempty"`
}

// AppEncryptionDeclarationRelationships defines model for AppEncryptionDeclaration.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclaration/relationships
type AppEncryptionDeclarationRelationships struct {
	App *Relationship `json:"app,omitempty"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationresponse
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included []App                    `json:"included,omitempty"`
	Links    DocumentLinks            `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationsresponse
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration `json:"data"`
	Included []App                      `json:"included,omitempty"`
	Links    PagedDocumentLinks         `json:"links"`
	Meta     *PagingInformation         `json:"meta,omitempty"`
}

// ListAppEncryptionDeclarationsQuery are query options for ListAppEncryptionDeclarations
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_encryption_declarations
type ListAppEncryptionDeclarationsQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FilterApp                       []string `url:"filter[app],omitempty"`
	FilterBuilds                    []string `url:"filter[builds],omitempty"`
	FilterPlatforms                 []string `url:"filter[platforms],omitempty"`
	Include                         []string `url:"include,omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	Cursor                          *string  `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationQuery are query options for GetAppEncryptionDeclaration
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_encryption_declaration_information
type GetAppEncryptionDeclarationQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	Include                         []string `url:"include,omitempty"`
}

// GetAppForEncryptionDeclarationQuery are query options for GetAppForEncryptionDeclaration
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_an_app_encryption_declaration
type GetAppForEncryptionDeclarationQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListAppEncryptionDeclarations finds and lists all available app encryption declarations.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_encryption_declarations
func (s *BuildsService) ListAppEncryptionDeclarations(ctx context.Context, params *ListAppEncryptionDeclarationsQuery) (*AppEncryptionDeclarationsResponse, *Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.client.get(ctx, "appEncryptionDeclarations", params, res)

	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_encryption_declaration_information
func (s *BuildsService) GetAppEncryptionDeclaration(ctx context.Context, id string, params *GetAppEncryptionDeclarationQuery) (*AppEncryptionDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_an_app_encryption_declaration
func (s *BuildsService) GetAppForAppEncryptionDeclaration(ctx context.Context, id string, params *GetAppForEncryptionDeclarationQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_builds_to_an_app_encryption_declaration
func (s *BuildsService) AssignBuildsToAppEncryptionDeclaration(ctx context.Context, id string, buildIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(buildIDs, "builds")
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)

	return s.client.post(ctx, url, newRequestBody(linkages.Data), nil)
}
