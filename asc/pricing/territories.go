package pricing

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// Territory defines model for Territory.
type Territory struct {
	Attributes *struct {
		Currency *string `json:"currency,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// TerritoryResponse defines model for TerritoryResponse.
type TerritoryResponse struct {
	Data  Territory            `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// TerritoriesResponse defines model for TerritoriesResponse.
type TerritoriesResponse struct {
	Data  []Territory               `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type ListTerritoriesQuery struct {
	Fields *struct {
		Territories *[]string `url:"territories,omitempty"`
	} `url:"fields,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// ListTerritories lists all territories where the App Store operates.
func (s *Service) ListTerritories(params *ListTerritoriesQuery) (*TerritoriesResponse, *common.Response, error) {
	res := new(TerritoriesResponse)
	resp, err := s.GetWithQuery("territories", params, res)
	return res, resp, err
}

// ListTerritoriesForApp gets a list of App Store territories where an app is or will be available.
func (s *Service) ListTerritoriesForApp(id string, params *ListTerritoriesQuery) (*TerritoriesResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/availableTerritories", id)
	res := new(TerritoriesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListTerritoriesForEULA lists all the App Store territories to which a specific custom app license agreement applies.
func (s *Service) ListTerritoriesForEULA(id string, params *ListTerritoriesQuery) (*TerritoriesResponse, *common.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s/territories", id)
	res := new(TerritoriesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetTerritoryForAppPrice gets the territory in which a specific price point applies.
func (s *Service) GetTerritoryForAppPrice(id string, params *ListTerritoriesQuery) (*TerritoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
