package asc

import (
	"fmt"
)

// AppInfo defines model for AppInfo.
type AppInfo struct {
	Attributes *struct {
		AppStoreAgeRating *AppStoreAgeRating    `json:"appStoreAgeRating,omitempty"`
		AppStoreState     *AppStoreVersionState `json:"appStoreState,omitempty"`
		BrazilAgeRating   *BrazilAgeRating      `json:"brazilAgeRating,omitempty"`
		KidsAgeBand       *KidsAgeBand          `json:"kidsAgeBand,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppInfoLocalizations *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appInfoLocalizations,omitempty"`
		PrimaryCategory *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"primaryCategory,omitempty"`
		PrimarySubcategoryOne *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryOne,omitempty"`
		PrimarySubcategoryTwo *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryTwo,omitempty"`
		SecondaryCategory *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondaryCategory,omitempty"`
		SecondarySubcategoryOne *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryOne,omitempty"`
		SecondarySubcategoryTwo *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryTwo,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppInfoResponse defines model for AppInfoResponse.
type AppInfoResponse struct {
	Data     AppInfo        `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
type AppInfosResponse struct {
	Data     []AppInfo          `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppInfoUpdateRequest defines model for AppInfoUpdateRequest.
type AppInfoUpdateRequest struct {
	ID            string                             `json:"id"`
	Relationships *AppInfoUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppInfoUpdateRequestRelationships are relationships for AppInfoUpdateRequest
type AppInfoUpdateRequestRelationships struct {
	PrimaryCategory *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"secondarySubcategoryTwo,omitempty"`
}

// GetAppInfoQuery are query options for GetAppInfo
type GetAppInfoQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        []string `url:"fields[appCategories],omitempty"`
	Include                    []string `url:"include,omitempty"`
	LimitAppInfoLocalizations  int      `url:"limit[appInfoLocalizations],omitempty"`
}

// ListAppInfosForAppQuery are query options for ListAppInfosForApp
type ListAppInfosForAppQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        []string `url:"fields[appCategories],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Include                    []string `url:"include,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetAppInfo reads App Store information including your App Store state, age ratings, Brazil age rating, and kids' age band.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
func (s *AppsService) GetAppInfo(id string, params *GetAppInfoQuery) (*AppInfoResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListAppInfosForApp gets information about an app that is currently live on App Store, or that goes live with the next version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
func (s *AppsService) ListAppInfosForApp(id string, params *ListAppInfosForAppQuery) (*AppInfosResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/appInfos", id)
	res := new(AppInfosResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// UpdateAppInfo updates the App Store categories and sub-categories for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_info
func (s *AppsService) UpdateAppInfo(id string, body *AppInfoUpdateRequest) (*AppInfoResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}
