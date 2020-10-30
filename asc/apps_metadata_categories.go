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

// AppCategory defines model for AppCategory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcategory
type AppCategory struct {
	Attributes    *AppCategoryAttributes    `json:"attributes,omitempty"`
	ID            string                    `json:"id"`
	Links         ResourceLinks             `json:"links"`
	Relationships *AppCategoryRelationships `json:"relationships,omitempty"`
	Type          string                    `json:"type"`
}

// AppCategoryAttributes defines model for AppCategory.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcategory/attributes
type AppCategoryAttributes struct {
	Platforms []Platform `json:"platforms,omitempty"`
}

// AppCategoryRelationships defines model for AppCategory.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcategory/relationships
type AppCategoryRelationships struct {
	Parent        *Relationship      `json:"parent,omitempty"`
	Subcategories *PagedRelationship `json:"subcategories,omitempty"`
}

// AppCategoriesResponse defines model for AppCategoriesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcategoriesresponse
type AppCategoriesResponse struct {
	Data     []AppCategory                 `json:"data"`
	Included []AppCategoryResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks            `json:"links"`
	Meta     *PagingInformation            `json:"meta,omitempty"`
}

// AppCategoryResponse defines model for AppCategoryResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcategoryresponse
type AppCategoryResponse struct {
	Data     AppCategory                   `json:"data"`
	Included []AppCategoryResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                 `json:"links"`
}

// AppCategoryResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a AppCategoryResponse or AppCategoriesResponse.
type AppCategoryResponseIncluded included

// ListAppCategoriesQuery are query options for ListAppCategories
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_categories
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_subcategories_for_an_app_category
type ListSubcategoriesForAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// GetAppCategoryQuery are query options for GetAppCategory
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_category_information
type GetAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Include             []string `url:"include,omitempty"`
	LimitSubcategories  []string `url:"limit[subcategories],omitempty"`
}

// GetAppCategoryForAppInfoQuery are query options for GetAppCategoryForAppInfo
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_parent_information_of_an_app_category
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_category_information_of_an_app_info
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_category_information_of_an_app_info
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_subcategory_one_information_of_an_app_info
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_primary_subcategory_two_information_of_an_app_info
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_subcategory_one_information_of_an_app_info
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_secondary_subcategory_two_information_of_an_app_info
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

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in AppCategoryResponseIncluded.
func (i *AppCategoryResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// AppCategory returns the AppCategory stored within, if one is present.
func (i *AppCategoryResponseIncluded) AppCategory() *AppCategory {
	return extractIncludedAppCategory(i.inner)
}
