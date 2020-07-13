package builds

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// IconAssetType defines model for IconAssetType.
type IconAssetType string

// List of IconAssetType
const (
	AppStore         IconAssetType = "APP_STORE"
	MessagesAppStore IconAssetType = "MESSAGES_APP_STORE"
	TVOSHomeScreen   IconAssetType = "TV_OS_HOME_SCREEN"
	TVOSTopShelf     IconAssetType = "TV_OS_TOP_SHELF"
	WatchAppStore    IconAssetType = "WATCH_APP_STORE"
)

// BuildIcon defines model for BuildIcon.
type BuildIcon struct {
	Attributes *struct {
		IconAsset *apps.ImageAsset `json:"iconAsset,omitempty"`
		IconType  *IconAssetType   `json:"iconType,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// BuildIconsResponse defines model for BuildIconsResponse.
type BuildIconsResponse struct {
	Data  []BuildIcon               `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type ListIconsOptions struct {
	Fields *struct {
		BuildIcons *[]string `url:"buildIcons,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// ListIconsForBuild lists all the icons for various platforms delivered with a build.
func (s *Service) ListIconsForBuild(id string, params *ListIconsOptions) (*BuildIconsResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/icons", id)
	res := new(BuildIconsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
