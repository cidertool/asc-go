package submission

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// IDFADeclaration defines model for IdfaDeclaration.
type IDFADeclaration struct {
	Attributes *struct {
		AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
		AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
		HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
		ServesAds                             *bool `json:"servesAds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// IDFADeclarationCreateRequest defines model for IdfaDeclarationCreateRequest.
type IDFADeclarationCreateRequest struct {
	Data struct {
		Attributes struct {
			AttributesActionWithPreviousAd        bool `json:"attributesActionWithPreviousAd"`
			AttributesAppInstallationToPreviousAd bool `json:"attributesAppInstallationToPreviousAd"`
			HonorsLimitedAdTracking               bool `json:"honorsLimitedAdTracking"`
			ServesAds                             bool `json:"servesAds"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreVersion struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// IDFADeclarationUpdateRequest defines model for IdfaDeclarationUpdateRequest.
type IDFADeclarationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
			AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
			HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
			ServesAds                             *bool `json:"servesAds,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// IDFADeclarationResponse defines model for IdfaDeclarationResponse.
type IDFADeclarationResponse struct {
	Data  IDFADeclaration      `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

type GetIDFADeclarationForAppStoreVersionOptions struct {
	Fields *struct {
		IdfaDeclarations *[]string `url:"idfaDeclarations,omitempty"`
	} `url:"fields,omitempty"`
}

// CreateIDFADeclaration declares the IDFA usage for an App Store version.
func (s *Service) CreateIDFADeclaration(body *IDFADeclarationCreateRequest) (*IDFADeclarationResponse, *common.Response, error) {
	res := new(IDFADeclarationResponse)
	resp, err := s.Post("idfaDeclarations", body, res)
	return res, resp, err
}

// UpdateIDFADeclaration updates your declared IDFA usage.
func (s *Service) UpdateIDFADeclaration(id string, body *IDFADeclarationUpdateRequest) (*IDFADeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteIDFADeclaration deletes the IDFA declaration that is associated with a version.
func (s *Service) DeleteIDFADeclaration(id string) (*common.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	return s.Delete(url, nil)
}

// GetIDFADeclarationForAppStoreVersion reads your declared Advertising Identifier (IDFA) usage responses.
func (s *Service) GetIDFADeclarationForAppStoreVersion(id string, params *GetIDFADeclarationForAppStoreVersionOptions) (*IDFADeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/idfaDeclaration", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
