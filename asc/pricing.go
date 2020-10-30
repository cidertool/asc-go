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

// PricingService handles communication with pricing-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/app_prices
// https://developer.apple.com/documentation/appstoreconnectapi/territories
// https://developer.apple.com/documentation/appstoreconnectapi/app_price_reference_data
type PricingService service

// AppPrice defines model for AppPrice.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appprice
type AppPrice struct {
	ID            string                 `json:"id"`
	Links         ResourceLinks          `json:"links"`
	Relationships *AppPriceRelationships `json:"relationships,omitempty"`
	Type          string                 `json:"type"`
}

// AppPriceRelationships defines model for AppPrice.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appprice/relationships
type AppPriceRelationships struct {
	App       *Relationship `json:"app,omitempty"`
	PriceTier *Relationship `json:"priceTier,omitempty"`
}

// AppPriceResponse defines model for AppPriceResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppriceresponse
type AppPriceResponse struct {
	Data  AppPrice      `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppPricesResponse defines model for AppPricesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricesresponse
type AppPricesResponse struct {
	Data  []AppPrice         `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListPricesQuery are query options for ListPrices
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prices_for_an_app
type ListPricesQuery struct {
	FieldsAppPrices     []string `url:"fields[appPrices],omitempty"`
	FieldsApps          []string `url:"fields[apps],omitempty"`
	FieldsAppPriceTiers []string `url:"fields[appPriceTiers],omitempty"`
	Include             []string `url:"include,omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// GetPriceQuery are query options for GetPrice
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_information
type GetPriceQuery struct {
	FieldsAppPrices []string `url:"fields[appPrices],omitempty"`
	Include         []string `url:"include,omitempty"`
}

// ListPricesForApp gets current price tier of an app and any future planned price changes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prices_for_an_app
func (s *PricingService) ListPricesForApp(ctx context.Context, id string, params *ListPricesQuery) (*AppPricesResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/prices", id)
	res := new(AppPricesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetPrice reads current price and scheduled price changes for an app, including price tier and start date.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_information
func (s *PricingService) GetPrice(ctx context.Context, id string, params *GetPriceQuery) (*AppPriceResponse, *Response, error) {
	url := fmt.Sprintf("appPrices/%s", id)
	res := new(AppPriceResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
