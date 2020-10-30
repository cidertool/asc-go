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

// IconAssetType defines model for IconAssetType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/iconassettype
type IconAssetType string

const (
	// IconAssetTypeAppStore is an icon asset type for AppStore.
	IconAssetTypeAppStore IconAssetType = "APP_STORE"
	// IconAssetTypeMessagesAppStore is an icon asset type for MessagesAppStore.
	IconAssetTypeMessagesAppStore IconAssetType = "MESSAGES_APP_STORE"
	// IconAssetTypeTVOSHomeScreen is an icon asset type for TVOSHomeScreen.
	IconAssetTypeTVOSHomeScreen IconAssetType = "TV_OS_HOME_SCREEN"
	// IconAssetTypeTVOSTopShelf is an icon asset type for TVOSTopShelf.
	IconAssetTypeTVOSTopShelf IconAssetType = "TV_OS_TOP_SHELF"
	// IconAssetTypeWatchAppStore is an icon asset type for WatchAppStore.
	IconAssetTypeWatchAppStore IconAssetType = "WATCH_APP_STORE"
)

// BuildIcon defines model for BuildIcon.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildicon
type BuildIcon struct {
	Attributes *BuildIconAttributes `json:"attributes,omitempty"`
	ID         string               `json:"id"`
	Links      ResourceLinks        `json:"links"`
	Type       string               `json:"type"`
}

// BuildIconAttributes defines model for BuildIcon.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildicon/attributes
type BuildIconAttributes struct {
	IconAsset *ImageAsset    `json:"iconAsset,omitempty"`
	IconType  *IconAssetType `json:"iconType,omitempty"`
}

// BuildIconsResponse defines model for BuildIconsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildiconsresponse
type BuildIconsResponse struct {
	Data  []BuildIcon        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListIconsQuery are query options for ListIcons
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_icons_for_a_build
type ListIconsQuery struct {
	FieldsBuildIcons []string `url:"fields[buildIcons],omitempty"`
	Limit            int      `url:"limit,omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListIconsForBuild lists all the icons for various platforms delivered with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_icons_for_a_build
func (s *BuildsService) ListIconsForBuild(ctx context.Context, id string, params *ListIconsQuery) (*BuildIconsResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/icons", id)
	res := new(BuildIconsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
