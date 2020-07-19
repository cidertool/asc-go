package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
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
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppPreviewSets *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"appPreviewSets,omitempty"`
		AppScreenshotSets *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"appScreenshotSets,omitempty"`
		AppStoreVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionLocalizationCreateRequest defines model for AppStoreVersionLocalizationCreateRequest.
type AppStoreVersionLocalizationCreateRequest struct {
	Data struct {
		Attributes struct {
			Description     *string `json:"description,omitempty"`
			Keywords        *string `json:"keywords,omitempty"`
			Locale          string  `json:"locale"`
			MarketingURL    *string `json:"marketingUrl,omitempty"`
			PromotionalText *string `json:"promotionalText,omitempty"`
			SupportURL      *string `json:"supportUrl,omitempty"`
			WhatsNew        *string `json:"whatsNew,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreVersion struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionLocalizationResponse defines model for AppStoreVersionLocalizationResponse.
type AppStoreVersionLocalizationResponse struct {
	Data     AppStoreVersionLocalization `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    common.DocumentLinks        `json:"links"`
}

// AppStoreVersionLocalizationsResponse defines model for AppStoreVersionLocalizationsResponse.
type AppStoreVersionLocalizationsResponse struct {
	Data     []AppStoreVersionLocalization `json:"data"`
	Included *[]interface{}                `json:"included,omitempty"`
	Links    common.PagedDocumentLinks     `json:"links"`
	Meta     *common.PagingInformation     `json:"meta,omitempty"`
}

// AppStoreVersionLocalizationUpdateRequest defines model for AppStoreVersionLocalizationUpdateRequest.
type AppStoreVersionLocalizationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Description     *string `json:"description,omitempty"`
			Keywords        *string `json:"keywords,omitempty"`
			MarketingURL    *string `json:"marketingUrl,omitempty"`
			PromotionalText *string `json:"promotionalText,omitempty"`
			SupportURL      *string `json:"supportUrl,omitempty"`
			WhatsNew        *string `json:"whatsNew,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

type ListLocalizationsForAppStoreVersionQuery struct {
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              *int      `url:"limit,omitempty"`
	Cursor                             *string   `url:"cursor,omitempty"`
}

type GetAppStoreVersionLocalizationQuery struct {
	FieldsAppPreviewSets               *[]string `url:"fields[appPreviewSets],omitempty"`
	FieldsAppScreenshotSets            *[]string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                            *[]string `url:"include,omitempty"`
	LimitAppPreviewSets                *int      `url:"limit[appPreviewSets],omitempty"`
	LimitAppScreenshotSets             *int      `url:"limit[appScreenshotSets],omitempty"`
}

type ListAppScreenshotSetsForAppStoreVersionLocalizationQuery struct {
	FieldsAppScreenshotSets            *[]string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots               *[]string `url:"fields[appScreenshots],omitempty"`
	FieldsAppStoreVersionLocalizations *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Limit                              *int      `url:"limit,omitempty"`
	Include                            *[]string `url:"include,omitempty"`
	FilterScreenshotDisplayType        *[]string `url:"filter[screenshotDisplayType],omitempty"`
	Cursor                             *string   `url:"cursor,omitempty"`
}

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
func (s *Service) ListLocalizationsForAppStoreVersion(id string, params *ListLocalizationsForAppStoreVersionQuery) (*AppStoreVersionLocalizationsResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionLocalizations", id)
	res := new(AppStoreVersionLocalizationsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppStoreVersionLocalization reads localized version-level information.
func (s *Service) GetAppStoreVersionLocalization(id string, params *GetAppStoreVersionLocalizationQuery) (*AppStoreVersionLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppStoreVersionLocalization adds localized version-level information for a new locale.
func (s *Service) CreateAppStoreVersionLocalization(body *AppStoreVersionLocalizationCreateRequest) (*AppStoreVersionLocalizationResponse, *common.Response, error) {
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.Post("appStoreVersionLocalizations", body, res)
	return res, resp, err
}

// UpdateAppStoreVersionLocalization modifies localized version-level information for a particular language.
func (s *Service) UpdateAppStoreVersionLocalization(id string, body *AppStoreVersionLocalizationUpdateRequest) (*AppStoreVersionLocalizationResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	res := new(AppStoreVersionLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersionLocalization deletes a language from your version metadata.
func (s *Service) DeleteAppStoreVersionLocalization(id string) (*common.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s", id)
	return s.Delete(url, nil)
}

// ListAppScreenshotSetsForAppStoreVersionLocalization lists all screenshot sets for a specific localization.
func (s *Service) ListAppScreenshotSetsForAppStoreVersionLocalization(id string, params *ListAppScreenshotSetsForAppStoreVersionLocalizationQuery) (*AppScreenshotSetsResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appScreenshotSets", id)
	res := new(AppScreenshotSetsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppPreviewSetsForAppStoreVersionLocalization lists all app preview sets for a specific localization.
func (s *Service) ListAppPreviewSetsForAppStoreVersionLocalization(id string, params *ListAppPreviewSetsForAppStoreVersionLocalizationQuery) (*AppPreviewSetsResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersionLocalizations/%s/appPreviewSets", id)
	res := new(AppPreviewSetsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
