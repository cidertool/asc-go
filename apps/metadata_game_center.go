package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// GameCenterEnabledVersion defines model for GameCenterEnabledVersion.
type GameCenterEnabledVersion struct {
	Attributes *struct {
		IconAsset     *ImageAsset `json:"iconAsset,omitempty"`
		Platform      *Platform   `json:"platform,omitempty"`
		VersionString *string     `json:"versionString,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		CompatibleVersions *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"compatibleVersions,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesRequest defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.
type GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	Data []internal.RelationshipsData `json:"data"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data  []internal.RelationshipsData `json:"data"`
	Links internal.PagedDocumentLinks  `json:"links"`
	Meta  *internal.PagingInformation  `json:"meta,omitempty"`
}

// GameCenterEnabledVersionsResponse defines model for GameCenterEnabledVersionsResponse.
type GameCenterEnabledVersionsResponse struct {
	Data     []GameCenterEnabledVersion  `json:"data"`
	Included *[]GameCenterEnabledVersion `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// ListGameCenterEnabledVersionsForAppQuery are query options for ListGameCenterEnabledVersionsForApp
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

// ListCompatibleVersionsForGameCenterEnabledVersionQuery are query options for ListCompatibleVersionsForGameCenterEnabledVersion
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

// ListCompatibleVersionIDsForGameCenterEnabledVersionQuery are query options for ListCompatibleVersionIDsForGameCenterEnabledVersion
type ListCompatibleVersionIDsForGameCenterEnabledVersionQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// ListGameCenterEnabledVersionsForApp lists the versions for a given app that are enabled for Game Center
func (s *Service) ListGameCenterEnabledVersionsForApp(id string, params *ListGameCenterEnabledVersionsForAppQuery) (*GameCenterEnabledVersionsResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/gameCenterEnabledVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListCompatibleVersionsForGameCenterEnabledVersion lists the versions that are compatible with a given Game Center version
func (s *Service) ListCompatibleVersionsForGameCenterEnabledVersion(id string, params *ListCompatibleVersionsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionsResponse, *internal.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/compatibleVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListCompatibleVersionIDsForGameCenterEnabledVersion lists the version IDs that are compatible with a given Game Center version
func (s *Service) ListCompatibleVersionIDsForGameCenterEnabledVersion(id string, params *ListCompatibleVersionIDsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionCompatibleVersionsLinkagesResponse, *internal.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	res := new(GameCenterEnabledVersionCompatibleVersionsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateCompatibleVersionsForGameCenterEnabledVersion adds a relationship between a given version and a Game Center enabled version
func (s *Service) CreateCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Post(url, body, nil)
}

// UpdateCompatibleVersionsForGameCenterEnabledVersion updates the relationship between a given version and a Game Center enabled version
func (s *Service) UpdateCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Patch(url, body, nil)
}

// RemoveCompatibleVersionsForGameCenterEnabledVersion deletes the relationship between a given version and a Game Center enabled version
func (s *Service) RemoveCompatibleVersionsForGameCenterEnabledVersion(id string, body *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	return s.Delete(url, body)
}
