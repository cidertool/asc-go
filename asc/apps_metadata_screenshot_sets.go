package asc

import (
	"fmt"
	"net/http"
)

// ScreenshotDisplayType defines model for ScreenshotDisplayType.
type ScreenshotDisplayType string

// List of ScreenshotDisplayType
const (
	ScreenshotDisplayTypeAppAppleTV                ScreenshotDisplayType = "APP_APPLE_TV"
	ScreenshotDisplayTypeAppDesktop                ScreenshotDisplayType = "APP_DESKTOP"
	ScreenshotDisplayTypeAppiPad105                ScreenshotDisplayType = "APP_IPAD_105"
	ScreenshotDisplayTypeAppiPad97                 ScreenshotDisplayType = "APP_IPAD_97"
	ScreenshotDisplayTypeAppiPadPro129             ScreenshotDisplayType = "APP_IPAD_PRO_129"
	ScreenshotDisplayTypeAppiPadPro3Gen11          ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_11"
	ScreenshotDisplayTypeAppiPadPro3Gen129         ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_129"
	ScreenshotDisplayTypeAppiPhone35               ScreenshotDisplayType = "APP_IPHONE_35"
	ScreenshotDisplayTypeAppiPhone40               ScreenshotDisplayType = "APP_IPHONE_40"
	ScreenshotDisplayTypeAppiPhone47               ScreenshotDisplayType = "APP_IPHONE_47"
	ScreenshotDisplayTypeAppiPhone55               ScreenshotDisplayType = "APP_IPHONE_55"
	ScreenshotDisplayTypeAppiPhone58               ScreenshotDisplayType = "APP_IPHONE_58"
	ScreenshotDisplayTypeAppiPhone65               ScreenshotDisplayType = "APP_IPHONE_65"
	ScreenshotDisplayTypeAppWatchSeries3           ScreenshotDisplayType = "APP_WATCH_SERIES_3"
	ScreenshotDisplayTypeAppWatchSeries4           ScreenshotDisplayType = "APP_WATCH_SERIES_4"
	ScreenshotDisplayTypeiMessageAppIPad105        ScreenshotDisplayType = "IMESSAGE_APP_IPAD_105"
	ScreenshotDisplayTypeiMessageAppIPad97         ScreenshotDisplayType = "IMESSAGE_APP_IPAD_97"
	ScreenshotDisplayTypeiMessageAppIPadPro129     ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_129"
	ScreenshotDisplayTypeiMessageAppIPadPro3Gen11  ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_11"
	ScreenshotDisplayTypeiMessageAppIPadPro3Gen129 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_129"
	ScreenshotDisplayTypeiMessageAppIPhone40       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_40"
	ScreenshotDisplayTypeiMessageAppIPhone47       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_47"
	ScreenshotDisplayTypeiMessageAppIPhone55       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_55"
	ScreenshotDisplayTypeiMessageAppIPhone58       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_58"
	ScreenshotDisplayTypeiMessageAppIPhone65       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_65"
)

// AppScreenshotSet defines model for AppScreenshotSet.
type AppScreenshotSet struct {
	Attributes *struct {
		ScreenshotDisplayType *ScreenshotDisplayType `json:"screenshotDisplayType,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshots *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appScreenshots,omitempty"`
		AppStoreVersionLocalization *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionLocalization,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotSetCreateRequest defines model for AppScreenshotSetCreateRequest.
type AppScreenshotSetCreateRequest struct {
	Attributes    AppScreenshotSetCreateRequestAttributes    `json:"attributes"`
	Relationships AppScreenshotSetCreateRequestRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

// AppScreenshotSetCreateRequestAttributes are attributes for AppScreenshotSetCreateRequest
type AppScreenshotSetCreateRequestAttributes struct {
	ScreenshotDisplayType ScreenshotDisplayType `json:"screenshotDisplayType"`
}

// AppScreenshotSetCreateRequestRelationships are relationships for AppScreenshotSetCreateRequest
type AppScreenshotSetCreateRequestRelationships struct {
	AppStoreVersionLocalization struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersionLocalization"`
}

// AppScreenshotSetResponse defines model for AppScreenshotSetResponse.
type AppScreenshotSetResponse struct {
	Data     AppScreenshotSet `json:"data"`
	Included *[]AppScreenshot `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}

// AppScreenshotSetsResponse defines model for AppScreenshotSetsResponse.
type AppScreenshotSetsResponse struct {
	Data     []AppScreenshotSet `json:"data"`
	Included *[]AppScreenshot   `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppScreenshotSetAppScreenshotsLinkagesResponse defines model for AppScreenshotSetAppScreenshotsLinkagesResponse.
type AppScreenshotSetAppScreenshotsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// GetAppScreenshotSetQuery are query options for GetAppScreenshotSet
type GetAppScreenshotSetQuery struct {
	FieldsAppScreenshots    *[]string `url:"fields[appScreenshots],omitempty"`
	FieldsAppScreenshotSets *[]string `url:"fields[appScreenshotSets],omitempty"`
	Include                 *[]string `url:"include,omitempty"`
	LimitAppScreenshots     *int      `url:"limit[appScreenshots],omitempty"`
}

// ListAppScreenshotsForSetQuery are query options for ListAppScreenshotsForSet
type ListAppScreenshotsForSetQuery struct {
	FieldsAppScreenshotSets *[]string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots    *[]string `url:"fields[appScreenshots],omitempty"`
	Limit                   *int      `url:"limit,omitempty"`
	Include                 *[]string `url:"include,omitempty"`
	Cursor                  *string   `url:"cursor,omitempty"`
}

// ListAppScreenshotIDsForSetQuery are query options for ListAppScreenshotIDsForSet
type ListAppScreenshotIDsForSetQuery struct {
	Limit *int `url:"limit,omitempty"`
}

// GetAppScreenshotSet gets an app screenshot set including its display target, language, and the screenshot it contains.
func (s *AppsService) GetAppScreenshotSet(id string, params *GetAppScreenshotSetQuery) (*AppScreenshotSetResponse, *http.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	res := new(AppScreenshotSetResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppScreenshotSet adds a new screenshot set to an App Store version localization for a specific screenshot type and display size.
func (s *AppsService) CreateAppScreenshotSet(id string, body *AppScreenshotSetCreateRequest) (*AppScreenshotSetResponse, *http.Response, error) {
	res := new(AppScreenshotSetResponse)
	resp, err := s.client.Post("appScreenshotSets", body, res)
	return res, resp, err
}

// DeleteAppScreenshotSet deletes an app screenshot set and all of its screenshots.
func (s *AppsService) DeleteAppScreenshotSet(id string) (*http.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	return s.client.Delete(url, nil)
}

// ListAppScreenshotsForSet lists all ordered screenshots in a screenshot set.
func (s *AppsService) ListAppScreenshotsForSet(id string, params *ListAppScreenshotsForSetQuery) (*AppScreenshotsResponse, *http.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/appScreenshots", id)
	res := new(AppScreenshotsResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppScreenshotIDsForSet gets the ordered screenshot IDs in a screenshot set.
func (s *AppsService) ListAppScreenshotIDsForSet(id string, params *ListAppScreenshotIDsForSetQuery) (*AppScreenshotSetAppScreenshotsLinkagesResponse, *http.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	res := new(AppScreenshotSetAppScreenshotsLinkagesResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// ReplaceAppScreenshotsForSet changes the order of the screenshots in a screenshot set.
func (s *AppsService) ReplaceAppScreenshotsForSet(id string, linkages *[]RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	return s.client.Patch(url, linkages, nil)
}
