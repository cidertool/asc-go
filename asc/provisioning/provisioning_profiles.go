package provisioning

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// Profile defines model for Profile.
type Profile struct {
	Attributes *struct {
		CreatedDate    *time.Time        `json:"createdDate,omitempty"`
		ExpirationDate *time.Time        `json:"expirationDate,omitempty"`
		Name           *string           `json:"name,omitempty"`
		Platform       *BundleIDPlatform `json:"platform,omitempty"`
		ProfileContent *string           `json:"profileContent,omitempty"`
		ProfileState   *string           `json:"profileState,omitempty"`
		ProfileType    *string           `json:"profileType,omitempty"`
		UUID           *string           `json:"uuid,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		BundleID *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"bundleId,omitempty"`
		Certificates *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"certificates,omitempty"`
		Devices *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"devices,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// ProfileCreateRequest defines model for ProfileCreateRequest.
type ProfileCreateRequest struct {
	Data struct {
		Attributes struct {
			Name        string `json:"name"`
			ProfileType string `json:"profileType"`
		} `json:"attributes"`
		Relationships struct {
			BundleID struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"bundleId"`
			Certificates struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"certificates"`
			Devices *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"devices,omitempty"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// ProfileResponse defines model for ProfileResponse.
type ProfileResponse struct {
	Data     Profile              `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// ProfilesResponse defines model for ProfilesResponse.
type ProfilesResponse struct {
	Data     []Profile                 `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListProfileQuery struct {
	FieldsCertificates *[]string `url:"fields[certificates],omitempty"`
	FieldsDevices      *[]string `url:"fields[devices],omitempty"`
	FieldsProfiles     *[]string `url:"fields[profiles],omitempty"`
	FieldsID           *[]string `url:"fields[id],omitempty"`
	FieldsName         *[]string `url:"fields[name],omitempty"`
	FieldsBundleIDs    *[]string `url:"fields[bundleIds],omitempty"`
	Include            *[]string `url:"include,omitempty"`
	Limit              *int      `url:"limit,omitempty"`
	LimitCertificates  *int      `url:"limit[certificates],omitempty"`
	LimitDevices       *int      `url:"limit[devices],omitempty"`
	Sort               *[]string `url:"sort,omitempty"`
	FilterProfileState *[]string `url:"filter[profileState],omitempty"`
	FilterProfileType  *[]string `url:"filter[profileType],omitempty"`
	Cursor             *string   `url:"cursor,omitempty"`
}

type GetProfileQuery struct {
	FieldsCertificates *[]string `url:"fields[certificates],omitempty"`
	FieldsDevices      *[]string `url:"fields[devices],omitempty"`
	FieldsProfiles     *[]string `url:"fields[profiles],omitempty"`
	FieldsBundleIds    *[]string `url:"fields[bundleIds],omitempty"`
	LimitCertificates  *int      `url:"limit[certificates],omitempty"`
	LimitDevices       *int      `url:"limit[devices],omitempty"`
	Include            *[]string `url:"include,omitempty"`
}

type GetBundleIDForProfileQuery struct {
	FieldsCertificates *[]string `url:"fields[certificates],omitempty"`
}

type ListCertificatesForProfileQuery struct {
	FieldsCertificates *[]string `url:"fields[certificates],omitempty"`
	Limit              *int      `url:"limit,omitempty"`
	Cursor             *string   `url:"cursor,omitempty"`
}

type ListDevicesInProfileQuery struct {
	FieldsDevices *[]string `url:"fields[devices],omitempty"`
	Limit         *int      `url:"limit,omitempty"`
	Cursor        *string   `url:"cursor,omitempty"`
}

// CreateProfile creates a new provisioning profile.
func (s *Service) CreateProfile(body *ProfileCreateRequest) (*ProfileResponse, *common.Response, error) {
	res := new(ProfileResponse)
	resp, err := s.Post("profiles", body, res)
	return res, resp, err
}

// DeleteProfile deletes a provisioning profile that is used for app development or distribution.
func (s *Service) DeleteProfile(id string) (*common.Response, error) {
	url := fmt.Sprintf("profiles/%s", id)
	return s.Delete(url, nil)
}

// ListProfiles finds and list provisioning profiles and download their data.
func (s *Service) ListProfiles(params *ListProfileQuery) (*ProfilesResponse, *common.Response, error) {
	res := new(ProfilesResponse)
	resp, err := s.GetWithQuery("profiles", params, res)
	return res, resp, err
}

// GetProfile gets information for a specific provisioning profile and download its data.
func (s *Service) GetProfile(id string, params *GetProfileQuery) (*ProfileResponse, *common.Response, error) {
	url := fmt.Sprintf("profiles/%s", id)
	res := new(ProfileResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBundleIDForProfile gets the bundle ID information for a specific provisioning profile.
func (s *Service) GetBundleIDForProfile(id string, params *GetBundleIDForProfileQuery) (*BundleIDResponse, *common.Response, error) {
	url := fmt.Sprintf("profiles/%s/bundleId", id)
	res := new(BundleIDResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListCertificatesInProfile gets a list of all certificates and their data for a specific provisioning profile.
func (s *Service) ListCertificatesInProfile(id string, params *ListCertificatesForProfileQuery) (*CertificatesResponse, *common.Response, error) {
	url := fmt.Sprintf("profiles/%s/certificates", id)
	res := new(CertificatesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListDevicesInProfile gets a list of all devices for a specific provisioning profile.
func (s *Service) ListDevicesInProfile(id string, params *ListDevicesInProfileQuery) (*DevicesResponse, *common.Response, error) {
	url := fmt.Sprintf("profiles/%s/devices", id)
	res := new(DevicesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
