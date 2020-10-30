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

// BetaAppReviewDetail defines model for BetaAppReviewDetail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetail
type BetaAppReviewDetail struct {
	Attributes    *BetaAppReviewDetailAttributes    `json:"attributes,omitempty"`
	ID            string                            `json:"id"`
	Links         ResourceLinks                     `json:"links"`
	Relationships *BetaAppReviewDetailRelationships `json:"relationships,omitempty"`
	Type          string                            `json:"type"`
}

// BetaAppReviewDetailAttributes defines model for BetaAppReviewDetail.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetail/attributes
type BetaAppReviewDetailAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// BetaAppReviewDetailRelationships defines model for BetaAppReviewDetail.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetail/relationships
type BetaAppReviewDetailRelationships struct {
	App *Relationship `json:"app,omitempty"`
}

// BetaAppReviewDetailUpdateRequest defines model for BetaAppReviewDetailUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetailupdaterequest/data
type betaAppReviewDetailUpdateRequest struct {
	Attributes *BetaAppReviewDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// BetaAppReviewDetailUpdateRequestAttributes are attributes for BetaAppReviewDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetailupdaterequest/data/attributes
type BetaAppReviewDetailUpdateRequestAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// BetaAppReviewDetailResponse defines model for BetaAppReviewDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetailresponse
type BetaAppReviewDetailResponse struct {
	Data     BetaAppReviewDetail `json:"data"`
	Included []App               `json:"included,omitempty"`
	Links    DocumentLinks       `json:"links"`
}

// BetaAppReviewDetailsResponse defines model for BetaAppReviewDetailsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaappreviewdetailsresponse
type BetaAppReviewDetailsResponse struct {
	Data     []BetaAppReviewDetail `json:"data"`
	Included []App                 `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// ListBetaAppReviewDetailsQuery defines model for ListBetaAppReviewDetails
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_review_details
type ListBetaAppReviewDetailsQuery struct {
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails []string `url:"fields[betaAppReviewDetails],omitempty"`
	FilterApp                  []string `url:"filter[app],omitempty"`
	Include                    []string `url:"include,omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetBetaAppReviewDetailQuery defines model for GetBetaAppReviewDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_review_detail_information
type GetBetaAppReviewDetailQuery struct {
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails []string `url:"fields[betaAppReviewDetails],omitempty"`
	Include                    []string `url:"include,omitempty"`
}

// GetAppForBetaAppReviewDetailQuery defines model for GetAppForBetaAppReviewDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_app_review_detail
type GetAppForBetaAppReviewDetailQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// GetBetaAppReviewDetailsForAppQuery defines model for GetBetaAppReviewDetailsForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_app_review_details_resource_of_an_app
type GetBetaAppReviewDetailsForAppQuery struct {
	FieldsBetaAppReviewDetails []string `url:"fields[betaAppReviewDetails],omitempty"`
}

// ListBetaAppReviewDetails finds and lists beta app review details for all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_review_details
func (s *TestflightService) ListBetaAppReviewDetails(ctx context.Context, params *ListBetaAppReviewDetailsQuery) (*BetaAppReviewDetailsResponse, *Response, error) {
	res := new(BetaAppReviewDetailsResponse)
	resp, err := s.client.get(ctx, "betaAppReviewDetails", params, res)

	return res, resp, err
}

// GetBetaAppReviewDetail gets beta app review details for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_review_detail_information
func (s *TestflightService) GetBetaAppReviewDetail(ctx context.Context, id string, params *GetBetaAppReviewDetailQuery) (*BetaAppReviewDetailResponse, *Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppForBetaAppReviewDetail gets the app information for a specific beta app review details resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_app_review_detail
func (s *TestflightService) GetAppForBetaAppReviewDetail(ctx context.Context, id string, params *GetAppForBetaAppReviewDetailQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetBetaAppReviewDetailsForApp gets the beta app review details for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_app_review_details_resource_of_an_app
func (s *TestflightService) GetBetaAppReviewDetailsForApp(ctx context.Context, id string, params *GetBetaAppReviewDetailsForAppQuery) (*BetaAppReviewDetailResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppReviewDetail", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateBetaAppReviewDetail updates the details for a specific app's beta app review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_app_review_detail
func (s *TestflightService) UpdateBetaAppReviewDetail(ctx context.Context, id string, attributes *BetaAppReviewDetailUpdateRequestAttributes) (*BetaAppReviewDetailResponse, *Response, error) {
	req := betaAppReviewDetailUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "betaAppReviewDetails",
	}
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
