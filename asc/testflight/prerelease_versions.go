package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// PrereleaseVersion defines model for PrereleaseVersion.
type PrereleaseVersion struct {
	Attributes *struct {
		Platform *apps.Platform `json:"platform,omitempty"`
		Version  *string        `json:"version,omitempty"`
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
		Builds *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// PrereleaseVersionResponse defines model for PrereleaseVersionResponse.
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion    `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// PrereleaseVersionsResponse defines model for PreReleaseVersionsResponse.
type PrereleaseVersionsResponse struct {
	Data     []PrereleaseVersion       `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListPrereleaseVersionsQuery struct {
	Fields *struct {
		Apps               *[]string `url:"apps,omitempty"`
		Builds             *[]string `url:"builds,omitempty"`
		PreReleaseVersions *[]string `url:"preReleaseVersions,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App                   *[]string `url:"app,omitempty"`
		Builds                *[]string `url:"builds,omitempty"`
		BuildsExpired         *[]string `url:"builds.expired,omitempty"`
		BuildsProcessingState *[]string `url:"builds.processingState,omitempty"`
		Platform              *[]string `url:"platform,omitempty"`
		Version               *[]string `url:"version,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	Sort     *[]string `url:"sort,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		Builds *int `url:"builds,omitempty"`
	} `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type GetPrereleaseVersionQuery struct {
	Fields *struct {
		Apps               *[]string `url:"apps,omitempty"`
		Builds             *[]string `url:"builds,omitempty"`
		PreReleaseVersions *[]string `url:"preReleaseVersions,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		Builds *int `url:"builds,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAppForPrereleaseVersionQuery struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type ListPrereleaseVersionsForAppQuery struct {
	Fields *struct {
		PreReleaseVersions *[]string `url:"preReleaseVersions,omitempty"`
	} `url:"fields,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListBuildsForPrereleaseVersionQuery struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type GetPrereleaseVersionForBuildQuery struct {
	Fields *struct {
		PreReleaseVersions *[]string `url:"preReleaseVersions,omitempty"`
	} `url:"fields,omitempty"`
}

// ListPrereleaseVersions gets a list of prerelease versions for all apps.
func (s *Service) ListPrereleaseVersions(params *ListPrereleaseVersionsQuery) (*PrereleaseVersionsResponse, *common.Response, error) {
	res := new(PrereleaseVersionsResponse)
	resp, err := s.GetWithQuery("preReleaseVersions", params, res)
	return res, resp, err
}

// GetPrereleaseVersion gets information about a specific prerelease version.
func (s *Service) GetPrereleaseVersion(id string, params *GetPrereleaseVersionQuery) (*PrereleaseVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForPrereleaseVersion gets the app information for a specific prerelease version.
func (s *Service) GetAppForPrereleaseVersion(id string, params *GetAppForPrereleaseVersionQuery) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListPrereleaseVersionsForApp gets a list of prerelease versions associated with a specific app.
func (s *Service) ListPrereleaseVersionsForApp(id string, params *ListPrereleaseVersionsForAppQuery) (*PrereleaseVersionsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/preReleaseVersions", id)
	res := new(PrereleaseVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBuildsForPrereleaseVersion gets a list of builds of a specific prerelease version.
func (s *Service) ListBuildsForPrereleaseVersion(id string, params *ListBuildsForPrereleaseVersionQuery) (*builds.BuildsResponse, *common.Response, error) {
	url := fmt.Sprintf("preReleaseVersions/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPrereleaseVersionForBuild gets the prerelease version for a specific build.
func (s *Service) GetPrereleaseVersionForBuild(id string, params *GetPrereleaseVersionForBuildQuery) (*PrereleaseVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/preReleaseVersion", id)
	res := new(PrereleaseVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
