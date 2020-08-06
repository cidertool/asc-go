package asc

import (
	"fmt"
)

// BetaLicenseAgreement defines model for BetaLicenseAgreement.
type BetaLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
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
	Data     []BetaLicenseAgreement `json:"data"`
	Included *[]App                 `json:"included,omitempty"`
	Links    PagedDocumentLinks     `json:"links"`
	Meta     *PagingInformation     `json:"meta,omitempty"`
}

// BetaLicenseAgreementResponse defines model for BetaLicenseAgreementResponse.
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement `json:"data"`
	Included *[]App               `json:"included,omitempty"`
	Links    DocumentLinks        `json:"links"`
}

// ListBetaLicenseAgreementsQuery defines model for ListBetaLicenseAgreements
type ListBetaLicenseAgreementsQuery struct {
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
	FilterApp                   []string `url:"filter[app],omitempty"`
	Include                     []string `url:"include,omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	Cursor                      string   `url:"cursor,omitempty"`
}

// GetBetaLicenseAgreementQuery defines model for GetBetaLicenseAgreement
type GetBetaLicenseAgreementQuery struct {
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
	Include                     []string `url:"include,omitempty"`
}

// GetAppForBetaLicenseAgreementQuery defines model for GetAppForBetaLicenseAgreement
type GetAppForBetaLicenseAgreementQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// GetBetaLicenseAgreementForAppQuery defines model for GetBetaLicenseAgreementForApp
type GetBetaLicenseAgreementForAppQuery struct {
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
}

// ListBetaLicenseAgreements finds and lists beta license agreements for all apps.
func (s *TestflightService) ListBetaLicenseAgreements(params *ListBetaLicenseAgreementsQuery) (*BetaLicenseAgreementsResponse, *Response, error) {
	res := new(BetaLicenseAgreementsResponse)
	resp, err := s.client.get("betaLicenseAgreements", params, res)
	return res, resp, err
}

// GetBetaLicenseAgreement gets a specific beta license agreement.
func (s *TestflightService) GetBetaLicenseAgreement(id string, params *GetBetaLicenseAgreementQuery) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForBetaLicenseAgreement gets the app information for a specific beta license agreement.
func (s *TestflightService) GetAppForBetaLicenseAgreement(id string, params *GetAppForBetaLicenseAgreementQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBetaLicenseAgreementForApp gets the beta license agreement for a specific app.
func (s *TestflightService) GetBetaLicenseAgreementForApp(id string, params *GetBetaLicenseAgreementForAppQuery) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/betaLicenseAgreement", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// UpdateBetaLicenseAgreement updates the text for your beta license agreement.
func (s *TestflightService) UpdateBetaLicenseAgreement(id string, body *BetaLicenseAgreementUpdateRequest) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}
