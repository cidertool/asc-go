package asc

import (
	"fmt"
	"net/http"
)

// IDFADeclaration defines model for IDFADeclaration.
type IDFADeclaration struct {
	Attributes *struct {
		AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
		AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
		HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
		ServesAds                             *bool `json:"servesAds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// IDFADeclarationCreateRequest defines model for IDFADeclarationCreateRequest.
type IDFADeclarationCreateRequest struct {
	Attributes    IDFADeclarationCreateRequestAttributes    `json:"attributes"`
	Relationships IDFADeclarationCreateRequestRelationships `json:"relationships"`
	Type          string                                    `json:"type"`
}

// IDFADeclarationCreateRequestAttributes are attributes for IDFADeclarationCreateRequest
type IDFADeclarationCreateRequestAttributes struct {
	AttributesActionWithPreviousAd        bool `json:"attributesActionWithPreviousAd"`
	AttributesAppInstallationToPreviousAd bool `json:"attributesAppInstallationToPreviousAd"`
	HonorsLimitedAdTracking               bool `json:"honorsLimitedAdTracking"`
	ServesAds                             bool `json:"servesAds"`
}

// IDFADeclarationCreateRequestRelationships are relationships for IDFADeclarationCreateRequest
type IDFADeclarationCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// IDFADeclarationUpdateRequest defines model for IDFADeclarationUpdateRequest.
type IDFADeclarationUpdateRequest struct {
	Attributes *IDFADeclarationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// IDFADeclarationUpdateRequestAttributes are attributes for IDFADeclarationUpdateRequest
type IDFADeclarationUpdateRequestAttributes struct {
	AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
	AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
	HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
	ServesAds                             *bool `json:"servesAds,omitempty"`
}

// IDFADeclarationResponse defines model for IDFADeclarationResponse.
type IDFADeclarationResponse struct {
	Data  IDFADeclaration `json:"data"`
	Links DocumentLinks   `json:"links"`
}

// GetIDFADeclarationForAppStoreVersionQuery are query options for GetIDFADeclarationForAppStoreVersion
type GetIDFADeclarationForAppStoreVersionQuery struct {
	FieldsIDFADeclarations *[]string `url:"fields[idfaDeclarations],omitempty"`
}

// CreateIDFADeclaration declares the IDFA usage for an App Store version.
func (s *SubmissionService) CreateIDFADeclaration(body *IDFADeclarationCreateRequest) (*IDFADeclarationResponse, *http.Response, error) {
	res := new(IDFADeclarationResponse)
	resp, err := s.client.Post("idfaDeclarations", body, res)
	return res, resp, err
}

// UpdateIDFADeclaration updates your declared IDFA usage.
func (s *SubmissionService) UpdateIDFADeclaration(id string, body *IDFADeclarationUpdateRequest) (*IDFADeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.client.Patch(url, body, res)
	return res, resp, err
}

// DeleteIDFADeclaration deletes the IDFA declaration that is associated with a version.
func (s *SubmissionService) DeleteIDFADeclaration(id string) (*http.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	return s.client.Delete(url, nil)
}

// GetIDFADeclarationForAppStoreVersion reads your declared Advertising Identifier (IDFA) usage responses.
func (s *SubmissionService) GetIDFADeclarationForAppStoreVersion(id string, params *GetIDFADeclarationForAppStoreVersionQuery) (*IDFADeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/idfaDeclaration", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}
