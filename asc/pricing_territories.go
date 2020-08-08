package asc

import (
	"context"
	"fmt"
)

// Territory defines model for Territory.
type Territory struct {
	Attributes *struct {
		Currency *string `json:"currency,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// TerritoryResponse defines model for TerritoryResponse.
type TerritoryResponse struct {
	Data  Territory     `json:"data"`
	Links DocumentLinks `json:"links"`
}

// TerritoriesResponse defines model for TerritoriesResponse.
type TerritoriesResponse struct {
	Data  []Territory        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListTerritoriesQuery are query options for ListTerritories
type ListTerritoriesQuery struct {
	FieldsTerritories []string `url:"fields[territories],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// ListTerritories lists all territories where the App Store operates.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_territories
func (s *PricingService) ListTerritories(ctx context.Context, params *ListTerritoriesQuery) (*TerritoriesResponse, *Response, error) {
	res := new(TerritoriesResponse)
	resp, err := s.client.get(ctx, "territories", params, res)
	return res, resp, err
}

// ListTerritoriesForApp gets a list of App Store territories where an app is or will be available.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_available_territories_for_an_app
func (s *PricingService) ListTerritoriesForApp(ctx context.Context, id string, params *ListTerritoriesQuery) (*TerritoriesResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/availableTerritories", id)
	res := new(TerritoriesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListTerritoriesForEULA lists all the App Store territories to which a specific custom app license agreement applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_territories_for_an_end_user_license_agreement
func (s *PricingService) ListTerritoriesForEULA(ctx context.Context, id string, params *ListTerritoriesQuery) (*TerritoriesResponse, *Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s/territories", id)
	res := new(TerritoriesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetTerritoryForAppPrice gets the territory in which a specific price point applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_territory_information_of_an_app_price_point
func (s *PricingService) GetTerritoryForAppPrice(ctx context.Context, id string, params *ListTerritoriesQuery) (*TerritoryResponse, *Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
