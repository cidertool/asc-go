package provisioning

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BundleIDPlatform defines model for BundleIdPlatform.
type BundleIDPlatform string

const (
	// IOS is a string that represents iOS.
	IOS BundleIDPlatform = "IOS"
	// MacOS is a string that represents macOS.
	MacOS BundleIDPlatform = "MAC_OS"
)

// BundleID defines model for BundleId.
type BundleID struct {
	Attributes *struct {
		IDentifier *string           `json:"identifier,omitempty"`
		Name       *string           `json:"name,omitempty"`
		Platform   *BundleIDPlatform `json:"platform,omitempty"`
		SeedID     *string           `json:"seedId,omitempty"`
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
		BundleIDCapabilities *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"bundleIdCapabilities,omitempty"`
		Profiles *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"profiles,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BundleIDCreateRequest defines model for BundleIdCreateRequest.
type BundleIDCreateRequest struct {
	Data struct {
		Attributes struct {
			Identifier string           `json:"identifier"`
			Name       string           `json:"name"`
			Platform   BundleIDPlatform `json:"platform"`
			SeedID     *string          `json:"seedId,omitempty"`
		} `json:"attributes"`
		Type string `json:"type"`
	} `json:"data"`
}

// BundleIDUpdateRequest defines model for BundleIdUpdateRequest.
type BundleIDUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Name *string `json:"name,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BundleIDResponse defines model for BundleIdResponse.
type BundleIDResponse struct {
	Data     BundleID             `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BundleIDsResponse defines model for BundleIdsResponse.
type BundleIDsResponse struct {
	Data     []BundleID                `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBundleIDsOptions struct {
	Fields *struct {
		BundleIds            *[]string `url:"bundleIds,omitempty"`
		Profiles             *[]string `url:"profiles,omitempty"`
		BundleIDCapabilities *[]string `url:"bundleIdCapabilities,omitempty"`
		Apps                 *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		Profiles             *int `url:"profiles,omitempty"`
		BundleIDCapabilities *int `url:"bundleIdCapabilities,omitempty"`
	} `url:"limit,omitempty"`
	Sort   *[]string `url:"sort,omitempty"`
	Filter *struct {
		ID         *[]string `url:"id,omitempty"`
		Identifier *[]string `url:"identifier,omitempty"`
		Name       *[]string `url:"name,omitempty"`
		Platform   *[]string `url:"platform,omitempty"`
		SeedID     *[]string `url:"seedId,omitempty"`
	} `url:"filter,omitempty"`
}

type GetBundleIDOptions struct {
	Fields *struct {
		BundleIds            *[]string `url:"bundleIds,omitempty"`
		Profiles             *[]string `url:"profiles,omitempty"`
		BundleIDCapabilities *[]string `url:"bundleIdCapabilities,omitempty"`
		Apps                 *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
	Limit *struct {
		Profiles             *int `url:"profiles,omitempty"`
		BundleIDCapabilities *int `url:"bundleIdCapabilities,omitempty"`
	} `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetAppForBundleIDOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type ListProfilesForBundleIDOptions struct {
	Fields *struct {
		Profiles *[]string `url:"profiles,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListCapabilitiesForBundleIDOptions struct {
	Fields *struct {
		BundleIDCapabilities *[]string `url:"bundleIdCapabilities,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// CreateBundleID registers a new bundle ID for app development.
func (s *Service) CreateBundleID(body *BundleIDCreateRequest) (*BundleIDResponse, *common.Response, error) {
	res := new(BundleIDResponse)
	resp, err := s.Post("bundleIds", body, res)
	return res, resp, err
}

// UpdateBundleID updates a specific bundle ID’s name.
func (s *Service) UpdateBundleID(id string, body *BundleIDUpdateRequest) (*BundleIDResponse, *common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBundleID deletes a bundle ID that is used for app development.
func (s *Service) DeleteBundleID(id string) (*common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	return s.Delete(url, nil)
}

// ListBundleIDs finds and lists bundle IDs that are registered to your team.
func (s *Service) ListBundleIDs(params *ListBundleIDsOptions) (*BundleIDsResponse, *common.Response, error) {
	res := new(BundleIDsResponse)
	resp, err := s.Get("bundleIds", params, res)
	return res, resp, err
}

// GetBundleID gets information about a specific bundle ID.
func (s *Service) GetBundleID(id string, params *GetBundleIDOptions) (*BundleIDResponse, *common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForBundleID gets app information for a specific bundle identifier.
func (s *Service) GetAppForBundleID(id string, params *GetAppForBundleIDOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListProfilesForBundleID gets a list of all profiles for a specific bundle ID.
func (s *Service) ListProfilesForBundleID(id string, params *ListProfilesForBundleIDOptions) (*ProfilesResponse, *common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/profiles", id)
	res := new(ProfilesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListCapabilitiesForBundleID gets a list of all capabilities for a specific bundle ID.
func (s *Service) ListCapabilitiesForBundleID(id string, params *ListCapabilitiesForBundleIDOptions) (*BundleIDCapabilitiesResponse, *common.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/bundleIdCapabilities", id)
	res := new(BundleIDCapabilitiesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
