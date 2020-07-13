package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaLicenseAgreement defines model for BetaLicenseAgreement.
type BetaLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
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
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaLicenseAgreementUpdateRequest defines model for BetaLicenseAgreementUpdateRequest.
type BetaLicenseAgreementUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AgreementText *string `json:"agreementText,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaLicenseAgreementsResponse defines model for BetaLicenseAgreementsResponse.
type BetaLicenseAgreementsResponse struct {
	Data     []BetaLicenseAgreement    `json:"data"`
	Included *[]apps.App               `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// BetaLicenseAgreementResponse defines model for BetaLicenseAgreementResponse.
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement `json:"data"`
	Included *[]apps.App          `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

type ListBetaLicenseAgreementsOptions struct {
	Fields *struct {
		Apps                  *[]string `url:"apps,omitempty"`
		BetaLicenseAgreements *[]string `url:"betaLicenseAgreements,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App *[]string `url:"app,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
}

type GetBetaLicenseAgreementOptions struct {
	Fields *struct {
		Apps                  *[]string `url:"apps,omitempty"`
		BetaLicenseAgreements *[]string `url:"betaLicenseAgreements,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetAppForBetaLicenseAgreementOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type GetBetaLicenseAgreementForAppOptions struct {
	Fields *struct {
		BetaLicenseAgreements *[]string `url:"betaLicenseAgreements,omitempty"`
	} `url:"fields,omitempty"`
}

// ListBetaLicenseAgreements finds and lists beta license agreements for all apps.
func (s *Service) ListBetaLicenseAgreements(params *ListBetaLicenseAgreementsOptions) (*BetaLicenseAgreementsResponse, *common.Response, error) {
	res := new(BetaLicenseAgreementsResponse)
	resp, err := s.Get("betaLicenseAgreements", params, res)
	return res, resp, err
}

// GetBetaLicenseAgreement gets a specific beta license agreement.
func (s *Service) GetBetaLicenseAgreement(id string, params *GetBetaLicenseAgreementOptions) (*BetaLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForBetaLicenseAgreement gets the app information for a specific beta license agreement.
func (s *Service) GetAppForBetaLicenseAgreement(id string, params *GetAppForBetaLicenseAgreementOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetBetaLicenseAgreementForApp gets the beta license agreement for a specific app.
func (s *Service) GetBetaLicenseAgreementForApp(id string, params *GetBetaLicenseAgreementForAppOptions) (*BetaLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/betaLicenseAgreement", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// UpdateBetaLicenseAgreement updates the text for your beta license agreement.
func (s *Service) UpdateBetaLicenseAgreement(id string, body *BetaLicenseAgreementUpdateRequest) (*BetaLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
