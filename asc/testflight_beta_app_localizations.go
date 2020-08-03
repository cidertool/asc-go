package asc

import (
	"fmt"
	"net/http"
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
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppLocalizationCreateRequest defines model for BetaAppLocalizationCreateRequest.
type BetaAppLocalizationCreateRequest struct {
	Attributes    BetaAppLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships BetaAppLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                        `json:"type"`
}

// BetaAppLocalizationCreateRequestAttributes are attributes for BetaAppLocalizationCreateRequest
type BetaAppLocalizationCreateRequestAttributes struct {
	Description       *string `json:"description,omitempty"`
	FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
	Locale            string  `json:"locale"`
	MarketingURL      *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
}

// BetaAppLocalizationCreateRequestRelationships are relationships for BetaAppLocalizationCreateRequest
type BetaAppLocalizationCreateRequestRelationships struct {
	App struct {
		Data RelationshipsData `json:"data"`
	} `json:"app"`
}

// BetaAppLocalizationResponse defines model for BetaAppLocalizationResponse.
type BetaAppLocalizationResponse struct {
	Data     BetaAppLocalization `json:"data"`
	Included *[]App              `json:"included,omitempty"`
	Links    DocumentLinks       `json:"links"`
}

// BetaAppLocalizationUpdateRequest defines model for BetaAppLocalizationUpdateRequest.
type BetaAppLocalizationUpdateRequest struct {
	Attributes *BetaAppLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// BetaAppLocalizationUpdateRequestAttributes are attributes for BetaAppLocalizationUpdateRequest
type BetaAppLocalizationUpdateRequestAttributes struct {
	Description       *string `json:"description,omitempty"`
	FeedbackEmail     *string `json:"feedbackEmail,omitempty"`
	MarketingURL      *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	TVOSPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
}

// BetaAppLocalizationsResponse defines model for BetaAppLocalizationsResponse.
type BetaAppLocalizationsResponse struct {
	Data     []BetaAppLocalization `json:"data"`
	Included *[]App                `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// ListBetaAppLocalizationsQuery defines model for ListBetaAppLocalizations
type ListBetaAppLocalizationsQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppLocalizations *[]string `url:"fields[betaAppLocalizations],omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	FilterApp                  *[]string `url:"filter[app],omitempty"`
	FilterLocale               *[]string `url:"filter[locale],omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// GetBetaAppLocalizationQuery defines model for GetBetaAppLocalization
type GetBetaAppLocalizationQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppLocalizations *[]string `url:"fields[betaAppLocalizations],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
}

// GetAppForBetaAppLocalizationQuery defines model for GetAppForBetaAppLocalization
type GetAppForBetaAppLocalizationQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// ListBetaAppLocalizationsForAppQuery defines model for ListBetaAppLocalizationsForApp
type ListBetaAppLocalizationsForAppQuery struct {
	FieldsBetaAppLocalizations *[]string `url:"fields[betaAppLocalizations],omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// ListBetaAppLocalizations finds and lists beta app localizations for all apps and locales.
func (s *TestflightService) ListBetaAppLocalizations(params *ListBetaAppLocalizationsQuery) (*BetaAppLocalizationsResponse, *http.Response, error) {
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.client.get("betaAppLocalizations", params, res)
	return res, resp, err
}

// GetBetaAppLocalization gets localized beta app information for a specific app and locale.
func (s *TestflightService) GetBetaAppLocalization(id string, params *GetBetaAppLocalizationQuery) (*BetaAppLocalizationResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForBetaAppLocalization gets the app information associated with a specific beta app localization.
func (s *TestflightService) GetAppForBetaAppLocalization(id string, params *GetAppForBetaAppLocalizationQuery) (*AppResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBetaAppLocalizationsForApp gets a list of localized beta test information for a specific app.
func (s *TestflightService) ListBetaAppLocalizationsForApp(id string, params *ListBetaAppLocalizationsForAppQuery) (*BetaAppLocalizationsResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppLocalizations", id)
	res := new(BetaAppLocalizationsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateBetaAppLocalization creates localized descriptive information for an app.
func (s *TestflightService) CreateBetaAppLocalization(body *BetaAppLocalizationCreateRequest) (*BetaAppLocalizationResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations")
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.post(url, body, res)
	return res, resp, err
}

// UpdateBetaAppLocalization updates the localized What’s New text for a specific app and locale.
func (s *TestflightService) UpdateBetaAppLocalization(id string, body *BetaAppLocalizationUpdateRequest) (*BetaAppLocalizationResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	res := new(BetaAppLocalizationResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteBetaAppLocalization deletes a beta app localization associated with an app.
func (s *TestflightService) DeleteBetaAppLocalization(id string) (*http.Response, error) {
	url := fmt.Sprintf("betaAppLocalizations/%s", id)
	return s.client.delete(url, nil)
}
