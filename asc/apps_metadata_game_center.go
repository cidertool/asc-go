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

// GameCenterEnabledVersion defines model for GameCenterEnabledVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterenabledversion
type GameCenterEnabledVersion struct {
	Attributes    *GameCenterEnabledVersionAttributes    `json:"attributes,omitempty"`
	ID            string                                 `json:"id"`
	Links         ResourceLinks                          `json:"links"`
	Relationships *GameCenterEnabledVersionRelationships `json:"relationships,omitempty"`
	Type          string                                 `json:"type"`
}

// GameCenterEnabledVersionAttributes defines model for GameCenterEnabledVersion.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterenabledversion/attributes
type GameCenterEnabledVersionAttributes struct {
	IconAsset     *ImageAsset `json:"iconAsset,omitempty"`
	Platform      *Platform   `json:"platform,omitempty"`
	VersionString *string     `json:"versionString,omitempty"`
}

// GameCenterEnabledVersionRelationships defines model for GameCenterEnabledVersion.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterenabledversion/relationships
type GameCenterEnabledVersionRelationships struct {
	App                *Relationship      `json:"app,omitempty"`
	CompatibleVersions *PagedRelationship `json:"compatibleVersions,omitempty"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse defines model for GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterenabledversioncompatibleversionslinkagesresponse
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GameCenterEnabledVersionsResponse defines model for GameCenterEnabledVersionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterenabledversionsresponse
type GameCenterEnabledVersionsResponse struct {
	Data     []GameCenterEnabledVersion `json:"data"`
	Included []GameCenterEnabledVersion `json:"included,omitempty"`
	Links    PagedDocumentLinks         `json:"links"`
	Meta     *PagingInformation         `json:"meta,omitempty"`
}

// ListGameCenterEnabledVersionsForAppQuery are query options for ListGameCenterEnabledVersionsForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_compatible_versions_for_a_game_center_enabled_version
type ListGameCenterEnabledVersionsForAppQuery struct {
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FieldsGameCenterEnabledVersions []string `url:"fields[gameCenterEnabledVersions],omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	Include                         []string `url:"include,omitempty"`
	Sort                            []string `url:"sort,omitempty"`
	FilterID                        []string `url:"filter[id],omitempty"`
	FilterPlatform                  []string `url:"filter[platform],omitempty"`
	FilterVersionString             []string `url:"filter[versionString],omitempty"`
	Cursor                          string   `url:"cursor,omitempty"`
}

// ListCompatibleVersionsForGameCenterEnabledVersionQuery are query options for ListCompatibleVersionsForGameCenterEnabledVersion.
type ListCompatibleVersionsForGameCenterEnabledVersionQuery struct {
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FieldsGameCenterEnabledVersions []string `url:"fields[gameCenterEnabledVersions],omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	Include                         []string `url:"include,omitempty"`
	Sort                            []string `url:"sort,omitempty"`
	FilterApp                       []string `url:"filter[app],omitempty"`
	FilterID                        []string `url:"filter[id],omitempty"`
	FilterPlatform                  []string `url:"filter[platform],omitempty"`
	FilterVersionString             []string `url:"filter[versionString],omitempty"`
	Cursor                          string   `url:"cursor,omitempty"`
}

// ListCompatibleVersionIDsForGameCenterEnabledVersionQuery are query options for ListCompatibleVersionIDsForGameCenterEnabledVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_compatible_version_ids_for_a_game_center_enabled_version
type ListCompatibleVersionIDsForGameCenterEnabledVersionQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListGameCenterEnabledVersionsForApp lists the versions for a given app that are enabled for Game Center
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_game_center_enabled_versions_for_an_app
func (s *AppsService) ListGameCenterEnabledVersionsForApp(ctx context.Context, id string, params *ListGameCenterEnabledVersionsForAppQuery) (*GameCenterEnabledVersionsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/gameCenterEnabledVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListCompatibleVersionsForGameCenterEnabledVersion lists the versions that are compatible with a given Game Center version
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_compatible_versions_for_a_game_center_enabled_version
func (s *AppsService) ListCompatibleVersionsForGameCenterEnabledVersion(ctx context.Context, id string, params *ListCompatibleVersionsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/compatibleVersions", id)
	res := new(GameCenterEnabledVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListCompatibleVersionIDsForGameCenterEnabledVersion lists the version IDs that are compatible with a given Game Center version
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_compatible_version_ids_for_a_game_center_enabled_version
func (s *AppsService) ListCompatibleVersionIDsForGameCenterEnabledVersion(ctx context.Context, id string, params *ListCompatibleVersionIDsForGameCenterEnabledVersionQuery) (*GameCenterEnabledVersionCompatibleVersionsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)
	res := new(GameCenterEnabledVersionCompatibleVersionsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateCompatibleVersionsForGameCenterEnabledVersion adds a relationship between a given version and a Game Center enabled version
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_compatible_versions_to_a_game_center_enabled_version
func (s *AppsService) CreateCompatibleVersionsForGameCenterEnabledVersion(ctx context.Context, id string, gameCenterCompatibleVersionIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterCompatibleVersionIDs, "gameCenterEnabledVersions")
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)

	return s.client.post(ctx, url, newRequestBody(linkages.Data), nil)
}

// UpdateCompatibleVersionsForGameCenterEnabledVersion updates the relationship between a given version and a Game Center enabled version
//
// https://developer.apple.com/documentation/appstoreconnectapi/replace_all_compatible_versions_for_a_game_center_enabled_version
func (s *AppsService) UpdateCompatibleVersionsForGameCenterEnabledVersion(ctx context.Context, id string, gameCenterCompatibleVersionIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterCompatibleVersionIDs, "gameCenterEnabledVersions")
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}

// RemoveCompatibleVersionsForGameCenterEnabledVersion deletes the relationship between a given version and a Game Center enabled version
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_compatible_versions_from_a_game_center_enabled_version
func (s *AppsService) RemoveCompatibleVersionsForGameCenterEnabledVersion(ctx context.Context, id string, gameCenterCompatibleVersionIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterCompatibleVersionIDs, "gameCenterEnabledVersions")
	url := fmt.Sprintf("gameCenterEnabledVersions/%s/relationships/compatibleVersions", id)

	return s.client.delete(ctx, url, newRequestBody(linkages.Data))
}
