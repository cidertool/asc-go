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
}

// AppInfoResponse defines model for AppInfoResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinforesponse
type AppInfoResponse struct {
	Data     AppInfo        `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfosresponse
type AppInfosResponse struct {
	Data     []AppInfo          `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppInfoUpdateRequest defines model for AppInfoUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest
type AppInfoUpdateRequest struct {
	ID            string                             `json:"id"`
	Relationships *AppInfoUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppInfoUpdateRequestRelationships are relationships for AppInfoUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest/data/relationships
type AppInfoUpdateRequestRelationships struct {
	PrimaryCategory         *RelationshipDeclaration `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   *RelationshipDeclaration `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   *RelationshipDeclaration `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       *RelationshipDeclaration `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *RelationshipDeclaration `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *RelationshipDeclaration `json:"secondarySubcategoryTwo,omitempty"`
}

// GetAppInfoQuery are query options for GetAppInfo
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
type GetAppInfoQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        []string `url:"fields[appCategories],omitempty"`
	Include                    []string `url:"include,omitempty"`
	LimitAppInfoLocalizations  int      `url:"limit[appInfoLocalizations],omitempty"`
}

// ListAppInfosForAppQuery are query options for ListAppInfosForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
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
func (s *AppsService) UpdateAppInfo(ctx context.Context, id string, body *AppInfoUpdateRequest) (*AppInfoResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}
