package asc

import (
	"context"
	"fmt"
)

// ScreenshotDisplayType defines model for ScreenshotDisplayType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/screenshotdisplaytype
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotset
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest
type AppScreenshotSetCreateRequest struct {
	Attributes    AppScreenshotSetCreateRequestAttributes    `json:"attributes"`
	Relationships AppScreenshotSetCreateRequestRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

// AppScreenshotSetCreateRequestAttributes are attributes for AppScreenshotSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest/data/attributes
type AppScreenshotSetCreateRequestAttributes struct {
	ScreenshotDisplayType ScreenshotDisplayType `json:"screenshotDisplayType"`
}

// AppScreenshotSetCreateRequestRelationships are relationships for AppScreenshotSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest/data/relationships
type AppScreenshotSetCreateRequestRelationships struct {
	AppStoreVersionLocalization struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersionLocalization"`
}

// AppScreenshotSetResponse defines model for AppScreenshotSetResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetresponse
type AppScreenshotSetResponse struct {
	Data     AppScreenshotSet `json:"data"`
	Included *[]AppScreenshot `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}

// AppScreenshotSetsResponse defines model for AppScreenshotSetsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetsresponse
type AppScreenshotSetsResponse struct {
	Data     []AppScreenshotSet `json:"data"`
	Included *[]AppScreenshot   `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppScreenshotSetAppScreenshotsLinkagesResponse defines model for AppScreenshotSetAppScreenshotsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetappscreenshotslinkagesresponse
type AppScreenshotSetAppScreenshotsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// GetAppScreenshotSetQuery are query options for GetAppScreenshotSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_set_information
type GetAppScreenshotSetQuery struct {
	FieldsAppScreenshots    []string `url:"fields[appScreenshots],omitempty"`
	FieldsAppScreenshotSets []string `url:"fields[appScreenshotSets],omitempty"`
	Include                 []string `url:"include,omitempty"`
	LimitAppScreenshots     int      `url:"limit[appScreenshots],omitempty"`
}

// ListAppScreenshotsForSetQuery are query options for ListAppScreenshotsForSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_screenshots_for_an_app_screenshot_set
type ListAppScreenshotsForSetQuery struct {
	FieldsAppScreenshotSets []string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots    []string `url:"fields[appScreenshots],omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Include                 []string `url:"include,omitempty"`
	Cursor                  string   `url:"cursor,omitempty"`
}

// ListAppScreenshotIDsForSetQuery are query options for ListAppScreenshotIDsForSet
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_screenshot_ids_for_an_app_screenshot_set
type ListAppScreenshotIDsForSetQuery struct {
	Limit int `url:"limit,omitempty"`
}

// GetAppScreenshotSet gets an app screenshot set including its display target, language, and the screenshot it contains.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_set_information
func (s *AppsService) GetAppScreenshotSet(ctx context.Context, id string, params *GetAppScreenshotSetQuery) (*AppScreenshotSetResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	res := new(AppScreenshotSetResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateAppScreenshotSet adds a new screenshot set to an App Store version localization for a specific screenshot type and display size.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_screenshot_set
func (s *AppsService) CreateAppScreenshotSet(ctx context.Context, body *AppScreenshotSetCreateRequest) (*AppScreenshotSetResponse, *Response, error) {
	res := new(AppScreenshotSetResponse)
	resp, err := s.client.post(ctx, "appScreenshotSets", body, res)
	return res, resp, err
}

// DeleteAppScreenshotSet deletes an app screenshot set and all of its screenshots.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_screenshot_set
func (s *AppsService) DeleteAppScreenshotSet(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	return s.client.delete(ctx, url, nil)
}

// ListAppScreenshotsForSet lists all ordered screenshots in a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_screenshots_for_an_app_screenshot_set
func (s *AppsService) ListAppScreenshotsForSet(ctx context.Context, id string, params *ListAppScreenshotsForSetQuery) (*AppScreenshotsResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/appScreenshots", id)
	res := new(AppScreenshotsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListAppScreenshotIDsForSet gets the ordered screenshot IDs in a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_screenshot_ids_for_an_app_screenshot_set
func (s *AppsService) ListAppScreenshotIDsForSet(ctx context.Context, id string, params *ListAppScreenshotIDsForSetQuery) (*AppScreenshotSetAppScreenshotsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	res := new(AppScreenshotSetAppScreenshotsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ReplaceAppScreenshotsForSet changes the order of the screenshots in a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/replace_all_app_screenshots_for_an_app_screenshot_set
func (s *AppsService) ReplaceAppScreenshotsForSet(ctx context.Context, id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	return s.client.patch(ctx, url, linkages, nil)
}
