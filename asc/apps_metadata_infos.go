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

// AppInfo defines model for AppInfo.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo
type AppInfo struct {
	Attributes    *AppInfoAttributes    `json:"attributes,omitempty"`
	ID            string                `json:"id"`
	Links         ResourceLinks         `json:"links"`
	Relationships *AppInfoRelationships `json:"relationships,omitempty"`
	Type          string                `json:"type"`
}

// AppInfoAttributes defines model for AppInfo.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo/attributes
type AppInfoAttributes struct {
	AppStoreAgeRating *AppStoreAgeRating    `json:"appStoreAgeRating,omitempty"`
	AppStoreState     *AppStoreVersionState `json:"appStoreState,omitempty"`
	BrazilAgeRating   *BrazilAgeRating      `json:"brazilAgeRating,omitempty"`
	KidsAgeBand       *KidsAgeBand          `json:"kidsAgeBand,omitempty"`
}

// AppInfoRelationships defines model for AppInfo.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo/relationships
type AppInfoRelationships struct {
	App                     *Relationship      `json:"app,omitempty"`
	AppInfoLocalizations    *PagedRelationship `json:"appInfoLocalizations,omitempty"`
	PrimaryCategory         *Relationship      `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   *Relationship      `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   *Relationship      `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       *Relationship      `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *Relationship      `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *Relationship      `json:"secondarySubcategoryTwo,omitempty"`
	AgeRatingDeclaration    *Relationship      `json:"ageRatingDeclarations,omitempty"`
}

// AppInfoResponse defines model for AppInfoResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinforesponse
type AppInfoResponse struct {
	Data     AppInfo                   `json:"data"`
	Included []AppInfoResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks             `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfosresponse
type AppInfosResponse struct {
	Data     []AppInfo                 `json:"data"`
	Included []AppInfoResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks        `json:"links"`
	Meta     *PagingInformation        `json:"meta,omitempty"`
}

// AppInfoResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a AppInfoResponse or AppInfosResponse.
type AppInfoResponseIncluded included

// appInfoUpdateRequest defines model for AppInfoUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest/data
type appInfoUpdateRequest struct {
	ID            string                             `json:"id"`
	Relationships *appInfoUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppInfoUpdateRequestRelationships are relationships for AppInfoUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest/data/relationships
type appInfoUpdateRequestRelationships struct {
	PrimaryCategory         *relationshipDeclaration `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   *relationshipDeclaration `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   *relationshipDeclaration `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       *relationshipDeclaration `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *relationshipDeclaration `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *relationshipDeclaration `json:"secondarySubcategoryTwo,omitempty"`
}

// AppInfoUpdateRequestRelationships is a public-facing options object for AppInfoUpdateRequest relationships.
type AppInfoUpdateRequestRelationships struct {
	PrimaryCategoryID         *string
	PrimarySubcategoryOneID   *string
	PrimarySubcategoryTwoID   *string
	SecondaryCategoryID       *string
	SecondarySubcategoryOneID *string
	SecondarySubcategoryTwoID *string
}

// GetAppInfoQuery are query options for GetAppInfo
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
type GetAppInfoQuery struct {
	FieldsAppInfos              []string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations  []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories         []string `url:"fields[appCategories],omitempty"`
	Include                     []string `url:"include,omitempty"`
	LimitAppInfoLocalizations   int      `url:"limit[appInfoLocalizations],omitempty"`
	FieldsAgeRatingDeclarations []string `url:"fields[ageRatingDeclarations],omitEmpty"`
}

// ListAppInfosForAppQuery are query options for ListAppInfosForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
type ListAppInfosForAppQuery struct {
	FieldsAppInfos              []string `url:"fields[appInfos],omitempty"`
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsAppInfoLocalizations  []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories         []string `url:"fields[appCategories],omitempty"`
	FieldsAgeRatingDeclarations []string `url:"fields[ageRatingDeclarations],omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	Include                     []string `url:"include,omitempty"`
	Cursor                      string   `url:"cursor,omitempty"`
}

// GetAgeRatingDeclarationForAppInfoQuery are query options for GetAgeRatingDeclarationForInfo
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appinfos_id_ageratingdeclaration
type GetAgeRatingDeclarationForAppInfoQuery struct {
	FieldsAgeRatingDeclarations []string `url:"fields[ageRatingDeclarations],omitempty"`
}

// GetAppInfo reads App Store information including your App Store state, age ratings, Brazil age rating, and kids' age band.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
func (s *AppsService) GetAppInfo(ctx context.Context, id string, params *GetAppInfoQuery) (*AppInfoResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListAppInfosForApp gets information about an app that is currently live on App Store, or that goes live with the next version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
func (s *AppsService) ListAppInfosForApp(ctx context.Context, id string, params *ListAppInfosForAppQuery) (*AppInfosResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/appInfos", id)
	res := new(AppInfosResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateAppInfo updates the App Store categories and sub-categories for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_info
func (s *AppsService) UpdateAppInfo(ctx context.Context, id string, relationships *AppInfoUpdateRequestRelationships) (*AppInfoResponse, *Response, error) {
	req := appInfoUpdateRequest{
		ID:   id,
		Type: "appInfos",
	}

	if relationships != nil {
		req.Relationships = &appInfoUpdateRequestRelationships{
			PrimaryCategory:         newRelationshipDeclaration(relationships.PrimaryCategoryID, "appCategories"),
			PrimarySubcategoryOne:   newRelationshipDeclaration(relationships.PrimarySubcategoryOneID, "appCategories"),
			PrimarySubcategoryTwo:   newRelationshipDeclaration(relationships.PrimarySubcategoryTwoID, "appCategories"),
			SecondaryCategory:       newRelationshipDeclaration(relationships.SecondaryCategoryID, "appCategories"),
			SecondarySubcategoryOne: newRelationshipDeclaration(relationships.SecondarySubcategoryOneID, "appCategories"),
			SecondarySubcategoryTwo: newRelationshipDeclaration(relationships.SecondarySubcategoryTwoID, "appCategories"),
		}
	}

	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in AppInfoResponseIncluded.
func (i *AppInfoResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// AppInfoLocalization returns the AppInfoLocalization stored within, if one is present.
func (i *AppInfoResponseIncluded) AppInfoLocalization() *AppInfoLocalization {
	return extractIncludedAppInfoLocalization(i.inner)
}

// AppCategory returns the AppCategory stored within, if one is present.
func (i *AppInfoResponseIncluded) AppCategory() *AppCategory {
	return extractIncludedAppCategory(i.inner)
}

// GetAgeRatingDeclarationForAppInfo gets the age-related information declared for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appinfos_id_ageratingdeclaration
func (s *AppsService) GetAgeRatingDeclarationForAppInfo(ctx context.Context, id string, params *GetAgeRatingDeclarationForAppInfoQuery) (*AgeRatingDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/ageRatingDeclaration", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
