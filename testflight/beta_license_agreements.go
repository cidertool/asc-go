package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal"
)

// BetaLicenseAgreement defines model for BetaLicenseAgreement.
type BetaLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
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
	Data     []BetaLicenseAgreement      `json:"data"`
	Included *[]apps.App                 `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// BetaLicenseAgreementResponse defines model for BetaLicenseAgreementResponse.
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement   `json:"data"`
	Included *[]apps.App            `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
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
func (s *Service) ListBetaLicenseAgreements(params *ListBetaLicenseAgreementsQuery) (*BetaLicenseAgreementsResponse, *internal.Response, error) {
	res := new(BetaLicenseAgreementsResponse)
	resp, err := s.GetWithQuery("betaLicenseAgreements", params, res)
	return res, resp, err
}

// GetBetaLicenseAgreement gets a specific beta license agreement.
func (s *Service) GetBetaLicenseAgreement(id string, params *GetBetaLicenseAgreementQuery) (*BetaLicenseAgreementResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaLicenseAgreement gets the app information for a specific beta license agreement.
func (s *Service) GetAppForBetaLicenseAgreement(id string, params *GetAppForBetaLicenseAgreementQuery) (*apps.AppResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaLicenseAgreementForApp gets the beta license agreement for a specific app.
func (s *Service) GetBetaLicenseAgreementForApp(id string, params *GetBetaLicenseAgreementForAppQuery) (*BetaLicenseAgreementResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/betaLicenseAgreement", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBetaLicenseAgreement updates the text for your beta license agreement.
func (s *Service) UpdateBetaLicenseAgreement(id string, body *BetaLicenseAgreementUpdateRequest) (*BetaLicenseAgreementResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
