package apps

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/internal/services"
)

// AppStoreVersionLocalization defines model for AppStoreVersionLocalization.
type AppStoreVersionLocalization struct {
	Attributes *struct {
		Description     *string `json:"description,omitempty"`
		Keywords        *string `json:"keywords,omitempty"`
		Locale          *string `json:"locale,omitempty"`
		MarketingURL    *string `json:"marketingUrl,omitempty"`
		PromotionalText *string `json:"promotionalText,omitempty"`
		SupportURL      *string `json:"supportUrl,omitempty"`
		WhatsNew        *string `json:"whatsNew,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		AppPreviewSets *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"appPreviewSets,omitempty"`
		AppScreenshotSets *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"appScreenshotSets,omitempty"`
		AppStoreVersion *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionLocalizationCreateRequest defines model for AppStoreVersionLocalizationCreateRequest.
type AppStoreVersionLocalizationCreateRequest struct {
	Attributes    AppStoreVersionLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships AppStoreVersionLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                                `json:"type"`
}

// AppStoreVersionLocalizationCreateRequestAttributes are attributes for AppStoreVersionLocalizationCreateRequest
type AppStoreVersionLocalizationCreateRequestAttributes struct {
	Description     *string `json:"description,omitempty"`
	Keywords        *string `json:"keywords,omitempty"`
	Locale          string  `json:"locale"`
	MarketingURL    *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportURL      *string `json:"supportUrl,omitempty"`
	WhatsNew        *string `json:"whatsNew,omitempty"`
}

// AppStoreVersionLocalizationCreateRequestRelationships are relationships for AppStoreVersionLocalizationCreateRequest
type AppStoreVersionLocalizationCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// AppStoreVersionLocalizationResponse defines model for AppStoreVersionLocalizationResponse.
type AppStoreVersionLocalizationResponse struct {
	Data     AppStoreVersionLocalization `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    services.DocumentLinks      `json:"links"`
}

// AppStoreVersionLocalizationsResponse defines model for AppStoreVersionLocalizationsResponse.
type AppStoreVersionLocalizationsResponse struct {
	Data     []AppStoreVersionLocalization `json:"data"`
	Included *[]interface{}                `json:"included,omitempty"`
	Links    services.PagedDocumentLinks   `json:"links"`
	Meta     *services.PagingInformation   `json:"meta,omitempty"`
}

// AppStoreVersionLocalizationUpdateRequest defines model for AppStoreVersionLocalizationUpdateRequest.
type AppStoreVersionLocalizationUpdateRequest struct {
	Attributes *AppStoreVersionLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                              `json:"id"`
	Type       string                                              `json:"type"`
}

// AppStoreVersionLocalizationUpdateRequestAttributes are attributes for AppStoreVersionLocalizationUpdateRequest
type AppStoreVersionLocalizationUpdateRequestAttributes struct {
	Description     *string `json:"description,omitempty"`
	Keywords        *string `json:"keywords,omitempty"`
	MarketingURL    *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportURL      *string `json:"supportUrl,omitempty"`
	WhatsNew        *string `json:"whatsNew,omitempty"`
}

// ListLocalizationsForAppStoreVersionQuery are query options for ListLocalizationsForAppStoreVersion
type ListLocalizationsForAppStoreVersionQuery struct {
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              *int      `url:"limit,omitempty"`
	Cursor                             *string   `url:"cursor,omitempty"`
}

// GetAppStoreVersionLocalizationQuery are query options for GetAppStoreVersionLocalization
type GetAppStoreVersionLocalizationQuery struct {
	FieldsAppPreviewSets               *[]string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppScreenshotSets            *[]string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                            *[]string `url:"include,omitempty"`
	LimitAppPreviewSets                *int      `url:"limit[appPreviewSets],omitempty"`
	LimitAppScreenshotSets             *int      `url:"limit[appScreenshotSets],omitempty"`
}

// ListAppScreenshotSetsForAppStoreVersionLocalizationQuery are query options for ListAppScreenshotSetsForAppStoreVersionLocalization
type ListAppScreenshotSetsForAppStoreVersionLocalizationQuery struct {
	FieldsAppScreenshotSets            *[]string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots               *[]string `url:"fields[appScreenshots],omitempty"`
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              *int      `url:"limit,omitempty"`
	Include                            *[]string `url:"include,omitempty"`
	FilterScreenshotDisplayType        *[]string `url:"filter[screenshotDisplayType],omitempty"`
	Cursor                             *string   `url:"cursor,omitempty"`
}

// ListAppPreviewSetsForAppStoreVersionLocalizationQuery are query options for ListAppPreviewSetsForAppStoreVersionLocalization
type ListAppPreviewSetsForAppStoreVersionLocalizationQuery struct {
	FieldsAppPreviewSets               *[]string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppPreviews                  *[]string `url:"fields[appPreviews],omitempty"`
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              *int      `url:"limit,omitempty"`
	Include                            *[]string `url:"include,omitempty"`
	FilterPreviewType                  *[]string `url:"filter[previewType],omitempty"`
	Cursor                             *string   `url:"cursor,omitempty"`
}

// ListLocalizationsForAppStoreVersion gets a list of localized, version-level information about an app, for all locales.
func (s *Service) ListLocalizationsForAppStoreVersion(id string, params *ListLocalizationsForAppStoreVersionQuery) (*AppStoreVersionLocalizationsResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionLocalizations", id)
	res := new(AppStoreVersionLocalizationsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppStoreVersionLocalization reads localized version-level information.
func (s *Service) GetAppStoreVersionLocalization(id string, params *GetAppStoreVersionLocalizationQuery) (*AppStoreVersionLocalizationResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppStoreVersionLocalization adds localized version-level information for a new locale.
func (s *Service) CreateAppStoreVersionLocalization(body *AppStoreVersionLocalizationCreateRequest) (*AppStoreVersionLocalizationResponse, *http.Response, error) {
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.Post("appStoreVersionLocalizations", body, res)
	return res, resp, err
}

// UpdateAppStoreVersionLocalization modifies localized version-level information for a particular language.
func (s *Service) UpdateAppStoreVersionLocalization(id string, body *AppStoreVersionLocalizationUpdateRequest) (*AppStoreVersionLocalizationResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersionLocalization deletes a language from your version metadata.
func (s *Service) DeleteAppStoreVersionLocalization(id string) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	return s.Delete(url, nil)
}

// ListAppScreenshotSetsForAppStoreVersionLocalization lists all screenshot sets for a specific localization.
func (s *Service) ListAppScreenshotSetsForAppStoreVersionLocalization(id string, params *ListAppScreenshotSetsForAppStoreVersionLocalizationQuery) (*AppScreenshotSetsResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appScreenshotSets", id)
	res := new(AppScreenshotSetsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppPreviewSetsForAppStoreVersionLocalization lists all app preview sets for a specific localization.
func (s *Service) ListAppPreviewSetsForAppStoreVersionLocalization(id string, params *ListAppPreviewSetsForAppStoreVersionLocalizationQuery) (*AppPreviewSetsResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appPreviewSets", id)
	res := new(AppPreviewSetsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
