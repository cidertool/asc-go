package asc

import (
	"context"
	"fmt"
	"time"
)

// Profile defines model for Profile.
//
// https://developer.apple.com/documentation/appstoreconnectapi/profile
type Profile struct {
	Attributes    *ProfileAttributes    `json:"attributes,omitempty"`
	ID            string                `json:"id"`
	Links         ResourceLinks         `json:"links"`
	Relationships *ProfileRelationships `json:"relationships,omitempty"`
	Type          string                `json:"type"`
}

// ProfileAttributes defines model for Profile.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/profile/attributes
type ProfileAttributes struct {
	CreatedDate    *time.Time        `json:"createdDate,omitempty"`
	ExpirationDate *time.Time        `json:"expirationDate,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Platform       *BundleIDPlatform `json:"platform,omitempty"`
	ProfileContent *string           `json:"profileContent,omitempty"`
	ProfileState   *string           `json:"profileState,omitempty"`
	ProfileType    *string           `json:"profileType,omitempty"`
	UUID           *string           `json:"uuid,omitempty"`
}

// ProfileRelationships defines model for Profile.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/profile/relationships
type ProfileRelationships struct {
	BundleID     *Relationship      `json:"bundleId,omitempty"`
	Certificates *PagedRelationship `json:"certificates,omitempty"`
	Devices      *PagedRelationship `json:"devices,omitempty"`
}

// ProfileCreateRequest defines model for ProfileCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/profilecreaterequest
type ProfileCreateRequest struct {
	Attributes    ProfileCreateRequestAttributes    `json:"attributes"`
	Relationships ProfileCreateRequestRelationships `json:"relationships"`
	Type          string                            `json:"type"`
}

// ProfileCreateRequestAttributes are attributes for ProfileCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/profilecreaterequest/data/attributes
type ProfileCreateRequestAttributes struct {
	Name        string `json:"name"`
	ProfileType string `json:"profileType"`
}

// ProfileCreateRequestRelationships are relationships for ProfileCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/profilecreaterequest/data/relationships
type ProfileCreateRequestRelationships struct {
	BundleID     RelationshipDeclaration       `json:"bundleId"`
	Certificates PagedRelationshipDeclaration  `json:"certificates"`
	Devices      *PagedRelationshipDeclaration `json:"devices,omitempty"`
}

// ProfileResponse defines model for ProfileResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/profileresponse
type ProfileResponse struct {
	Data     Profile        `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// ProfilesResponse defines model for ProfilesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/profilesresponse
type ProfilesResponse struct {
	Data     []Profile          `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListProfilesQuery are query options for ListProfile
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_and_download_profiles
type ListProfilesQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
	FieldsDevices      []string `url:"fields[devices],omitempty"`
	FieldsProfiles     []string `url:"fields[profiles],omitempty"`
	FieldsID           []string `url:"fields[id],omitempty"`
	FieldsName         []string `url:"fields[name],omitempty"`
	FieldsBundleIDs    []string `url:"fields[bundleIds],omitempty"`
	Include            []string `url:"include,omitempty"`
	Limit              int      `url:"limit,omitempty"`
	LimitCertificates  int      `url:"limit[certificates],omitempty"`
	LimitDevices       int      `url:"limit[devices],omitempty"`
	Sort               []string `url:"sort,omitempty"`
	FilterProfileState []string `url:"filter[profileState],omitempty"`
	FilterProfileType  []string `url:"filter[profileType],omitempty"`
	Cursor             string   `url:"cursor,omitempty"`
}

// GetProfileQuery are query options for GetProfile
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_and_download_profile_information
type GetProfileQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
	FieldsDevices      []string `url:"fields[devices],omitempty"`
	FieldsProfiles     []string `url:"fields[profiles],omitempty"`
	FieldsBundleIds    []string `url:"fields[bundleIds],omitempty"`
	LimitCertificates  int      `url:"limit[certificates],omitempty"`
	LimitDevices       int      `url:"limit[devices],omitempty"`
	Include            []string `url:"include,omitempty"`
}

// GetBundleIDForProfileQuery are query options for GetBundleIDForProfile
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_bundle_id_in_a_profile
type GetBundleIDForProfileQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
}

// ListCertificatesForProfileQuery are query options for ListCertificatesForProfile
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_certificates_in_a_profile
type ListCertificatesForProfileQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
	Limit              int      `url:"limit,omitempty"`
	Cursor             string   `url:"cursor,omitempty"`
}

// ListDevicesInProfileQuery are query options for ListDevicesInProfile
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_devices_in_a_profile
type ListDevicesInProfileQuery struct {
	FieldsDevices []string `url:"fields[devices],omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Cursor        string   `url:"cursor,omitempty"`
}

// CreateProfile creates a new provisioning profile.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_profile
func (s *ProvisioningService) CreateProfile(ctx context.Context, body *ProfileCreateRequest) (*ProfileResponse, *Response, error) {
	res := new(ProfileResponse)
	resp, err := s.client.post(ctx, "profiles", body, res)
	return res, resp, err
}

// DeleteProfile deletes a provisioning profile that is used for app development or distribution.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_profile
func (s *ProvisioningService) DeleteProfile(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("profiles/%s", id)
	return s.client.delete(ctx, url, nil)
}

// ListProfiles finds and list provisioning profiles and download their data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_and_download_profiles
func (s *ProvisioningService) ListProfiles(ctx context.Context, params *ListProfilesQuery) (*ProfilesResponse, *Response, error) {
	res := new(ProfilesResponse)
	resp, err := s.client.get(ctx, "profiles", params, res)
	return res, resp, err
}

// GetProfile gets information for a specific provisioning profile and download its data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_and_download_profile_information
func (s *ProvisioningService) GetProfile(ctx context.Context, id string, params *GetProfileQuery) (*ProfileResponse, *Response, error) {
	url := fmt.Sprintf("profiles/%s", id)
	res := new(ProfileResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBundleIDForProfile gets the bundle ID information for a specific provisioning profile.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_bundle_id_in_a_profile
func (s *ProvisioningService) GetBundleIDForProfile(ctx context.Context, id string, params *GetBundleIDForProfileQuery) (*BundleIDResponse, *Response, error) {
	url := fmt.Sprintf("profiles/%s/bundleId", id)
	res := new(BundleIDResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListCertificatesInProfile gets a list of all certificates and their data for a specific provisioning profile.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_certificates_in_a_profile
func (s *ProvisioningService) ListCertificatesInProfile(ctx context.Context, id string, params *ListCertificatesForProfileQuery) (*CertificatesResponse, *Response, error) {
	url := fmt.Sprintf("profiles/%s/certificates", id)
	res := new(CertificatesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListDevicesInProfile gets a list of all devices for a specific provisioning profile.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_devices_in_a_profile
func (s *ProvisioningService) ListDevicesInProfile(ctx context.Context, id string, params *ListDevicesInProfileQuery) (*DevicesResponse, *Response, error) {
	url := fmt.Sprintf("profiles/%s/devices", id)
	res := new(DevicesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
