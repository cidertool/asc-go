package asc

import (
	"context"
	"fmt"
)

// AppStoreReviewDetail defines model for AppStoreReviewDetail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetail
type AppStoreReviewDetail struct {
	Attributes    *AppStoreReviewDetailAttributes    `json:"attributes,omitempty"`
	ID            string                             `json:"id"`
	Links         ResourceLinks                      `json:"links"`
	Relationships *AppStoreReviewDetailRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppStoreReviewDetailAttributes defines model for AppStoreReviewDetail.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetail/attributes
type AppStoreReviewDetailAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// AppStoreReviewDetailRelationships defines model for AppStoreReviewDetail.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetail/relationships
type AppStoreReviewDetailRelationships struct {
	AppStoreReviewAttachments *PagedRelationship `json:"appStoreReviewAttachments,omitempty"`
	AppStoreVersion           *Relationship      `json:"appStoreVersion,omitempty"`
}

// AppStoreReviewDetailCreateRequest defines model for AppStoreReviewDetailCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailcreaterequest
type appStoreReviewDetailCreateRequest struct {
	Attributes    *AppStoreReviewDetailCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships appStoreReviewDetailCreateRequestRelationships `json:"relationships"`
	Type          string                                         `json:"type"`
}

// AppStoreReviewDetailCreateRequestAttributes are attributes for AppStoreReviewDetailCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailcreaterequest/data/attributes
type AppStoreReviewDetailCreateRequestAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// AppStoreReviewDetailCreateRequestRelationships are relationships for AppStoreReviewDetailCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailcreaterequest/data/relationships
type appStoreReviewDetailCreateRequestRelationships struct {
	AppStoreVersion RelationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreReviewDetailUpdateRequest defines model for AppStoreReviewDetailUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailupdaterequest
type AppStoreReviewDetailUpdateRequest struct {
	Attributes *AppStoreReviewDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// AppStoreReviewDetailUpdateRequestAttributes are attributes for AppStoreReviewDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailupdaterequest/data/attributes
type AppStoreReviewDetailUpdateRequestAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// AppStoreReviewDetailResponse defines model for AppStoreReviewDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailresponse
type AppStoreReviewDetailResponse struct {
	Data     AppStoreReviewDetail       `json:"data"`
	Included []AppStoreReviewAttachment `json:"included,omitempty"`
	Links    DocumentLinks              `json:"links"`
}

// GetReviewDetailQuery are query options for GetReviewDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_detail_information
type GetReviewDetailQuery struct {
	FieldsAppStoreReviewDetails     []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAppStoreReviewAttachments []string `url:"fields[appStoreReviewAttachments],omitempty"`
	Include                         []string `url:"include,omitempty"`
	LimitAppStoreReviewAttachments  int      `url:"limit[appStoreReviewAttachments],omitempty"`
}

// GetAppStoreReviewDetailsForAppStoreVersionQuery are query options for GetAppStoreReviewDetailsForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_review_details_resource_information_of_an_app_store_version
type GetAppStoreReviewDetailsForAppStoreVersionQuery struct {
	FieldsAppStoreReviewAttachments []string `url:"fields[appStoreReviewAttachments],omitempty"`
	FieldsAppStoreReviewDetails     []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAppStoreVersions          []string `url:"fields[appStoreVersions],omitempty"`
	Include                         []string `url:"include,omitempty"`
}

// CreateReviewDetail adds App Store review details to an App Store version, including contact and demo account information.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_review_detail
func (s *SubmissionService) CreateReviewDetail(ctx context.Context, attributes *AppStoreReviewDetailCreateRequestAttributes, appStoreVersionID string) (*AppStoreReviewDetailResponse, *Response, error) {
	req := appStoreReviewDetailCreateRequest{
		Attributes: attributes,
		Relationships: appStoreReviewDetailCreateRequestRelationships{
			AppStoreVersion: RelationshipDeclaration{
				Data: &RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreReviewDetails",
	}
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.client.post(ctx, "appStoreReviewDetails", req, res)
	return res, resp, err
}

// GetReviewDetail gets App Review details you provided, including contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_detail_information
func (s *SubmissionService) GetReviewDetail(ctx context.Context, id string, params *GetReviewDetailQuery) (*AppStoreReviewDetailResponse, *Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetReviewDetailsForAppStoreVersion gets the details you provide to App Review so they can test your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_review_details_resource_information_of_an_app_store_version
func (s *SubmissionService) GetReviewDetailsForAppStoreVersion(ctx context.Context, id string, params *GetAppStoreReviewDetailsForAppStoreVersionQuery) (*AppStoreReviewDetailResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreReviewDetail", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateReviewDetail update the app store review details, including the contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_review_detail
func (s *SubmissionService) UpdateReviewDetail(ctx context.Context, id string, body AppStoreReviewDetailUpdateRequest) (*AppStoreReviewDetailResponse, *Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s", id)
	res := new(AppStoreReviewDetailResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}
