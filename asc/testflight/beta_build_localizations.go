package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaBuildLocalization defines model for BetaBuildLocalization.
type BetaBuildLocalization struct {
	Attributes *struct {
		Locale   *string `json:"locale,omitempty"`
		WhatsNew *string `json:"whatsNew,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaBuildLocalizationResponse defines model for BetaBuildLocalizationResponse.
type BetaBuildLocalizationResponse struct {
	Data     BetaBuildLocalization `json:"data"`
	Included *[]builds.Build       `json:"included,omitempty"`
	Links    common.DocumentLinks  `json:"links"`
}

// BetaBuildLocalizationCreateRequest defines model for BetaBuildLocalizationCreateRequest.
type BetaBuildLocalizationCreateRequest struct {
	Data struct {
		Attributes struct {
			Locale   string  `json:"locale"`
			WhatsNew *string `json:"whatsNew,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			Build struct {
				Data common.RelationshipsData `json:"data"`
			} `json:"build"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaBuildLocalizationUpdateRequest defines model for BetaBuildLocalizationUpdateRequest.
type BetaBuildLocalizationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			WhatsNew *string `json:"whatsNew,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaBuildLocalizationsResponse defines model for BetaBuildLocalizationsResponse.
type BetaBuildLocalizationsResponse struct {
	Data     []BetaBuildLocalization   `json:"data"`
	Included *[]builds.Build           `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaBuildLocalizationsQuery struct {
	FieldsBuilds                 *[]string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        *int      `url:"limit,omitempty"`
	Include                      *[]string `url:"include,omitempty"`
	FilterBuild                  *[]string `url:"filter[build],omitempty"`
	FilterLocale                 *[]string `url:"filter[locale],omitempty"`
}

type GetBetaBuildLocalizationQuery struct {
	FieldsBuilds                 *[]string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Include                      *[]string `url:"include,omitempty"`
}

type GetBuildForBetaBuildLocalizationQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

type ListBetaBuildLocalizationsForBuildQuery struct {
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        *int      `url:"limit,omitempty"`
}

// ListBetaBuildLocalizations finds and lists beta build localizations for all builds and locales.
func (s *Service) ListBetaBuildLocalizations(params *ListBetaBuildLocalizationsQuery) (*BetaBuildLocalizationsResponse, *common.Response, error) {
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.GetWithQuery("betaBuildLocalizations", params, res)
	return res, resp, err
}

// GetBetaBuildLocalization gets localized beta build information for a specific build and locale.
func (s *Service) GetBetaBuildLocalization(id string, params *GetBetaBuildLocalizationQuery) (*BetaBuildLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBetaBuildLocalization gets the build information associated with a specific beta build localization.
func (s *Service) GetBuildForBetaBuildLocalization(id string, params *GetBuildForBetaBuildLocalizationQuery) (*builds.BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaBuildLocalizationsForBuild gets a list of localized beta test information for a specific build.
func (s *Service) ListBetaBuildLocalizationsForBuild(id string, params *ListBetaBuildLocalizationsForBuildQuery) (*BetaBuildLocalizationsResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/betaBuildLocalizations", id)
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateBetaBuildLocalization creates localized descriptive information for an build.
func (s *Service) CreateBetaBuildLocalization(body *BetaBuildLocalizationCreateRequest) (*BetaBuildLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations")
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateBetaBuildLocalization updates the localized What’s New text for a specific build and locale.
func (s *Service) UpdateBetaBuildLocalization(id string, body *BetaBuildLocalizationUpdateRequest) (*BetaBuildLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBetaBuildLocalization deletes a beta build localization associated with an build.
func (s *Service) DeleteBetaBuildLocalization(id string) (*common.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	return s.Delete(url, nil)
}
