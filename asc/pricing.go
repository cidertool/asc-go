package asc

import (
	"fmt"
	"net/http"
)

// PricingService handles communication with pricing-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/app_prices
// https://developer.apple.com/documentation/appstoreconnectapi/territories
// https://developer.apple.com/documentation/appstoreconnectapi/app_price_reference_data
type PricingService service

// AppPrice defines model for AppPrice.
type AppPrice struct {
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		PriceTier *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceResponse defines model for AppPriceResponse.
type AppPriceResponse struct {
	Data  AppPrice      `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppPricesResponse defines model for AppPricesResponse.
type AppPricesResponse struct {
	Data  []AppPrice         `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListPricesQuery are query options for ListPrices
type ListPricesQuery struct {
	FieldsAppPrices     []string `url:"fields[appPrices],omitempty"`
	FieldsApps          []string `url:"fields[apps],omitempty"`
	FieldsAppPriceTiers []string `url:"fields[appPriceTiers],omitempty"`
	Include             []string `url:"include,omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// GetPriceQuery are query options for GetPrice
type GetPriceQuery struct {
	FieldsAppPrices []string `url:"fields[appPrices],omitempty"`
	Include         []string `url:"include,omitempty"`
}

// ListPricesForApp gets current price tier of an app and any future planned price changes.
func (s *PricingService) ListPricesForApp(id string, params *ListPricesQuery) (*AppPricesResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/prices", id)
	res := new(AppPricesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetPrice reads current price and scheduled price changes for an app, including price tier and start date.
func (s *PricingService) GetPrice(id string, params *GetPriceQuery) (*AppPriceResponse, *http.Response, error) {
	url := fmt.Sprintf("appPrices/%s", id)
	res := new(AppPriceResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
