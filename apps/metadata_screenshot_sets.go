package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// ScreenshotDisplayType defines model for ScreenshotDisplayType.
type ScreenshotDisplayType string

// List of ScreenshotDisplayType
const (
	AppAppleTV                ScreenshotDisplayType = "APP_APPLE_TV"
	AppDesktop                ScreenshotDisplayType = "APP_DESKTOP"
	AppIPad105                ScreenshotDisplayType = "APP_IPAD_105"
	AppIPad97                 ScreenshotDisplayType = "APP_IPAD_97"
	AppIPadPro129             ScreenshotDisplayType = "APP_IPAD_PRO_129"
	AppIPadPro3Gen11          ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_11"
	AppIPadPro3Gen129         ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_129"
	AppIPhone35               ScreenshotDisplayType = "APP_IPHONE_35"
	AppIPhone40               ScreenshotDisplayType = "APP_IPHONE_40"
	AppIPhone47               ScreenshotDisplayType = "APP_IPHONE_47"
	AppIPhone55               ScreenshotDisplayType = "APP_IPHONE_55"
	AppIPhone58               ScreenshotDisplayType = "APP_IPHONE_58"
	AppIPhone65               ScreenshotDisplayType = "APP_IPHONE_65"
	AppWatchSeries3           ScreenshotDisplayType = "APP_WATCH_SERIES_3"
	AppWatchSeries4           ScreenshotDisplayType = "APP_WATCH_SERIES_4"
	IMessageAppIPad105        ScreenshotDisplayType = "IMESSAGE_APP_IPAD_105"
	IMessageAppIPad97         ScreenshotDisplayType = "IMESSAGE_APP_IPAD_97"
	IMessageAppIPadPro129     ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_129"
	IMessageAppIPadPro3Gen11  ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_11"
	IMessageAppIPadPro3Gen129 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_129"
	IMessageAppIPhone40       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_40"
	IMessageAppIPhone47       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_47"
	IMessageAppIPhone55       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_55"
	IMessageAppIPhone58       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_58"
	IMessageAppIPhone65       ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_65"
)

// AppScreenshotSet defines model for AppScreenshotSet.
type AppScreenshotSet struct {
	Attributes *struct {
		ScreenshotDisplayType *ScreenshotDisplayType `json:"screenshotDisplayType,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshots *struct {
			Data  *[]types.RelationshipsData `json:"data,omitempty"`
			Links *types.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *types.PagingInformation   `json:"meta,omitempty"`
		} `json:"appScreenshots,omitempty"`
		AppStoreVersionLocalization *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionLocalization,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotSetCreateRequest defines model for AppScreenshotSetCreateRequest.
type AppScreenshotSetCreateRequest struct {
	Attributes struct {
		ScreenshotDisplayType ScreenshotDisplayType `json:"screenshotDisplayType"`
	} `json:"attributes"`
	Relationships struct {
		AppStoreVersionLocalization struct {
			Data types.RelationshipsData `json:"data"`
		} `json:"appStoreVersionLocalization"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// AppScreenshotSetResponse defines model for AppScreenshotSetResponse.
type AppScreenshotSetResponse struct {
	Data     AppScreenshotSet    `json:"data"`
	Included *[]AppScreenshot    `json:"included,omitempty"`
	Links    types.DocumentLinks `json:"links"`
}

// AppScreenshotSetsResponse defines model for AppScreenshotSetsResponse.
type AppScreenshotSetsResponse struct {
	Data     []AppScreenshotSet       `json:"data"`
	Included *[]AppScreenshot         `json:"included,omitempty"`
	Links    types.PagedDocumentLinks `json:"links"`
	Meta     *types.PagingInformation `json:"meta,omitempty"`
}

// AppScreenshotSetAppScreenshotsLinkagesResponse defines model for AppScreenshotSetAppScreenshotsLinkagesResponse.
type AppScreenshotSetAppScreenshotsLinkagesResponse struct {
	Data  []types.RelationshipsData `json:"data"`
	Links types.PagedDocumentLinks  `json:"links"`
	Meta  *types.PagingInformation  `json:"meta,omitempty"`
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
func (s *Service) GetAppScreenshotSet(id string, params *GetAppScreenshotSetQuery) (*AppScreenshotSetResponse, *services.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	res := new(AppScreenshotSetResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppScreenshotSet adds a new screenshot set to an App Store version localization for a specific screenshot type and display size.
func (s *Service) CreateAppScreenshotSet(id string, body *AppScreenshotSetCreateRequest) (*AppScreenshotSetResponse, *services.Response, error) {
	res := new(AppScreenshotSetResponse)
	resp, err := s.Post("appScreenshotSets", body, res)
	return res, resp, err
}

// DeleteAppScreenshotSet deletes an app screenshot set and all of its screenshots.
func (s *Service) DeleteAppScreenshotSet(id string) (*services.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s", id)
	return s.Delete(url, nil)
}

// ListAppScreenshotsForSet lists all ordered screenshots in a screenshot set.
func (s *Service) ListAppScreenshotsForSet(id string, params *ListAppScreenshotsForSetQuery) (*AppScreenshotsResponse, *services.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/appScreenshots", id)
	res := new(AppScreenshotsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppScreenshotIDsForSet gets the ordered screenshot IDs in a screenshot set.
func (s *Service) ListAppScreenshotIDsForSet(id string, params *ListAppScreenshotIDsForSetQuery) (*AppScreenshotSetAppScreenshotsLinkagesResponse, *services.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	res := new(AppScreenshotSetAppScreenshotsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ReplaceAppScreenshotsForSet changes the order of the screenshots in a screenshot set.
func (s *Service) ReplaceAppScreenshotsForSet(id string, linkages *[]types.RelationshipsData) (*services.Response, error) {
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)
	return s.Patch(url, linkages, nil)
}
