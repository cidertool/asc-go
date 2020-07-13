package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppCategory defines model for AppCategory.
type AppCategory struct {
	Attributes *struct {
		Platforms *[]Platform `json:"platforms,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		Parent *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"parent,omitempty"`
		Subcategories *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"subcategories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppCategoriesResponse defines model for AppCategoriesResponse.
type AppCategoriesResponse struct {
	Data     []AppCategory             `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// AppCategoryResponse defines model for AppCategoryResponse.
type AppCategoryResponse struct {
	Data     AppCategory          `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

type ListAppCategoriesOptions struct {
	Exists *struct {
		Parent *[]string `url:"parent,omitempty"`
	} `url:"exists,omitempty"`
	Fields *struct {
		AppCategories *[]string `url:"appCategories,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		Platforms *[]string `url:"platforms,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		Subcategories *[]string `url:"subcategories,omitempty"`
	} `url:"limit,omitempty"`
}

type ListSubcategoriesForAppCategoryOptions struct {
	Fields *struct {
		AppCategories *[]string `url:"appCategories,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type GetAppCategoryOptions struct {
	Fields *struct {
		AppCategories *[]string `url:"appCategories,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		Subcategories *[]string `url:"subcategories,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAppCategoryForAppInfoOptions struct {
	Fields *struct {
		AppCategories *[]string `url:"appCategories,omitempty"`
	} `url:"fields,omitempty"`
}

// ListAppCategories lists all categories on the App Store, including the category and subcategory hierarchy.
func (s *Service) ListAppCategories(params *ListAppCategoriesOptions) (*AppCategoriesResponse, *common.Response, error) {
	res := new(AppCategoriesResponse)
	resp, err := s.Get("appCategories", params, res)
	return res, resp, err
}

// ListSubcategoriesForAppCategory lists all App Store subcategories that belong to a specific category.
func (s *Service) ListSubcategoriesForAppCategory(id string, params *ListSubcategoriesForAppCategoryOptions) (*AppCategoriesResponse, *common.Response, error) {
	url := fmt.Sprintf("appCategories/%s/subcategories", id)
	res := new(AppCategoriesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppCategory gets a specific app category.
func (s *Service) GetAppCategory(id string, params *GetAppCategoryOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appCategories/%s", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetParentCategoryForAppCategory gets the App Store category to which a specific subcategory belongs.
func (s *Service) GetParentCategoryForAppCategory(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appCategories/%s/parent", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetPrimaryCategoryForAppInfo gets an app’s primary App Store category.
func (s *Service) GetPrimaryCategoryForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetSecondaryCategoryForAppInfo gets an app’s secondary App Store category.
func (s *Service) GetSecondaryCategoryForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s primary category.
func (s *Service) GetPrimarySubcategoryOneForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s primary category.
func (s *Service) GetPrimarySubcategoryTwoForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s secondary category.
func (s *Service) GetSecondarySubcategoryOneForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s secondary category.
func (s *Service) GetSecondarySubcategoryTwoForAppInfo(id string, params *GetAppCategoryForAppInfoOptions) (*AppCategoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
