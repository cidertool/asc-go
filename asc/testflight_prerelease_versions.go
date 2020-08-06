package asc

import (
	"fmt"
)

// PrereleaseVersion defines model for PrereleaseVersion.
type PrereleaseVersion struct {
	Attributes *struct {
		Platform *Platform `json:"platform,omitempty"`
		Version  *string   `json:"version,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		Builds *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// PrereleaseVersionResponse defines model for PrereleaseVersionResponse.
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion `json:"data"`
	Included *[]interface{}    `json:"included,omitempty"`
	Links    DocumentLinks     `json:"links"`
}

// PrereleaseVersionsResponse defines model for PreReleaseVersionsResponse.
type PrereleaseVersionsResponse struct {
	Data     []PrereleaseVersion `json:"data"`
	Included *[]interface{}      `json:"included,omitempty"`
	Links    PagedDocumentLinks  `json:"links"`
	Meta     *PagingInformation  `json:"meta,omitempty"`
}

// ListPrereleaseVersionsQuery defines model for ListPrereleaseVersions
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
type GetPrereleaseVersionQuery struct {
	FieldsApps               []string `url:"fields[apps],omitempty"`
	FieldsBuilds             []string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
	Include                  []string `url:"include,omitempty"`
	LimitBuilds              int      `url:"limit[builds],omitempty"`
}

// GetAppForPrereleaseVersionQuery defines model for GetAppForPrereleaseVersion
type GetAppForPrereleaseVersionQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListPrereleaseVersionsForAppQuery defines model for ListPrereleaseVersionsForApp
type ListPrereleaseVersionsForAppQuery struct {
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
	Limit                    int      `url:"limit,omitempty"`
	Cursor                   string   `url:"cursor,omitempty"`
}

// ListBuildsForPrereleaseVersionQuery defines model for ListBuildsForPrereleaseVersion
type ListBuildsForPrereleaseVersionQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// GetPrereleaseVersionForBuildQuery defines model for GetPrereleaseVersionForBuild
type GetPrereleaseVersionForBuildQuery struct {
	FieldsPreReleaseVersions []string `url:"fields[preReleaseVersions],omitempty"`
}

// ListPrereleaseVersions gets a list of prerelease versions for all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_prerelease_versions
func (s *TestflightService) ListPrereleaseVersions(params *ListPrereleaseVersionsQuery) (*PrereleaseVersionsResponse, *Response, error) {
	res := new(PrereleaseVersionsResponse)
	resp, err := s.client.get("preReleaseVersions", params, res)
	return res, resp, err
}

// GetPrereleaseVersion gets information about a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_prerelease_version_information
func (s *TestflightService) GetPrereleaseVersion(id string, params *GetPrereleaseVersionQuery) (*PrereleaseVersionResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForPrereleaseVersion gets the app information for a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_prerelease_version
func (s *TestflightService) GetAppForPrereleaseVersion(id string, params *GetAppForPrereleaseVersionQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListPrereleaseVersionsForApp gets a list of prerelease versions associated with a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prerelease_versions_for_an_app
func (s *TestflightService) ListPrereleaseVersionsForApp(id string, params *ListPrereleaseVersionsForAppQuery) (*PrereleaseVersionsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/preReleaseVersions", id)
	res := new(PrereleaseVersionsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBuildsForPrereleaseVersion gets a list of builds of a specific prerelease version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_a_prerelease_version
func (s *TestflightService) ListBuildsForPrereleaseVersion(id string, params *ListBuildsForPrereleaseVersionQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetPrereleaseVersionForBuild gets the prerelease version for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_prerelease_version_of_a_build
func (s *TestflightService) GetPrereleaseVersionForBuild(id string, params *GetPrereleaseVersionForBuildQuery) (*PrereleaseVersionResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/preReleaseVersion", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
