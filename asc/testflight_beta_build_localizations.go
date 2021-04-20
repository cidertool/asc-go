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

// BetaBuildLocalization defines model for BetaBuildLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalization
type BetaBuildLocalization struct {
	Attributes    *BetaBuildLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                              `json:"id"`
	Links         ResourceLinks                       `json:"links"`
	Relationships *BetaBuildLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                              `json:"type"`
}

// BetaBuildLocalizationAttributes defines model for BetaBuildLocalization.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalization/attributes
type BetaBuildLocalizationAttributes struct {
	Locale   *string `json:"locale,omitempty"`
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationRelationships defines model for BetaBuildLocalization.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalization/relationships
type BetaBuildLocalizationRelationships struct {
	Build *Relationship `json:"build,omitempty"`
}

// BetaBuildLocalizationResponse defines model for BetaBuildLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationresponse
type BetaBuildLocalizationResponse struct {
	Data     BetaBuildLocalization `json:"data"`
	Included []Build               `json:"included,omitempty"`
	Links    DocumentLinks         `json:"links"`
}

// BetaBuildLocalizationCreateRequest defines model for BetaBuildLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest/data
type betaBuildLocalizationCreateRequest struct {
	Attributes    betaBuildLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships betaBuildLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// BetaBuildLocalizationCreateRequestAttributes are attributes for BetaBuildLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest/data/attributes
type betaBuildLocalizationCreateRequestAttributes struct {
	Locale   string  `json:"locale"`
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationCreateRequestRelationships are relationships for BetaBuildLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest/data/relationships
type betaBuildLocalizationCreateRequestRelationships struct {
	Build relationshipDeclaration `json:"build"`
}

// BetaBuildLocalizationUpdateRequest defines model for BetaBuildLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationupdaterequest/data
type betaBuildLocalizationUpdateRequest struct {
	Attributes *betaBuildLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                        `json:"id"`
	Type       string                                        `json:"type"`
}

// BetaBuildLocalizationUpdateRequestAttributes are attributes for BetaBuildLocalizationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationupdaterequest/data/attributes
type betaBuildLocalizationUpdateRequestAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationsResponse defines model for BetaBuildLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationsresponse
type BetaBuildLocalizationsResponse struct {
	Data     []BetaBuildLocalization `json:"data"`
	Included []Build                 `json:"included,omitempty"`
	Links    PagedDocumentLinks      `json:"links"`
	Meta     *PagingInformation      `json:"meta,omitempty"`
}

// ListBetaBuildLocalizationsQuery defines model for ListBetaBuildLocalizations
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_build_localizations
type ListBetaBuildLocalizationsQuery struct {
	FieldsBuilds                 []string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations []string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        int      `url:"limit,omitempty"`
	Include                      []string `url:"include,omitempty"`
	FilterBuild                  []string `url:"filter[build],omitempty"`
	FilterLocale                 []string `url:"filter[locale],omitempty"`
}

// GetBetaBuildLocalizationQuery defines model for GetBetaBuildLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_build_localization_information
type GetBetaBuildLocalizationQuery struct {
	FieldsBuilds                 []string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations []string `url:"fields[betaBuildLocalizations],omitempty"`
	Include                      []string `url:"include,omitempty"`
}

// GetBuildForBetaBuildLocalizationQuery defines model for GetBuildForBetaBuildLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_beta_build_localization
type GetBuildForBetaBuildLocalizationQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
}

// ListBetaBuildLocalizationsForBuildQuery defines model for ListBetaBuildLocalizationsForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_build_localizations_of_a_build
type ListBetaBuildLocalizationsForBuildQuery struct {
	FieldsBetaBuildLocalizations []string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        int      `url:"limit,omitempty"`
}

// ListBetaBuildLocalizations finds and lists beta build localizations for all builds and locales.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_build_localizations
func (s *TestflightService) ListBetaBuildLocalizations(ctx context.Context, params *ListBetaBuildLocalizationsQuery) (*BetaBuildLocalizationsResponse, *Response, error) {
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.client.get(ctx, "betaBuildLocalizations", params, res)

	return res, resp, err
}

// GetBetaBuildLocalization gets localized beta build information for a specific build and locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_build_localization_information
func (s *TestflightService) GetBetaBuildLocalization(ctx context.Context, id string, params *GetBetaBuildLocalizationQuery) (*BetaBuildLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetBuildForBetaBuildLocalization gets the build information associated with a specific beta build localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_a_beta_build_localization
func (s *TestflightService) GetBuildForBetaBuildLocalization(ctx context.Context, id string, params *GetBuildForBetaBuildLocalizationQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBetaBuildLocalizationsForBuild gets a list of localized beta test information for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_build_localizations_of_a_build
func (s *TestflightService) ListBetaBuildLocalizationsForBuild(ctx context.Context, id string, params *ListBetaBuildLocalizationsForBuildQuery) (*BetaBuildLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/betaBuildLocalizations", id)
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateBetaBuildLocalization creates localized descriptive information for an build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_beta_build_localization
func (s *TestflightService) CreateBetaBuildLocalization(ctx context.Context, locale string, whatsNew *string, buildID string) (*BetaBuildLocalizationResponse, *Response, error) {
	req := betaBuildLocalizationCreateRequest{
		Attributes: betaBuildLocalizationCreateRequestAttributes{
			Locale:   locale,
			WhatsNew: whatsNew,
		},
		Relationships: betaBuildLocalizationCreateRequestRelationships{
			Build: relationshipDeclaration{
				Data: RelationshipData{
					ID:   buildID,
					Type: "builds",
				},
			},
		},
		Type: "betaBuildLocalizations",
	}
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.client.post(ctx, "betaBuildLocalizations", newRequestBody(req), res)

	return res, resp, err
}

// UpdateBetaBuildLocalization updates the localized What’s New text for a specific build and locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_build_localization
func (s *TestflightService) UpdateBetaBuildLocalization(ctx context.Context, id string, whatsNew *string) (*BetaBuildLocalizationResponse, *Response, error) {
	req := betaBuildLocalizationUpdateRequest{
		ID:   id,
		Type: "betaBuildLocalizations",
	}

	if whatsNew != nil {
		req.Attributes = &betaBuildLocalizationUpdateRequestAttributes{
			WhatsNew: whatsNew,
		}
	}

	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteBetaBuildLocalization deletes a beta build localization associated with an build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_beta_build_localization
func (s *TestflightService) DeleteBetaBuildLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)

	return s.client.delete(ctx, url, nil)
}
