package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// PrereleaseVersion defines model for PrereleaseVersion.
type PrereleaseVersion struct {
	Attributes *struct {
		Platform *apps.Platform `json:"platform,omitempty"`
		Version  *string        `json:"version,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		Builds *struct {
			Data  *[]types.RelationshipsData `json:"data,omitempty"`
			Links *types.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *types.PagingInformation   `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// PrereleaseVersionResponse defines model for PrereleaseVersionResponse.
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion   `json:"data"`
	Included *[]interface{}      `json:"included,omitempty"`
	Links    types.DocumentLinks `json:"links"`
}

// PrereleaseVersionsResponse defines model for PreReleaseVersionsResponse.
type PrereleaseVersionsResponse struct {
	Data     []PrereleaseVersion      `json:"data"`
	Included *[]interface{}           `json:"included,omitempty"`
	Links    types.PagedDocumentLinks `json:"links"`
	Meta     *types.PagingInformation `json:"meta,omitempty"`
}

// ListPrereleaseVersionsQuery defines model for ListPrereleaseVersions
type ListPrereleaseVersionsQuery struct {
	FieldsApps                  *[]string `url:"fields[apps],omitempty"`
	FieldsBuilds                *[]string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions    *[]string `url:"fields[preReleaseVersions],omitempty"`
	FilterApp                   *[]string `url:"filter[app],omitempty"`
	FilterBuilds                *[]string `url:"filter[builds],omitempty"`
	FilterBuildsExpired         *[]string `url:"filter[builds.expired],omitempty"`
	FilterBuildsProcessingState *[]string `url:"filter[builds.processingState],omitempty"`
	FilterPlatform              *[]string `url:"filter[platform],omitempty"`
	FilterVersion               *[]string `url:"filter[version],omitempty"`
	Include                     *[]string `url:"include,omitempty"`
	Sort                        *[]string `url:"sort,omitempty"`
	Limit                       *int      `url:"limit,omitempty"`
	LimitBuilds                 *int      `url:"limit[builds],omitempty"`
	Cursor                      *string   `url:"cursor,omitempty"`
}

// GetPrereleaseVersionQuery defines model for GetPrereleaseVersion
type GetPrereleaseVersionQuery struct {
	FieldsApps               *[]string `url:"fields[apps],omitempty"`
	FieldsBuilds             *[]string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions *[]string `url:"fields[preReleaseVersions],omitempty"`
	Include                  *[]string `url:"include,omitempty"`
	LimitBuilds              *int      `url:"limit[builds],omitempty"`
}

// GetAppForPrereleaseVersionQuery defines model for GetAppForPrereleaseVersion
type GetAppForPrereleaseVersionQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// ListPrereleaseVersionsForAppQuery defines model for ListPrereleaseVersionsForApp
type ListPrereleaseVersionsForAppQuery struct {
	FieldsPreReleaseVersions *[]string `url:"fields[preReleaseVersions],omitempty"`
	Limit                    *int      `url:"limit,omitempty"`
	Cursor                   *string   `url:"cursor,omitempty"`
}

// ListBuildsForPrereleaseVersionQuery defines model for ListBuildsForPrereleaseVersion
type ListBuildsForPrereleaseVersionQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
	Limit        *int      `url:"limit,omitempty"`
	Cursor       *string   `url:"cursor,omitempty"`
}

// GetPrereleaseVersionForBuildQuery defines model for GetPrereleaseVersionForBuild
type GetPrereleaseVersionForBuildQuery struct {
	FieldsPreReleaseVersions *[]string `url:"fields[preReleaseVersions],omitempty"`
}

// ListPrereleaseVersions gets a list of prerelease versions for all apps.
func (s *Service) ListPrereleaseVersions(params *ListPrereleaseVersionsQuery) (*PrereleaseVersionsResponse, *services.Response, error) {
	res := new(PrereleaseVersionsResponse)
	resp, err := s.GetWithQuery("preReleaseVersions", params, res)
	return res, resp, err
}

// GetPrereleaseVersion gets information about a specific prerelease version.
func (s *Service) GetPrereleaseVersion(id string, params *GetPrereleaseVersionQuery) (*PrereleaseVersionResponse, *services.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForPrereleaseVersion gets the app information for a specific prerelease version.
func (s *Service) GetAppForPrereleaseVersion(id string, params *GetAppForPrereleaseVersionQuery) (*apps.AppResponse, *services.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListPrereleaseVersionsForApp gets a list of prerelease versions associated with a specific app.
func (s *Service) ListPrereleaseVersionsForApp(id string, params *ListPrereleaseVersionsForAppQuery) (*PrereleaseVersionsResponse, *services.Response, error) {
	url := fmt.Sprintf("apps/%s/preReleaseVersions", id)
	res := new(PrereleaseVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBuildsForPrereleaseVersion gets a list of builds of a specific prerelease version.
func (s *Service) ListBuildsForPrereleaseVersion(id string, params *ListBuildsForPrereleaseVersionQuery) (*builds.BuildsResponse, *services.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPrereleaseVersionForBuild gets the prerelease version for a specific build.
func (s *Service) GetPrereleaseVersionForBuild(id string, params *GetPrereleaseVersionForBuildQuery) (*PrereleaseVersionResponse, *services.Response, error) {
	url := fmt.Sprintf("builds/%s/preReleaseVersion", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
