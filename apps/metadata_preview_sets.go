package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// AppPreviewSet defines model for AppPreviewSet.
type AppPreviewSet struct {
	Attributes *struct {
		PreviewType *PreviewType `json:"previewType,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		AppPreviews *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"appPreviews,omitempty"`
		AppStoreVersionLocalization *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionLocalization,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreviewSetCreateRequest defines model for AppPreviewSetCreateRequest.
type AppPreviewSetCreateRequest struct {
	Data struct {
		Attributes struct {
			PreviewType PreviewType `json:"previewType"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreVersionLocalization struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"appStoreVersionLocalization"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreviewSetResponse defines model for AppPreviewSetResponse.
type AppPreviewSetResponse struct {
	Data     AppPreviewSet          `json:"data"`
	Included *[]AppPreview          `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
}

// AppPreviewSetsResponse defines model for AppPreviewSetsResponse.
type AppPreviewSetsResponse struct {
	Data     []AppPreviewSet             `json:"data"`
	Included *[]AppPreview               `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// AppPreviewSetAppPreviewsLinkagesRequest defines model for AppPreviewSetAppPreviewsLinkagesRequest.
type AppPreviewSetAppPreviewsLinkagesRequest struct {
	Data []internal.RelationshipsData `json:"data"`
}

// AppPreviewSetAppPreviewsLinkagesResponse defines model for AppPreviewSetAppPreviewsLinkagesResponse.
type AppPreviewSetAppPreviewsLinkagesResponse struct {
	Data  []internal.RelationshipsData `json:"data"`
	Links internal.PagedDocumentLinks  `json:"links"`
	Meta  *internal.PagingInformation  `json:"meta,omitempty"`
}

type GetAppPreviewSetQuery struct {
	FieldsAppPreviews    *[]string `url:"fields[appPreviews],omitempty"`
	FieldsAppPreviewSets *[]string `url:"fields[appPreviewSets],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	LimitAppPreviews     *int      `url:"limit[appPreviews],omitempty"`
}

type ListAppPreviewsForSetQuery struct {
	FieldsAppPreviewSets *[]string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppPreviews    *[]string `url:"fields[appPreviews],omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	Include              *[]string `url:"include,omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

type ListAppPreviewIDsForSetQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// GetAppPreviewSet gets an app preview set including its display target, language, and the preview it contains.
func (s *Service) GetAppPreviewSet(id string, params *GetAppPreviewSetQuery) (*AppPreviewSetResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s", id)
	res := new(AppPreviewSetResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppPreviewSet adds a new preview set to an App Store version localization for a specific preview type and display size.
func (s *Service) CreateAppPreviewSet(id string, body *AppPreviewSetCreateRequest) (*AppPreviewSetResponse, *internal.Response, error) {
	res := new(AppPreviewSetResponse)
	resp, err := s.Post("appPreviewSets", body, res)
	return res, resp, err
}

// DeleteAppPreviewSet deletes an app preview set and all of its previews.
func (s *Service) DeleteAppPreviewSet(id string) (*internal.Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s", id)
	return s.Delete(url, nil)
}

// ListAppPreviewsForSet lists all ordered previews in a preview set.
func (s *Service) ListAppPreviewsForSet(id string, params *ListAppPreviewsForSetQuery) (*AppPreviewsResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/appPreviews", id)
	res := new(AppPreviewsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppPreviewIDsForSet gets the ordered preview IDs in a preview set.
func (s *Service) ListAppPreviewIDsForSet(id string, params *ListAppPreviewIDsForSetQuery) (*AppPreviewSetAppPreviewsLinkagesResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/relationships/appPreviews", id)
	res := new(AppPreviewSetAppPreviewsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ReplaceAppPreviewsForSet changes the order of the previews in a preview set.
func (s *Service) ReplaceAppPreviewsForSet(id string, body *AppPreviewSetAppPreviewsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("appPreviewSets/%s/relationships/appPreviews", id)
	return s.Patch(url, body, nil)
}
