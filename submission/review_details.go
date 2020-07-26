package submission

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// AppStoreReviewDetail defines model for AppStoreReviewDetail.
type AppStoreReviewDetail struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreReviewAttachments *struct {
			Data  *[]types.RelationshipsData `json:"data,omitempty"`
			Links *types.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *types.PagingInformation   `json:"meta,omitempty"`
		} `json:"appStoreReviewAttachments,omitempty"`
		AppStoreVersion *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreReviewDetailCreateRequest defines model for AppStoreReviewDetailCreateRequest.
type AppStoreReviewDetailCreateRequest struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		AppStoreVersion struct {
			Data types.RelationshipsData `json:"data"`
		} `json:"appStoreVersion"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// AppStoreReviewDetailUpdateRequest defines model for AppStoreReviewDetailUpdateRequest.
type AppStoreReviewDetailUpdateRequest struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	ID   string `json:"id"`
	Type string `json:"type"`
}

// AppStoreReviewDetailResponse defines model for AppStoreReviewDetailResponse.
type AppStoreReviewDetailResponse struct {
	Data     AppStoreReviewDetail        `json:"data"`
	Included *[]AppStoreReviewAttachment `json:"included,omitempty"`
	Links    types.DocumentLinks         `json:"links"`
}

// GetReviewDetailQuery are query options for GetReviewDetail
type GetReviewDetailQuery struct {
	FieldsAppStoreReviewDetails     *[]string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAppStoreReviewAttachments *[]string `url:"fields[appStoreReviewAttachments],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	LimitAppStoreReviewAttachments  *int      `url:"limit[appStoreReviewAttachments],omitempty"`
}

// GetAppStoreReviewDetailsForAppStoreVersionQuery are query options for GetAppStoreReviewDetailsForAppStoreVersion
type GetAppStoreReviewDetailsForAppStoreVersionQuery struct {
	FieldsAppStoreReviewAttachments *[]string `url:"fields[appStoreReviewAttachments],omitempty"`
	FieldsAppStoreReviewDetails     *[]string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAppStoreVersions          *[]string `url:"fields[appStoreVersions],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
}

// CreateReviewDetail adds App Store review details to an App Store version, including contact and demo account information.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_review_detail
func (s *Service) CreateReviewDetail(body *AppStoreReviewDetailCreateRequest) (*AppStoreReviewDetailResponse, *services.Response, error) {
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.Post("appStoreReviewDetails", body, res)
	return res, resp, err
}

// GetReviewDetail gets App Review details you provided, including contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_detail_information
func (s *Service) GetReviewDetail(id string, params *GetReviewDetailQuery) (*AppStoreReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetReviewDetailsForAppStoreVersion gets the details you provide to App Review so they can test your app.
func (s *Service) GetReviewDetailsForAppStoreVersion(id string, params *GetAppStoreReviewDetailsForAppStoreVersionQuery) (*AppStoreReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreReviewDetail", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateReviewDetail update the app store review details, including the contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_review_detail
func (s *Service) UpdateReviewDetail(id string, body *AppStoreReviewDetailUpdateRequest) (*AppStoreReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
