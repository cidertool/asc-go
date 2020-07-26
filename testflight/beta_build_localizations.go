package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// BetaBuildLocalization defines model for BetaBuildLocalization.
type BetaBuildLocalization struct {
	Attributes *struct {
		Locale   *string `json:"locale,omitempty"`
		WhatsNew *string `json:"whatsNew,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaBuildLocalizationResponse defines model for BetaBuildLocalizationResponse.
type BetaBuildLocalizationResponse struct {
	Data     BetaBuildLocalization `json:"data"`
	Included *[]builds.Build       `json:"included,omitempty"`
	Links    types.DocumentLinks   `json:"links"`
}

// BetaBuildLocalizationCreateRequest defines model for BetaBuildLocalizationCreateRequest.
type BetaBuildLocalizationCreateRequest struct {
	Attributes    BetaBuildLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships BetaBuildLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

type BetaBuildLocalizationCreateRequestAttributes struct {
	Locale   string  `json:"locale"`
	WhatsNew *string `json:"whatsNew,omitempty"`
}

type BetaBuildLocalizationCreateRequestRelationships struct {
	Build struct {
		Data types.RelationshipsData `json:"data"`
	} `json:"build"`
}

// BetaBuildLocalizationUpdateRequest defines model for BetaBuildLocalizationUpdateRequest.
type BetaBuildLocalizationUpdateRequest struct {
	Attributes *BetaBuildLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                        `json:"id"`
	Type       string                                        `json:"type"`
}

type BetaBuildLocalizationUpdateRequestAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// BetaBuildLocalizationsResponse defines model for BetaBuildLocalizationsResponse.
type BetaBuildLocalizationsResponse struct {
	Data     []BetaBuildLocalization  `json:"data"`
	Included *[]builds.Build          `json:"included,omitempty"`
	Links    types.PagedDocumentLinks `json:"links"`
	Meta     *types.PagingInformation `json:"meta,omitempty"`
}

// ListBetaBuildLocalizationsQuery defines model for ListBetaBuildLocalizations
type ListBetaBuildLocalizationsQuery struct {
	FieldsBuilds                 *[]string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        *int      `url:"limit,omitempty"`
	Include                      *[]string `url:"include,omitempty"`
	FilterBuild                  *[]string `url:"filter[build],omitempty"`
	FilterLocale                 *[]string `url:"filter[locale],omitempty"`
}

// GetBetaBuildLocalizationQuery defines model for GetBetaBuildLocalization
type GetBetaBuildLocalizationQuery struct {
	FieldsBuilds                 *[]string `url:"fields[builds],omitempty"`
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Include                      *[]string `url:"include,omitempty"`
}

// GetBuildForBetaBuildLocalizationQuery defines model for GetBuildForBetaBuildLocalization
type GetBuildForBetaBuildLocalizationQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

// ListBetaBuildLocalizationsForBuildQuery defines model for ListBetaBuildLocalizationsForBuild
type ListBetaBuildLocalizationsForBuildQuery struct {
	FieldsBetaBuildLocalizations *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	Limit                        *int      `url:"limit,omitempty"`
}

// ListBetaBuildLocalizations finds and lists beta build localizations for all builds and locales.
func (s *Service) ListBetaBuildLocalizations(params *ListBetaBuildLocalizationsQuery) (*BetaBuildLocalizationsResponse, *services.Response, error) {
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.GetWithQuery("betaBuildLocalizations", params, res)
	return res, resp, err
}

// GetBetaBuildLocalization gets localized beta build information for a specific build and locale.
func (s *Service) GetBetaBuildLocalization(id string, params *GetBetaBuildLocalizationQuery) (*BetaBuildLocalizationResponse, *services.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBetaBuildLocalization gets the build information associated with a specific beta build localization.
func (s *Service) GetBuildForBetaBuildLocalization(id string, params *GetBuildForBetaBuildLocalizationQuery) (*builds.BuildResponse, *services.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaBuildLocalizationsForBuild gets a list of localized beta test information for a specific build.
func (s *Service) ListBetaBuildLocalizationsForBuild(id string, params *ListBetaBuildLocalizationsForBuildQuery) (*BetaBuildLocalizationsResponse, *services.Response, error) {
	url := fmt.Sprintf("builds/%s/betaBuildLocalizations", id)
	res := new(BetaBuildLocalizationsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateBetaBuildLocalization creates localized descriptive information for an build.
func (s *Service) CreateBetaBuildLocalization(body *BetaBuildLocalizationCreateRequest) (*BetaBuildLocalizationResponse, *services.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations")
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateBetaBuildLocalization updates the localized What’s New text for a specific build and locale.
func (s *Service) UpdateBetaBuildLocalization(id string, body *BetaBuildLocalizationUpdateRequest) (*BetaBuildLocalizationResponse, *services.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	res := new(BetaBuildLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBetaBuildLocalization deletes a beta build localization associated with an build.
func (s *Service) DeleteBetaBuildLocalization(id string) (*services.Response, error) {
	url := fmt.Sprintf("betaBuildLocalizations/%s", id)
	return s.Delete(url, nil)
}
