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

// AppStoreVersionLocalization defines model for AppStoreVersionLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalization
type AppStoreVersionLocalization struct {
	Attributes    *AppStoreVersionLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                                    `json:"id"`
	Links         ResourceLinks                             `json:"links"`
	Relationships *AppStoreVersionLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                                    `json:"type"`
}

// AppStoreVersionLocalizationAttributes defines model for AppStoreVersionLocalization.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalization/attributes
type AppStoreVersionLocalizationAttributes struct {
	Description     *string `json:"description,omitempty"`
	Keywords        *string `json:"keywords,omitempty"`
	Locale          *string `json:"locale,omitempty"`
	MarketingURL    *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportURL      *string `json:"supportUrl,omitempty"`
	WhatsNew        *string `json:"whatsNew,omitempty"`
}

// AppStoreVersionLocalizationRelationships defines model for AppStoreVersionLocalization.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalization/relationships
type AppStoreVersionLocalizationRelationships struct {
	AppPreviewSets    *PagedRelationship `json:"appPreviewSets,omitempty"`
	AppScreenshotSets *PagedRelationship `json:"appScreenshotSets,omitempty"`
	AppStoreVersion   *Relationship      `json:"appStoreVersion,omitempty"`
}

// AppStoreVersionLocalizationCreateRequest defines model for AppStoreVersionLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationcreaterequest/data
type appStoreVersionLocalizationCreateRequest struct {
	Attributes    AppStoreVersionLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships appStoreVersionLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                                `json:"type"`
}

// AppStoreVersionLocalizationCreateRequestAttributes are attributes for AppStoreVersionLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationcreaterequest/data/attributes
type AppStoreVersionLocalizationCreateRequestAttributes struct {
	Description     *string `json:"description,omitempty"`
	Keywords        *string `json:"keywords,omitempty"`
	Locale          string  `json:"locale"`
	MarketingURL    *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportURL      *string `json:"supportUrl,omitempty"`
	WhatsNew        *string `json:"whatsNew,omitempty"`
}

// AppStoreVersionLocalizationCreateRequestRelationships are relationships for AppStoreVersionLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationcreaterequest/data/relationships
type appStoreVersionLocalizationCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreVersionLocalizationResponse defines model for AppStoreVersionLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationresponse
type AppStoreVersionLocalizationResponse struct {
	Data     AppStoreVersionLocalization                   `json:"data"`
	Included []AppStoreVersionLocalizationResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                 `json:"links"`
}

// AppStoreVersionLocalizationsResponse defines model for AppStoreVersionLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationsresponse
type AppStoreVersionLocalizationsResponse struct {
	Data     []AppStoreVersionLocalization                 `json:"data"`
	Included []AppStoreVersionLocalizationResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                            `json:"links"`
	Meta     *PagingInformation                            `json:"meta,omitempty"`
}

// AppStoreVersionLocalizationResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a AppStoreVersionLocalizationResponse or AppStoreVersionLocalizationsResponse.
type AppStoreVersionLocalizationResponseIncluded included

// AppStoreVersionLocalizationUpdateRequest defines model for AppStoreVersionLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationupdaterequest/data
type appStoreVersionLocalizationUpdateRequest struct {
	Attributes *AppStoreVersionLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                              `json:"id"`
	Type       string                                              `json:"type"`
}

// AppStoreVersionLocalizationUpdateRequestAttributes are attributes for AppStoreVersionLocalizationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionlocalizationupdaterequest/data/attributes
type AppStoreVersionLocalizationUpdateRequestAttributes struct {
	Description     *string `json:"description,omitempty"`
	Keywords        *string `json:"keywords,omitempty"`
	MarketingURL    *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportURL      *string `json:"supportUrl,omitempty"`
	WhatsNew        *string `json:"whatsNew,omitempty"`
}

// ListLocalizationsForAppStoreVersionQuery are query options for ListLocalizationsForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_version_localizations_for_an_app_store_version
type ListLocalizationsForAppStoreVersionQuery struct {
	FieldsAppStoreVersionLocalizations []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              int      `url:"limit,omitempty"`
	Cursor                             string   `url:"cursor,omitempty"`
}

// GetAppStoreVersionLocalizationQuery are query options for GetAppStoreVersionLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_localization_information
type GetAppStoreVersionLocalizationQuery struct {
	FieldsAppPreviewSets               []string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppScreenshotSets            []string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppStoreVersionLocalizations []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                            []string `url:"include,omitempty"`
	LimitAppPreviewSets                int      `url:"limit[appPreviewSets],omitempty"`
	LimitAppScreenshotSets             int      `url:"limit[appScreenshotSets],omitempty"`
}

// ListAppScreenshotSetsForAppStoreVersionLocalizationQuery are query options for ListAppScreenshotSetsForAppStoreVersionLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_screenshot_sets_for_an_app_store_version_localization
type ListAppScreenshotSetsForAppStoreVersionLocalizationQuery struct {
	FieldsAppScreenshotSets            []string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots               []string `url:"fields[appScreenshots],omitempty"`
	FieldsAppStoreVersionLocalizations []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              int      `url:"limit,omitempty"`
	Include                            []string `url:"include,omitempty"`
	FilterScreenshotDisplayType        []string `url:"filter[screenshotDisplayType],omitempty"`
	Cursor                             string   `url:"cursor,omitempty"`
}

// ListAppPreviewSetsForAppStoreVersionLocalizationQuery are query options for ListAppPreviewSetsForAppStoreVersionLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_preview_sets_for_an_app_store_version_localization
type ListAppPreviewSetsForAppStoreVersionLocalizationQuery struct {
	FieldsAppPreviewSets               []string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppPreviews                  []string `url:"fields[appPreviews],omitempty"`
	FieldsAppStoreVersionLocalizations []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              int      `url:"limit,omitempty"`
	Include                            []string `url:"include,omitempty"`
	FilterPreviewType                  []string `url:"filter[previewType],omitempty"`
	Cursor                             string   `url:"cursor,omitempty"`
}

// ListLocalizationsForAppStoreVersion gets a list of localized, version-level information about an app, for all locales.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_version_localizations_for_an_app_store_version
func (s *AppsService) ListLocalizationsForAppStoreVersion(ctx context.Context, id string, params *ListLocalizationsForAppStoreVersionQuery) (*AppStoreVersionLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionLocalizations", id)
	res := new(AppStoreVersionLocalizationsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppStoreVersionLocalization reads localized version-level information.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_localization_information
func (s *AppsService) GetAppStoreVersionLocalization(ctx context.Context, id string, params *GetAppStoreVersionLocalizationQuery) (*AppStoreVersionLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateAppStoreVersionLocalization adds localized version-level information for a new locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_localization
func (s *AppsService) CreateAppStoreVersionLocalization(ctx context.Context, attributes AppStoreVersionLocalizationCreateRequestAttributes, appStoreVersionID string) (*AppStoreVersionLocalizationResponse, *Response, error) {
	req := appStoreVersionLocalizationCreateRequest{
		Attributes: attributes,
		Relationships: appStoreVersionLocalizationCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreVersionLocalizations",
	}
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.client.post(ctx, "appStoreVersionLocalizations", newRequestBody(req), res)

	return res, resp, err
}

// UpdateAppStoreVersionLocalization modifies localized version-level information for a particular language.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_version_localization
func (s *AppsService) UpdateAppStoreVersionLocalization(ctx context.Context, id string, attributes *AppStoreVersionLocalizationUpdateRequestAttributes) (*AppStoreVersionLocalizationResponse, *Response, error) {
	req := appStoreVersionLocalizationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "appStoreVersionLocalizations",
	}
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteAppStoreVersionLocalization deletes a language from your version metadata.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_localization
func (s *AppsService) DeleteAppStoreVersionLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListAppScreenshotSetsForAppStoreVersionLocalization lists all screenshot sets for a specific localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_screenshot_sets_for_an_app_store_version_localization
func (s *AppsService) ListAppScreenshotSetsForAppStoreVersionLocalization(ctx context.Context, id string, params *ListAppScreenshotSetsForAppStoreVersionLocalizationQuery) (*AppScreenshotSetsResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appScreenshotSets", id)
	res := new(AppScreenshotSetsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListAppPreviewSetsForAppStoreVersionLocalization lists all app preview sets for a specific localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_preview_sets_for_an_app_store_version_localization
func (s *AppsService) ListAppPreviewSetsForAppStoreVersionLocalization(ctx context.Context, id string, params *ListAppPreviewSetsForAppStoreVersionLocalizationQuery) (*AppPreviewSetsResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appPreviewSets", id)
	res := new(AppPreviewSetsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in AppStoreVersionLocalizationResponseIncluded.
func (i *AppStoreVersionLocalizationResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// AppScreenshotSet returns the AppScreenshotSet stored within, if one is present.
func (i *AppStoreVersionLocalizationResponseIncluded) AppScreenshotSet() *AppScreenshotSet {
	return extractIncludedAppScreenshotSet(i.inner)
}

// AppPreviewSet returns the AppPreviewSet stored within, if one is present.
func (i *AppStoreVersionLocalizationResponseIncluded) AppPreviewSet() *AppPreviewSet {
	return extractIncludedAppPreviewSet(i.inner)
}
