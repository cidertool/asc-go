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

// AppPreviewSet defines model for AppPreviewSet.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewset
type AppPreviewSet struct {
	Attributes    *AppPreviewSetAttributes    `json:"attributes,omitempty"`
	ID            string                      `json:"id"`
	Links         ResourceLinks               `json:"links"`
	Relationships *AppPreviewSetRelationships `json:"relationships,omitempty"`
	Type          string                      `json:"type"`
}

// AppPreviewSetAttributes defines model for AppPreviewSet.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewset/attributes
type AppPreviewSetAttributes struct {
	PreviewType *PreviewType `json:"previewType,omitempty"`
}

// AppPreviewSetRelationships defines model for AppPreviewSet.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewset/relationships
type AppPreviewSetRelationships struct {
	AppPreviews                 *PagedRelationship `json:"appPreviews,omitempty"`
	AppStoreVersionLocalization *Relationship      `json:"appStoreVersionLocalization,omitempty"`
}

// appPreviewSetCreateRequest defines model for AppPreviewSetCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetcreaterequest/data
type appPreviewSetCreateRequest struct {
	Attributes    appPreviewSetCreateRequestAttributes    `json:"attributes"`
	Relationships appPreviewSetCreateRequestRelationships `json:"relationships"`
	Type          string                                  `json:"type"`
}

// appPreviewSetCreateRequestAttributes are attributes for AppPreviewSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetcreaterequest/data/attributes
type appPreviewSetCreateRequestAttributes struct {
	PreviewType PreviewType `json:"previewType"`
}

// appPreviewSetCreateRequestRelationships are relationships for AppPreviewSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetcreaterequest/data/relationships
type appPreviewSetCreateRequestRelationships struct {
	AppStoreVersionLocalization relationshipDeclaration `json:"appStoreVersionLocalization"`
}

// AppPreviewSetResponse defines model for AppPreviewSetResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetresponse
type AppPreviewSetResponse struct {
	Data     AppPreviewSet `json:"data"`
	Included []AppPreview  `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// AppPreviewSetsResponse defines model for AppPreviewSetsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetsresponse
type AppPreviewSetsResponse struct {
	Data     []AppPreviewSet    `json:"data"`
	Included []AppPreview       `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppPreviewSetAppPreviewsLinkagesResponse defines model for AppPreviewSetAppPreviewsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsetapppreviewslinkagesresponse
type AppPreviewSetAppPreviewsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GetAppPreviewSetQuery are query options for GetAppPreviewSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_preview_set_information
type GetAppPreviewSetQuery struct {
	FieldsAppPreviews    []string `url:"fields[appPreviews],omitempty"`
	FieldsAppPreviewSets []string `url:"fields[appPreviewSets],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitAppPreviews     int      `url:"limit[appPreviews],omitempty"`
}

// ListAppPreviewsForSetQuery are query options for ListAppPreviewsForSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_previews_for_an_app_preview_set
type ListAppPreviewsForSetQuery struct {
	FieldsAppPreviewSets []string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppPreviews    []string `url:"fields[appPreviews],omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Include              []string `url:"include,omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// ListAppPreviewIDsForSetQuery are query options for ListAppPreviewIDsForSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_preview_ids_for_an_app_preview_set
type ListAppPreviewIDsForSetQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetAppPreviewSet gets an app preview set including its display target, language, and the preview it contains.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_preview_set_information
func (s *AppsService) GetAppPreviewSet(ctx context.Context, id string, params *GetAppPreviewSetQuery) (*AppPreviewSetResponse, *Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s", id)
	res := new(AppPreviewSetResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateAppPreviewSet adds a new preview set to an App Store version localization for a specific preview type and display size.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_preview_set
func (s *AppsService) CreateAppPreviewSet(ctx context.Context, previewType PreviewType, appStoreVersionLocalizationID string) (*AppPreviewSetResponse, *Response, error) {
	req := appPreviewSetCreateRequest{
		Attributes: appPreviewSetCreateRequestAttributes{
			PreviewType: previewType,
		},
		Relationships: appPreviewSetCreateRequestRelationships{
			AppStoreVersionLocalization: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionLocalizationID,
					Type: "appStoreVersionLocalizations",
				},
			},
		},
		Type: "appPreviewSets",
	}
	res := new(AppPreviewSetResponse)
	resp, err := s.client.post(ctx, "appPreviewSets", newRequestBody(req), res)

	return res, resp, err
}

// DeleteAppPreviewSet deletes an app preview set and all of its previews.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_preview_set
func (s *AppsService) DeleteAppPreviewSet(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListAppPreviewsForSet lists all ordered previews in a preview set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_previews_for_an_app_preview_set
func (s *AppsService) ListAppPreviewsForSet(ctx context.Context, id string, params *ListAppPreviewsForSetQuery) (*AppPreviewsResponse, *Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/appPreviews", id)
	res := new(AppPreviewsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListAppPreviewIDsForSet gets the ordered preview IDs in a preview set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_preview_ids_for_an_app_preview_set
func (s *AppsService) ListAppPreviewIDsForSet(ctx context.Context, id string, params *ListAppPreviewIDsForSetQuery) (*AppPreviewSetAppPreviewsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/relationships/appPreviews", id)
	res := new(AppPreviewSetAppPreviewsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ReplaceAppPreviewsForSet changes the order of the previews in a preview set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/replace_all_app_previews_for_an_app_preview_set
func (s *AppsService) ReplaceAppPreviewsForSet(ctx context.Context, id string, appPreviewIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appPreviewIDs, "appPreviews")
	url := fmt.Sprintf("appPreviewSets/%s/relationships/appPreviews", id)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}
