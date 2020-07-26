package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
	"github.com/aaronsky/asc-go/pricing"
)

// EndUserLicenseAgreement defines model for EndUserLicenseAgreement.
type EndUserLicenseAgreement struct {
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
		Territories *struct {
			Data  *[]types.RelationshipsData `json:"data,omitempty"`
			Links *types.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *types.PagingInformation   `json:"meta,omitempty"`
		} `json:"territories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// EndUserLicenseAgreementCreateRequest defines model for EndUserLicenseAgreementCreateRequest.
type EndUserLicenseAgreementCreateRequest struct {
	Attributes struct {
		AgreementText string `json:"agreementText"`
	} `json:"attributes"`
	Relationships struct {
		App struct {
			Data types.RelationshipsData `json:"data"`
		} `json:"app"`
		Territories struct {
			Data []types.RelationshipsData `json:"data"`
		} `json:"territories"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// EndUserLicenseAgreementUpdateRequest defines model for EndUserLicenseAgreementUpdateRequest.
type EndUserLicenseAgreementUpdateRequest struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string `json:"id"`
	Relationships *struct {
		Territories *struct {
			Data *[]types.RelationshipsData `json:"data,omitempty"`
		} `json:"territories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// EndUserLicenseAgreementResponse defines model for EndUserLicenseAgreementResponse.
type EndUserLicenseAgreementResponse struct {
	Data     EndUserLicenseAgreement `json:"data"`
	Included *[]pricing.Territory    `json:"included,omitempty"`
	Links    types.DocumentLinks     `json:"links"`
}

// GetEULAQuery are query options for GetEULA
type GetEULAQuery struct {
	FieldsEndUserLicenseAgreements *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsTerritories              *[]string `url:"fields[territories],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
	LimitTerritories               *int      `url:"limit[territories],omitempty"`
}

// GetEULAForAppQuery are query options for GetEULAForApp
type GetEULAForAppQuery struct {
	FieldsEndUserLicenseAgreements *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
}

// CreateEULA adds a custom end user license agreement (EULA) to an app and configure the territories to which it applies.
func (s *Service) CreateEULA(id string, body *EndUserLicenseAgreementCreateRequest) (*EndUserLicenseAgreementResponse, *services.Response, error) {
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.Post("endUserLicenseAgreements", body, res)
	return res, resp, err
}

// UpdateEULA updates the text or territories for your custom end user license agreement.
func (s *Service) UpdateEULA(id string, body *EndUserLicenseAgreementUpdateRequest) (*EndUserLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteEULA deletes the custom end user license agreement that is associated with an app.
func (s *Service) DeleteEULA(id string) (*services.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	return s.Delete(url, nil)
}

// GetEULA gets the custom end user license agreement associated with an app, and the territories it applies to.
func (s *Service) GetEULA(id string, params *GetEULAQuery) (*EndUserLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetEULAForApp gets the custom end user license agreement (EULA) for a specific app and the territories where the agreement applies.
func (s *Service) GetEULAForApp(id string, params *GetEULAForAppQuery) (*EndUserLicenseAgreementResponse, *services.Response, error) {
	url := fmt.Sprintf("apps/%s/endUserLicenseAgreement", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
