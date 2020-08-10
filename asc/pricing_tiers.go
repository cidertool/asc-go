package asc

import (
	"context"
	"fmt"
)

// AppPriceTier defines model for AppPriceTier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricetier
type AppPriceTier struct {
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		PricePoints *PagedRelationship `json:"pricePoints,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceTierResponse defines model for AppPriceTierResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricetierresponse
type AppPriceTierResponse struct {
	Data     AppPriceTier     `json:"data"`
	Included *[]AppPricePoint `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}

// AppPriceTiersResponse defines model for AppPriceTiersResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricetiersresponse
type AppPriceTiersResponse struct {
	Data     []AppPriceTier     `json:"data"`
	Included *[]AppPricePoint   `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppPricePoint defines model for AppPricePoint.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricepoint
type AppPricePoint struct {
	Attributes *struct {
		CustomerPrice *string `json:"customerPrice,omitempty"`
		Proceeds      *string `json:"proceeds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		PriceTier *Relationship `json:"priceTier,omitempty"`
		Territory *Relationship `json:"territory,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPricePointResponse defines model for AppPricePointResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricepointresponse
type AppPricePointResponse struct {
	Data     AppPricePoint `json:"data"`
	Included *[]Territory  `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// AppPricePointsResponse defines model for AppPricePointsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppricepointsresponse
type AppPricePointsResponse struct {
	Data     []AppPricePoint    `json:"data"`
	Included *[]Territory       `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListAppPriceTiersQuery are query options for ListAppPriceTiers
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_price_tiers
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_tier_information
type GetAppPriceTierQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  []string `url:"fields[appPriceTiers],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitPricePoints     int      `url:"limit[pricePoints],omitempty"`
}

// ListPricePointsForAppPriceTierQuery are query options for ListPricePointsForAppPriceTier
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_price_points_for_an_app_price_tier
type ListPricePointsForAppPriceTierQuery struct {
	FieldsAppPricePoints []string `url:"fields[appPricePoints],omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Cursor               string   `url:"cursor,omitempty"`
}

// ListAppPricePointsQuery are query options for ListAppPricePoints
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_price_points
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_territory_information_of_an_app_price_point
type GetTerritoryForAppPricePointQuery struct {
	FieldsTerritories []string `url:"fields[territories],omitempty"`
}

// GetAppPricePointQuery are query options for GetAppPricePoint
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_point_information
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
func (s *PricingService) ListAppPriceTiers(ctx context.Context, params *ListAppPriceTiersQuery) (*AppPriceTiersResponse, *Response, error) {
	res := new(AppPriceTiersResponse)
	resp, err := s.client.get(ctx, "appPriceTiers", params, res)
	return res, resp, err
}

// GetAppPriceTier reads available app price tiers.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_tier_information
func (s *PricingService) GetAppPriceTier(ctx context.Context, id string, params *GetAppPriceTierQuery) (*AppPriceTierResponse, *Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s", id)
	res := new(AppPriceTierResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListPricePointsForAppPriceTier lists price points across all App Store territories for a specific price tier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_price_points_for_an_app_price_tier
func (s *PricingService) ListPricePointsForAppPriceTier(ctx context.Context, id string, params *ListPricePointsForAppPriceTierQuery) (*AppPricePointsResponse, *Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s/pricePoints", id)
	res := new(AppPricePointsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListAppPricePoints lists all app price points available in App Store Connect, including related price tier, developer proceeds, and territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_price_points
func (s *PricingService) ListAppPricePoints(ctx context.Context, params *ListAppPricePointsQuery) (*AppPricePointsResponse, *Response, error) {
	res := new(AppPricePointsResponse)
	resp, err := s.client.get(ctx, "appPricePoints", params, res)
	return res, resp, err
}

// GetTerritoryForAppPricePoint gets the territory in which a specific price point applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_territory_information_of_an_app_price_point
func (s *PricingService) GetTerritoryForAppPricePoint(ctx context.Context, id string, params *GetTerritoryForAppPricePointQuery) (*TerritoryResponse, *Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppPricePoint reads the customer prices and your proceeds for a price tier.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_price_point_information
func (s *PricingService) GetAppPricePoint(ctx context.Context, id string, params *GetAppPricePointQuery) (*AppPricePointResponse, *Response, error) {
	url := fmt.Sprintf("appPricePoints/%s", id)
	res := new(AppPricePointResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
