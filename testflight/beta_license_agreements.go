package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// BetaLicenseAgreement defines model for BetaLicenseAgreement.
type BetaLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaLicenseAgreementUpdateRequest defines model for BetaLicenseAgreementUpdateRequest.
type BetaLicenseAgreementUpdateRequest struct {
	Attributes *BetaLicenseAgreementUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// BetaLicenseAgreementUpdateRequestAttributes are attributes for BetaLicenseAgreementUpdateRequest
type BetaLicenseAgreementUpdateRequestAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// BetaLicenseAgreementsResponse defines model for BetaLicenseAgreementsResponse.
type BetaLicenseAgreementsResponse struct {
	Data     []BetaLicenseAgreement   `json:"data"`
	Included *[]apps.App              `json:"included,omitempty"`
	Links    types.PagedDocumentLinks `json:"links"`
	Meta     *types.PagingInformation `json:"meta,omitempty"`
}

// BetaLicenseAgreementResponse defines model for BetaLicenseAgreementResponse.
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement `json:"data"`
	Included *[]apps.App          `json:"included,omitempty"`
	Links    types.DocumentLinks  `json:"links"`
}

// ListBetaLicenseAgreementsQuery defines model for ListBetaLicenseAgreements
type ListBetaLicenseAgreementsQuery struct {
	FieldsApps                  *[]string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements *[]string `url:"fields[betaLicenseAgreements],omitempty"`
	FilterApp                   *[]string `url:"filter[app],omitempty"`
	Include                     *[]string `url:"include,omitempty"`
	Limit                       *int      `url:"limit,omitempty"`
	Cursor                      *string   `url:"cursor,omitempty"`
}

// GetBetaLicenseAgreementQuery defines model for GetBetaLicenseAgreement
type GetBetaLicenseAgreementQuery struct {
	FieldsApps                  *[]string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements *[]string `url:"fields[betaLicenseAgreements],omitempty"`
	Include                     *[]string `url:"include,omitempty"`
}

// GetAppForBetaLicenseAgreementQuery defines model for GetAppForBetaLicenseAgreement
type GetAppForBetaLicenseAgreementQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// GetBetaLicenseAgreementForAppQuery defines model for GetBetaLicenseAgreementForApp
type GetBetaLicenseAgreementForAppQuery struct {
	FieldsBetaLicenseAgreements *[]string `url:"fields[betaLicenseAgreements],omitempty"`
}

// ListBetaLicenseAgreements finds and lists beta license agreements for all apps.
func (s *Service) ListBetaLicenseAgreements(params *ListBetaLicenseAgreementsQuery) (*BetaLicenseAgreementsResponse, *services.Response, error) {
	res := new(BetaLicenseAgreementsResponse)
	resp, err := s.GetWithQuery("betaLicenseAgreements", params, res)
	return res, resp, err
}

// GetBetaLicenseAgreement gets a specific beta license agreement.
func (s *Service) GetBetaLicenseAgreement(id string, params *GetBetaLicenseAgreementQuery) (*BetaLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaLicenseAgreement gets the app information for a specific beta license agreement.
func (s *Service) GetAppForBetaLicenseAgreement(id string, params *GetAppForBetaLicenseAgreementQuery) (*apps.AppResponse, *services.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaLicenseAgreementForApp gets the beta license agreement for a specific app.
func (s *Service) GetBetaLicenseAgreementForApp(id string, params *GetBetaLicenseAgreementForAppQuery) (*BetaLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("apps/%s/betaLicenseAgreement", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBetaLicenseAgreement updates the text for your beta license agreement.
func (s *Service) UpdateBetaLicenseAgreement(id string, body *BetaLicenseAgreementUpdateRequest) (*BetaLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
