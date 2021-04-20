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

// BetaAppLocalization defines model for BetaAppLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalization
type BetaAppLocalization struct {
	Attributes    *BetaAppLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                            `json:"id"`
	Links         ResourceLinks                     `json:"links"`
	Relationships *BetaAppLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                            `json:"type"`
}

// BetaAppLocalizationAttributes defines model for BetaAppLocalization.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalization/attributes
type BetaAppLocalizationAttributes struct {
	Description       *string `json:"description,omitempty"`
	FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
	Locale            *string `json:"locale,omitempty"`
	MarketingURL      *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
}

// BetaAppLocalizationRelationships defines model for BetaAppLocalization.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalization/relationships
type BetaAppLocalizationRelationships struct {
	App *Relationship `json:"app,omitempty"`
}

// betaAppLocalizationCreateRequest defines model for betaAppLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationcreaterequest/data
type betaAppLocalizationCreateRequest struct {
	Attributes    BetaAppLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships betaAppLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                        `json:"type"`
}

// BetaAppLocalizationCreateRequestAttributes are attributes for BetaAppLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationcreaterequest/data/attributes
type BetaAppLocalizationCreateRequestAttributes struct {
	Description       *string `json:"description,omitempty"`
	FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
	Locale            string  `json:"locale"`
	MarketingURL      *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
}

// betaAppLocalizationCreateRequestRelationships are relationships for BetaAppLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationcreaterequest/data/relationships
type betaAppLocalizationCreateRequestRelationships struct {
	App relationshipDeclaration `json:"app"`
}

// BetaAppLocalizationResponse defines model for BetaAppLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationresponse
type BetaAppLocalizationResponse struct {
	Data     BetaAppLocalization `json:"data"`
	Included []App               `json:"included,omitempty"`
	Links    DocumentLinks       `json:"links"`
}

// BetaAppLocalizationUpdateRequest defines model for BetaAppLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationupdaterequest/data
type betaAppLocalizationUpdateRequest struct {
	Attributes *BetaAppLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// BetaAppLocalizationUpdateRequestAttributes are attributes for BetaAppLocalizationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationupdaterequest/data/attributes
type BetaAppLocalizationUpdateRequestAttributes struct {
	Description       *string `json:"description,omitempty"`
	FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
	MarketingURL      *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
}

// BetaAppLocalizationsResponse defines model for BetaAppLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betaapplocalizationsresponse
type BetaAppLocalizationsResponse struct {
	Data     []BetaAppLocalization `json:"data"`
	Included []App                 `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// ListBetaAppLocalizationsQuery defines model for ListBetaAppLocalizations
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_localizations
type ListBetaAppLocalizationsQuery struct {
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsBetaAppLocalizations []string `url:"fields[betaAppLocalizations],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Include                    []string `url:"include,omitempty"`
	FilterApp                  []string `url:"filter[app],omitempty"`
	FilterLocale               []string `url:"filter[locale],omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetBetaAppLocalizationQuery defines model for GetBetaAppLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_localization_information
type GetBetaAppLocalizationQuery struct {
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsBetaAppLocalizations []string `url:"fields[betaAppLocalizations],omitempty"`
	Include                    []string `url:"include,omitempty"`
}

// GetAppForBetaAppLocalizationQuery defines model for GetAppForBetaAppLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_app_localization
type GetAppForBetaAppLocalizationQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListBetaAppLocalizationsForAppQuery defines model for ListBetaAppLocalizationsForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_app_localizations_of_an_app
type ListBetaAppLocalizationsForAppQuery struct {
	FieldsBetaAppLocalizations []string `url:"fields[betaAppLocalizations],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// ListBetaAppLocalizations finds and lists beta app localizations for all apps and locales.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_app_localizations
func (s *TestflightService) ListBetaAppLocalizations(ctx context.Context, params *ListBetaAppLocalizationsQuery) (*BetaAppLocalizationsResponse, *Response, error) {
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.client.get(ctx, "betaAppLocalizations", params, res)

	return res, resp, err
}

// GetBetaAppLocalization gets localized beta app information for a specific app and locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_app_localization_information
func (s *TestflightService) GetBetaAppLocalization(ctx context.Context, id string, params *GetBetaAppLocalizationQuery) (*BetaAppLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppForBetaAppLocalization gets the app information associated with a specific beta app localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_app_localization
func (s *TestflightService) GetAppForBetaAppLocalization(ctx context.Context, id string, params *GetAppForBetaAppLocalizationQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBetaAppLocalizationsForApp gets a list of localized beta test information for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_app_localizations_of_an_app
func (s *TestflightService) ListBetaAppLocalizationsForApp(ctx context.Context, id string, params *ListBetaAppLocalizationsForAppQuery) (*BetaAppLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppLocalizations", id)
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateBetaAppLocalization creates localized descriptive information for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_beta_app_localization
func (s *TestflightService) CreateBetaAppLocalization(ctx context.Context, attributes BetaAppLocalizationCreateRequestAttributes, appID string) (*BetaAppLocalizationResponse, *Response, error) {
	req := betaAppLocalizationCreateRequest{
		Attributes: attributes,
		Relationships: betaAppLocalizationCreateRequestRelationships{
			App: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appID,
					Type: "apps",
				},
			},
		},
		Type: "betaAppLocalizations",
	}
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.post(ctx, "betaAppLocalizations", newRequestBody(req), res)

	return res, resp, err
}

// UpdateBetaAppLocalization updates the localized What’s New text for a specific app and locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_app_localization
func (s *TestflightService) UpdateBetaAppLocalization(ctx context.Context, id string, attributes *BetaAppLocalizationUpdateRequestAttributes) (*BetaAppLocalizationResponse, *Response, error) {
	req := betaAppLocalizationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "betaAppLocalizations",
	}
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteBetaAppLocalization deletes a beta app localization associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_beta_app_localization
func (s *TestflightService) DeleteBetaAppLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)

	return s.client.delete(ctx, url, nil)
}
