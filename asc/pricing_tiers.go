package asc

import (
	"fmt"
)

// AppPriceTier defines model for AppPriceTier.
type AppPriceTier struct {
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		PricePoints *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"pricePoints,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceTierResponse defines model for AppPriceTierResponse.
type AppPriceTierResponse struct {
	Data     AppPriceTier     `json:"data"`
	Included *[]AppPricePoint `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}

// AppPriceTiersResponse defines model for AppPriceTiersResponse.
type AppPriceTiersResponse struct {
	Data     []AppPriceTier     `json:"data"`
	Included *[]AppPricePoint   `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppPricePoint defines model for AppPricePoint.
type AppPricePoint struct {
	Attributes *struct {
		CustomerPrice *string `json:"customerPrice,omitempty"`
		Proceeds      *string `json:"proceeds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		PriceTier *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
		Territory *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"territory,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPricePointResponse defines model for AppPricePointResponse.
type AppPricePointResponse struct {
	Data     AppPricePoint `json:"data"`
	Included *[]Territory  `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// AppPricePointsResponse defines model for AppPricePointsResponse.
type AppPricePointsResponse struct {
	Data     []AppPricePoint    `json:"data"`
	Included *[]Territory       `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListAppPriceTiersQuery are query options for ListAppPriceTiers
type ListAppPriceTiersQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  []string `url:"fields[appPriceTiers],omitempty"`
	FilterID             []string `url:"filter[id],omitempty"`
	Include              []string `url:"include,omitempty"`
	Limit                int      `url:"limit,omitempty"`
	LimitPricePoints     int      `url:"limit[pricePoints],omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// GetAppPriceTierQuery are query options for GetAppPriceTier
type GetAppPriceTierQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  []string `url:"fields[appPriceTiers],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitPricePoints     int      `url:"limit[pricePoints],omitempty"`
}

// ListPricePointsForAppPriceTierQuery are query options for ListPricePointsForAppPriceTier
type ListPricePointsForAppPriceTierQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// ListAppPricePointsQuery are query options for ListAppPricePoints
type ListAppPricePointsQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    []string `url:"fields[territories],omitempty"`
	FilterPriceTier      []string `url:"filter[priceTier],omitempty"`
	FilterTerritory      []string `url:"filter[territory],omitempty"`
	Include              []string `url:"include,omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// GetTerritoryForAppPricePointQuery are query options for GetTerritoryForAppPricePoint
type GetTerritoryForAppPricePointQuery struct {
	FieldsTerritories []string `url:"fields[territories],omitempty"`
}

// GetAppPricePointQuery are query options for GetAppPricePoint
type GetAppPricePointQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    []string `url:"fields[territories],omitempty"`
	FilterPriceTier      []string `url:"filter[priceTier],omitempty"`
	FilterTerritory      []string `url:"filter[territory],omitempty"`
	Include              []string `url:"include,omitempty"`
}

// ListAppPriceTiers lists all app price tiers available in App Store Connect, including related price points.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_price_tiers
func (s *PricingService) ListAppPriceTiers(params *ListAppPriceTiersQuery) (*AppPriceTiersResponse, *Response, error) {
	res := new(AppPriceTiersResponse)
	resp, err := s.client.get("appPriceTiers", params, res)
	return res, resp, err
}

// GetAppPriceTier reads available app price tiers.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_tier_information
func (s *PricingService) GetAppPriceTier(id string, params *GetAppPriceTierQuery) (*AppPriceTierResponse, *Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s", id)
	res := new(AppPriceTierResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListPricePointsForAppPriceTier lists price points across all App Store territories for a specific price tier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_price_points_for_an_app_price_tier
func (s *PricingService) ListPricePointsForAppPriceTier(id string, params *ListPricePointsForAppPriceTierQuery) (*AppPricePointsResponse, *Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s/pricePoints", id)
	res := new(AppPricePointsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListAppPricePoints lists all app price points available in App Store Connect, including related price tier, developer proceeds, and territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_price_points
func (s *PricingService) ListAppPricePoints(params *ListAppPricePointsQuery) (*AppPricePointsResponse, *Response, error) {
	res := new(AppPricePointsResponse)
	resp, err := s.client.get("appPricePoints", params, res)
	return res, resp, err
}

// GetTerritoryForAppPricePoint gets the territory in which a specific price point applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_territory_information_of_an_app_price_point
func (s *PricingService) GetTerritoryForAppPricePoint(id string, params *GetTerritoryForAppPricePointQuery) (*TerritoryResponse, *Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppPricePoint reads the customer prices and your proceeds for a price tier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_point_information
func (s *PricingService) GetAppPricePoint(id string, params *GetAppPricePointQuery) (*AppPricePointResponse, *Response, error) {
	url := fmt.Sprintf("appPricePoints/%s", id)
	res := new(AppPricePointResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
