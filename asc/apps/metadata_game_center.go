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
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		CompatibleVersions *struct {
			Data  *[]common.RelationshipsData `json:"data,omitempty"`
			Links *common.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *common.PagingInformation   `json:"meta,omitempty"`
		} `json:"compatibleVersions,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesRequest defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.
type GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	Data []common.RelationshipsData `json:"data"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data  []common.RelationshipsData `json:"data"`
	Links common.PagedDocumentLinks  `json:"links"`
	Meta  *common.PagingInformation  `json:"meta,omitempty"`
}

// GameCenterEnabledVersionsResponse defines model for GameCenterEnabledVersionsResponse.
type GameCenterEnabledVersionsResponse struct {
	Data     []GameCenterEnabledVersion  `json:"data"`
	Included *[]GameCenterEnabledVersion `json:"included,omitempty"`
	Links    common.PagedDocumentLinks   `json:"links"`
	Meta     *common.PagingInformation   `json:"meta,omitempty"`
}

type ListGameCenterEnabledVersionsForAppQuery struct {
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FieldsGameCenterEnabledVersions *[]string `url:"fields[gameCenterEnabledVersions],omitempty"`
	Limit                           *int      `url:"limit,omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	Sort                            *[]string `url:"sort,omitempty"`
	FilterID                        *[]string `url:"filter[id],omitempty"`
	FilterPlatform                  *[]string `url:"filter[platform],omitempty"`
	FilterVersionString             *[]string `url:"filter[versionString],omitempty"`
	Cursor                          *string   `url:"cursor,omitempty"`
}

type ListCompatibleVersionsForGameCenterEnabledVersionQuery struct {
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FieldsGameCenterEnabledVersions *[]string `url:"fields[gameCenterEnabledVersions],omitempty"`
	Limit                           *int      `url:"limit,omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	Sort                            *[]string `url:"sort,omitempty"`
	FilterApp                       *[]string `url:"filter[app],omitempty"`
	FilterID                        *[]string `url:"filter[id],omitempty"`
	FilterPlatform                  *[]string `url:"filter[platform],omitempty"`
	FilterVersionString             *[]string `url:"filter[versionString],omitempty"`
	Cursor                          *string   `url:"cursor,omitempty"`
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
