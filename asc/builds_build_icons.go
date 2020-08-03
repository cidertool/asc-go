package asc

import (
	"fmt"
	"net/http"
)

// IconAssetType defines model for IconAssetType.
type IconAssetType string

// List of IconAssetType
const (
	IconAssetTypeAppStore         IconAssetType = "APP_STORE"
	IconAssetTypeMessagesAppStore IconAssetType = "MESSAGES_APP_STORE"
	IconAssetTypeTVOSHomeScreen   IconAssetType = "TV_OS_HOME_SCREEN"
	IconAssetTypeTVOSTopShelf     IconAssetType = "TV_OS_TOP_SHELF"
	IconAssetTypeWatchAppStore    IconAssetType = "WATCH_APP_STORE"
)

// BuildIcon defines model for BuildIcon.
type BuildIcon struct {
	Attributes *struct {
		IconAsset *ImageAsset    `json:"iconAsset,omitempty"`
		IconType  *IconAssetType `json:"iconType,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// BuildIconsResponse defines model for BuildIconsResponse.
type BuildIconsResponse struct {
	Data  []BuildIcon        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListIconsQuery are query options for ListIcons
type ListIconsQuery struct {
	FieldsBuildIcons *[]string `url:"fields[buildIcons],omitempty"`
	Limit            *int      `url:"limit,omitempty"`
	Cursor           *string   `url:"cursor,omitempty"`
}

// ListIconsForBuild lists all the icons for various platforms delivered with a build.
func (s *BuildsService) ListIconsForBuild(id string, params *ListIconsQuery) (*BuildIconsResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/icons", id)
	res := new(BuildIconsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
