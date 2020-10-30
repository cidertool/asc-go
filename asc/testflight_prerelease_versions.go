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

// PrereleaseVersion defines model for PrereleaseVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/prereleaseversion
type PrereleaseVersion struct {
	Attributes    *PrereleaseVersionAttributes    `json:"attributes,omitempty"`
	ID            string                          `json:"id"`
	Links         ResourceLinks                   `json:"links"`
	Relationships *PrereleaseVersionRelationships `json:"relationships,omitempty"`
	Type          string                          `json:"type"`
}

// PrereleaseVersionAttributes defines model for PrereleaseVersion.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/prereleaseversion/attributes
type PrereleaseVersionAttributes struct {
	Platform *Platform `json:"platform,omitempty"`
	Version  *string   `json:"version,omitempty"`
}

// PrereleaseVersionRelationships defines model for PrereleaseVersion.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/prereleaseversion/relationships
type PrereleaseVersionRelationships struct {
	App    *Relationship      `json:"app,omitempty"`
	Builds *PagedRelationship `json:"builds,omitempty"`
}

// PrereleaseVersionResponse defines model for PrereleaseVersionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/prereleaseversionresponse
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion                   `json:"data"`
	Included []PrereleaseVersionResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                       `json:"links"`
}

// PrereleaseVersionsResponse defines model for PreReleaseVersionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/prereleaseversionsresponse
type PrereleaseVersionsResponse struct {
	Data     []PrereleaseVersion                 `json:"data"`
	Included []PrereleaseVersionResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                  `json:"links"`
	Meta     *PagingInformation                  `json:"meta,omitempty"`
}

// PrereleaseVersionResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a PrereleaseVersionResponse or PrereleaseVersionsResponse.
type PrereleaseVersionResponseIncluded included

// ListPrereleaseVersionsQuery defines model for ListPrereleaseVersions
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_prerelease_versions
type ListPrereleaseVersionsQuery struct {
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsBuilds                []string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions    []string `url:"fields[preReleaseVersions],omitempty"`
	FilterApp                   []string `url:"filter[app],omitempty"`
	FilterBuilds                []string `url:"filter[builds],omitempty"`
	FilterBuildsExpired         []string `url:"filter[builds.expired],omitempty"`
	FilterBuildsProcessingState []string `url:"filter[builds.processingState],omitempty"`
	FilterPlatform              []string `url:"filter[platform],omitempty"`
	FilterVersion               []string `url:"filter[version],omitempty"`
	Include                     []string `url:"include,omitempty"`
	Sort                        []string `url:"sort,omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	LimitBuilds                 int      `url:"limit[builds],omitempty"`
	Cursor                      string   `url:"cursor,omitempty"`
}

// GetPrereleaseVersionQuery defines model for GetPrereleaseVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_prerelease_version_information
type GetPrereleaseVersionQuery struct {
	FieldsApps               []string `url:"fields[apps],omitempty"`
	FieldsBuilds             []string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
	Include                  []string `url:"include,omitempty"`
	LimitBuilds              int      `url:"limit[builds],omitempty"`
}

// GetAppForPrereleaseVersionQuery defines model for GetAppForPrereleaseVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_prerelease_version
type GetAppForPrereleaseVersionQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListPrereleaseVersionsForAppQuery defines model for ListPrereleaseVersionsForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prerelease_versions_for_an_app
type ListPrereleaseVersionsForAppQuery struct {
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
	Limit                    int      `url:"limit,omitempty"`
	Cursor                   string   `url:"cursor,omitempty"`
}

// ListBuildsForPrereleaseVersionQuery defines model for ListBuildsForPrereleaseVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_a_prerelease_version
type ListBuildsForPrereleaseVersionQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// GetPrereleaseVersionForBuildQuery defines model for GetPrereleaseVersionForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_prerelease_version_of_a_build
type GetPrereleaseVersionForBuildQuery struct {
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
}

// ListPrereleaseVersions gets a list of prerelease versions for all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_prerelease_versions
func (s *TestflightService) ListPrereleaseVersions(ctx context.Context, params *ListPrereleaseVersionsQuery) (*PrereleaseVersionsResponse, *Response, error) {
	res := new(PrereleaseVersionsResponse)
	resp, err := s.client.get(ctx, "preReleaseVersions", params, res)

	return res, resp, err
}

// GetPrereleaseVersion gets information about a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_prerelease_version_information
func (s *TestflightService) GetPrereleaseVersion(ctx context.Context, id string, params *GetPrereleaseVersionQuery) (*PrereleaseVersionResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppForPrereleaseVersion gets the app information for a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_prerelease_version
func (s *TestflightService) GetAppForPrereleaseVersion(ctx context.Context, id string, params *GetAppForPrereleaseVersionQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListPrereleaseVersionsForApp gets a list of prerelease versions associated with a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prerelease_versions_for_an_app
func (s *TestflightService) ListPrereleaseVersionsForApp(ctx context.Context, id string, params *ListPrereleaseVersionsForAppQuery) (*PrereleaseVersionsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/preReleaseVersions", id)
	res := new(PrereleaseVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBuildsForPrereleaseVersion gets a list of builds of a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_a_prerelease_version
func (s *TestflightService) ListBuildsForPrereleaseVersion(ctx context.Context, id string, params *ListBuildsForPrereleaseVersionQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetPrereleaseVersionForBuild gets the prerelease version for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_prerelease_version_of_a_build
func (s *TestflightService) GetPrereleaseVersionForBuild(ctx context.Context, id string, params *GetPrereleaseVersionForBuildQuery) (*PrereleaseVersionResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/preReleaseVersion", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in PrereleaseVersionResponseIncluded.
func (i *PrereleaseVersionResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// Build returns the Build stored within, if one is present.
func (i *PrereleaseVersionResponseIncluded) Build() *Build {
	return extractIncludedBuild(i.inner)
}

// App returns the App stored within, if one is present.
func (i *PrereleaseVersionResponseIncluded) App() *App {
	return extractIncludedApp(i.inner)
}
