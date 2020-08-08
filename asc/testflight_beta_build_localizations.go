package asc

import (
	"context"
	"fmt"
)

// BetaBuildLocalization defines model for BetaBuildLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalization
type BetaBuildLocalization struct {
	Attributes *struct {
		Locale   *string `json:"locale,omitempty"`
		WhatsNew *string `json:"whatsNew,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaBuildLocalizationResponse defines model for BetaBuildLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationresponse
type BetaBuildLocalizationResponse struct {
	Data     BetaBuildLocalization `json:"data"`
	Included *[]Build              `json:"included,omitempty"`
	Links    DocumentLinks         `json:"links"`
}

// BetaBuildLocalizationCreateRequest defines model for BetaBuildLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest
type BetaBuildLocalizationCreateRequest struct {
	Attributes    BetaBuildLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships BetaBuildLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// BetaBuildLocalizationCreateRequestAttributes are attributes for BetaBuildLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest/data/attributes
type BetaBuildLocalizationCreateRequestAttributes struct {
	Locale   string  `json:"locale"`
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationCreateRequestRelationships are relationships for BetaBuildLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationcreaterequest/data/relationships
type BetaBuildLocalizationCreateRequestRelationships struct {
	Build struct {
		Data RelationshipsData `json:"data"`
	} `json:"build"`
}

// BetaBuildLocalizationUpdateRequest defines model for BetaBuildLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationupdaterequest
type BetaBuildLocalizationUpdateRequest struct {
	Attributes *BetaBuildLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                        `json:"id"`
	Type       string                                        `json:"type"`
}

// BetaBuildLocalizationUpdateRequestAttributes are attributes for BetaBuildLocalizationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationupdaterequest/data/attributes
type BetaBuildLocalizationUpdateRequestAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationsResponse defines model for BetaBuildLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betabuildlocalizationsresponse
type BetaBuildLocalizationsResponse struct {
	Data     []BetaBuildLocalization `json:"data"`
	Included *[]Build                `json:"included,omitempty"`
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
func (s *TestflightService) CreateBetaBuildLocalization(ctx context.Context, body *BetaBuildLocalizationCreateRequest) (*BetaBuildLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations")
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.client.post(ctx, url, body, res)
	return res, resp, err
}

// UpdateBetaBuildLocalization updates the localized What’s New text for a specific build and locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_build_localization
func (s *TestflightService) UpdateBetaBuildLocalization(ctx context.Context, id string, body *BetaBuildLocalizationUpdateRequest) (*BetaBuildLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeleteBetaBuildLocalization deletes a beta build localization associated with an build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_beta_build_localization
func (s *TestflightService) DeleteBetaBuildLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	return s.client.delete(ctx, url, nil)
}
