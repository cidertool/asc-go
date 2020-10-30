/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// Territory defines model for Territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/territory
type Territory struct {
	Attributes *TerritoryAttributes `json:"attributes,omitempty"`
	ID         string               `json:"id"`
	Links      ResourceLinks        `json:"links"`
	Type       string               `json:"type"`
}

// TerritoryAttributes defines model for Territory.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/territory/attributes
type TerritoryAttributes struct {
	Currency *string `json:"currency,omitempty"`
}

// TerritoryResponse defines model for TerritoryResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/territoryresponse
type TerritoryResponse struct {
	Data  Territory     `json:"data"`
	Links DocumentLinks `json:"links"`
}

// TerritoriesResponse defines model for TerritoriesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/territoriesresponse
type TerritoriesResponse struct {
	Data  []Territory        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListTerritoriesQuery are query options for ListTerritories
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_territories
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_available_territories_for_an_app
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_territories_for_an_end_user_license_agreement
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_territory_information_of_an_app_price_point
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
