package asc

import (
	"context"
	"fmt"
)

// AppCategory defines model for AppCategory.
type AppCategory struct {
	Attributes *struct {
		Platforms *[]Platform `json:"platforms,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Parent *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"parent,omitempty"`
		Subcategories *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"subcategories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppCategoriesResponse defines model for AppCategoriesResponse.
type AppCategoriesResponse struct {
	Data     []AppCategory      `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppCategoryResponse defines model for AppCategoryResponse.
type AppCategoryResponse struct {
	Data     AppCategory    `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// ListAppCategoriesQuery are query options for ListAppCategories
type ListAppCategoriesQuery struct {
	ExistsParent        []string `url:"exists[parent],omitempty"`
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	FilterPlatforms     []string `url:"filter[platforms],omitempty"`
	Include             []string `url:"include,omitempty"`
	Limit               int      `url:"limit,omitempty"`
	LimitSubcategories  []string `url:"limit[subcategories],omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// ListSubcategoriesForAppCategoryQuery are query options for ListSubcategoriesForAppCategory
type ListSubcategoriesForAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// GetAppCategoryQuery are query options for GetAppCategory
type GetAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Include             []string `url:"include,omitempty"`
	LimitSubcategories  []string `url:"limit[subcategories],omitempty"`
}

// GetAppCategoryForAppInfoQuery are query options for GetAppCategoryForAppInfo
type GetAppCategoryForAppInfoQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
}

// ListAppCategories lists all categories on the App Store, including the category and subcategory hierarchy.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_categories
func (s *AppsService) ListAppCategories(ctx context.Context, params *ListAppCategoriesQuery) (*AppCategoriesResponse, *Response, error) {
	res := new(AppCategoriesResponse)
	resp, err := s.client.get(ctx, "appCategories", params, res)
	return res, resp, err
}

// ListSubcategoriesForAppCategory lists all App Store subcategories that belong to a specific category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_subcategories_for_an_app_category
func (s *AppsService) ListSubcategoriesForAppCategory(ctx context.Context, id string, params *ListSubcategoriesForAppCategoryQuery) (*AppCategoriesResponse, *Response, error) {
	url := fmt.Sprintf("appCategories/%s/subcategories", id)
	res := new(AppCategoriesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppCategory gets a specific app category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_category_information
func (s *AppsService) GetAppCategory(ctx context.Context, id string, params *GetAppCategoryQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appCategories/%s", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetParentCategoryForAppCategory gets the App Store category to which a specific subcategory belongs.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_parent_information_of_an_app_category
func (s *AppsService) GetParentCategoryForAppCategory(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appCategories/%s/parent", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetPrimaryCategoryForAppInfo gets an app’s primary App Store category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_category_information_of_an_app_info
func (s *AppsService) GetPrimaryCategoryForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/primaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetSecondaryCategoryForAppInfo gets an app’s secondary App Store category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_category_information_of_an_app_info
func (s *AppsService) GetSecondaryCategoryForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s primary category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_subcategory_one_information_of_an_app_info
func (s *AppsService) GetPrimarySubcategoryOneForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s primary category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_subcategory_two_information_of_an_app_info
func (s *AppsService) GetPrimarySubcategoryTwoForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s secondary category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_subcategory_one_information_of_an_app_info
func (s *AppsService) GetSecondarySubcategoryOneForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s secondary category.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_subcategory_two_information_of_an_app_info
func (s *AppsService) GetSecondarySubcategoryTwoForAppInfo(ctx context.Context, id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
