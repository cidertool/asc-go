package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// GameCenterEnabledVersion defines model for GameCenterEnabledVersion.
type GameCenterEnabledVersion struct {
	Attributes *struct {
		IconAsset     *ImageAsset `json:"iconAsset,omitempty"`
		Platform      *Platform   `json:"platform,omitempty"`
		VersionString *string     `json:"versionString,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"app,omitempty"`
		CompatibleVersions *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"compatibleVersions,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesRequest defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.
type GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// GameCenterEnabledVersionsResponse defines model for GameCenterEnabledVersionsResponse.
type GameCenterEnabledVersionsResponse struct {
	Data     []GameCenterEnabledVersion  `json:"data"`
	Included *[]GameCenterEnabledVersion `json:"included,omitempty"`
	Links    common.PagedDocumentLinks   `json:"links"`
	Meta     *common.PagingInformation   `json:"meta,omitempty"`
}

type ListGameCenterEnabledVersionsForAppQuery struct {
	Fields *struct {
		Apps                      *[]string `url:"apps,omitempty"`
		GameCenterEnabledVersions *[]string `url:"gameCenterEnabledVersions,omitempty"`
	} `url:"fields,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Sort    *[]string `url:"sort,omitempty"`
	Filter  *struct {
		ID            *[]string `url:"id,omitempty"`
		Platform      *[]string `url:"platform,omitempty"`
		VersionString *[]string `url:"versionString,omitempty"`
	} `url:"filter,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListCompatibleVersionsForGameCenterEnabledVersionQuery struct {
	Fields *struct {
		Apps                      *[]string `url:"apps,omitempty"`
		GameCenterEnabledVersions *[]string `url:"gameCenterEnabledVersions,omitempty"`
	} `url:"fields,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Sort    *[]string `url:"sort,omitempty"`
	Filter  *struct {
		App           *[]string `url:"app,omitempty"`
		ID            *[]string `url:"id,omitempty"`
		Platform      *[]string `url:"platform,omitempty"`
		VersionString *[]string `url:"versionString,omitempty"`
	} `url:"filter,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListCompatibleVersionIDsForGameCenterEnabledVersionQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

func (s *Service) ListGameCenterEnabledVersionsForApp(id string, params *ListGameCenterEnabledVersionsForAppQuery) (*GameCenterEnabledVersionsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/gameCenterEnabledVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

func (s *Service) ListCompatibleVersionsForGameCenterEnabledVersion(id string, params *ListCompatibleVersionsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionsResponse, *common.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/compatibleVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

func (s *Service) ListCompatibleVersionIDsForGameCenterEnabledVersion(id string, params *ListCompatibleVersionIDsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionCompatibleVersionsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	res := new(GameCenterEnabledVersionCompatibleVersionsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

func (s *Service) CreateCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Post(url, body, nil)
}

func (s *Service) UpdateCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Patch(url, body, nil)
}

func (s *Service) RemoveCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Delete(url, body)
}
