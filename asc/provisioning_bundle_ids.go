package asc

import (
	"context"
	"fmt"
)

// BundleIDPlatform defines model for BundleIdPlatform.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidplatform
type BundleIDPlatform string

const (
	// BundleIDPlatformiOS is a string that represents iOS.
	BundleIDPlatformiOS BundleIDPlatform = "IOS"
	// BundleIDPlatformMacOS is a string that represents macOS.
	BundleIDPlatformMacOS BundleIDPlatform = "MAC_OS"
)

// BundleID defines model for BundleId.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleid
type BundleID struct {
	Attributes    *BundleIDAttributes    `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         ResourceLinks          `json:"links"`
	Relationships *BundleIDRelationships `json:"relationships,omitempty"`
	Type          string                 `json:"type"`
}

// BundleIDAttributes defines model for BundleId.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleid/attributes
type BundleIDAttributes struct {
	IDentifier *string           `json:"identifier,omitempty"`
	Name       *string           `json:"name,omitempty"`
	Platform   *BundleIDPlatform `json:"platform,omitempty"`
	SeedID     *string           `json:"seedId,omitempty"`
}

// BundleIDRelationships defines model for BundleId.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleid/relationships
type BundleIDRelationships struct {
	App                  *Relationship      `json:"app,omitempty"`
	BundleIDCapabilities *PagedRelationship `json:"bundleIdCapabilities,omitempty"`
	Profiles             *PagedRelationship `json:"profiles,omitempty"`
}

// BundleIDCreateRequest defines model for BundleIdCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcreaterequest
type BundleIDCreateRequest struct {
	Attributes BundleIDCreateRequestAttributes `json:"attributes"`
	Type       string                          `json:"type"`
}

// BundleIDCreateRequestAttributes are attributes for BundleIDCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcreaterequest/data/attributes
type BundleIDCreateRequestAttributes struct {
	Identifier string           `json:"identifier"`
	Name       string           `json:"name"`
	Platform   BundleIDPlatform `json:"platform"`
	SeedID     *string          `json:"seedId,omitempty"`
}

// BundleIDUpdateRequest defines model for BundleIdUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidupdaterequest
type BundleIDUpdateRequest struct {
	Attributes *BundleIDUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                           `json:"id"`
	Type       string                           `json:"type"`
}

// BundleIDUpdateRequestAttributes are attributes for BundleIDUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidupdaterequest/data/attributes
type BundleIDUpdateRequestAttributes struct {
	Name *string `json:"name,omitempty"`
}

// BundleIDResponse defines model for BundleIdResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidresponse
type BundleIDResponse struct {
	Data     BundleID       `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// BundleIDsResponse defines model for BundleIdsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidsresponse
type BundleIDsResponse struct {
	Data     []BundleID         `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListBundleIDsQuery are query options for ListBundleIDs
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_bundle_ids
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_bundle_id_information
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_bundle_id
type GetAppForBundleIDQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListProfilesForBundleIDQuery are query options for ListProfilesForBundleID
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_profiles_for_a_bundle_id
type ListProfilesForBundleIDQuery struct {
	FieldsProfiles []string `url:"fields[profiles],omitempty"`
	Limit          int      `url:"limit,omitempty"`
	Cursor         string   `url:"cursor,omitempty"`
}

// ListCapabilitiesForBundleIDQuery are query options for ListCapabilitiesForBundleID
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_capabilities_for_a_bundle_id
type ListCapabilitiesForBundleIDQuery struct {
	FieldsBundleIDCapabilities []string `url:"fields[bundleIdCapabilities],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

func (r *BundleIDCreateRequest) applyTypes() {
	if r == nil || r.Type != "" {
		return
	}
	r.Type = "bundleIds"
}

func (r *BundleIDUpdateRequest) applyTypes() {
	if r == nil || r.Type != "" {
		return
	}
	r.Type = "bundleIds"
}

// CreateBundleID registers a new bundle ID for app development.
//
// https://developer.apple.com/documentation/appstoreconnectapi/register_a_new_bundle_id
func (s *ProvisioningService) CreateBundleID(ctx context.Context, body *BundleIDCreateRequest) (*BundleIDResponse, *Response, error) {
	res := new(BundleIDResponse)
	resp, err := s.client.post(ctx, "bundleIds", body, res)
	return res, resp, err
}

// UpdateBundleID updates a specific bundle ID’s name.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_bundle_id
func (s *ProvisioningService) UpdateBundleID(ctx context.Context, id string, body *BundleIDUpdateRequest) (*BundleIDResponse, *Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeleteBundleID deletes a bundle ID that is used for app development.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_bundle_id
func (s *ProvisioningService) DeleteBundleID(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	return s.client.delete(ctx, url, nil)
}

// ListBundleIDs finds and lists bundle IDs that are registered to your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_bundle_ids
func (s *ProvisioningService) ListBundleIDs(ctx context.Context, params *ListBundleIDsQuery) (*BundleIDsResponse, *Response, error) {
	res := new(BundleIDsResponse)
	resp, err := s.client.get(ctx, "bundleIds", params, res)
	return res, resp, err
}

// GetBundleID gets information about a specific bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_bundle_id_information
func (s *ProvisioningService) GetBundleID(ctx context.Context, id string, params *GetBundleIDQuery) (*BundleIDResponse, *Response, error) {
	url := fmt.Sprintf("bundleIds/%s", id)
	res := new(BundleIDResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppForBundleID gets app information for a specific bundle identifier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_bundle_id
func (s *ProvisioningService) GetAppForBundleID(ctx context.Context, id string, params *GetAppForBundleIDQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("bundleIds/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListProfilesForBundleID gets a list of all profiles for a specific bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_profiles_for_a_bundle_id
func (s *ProvisioningService) ListProfilesForBundleID(ctx context.Context, id string, params *ListProfilesForBundleIDQuery) (*ProfilesResponse, *Response, error) {
	url := fmt.Sprintf("bundleIds/%s/profiles", id)
	res := new(ProfilesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListCapabilitiesForBundleID gets a list of all capabilities for a specific bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_capabilities_for_a_bundle_id
func (s *ProvisioningService) ListCapabilitiesForBundleID(ctx context.Context, id string, params *ListCapabilitiesForBundleIDQuery) (*BundleIDCapabilitiesResponse, *Response, error) {
	url := fmt.Sprintf("bundleIds/%s/bundleIdCapabilities", id)
	res := new(BundleIDCapabilitiesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
