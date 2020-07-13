package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaAppLocalization defines model for BetaAppLocalization.
type BetaAppLocalization struct {
	Attributes *struct {
		Description       *string `json:"description,omitempty"`
		FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
		Locale            *string `json:"locale,omitempty"`
		MarketingURL      *string `json:"marketingUrl,omitempty"`
		PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
		TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppLocalizationCreateRequest defines model for BetaAppLocalizationCreateRequest.
type BetaAppLocalizationCreateRequest struct {
	Data struct {
		Attributes struct {
			Description       *string `json:"description,omitempty"`
			FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
			Locale            string  `json:"locale"`
			MarketingURL      *string `json:"marketingUrl,omitempty"`
			PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
			TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			App struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"app"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaAppLocalizationResponse defines model for BetaAppLocalizationResponse.
type BetaAppLocalizationResponse struct {
	Data     BetaAppLocalization  `json:"data"`
	Included *[]apps.App          `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BetaAppLocalizationUpdateRequest defines model for BetaAppLocalizationUpdateRequest.
type BetaAppLocalizationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Description       *string `json:"description,omitempty"`
			FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
			MarketingURL      *string `json:"marketingUrl,omitempty"`
			PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
			TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaAppLocalizationsResponse defines model for BetaAppLocalizationsResponse.
type BetaAppLocalizationsResponse struct {
	Data     []BetaAppLocalization     `json:"data"`
	Included *[]apps.App               `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaAppLocalizationsOptions struct {
	Fields *struct {
		Apps                 *[]string `url:"apps,omitempty"`
		BetaAppLocalizations *[]string `url:"betaAppLocalizations,omitempty"`
	} `url:"fields,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Filter  *struct {
		App    *[]string `url:"app,omitempty"`
		Locale *[]string `url:"locale,omitempty"`
	} `url:"filter,omitempty"`
}

type GetBetaAppLocalizationOptions struct {
	Fields *struct {
		Apps                 *[]string `url:"apps,omitempty"`
		BetaAppLocalizations *[]string `url:"betaAppLocalizations,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetAppForBetaAppLocalizationOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type ListBetaAppLocalizationsForAppOptions struct {
	Fields *struct {
		BetaAppLocalizations *[]string `url:"betaAppLocalizations,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// ListBetaAppLocalizations finds and lists beta app localizations for all apps and locales.
func (s *Service) ListBetaAppLocalizations(params *ListBetaAppLocalizationsOptions) (*BetaAppLocalizationsResponse, *common.Response, error) {
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.Get("betaAppLocalizations", params, res)
	return res, resp, err
}

// GetBetaAppLocalization gets localized beta app information for a specific app and locale.
func (s *Service) GetBetaAppLocalization(id string, params *GetBetaAppLocalizationOptions) (*BetaAppLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForBetaAppLocalization gets the app information associated with a specific beta app localization.
func (s *Service) GetAppForBetaAppLocalization(id string, params *GetAppForBetaAppLocalizationOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaAppLocalizationsForApp gets a list of localized beta test information for a specific app.
func (s *Service) ListBetaAppLocalizationsForApp(id string, params *ListBetaAppLocalizationsForAppOptions) (*BetaAppLocalizationsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppLocalizations", id)
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// CreateBetaAppLocalization creates localized descriptive information for an app.
func (s *Service) CreateBetaAppLocalization(body *BetaAppLocalizationCreateRequest) (*BetaAppLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations")
	res := new(BetaAppLocalizationResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateBetaAppLocalization updates the localized What’s New text for a specific app and locale.
func (s *Service) UpdateBetaAppLocalization(id string, body *BetaAppLocalizationUpdateRequest) (*BetaAppLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBetaAppLocalization deletes a beta app localization associated with an app.
func (s *Service) DeleteBetaAppLocalization(id string) (*common.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	return s.Delete(url, nil)
}
