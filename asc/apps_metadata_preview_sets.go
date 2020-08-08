package asc

import (
	"context"
	"fmt"
)

// AppPreviewSet defines model for AppPreviewSet.
type AppPreviewSet struct {
	Attributes *struct {
		PreviewType *PreviewType `json:"previewType,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppPreviews *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appPreviews,omitempty"`
		AppStoreVersionLocalization *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionLocalization,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreviewSetCreateRequest defines model for AppPreviewSetCreateRequest.
type AppPreviewSetCreateRequest struct {
	Attributes    AppPreviewSetCreateRequestAttributes    `json:"attributes"`
	Relationships AppPreviewSetCreateRequestRelationships `json:"relationships"`
	Type          string                                  `json:"type"`
}

// AppPreviewSetCreateRequestAttributes are attributes for AppPreviewSetCreateRequest
type AppPreviewSetCreateRequestAttributes struct {
	PreviewType PreviewType `json:"previewType"`
}

// AppPreviewSetCreateRequestRelationships are relationships for AppPreviewSetCreateRequest
type AppPreviewSetCreateRequestRelationships struct {
	AppStoreVersionLocalization struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersionLocalization"`
}

// AppPreviewSetResponse defines model for AppPreviewSetResponse.
type AppPreviewSetResponse struct {
	Data     AppPreviewSet `json:"data"`
	Included *[]AppPreview `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// AppPreviewSetsResponse defines model for AppPreviewSetsResponse.
type AppPreviewSetsResponse struct {
	Data     []AppPreviewSet    `json:"data"`
	Included *[]AppPreview      `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppPreviewSetAppPreviewsLinkagesResponse defines model for AppPreviewSetAppPreviewsLinkagesResponse.
type AppPreviewSetAppPreviewsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// GetAppPreviewSetQuery are query options for GetAppPreviewSet
type GetAppPreviewSetQuery struct {
	FieldsAppPreviews    []string `url:"fields[appPreviews],omitempty"`
	FieldsAppPreviewSets []string `url:"fields[appPreviewSets],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitAppPreviews     int      `url:"limit[appPreviews],omitempty"`
}

// ListAppPreviewsForSetQuery are query options for ListAppPreviewsForSet
type ListAppPreviewsForSetQuery struct {
	FieldsAppPreviewSets []string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppPreviews    []string `url:"fields[appPreviews],omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Include              []string `url:"include,omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// ListAppPreviewIDsForSetQuery are query options for ListAppPreviewIDsForSet
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
func (s *AppsService) CreateAppPreviewSet(ctx context.Context, body *AppPreviewSetCreateRequest) (*AppPreviewSetResponse, *Response, error) {
	res := new(AppPreviewSetResponse)
	resp, err := s.client.post(ctx, "appPreviewSets", body, res)
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
func (s *AppsService) ReplaceAppPreviewsForSet(ctx context.Context, id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/relationships/appPreviews", id)
	return s.client.patch(ctx, url, linkages, nil)
}
