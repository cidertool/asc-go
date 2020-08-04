package asc

import (
	"fmt"
	"net/http"
)

// BundleIDPlatform defines model for BundleIdPlatform.
type BundleIDPlatform string

const (
	// BundleIDPlatformiOS is a string that represents iOS.
	BundleIDPlatformiOS BundleIDPlatform = "IOS"
	// BundleIDPlatformMacOS is a string that represents macOS.
	BundleIDPlatformMacOS BundleIDPlatform = "MAC_OS"
)

// BundleID defines model for BundleId.
type BundleID struct {
	Attributes *struct {
		IDentifier *string           `json:"identifier,omitempty"`
		Name       *string           `json:"name,omitempty"`
		Platform   *BundleIDPlatform `json:"platform,omitempty"`
		SeedID     *string           `json:"seedId,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		BundleIDCapabilities *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"bundleIdCapabilities,omitempty"`
		Profiles *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"profiles,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BundleIDCreateRequest defines model for BundleIdCreateRequest.
type BundleIDCreateRequest struct {
	Attributes BundleIDCreateRequestAttributes `json:"attributes"`
	Type       string                          `json:"type"`
}

// BundleIDCreateRequestAttributes are attributes for BundleIDCreateRequest
type BundleIDCreateRequestAttributes struct {
	Identifier string           `json:"identifier"`
	Name       string           `json:"name"`
	Platform   BundleIDPlatform `json:"platform"`
	SeedID     *string          `json:"seedId,omitempty"`
}

// BundleIDUpdateRequest defines model for BundleIdUpdateRequest.
type BundleIDUpdateRequest struct {
	Attributes *BundleIDUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                           `json:"id"`
	Type       string                           `json:"type"`
}

// BundleIDUpdateRequestAttributes are attributes for BundleIDUpdateRequest
type BundleIDUpdateRequestAttributes struct {
	Name *string `json:"name,omitempty"`
}

// BundleIDResponse defines model for BundleIdResponse.
type BundleIDResponse struct {
	Data     BundleID       `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// BundleIDsResponse defines model for BundleIdsResponse.
type BundleIDsResponse struct {
	Data     []BundleID         `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListBundleIDsQuery are query options for ListBundleIDs
type ListBundleIDsQuery struct {
	FieldsBundleIds            []string `url:"fields[bundleIds],omitempty"`
	FieldsProfiles             []string `url:"fields[profiles],omitempty"`
	FieldsBundleIDCapabilities []string `url:"fields[bundleIdCapabilities],omitempty"`
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	Include                    []string `url:"include,omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	LimitProfiles              int      `url:"limit[profiles],omitempty"`
	LimitBundleIDCapabilities  int      `url:"limit[bundleIdCapabilities],omitempty"`
	Sort                       []string `url:"sort,omitempty"`
	FilterID                   []string `url:"filter[id],omitempty"`
	FilterIdentifier           []string `url:"filter[identifier],omitempty"`
	FilterName                 []string `url:"filter[name],omitempty"`
	FilterPlatform             []string `url:"filter[platform],omitempty"`
	FilterSeedID               []string `url:"filter[seedId],omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetBundleIDQuery are query options for GetBundleID
type GetBundleIDQuery struct {
	FieldsBundleIds            []string `url:"fields[bundleIds],omitempty"`
	FieldsProfiles             []string `url:"fields[profiles],omitempty"`
	FieldsBundleIDCapabilities []string `url:"fields[bundleIdCapabilities],omitempty"`
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	LimitProfiles              int      `url:"limit[profiles],omitempty"`
	LimitBundleIDCapabilities  int      `url:"limit[bundleIdCapabilities],omitempty"`
	Include                    []string `url:"include,omitempty"`
}

// GetAppForBundleIDQuery are query options for GetAppForBundleID
type GetAppForBundleIDQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListProfilesForBundleIDQuery are query options for ListProfilesForBundleID
type ListProfilesForBundleIDQuery struct {
	FieldsProfiles []string `url:"fields[profiles],omitempty"`
	Limit          int      `url:"limit,omitempty"`
	Cursor         string   `url:"cursor,omitempty"`
}

// ListCapabilitiesForBundleIDQuery are query options for ListCapabilitiesForBundleID
type ListCapabilitiesForBundleIDQuery struct {
	FieldsBundleIDCapabilities []string `url:"fields[bundleIdCapabilities],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// CreateBundleID registers a new bundle ID for app development.
func (s *ProvisioningService) CreateBundleID(body *BundleIDCreateRequest) (*BundleIDResponse, *http.Response, error) {
	res := new(BundleIDResponse)
	resp, err := s.client.post("bundleIds", body, res)
	return res, resp, err
}

// UpdateBundleID updates a specific bundle ID’s name.
func (s *ProvisioningService) UpdateBundleID(id string, body *BundleIDUpdateRequest) (*BundleIDResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteBundleID deletes a bundle ID that is used for app development.
func (s *ProvisioningService) DeleteBundleID(id string) (*http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	return s.client.delete(url, nil)
}

// ListBundleIDs finds and lists bundle IDs that are registered to your team.
func (s *ProvisioningService) ListBundleIDs(params *ListBundleIDsQuery) (*BundleIDsResponse, *http.Response, error) {
	res := new(BundleIDsResponse)
	resp, err := s.client.get("bundleIds", params, res)
	return res, resp, err
}

// GetBundleID gets information about a specific bundle ID.
func (s *ProvisioningService) GetBundleID(id string, params *GetBundleIDQuery) (*BundleIDResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForBundleID gets app information for a specific bundle identifier.
func (s *ProvisioningService) GetAppForBundleID(id string, params *GetAppForBundleIDQuery) (*AppResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListProfilesForBundleID gets a list of all profiles for a specific bundle ID.
func (s *ProvisioningService) ListProfilesForBundleID(id string, params *ListProfilesForBundleIDQuery) (*ProfilesResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/profiles", id)
	res := new(ProfilesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListCapabilitiesForBundleID gets a list of all capabilities for a specific bundle ID.
func (s *ProvisioningService) ListCapabilitiesForBundleID(id string, params *ListCapabilitiesForBundleIDQuery) (*BundleIDCapabilitiesResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIds/%s/bundleIdCapabilities", id)
	res := new(BundleIDCapabilitiesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
