/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// ScreenshotDisplayType defines model for ScreenshotDisplayType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/screenshotdisplaytype
type ScreenshotDisplayType string

const (
	// ScreenshotDisplayTypeAppAppleTV is a screenshot display type for AppAppleTV.
	ScreenshotDisplayTypeAppAppleTV ScreenshotDisplayType = "APP_APPLE_TV"
	// ScreenshotDisplayTypeAppDesktop is a screenshot display type for AppDesktop.
	ScreenshotDisplayTypeAppDesktop ScreenshotDisplayType = "APP_DESKTOP"
	// ScreenshotDisplayTypeAppiPad105 is a screenshot display type for AppiPad105.
	ScreenshotDisplayTypeAppiPad105 ScreenshotDisplayType = "APP_IPAD_105"
	// ScreenshotDisplayTypeAppiPad97 is a screenshot display type for AppiPad97.
	ScreenshotDisplayTypeAppiPad97 ScreenshotDisplayType = "APP_IPAD_97"
	// ScreenshotDisplayTypeAppiPadPro129 is a screenshot display type for AppiPadPro129.
	ScreenshotDisplayTypeAppiPadPro129 ScreenshotDisplayType = "APP_IPAD_PRO_129"
	// ScreenshotDisplayTypeAppiPadPro3Gen11 is a screenshot display type for AppiPadPro3Gen11.
	ScreenshotDisplayTypeAppiPadPro3Gen11 ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_11"
	// ScreenshotDisplayTypeAppiPadPro3Gen129 is a screenshot display type for AppiPadPro3Gen129.
	ScreenshotDisplayTypeAppiPadPro3Gen129 ScreenshotDisplayType = "APP_IPAD_PRO_3GEN_129"
	// ScreenshotDisplayTypeAppiPhone35 is a screenshot display type for AppiPhone35.
	ScreenshotDisplayTypeAppiPhone35 ScreenshotDisplayType = "APP_IPHONE_35"
	// ScreenshotDisplayTypeAppiPhone40 is a screenshot display type for AppiPhone40.
	ScreenshotDisplayTypeAppiPhone40 ScreenshotDisplayType = "APP_IPHONE_40"
	// ScreenshotDisplayTypeAppiPhone47 is a screenshot display type for AppiPhone47.
	ScreenshotDisplayTypeAppiPhone47 ScreenshotDisplayType = "APP_IPHONE_47"
	// ScreenshotDisplayTypeAppiPhone55 is a screenshot display type for AppiPhone55.
	ScreenshotDisplayTypeAppiPhone55 ScreenshotDisplayType = "APP_IPHONE_55"
	// ScreenshotDisplayTypeAppiPhone58 is a screenshot display type for AppiPhone58.
	ScreenshotDisplayTypeAppiPhone58 ScreenshotDisplayType = "APP_IPHONE_58"
	// ScreenshotDisplayTypeAppiPhone65 is a screenshot display type for AppiPhone65.
	ScreenshotDisplayTypeAppiPhone65 ScreenshotDisplayType = "APP_IPHONE_65"
	// ScreenshotDisplayTypeAppWatchSeries3 is a screenshot display type for AppWatchSeries3.
	ScreenshotDisplayTypeAppWatchSeries3 ScreenshotDisplayType = "APP_WATCH_SERIES_3"
	// ScreenshotDisplayTypeAppWatchSeries4 is a screenshot display type for AppWatchSeries4.
	ScreenshotDisplayTypeAppWatchSeries4 ScreenshotDisplayType = "APP_WATCH_SERIES_4"
	// ScreenshotDisplayTypeiMessageAppIPad105 is a screenshot display type for iMessageAppIPad105.
	ScreenshotDisplayTypeiMessageAppIPad105 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_105"
	// ScreenshotDisplayTypeiMessageAppIPad97 is a screenshot display type for iMessageAppIPad97.
	ScreenshotDisplayTypeiMessageAppIPad97 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_97"
	// ScreenshotDisplayTypeiMessageAppIPadPro129 is a screenshot display type for iMessageAppIPadPro129.
	ScreenshotDisplayTypeiMessageAppIPadPro129 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_129"
	// ScreenshotDisplayTypeiMessageAppIPadPro3Gen11 is a screenshot display type for iMessageAppIPadPro3Gen11.
	ScreenshotDisplayTypeiMessageAppIPadPro3Gen11 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_11"
	// ScreenshotDisplayTypeiMessageAppIPadPro3Gen129 is a screenshot display type for iMessageAppIPadPro3Gen129.
	ScreenshotDisplayTypeiMessageAppIPadPro3Gen129 ScreenshotDisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_129"
	// ScreenshotDisplayTypeiMessageAppIPhone40 is a screenshot display type for iMessageAppIPhone40.
	ScreenshotDisplayTypeiMessageAppIPhone40 ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_40"
	// ScreenshotDisplayTypeiMessageAppIPhone47 is a screenshot display type for iMessageAppIPhone47.
	ScreenshotDisplayTypeiMessageAppIPhone47 ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_47"
	// ScreenshotDisplayTypeiMessageAppIPhone55 is a screenshot display type for iMessageAppIPhone55.
	ScreenshotDisplayTypeiMessageAppIPhone55 ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_55"
	// ScreenshotDisplayTypeiMessageAppIPhone58 is a screenshot display type for iMessageAppIPhone58.
	ScreenshotDisplayTypeiMessageAppIPhone58 ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_58"
	// ScreenshotDisplayTypeiMessageAppIPhone65 is a screenshot display type for iMessageAppIPhone65.
	ScreenshotDisplayTypeiMessageAppIPhone65 ScreenshotDisplayType = "IMESSAGE_APP_IPHONE_65"
)

// AppScreenshotSet defines model for AppScreenshotSet.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotset
type AppScreenshotSet struct {
	Attributes    *AppScreenshotSetAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Links         ResourceLinks                  `json:"links"`
	Relationships *AppScreenshotSetRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

// AppScreenshotSetAttributes defines model for AppScreenshotSet.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotset/attributes
type AppScreenshotSetAttributes struct {
	ScreenshotDisplayType *ScreenshotDisplayType `json:"screenshotDisplayType,omitempty"`
}

// AppScreenshotSetRelationships defines model for AppScreenshotSet.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotset/relationships
type AppScreenshotSetRelationships struct {
	AppScreenshots              *PagedRelationship `json:"appScreenshots,omitempty"`
	AppStoreVersionLocalization *Relationship      `json:"appStoreVersionLocalization,omitempty"`
}

// appScreenshotSetCreateRequest defines model for AppScreenshotSetCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest/data
type appScreenshotSetCreateRequest struct {
	Attributes    appScreenshotSetCreateRequestAttributes    `json:"attributes"`
	Relationships appScreenshotSetCreateRequestRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

// appScreenshotSetCreateRequestAttributes are attributes for AppScreenshotSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest/data/attributes
type appScreenshotSetCreateRequestAttributes struct {
	ScreenshotDisplayType ScreenshotDisplayType `json:"screenshotDisplayType"`
}

// appScreenshotSetCreateRequestRelationships are relationships for AppScreenshotSetCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetcreaterequest/data/relationships
type appScreenshotSetCreateRequestRelationships struct {
	AppStoreVersionLocalization relationshipDeclaration `json:"appStoreVersionLocalization"`
}

// AppScreenshotSetResponse defines model for AppScreenshotSetResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetresponse
type AppScreenshotSetResponse struct {
	Data     AppScreenshotSet `json:"data"`
	Included []AppScreenshot  `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}

// AppScreenshotSetsResponse defines model for AppScreenshotSetsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetsresponse
type AppScreenshotSetsResponse struct {
	Data     []AppScreenshotSet `json:"data"`
	Included []AppScreenshot    `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppScreenshotSetAppScreenshotsLinkagesResponse defines model for AppScreenshotSetAppScreenshotsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsetappscreenshotslinkagesresponse
type AppScreenshotSetAppScreenshotsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
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
func (s *AppsService) CreateAppScreenshotSet(ctx context.Context, screenshotDisplayType ScreenshotDisplayType, appStoreVersionLocalizationID string) (*AppScreenshotSetResponse, *Response, error) {
	req := appScreenshotSetCreateRequest{
		Attributes: appScreenshotSetCreateRequestAttributes{
			ScreenshotDisplayType: screenshotDisplayType,
		},
		Relationships: appScreenshotSetCreateRequestRelationships{
			AppStoreVersionLocalization: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionLocalizationID,
					Type: "appStoreVersionLocalizations",
				},
			},
		},
		Type: "appScreenshotSets",
	}
	res := new(AppScreenshotSetResponse)
	resp, err := s.client.post(ctx, "appScreenshotSets", newRequestBody(req), res)

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
func (s *AppsService) ReplaceAppScreenshotsForSet(ctx context.Context, id string, appScreenshotIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appScreenshotIDs, "appScreenshots")
	url := fmt.Sprintf("appScreenshotSets/%s/relationships/appScreenshots", id)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}
