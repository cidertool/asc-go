package pricing

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
)

// AppPriceTier defines model for AppPriceTier.
type AppPriceTier struct {
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		PricePoints *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"pricePoints,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceTierResponse defines model for AppPriceTierResponse.
type AppPriceTierResponse struct {
	Data     AppPriceTier           `json:"data"`
	Included *[]AppPricePoint       `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// AppPriceTiersResponse defines model for AppPriceTiersResponse.
type AppPriceTiersResponse struct {
	Data     []AppPriceTier              `json:"data"`
	Included *[]AppPricePoint            `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// AppPricePoint defines model for AppPricePoint.
type AppPricePoint struct {
	Attributes *struct {
		CustomerPrice *string `json:"customerPrice,omitempty"`
		Proceeds      *string `json:"proceeds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		PriceTier *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
		Territory *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"territory,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPricePointResponse defines model for AppPricePointResponse.
type AppPricePointResponse struct {
	Data     AppPricePoint          `json:"data"`
	Included *[]Territory           `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// AppPricePointsResponse defines model for AppPricePointsResponse.
type AppPricePointsResponse struct {
	Data     []AppPricePoint             `json:"data"`
	Included *[]Territory                `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// ListAppPriceTiersQuery are query options for ListAppPriceTiers
type ListAppPriceTiersQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  *[]string `url:"fields[appPriceTiers],omitempty"`
	FilterID             *[]string `url:"filter[id],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	LimitPricePoints     *int      `url:"limit[pricePoints],omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

// GetAppPriceTierQuery are query options for GetAppPriceTier
type GetAppPriceTierQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  *[]string `url:"fields[appPriceTiers],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	LimitPricePoints     *int      `url:"limit[pricePoints],omitempty"`
}

// ListPricePointsForAppPriceTierQuery are query options for ListPricePointsForAppPriceTier
type ListPricePointsForAppPriceTierQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

// ListAppPricePointsQuery are query options for ListAppPricePoints
type ListAppPricePointsQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    *[]string `url:"fields[territories],omitempty"`
	FilterPriceTier      *[]string `url:"filter[priceTier],omitempty"`
	FilterTerritory      *[]string `url:"filter[territory],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

// GetTerritoryForAppPricePointQuery are query options for GetTerritoryForAppPricePoint
type GetTerritoryForAppPricePointQuery struct {
	FieldsTerritories *[]string `url:"fields[territories],omitempty"`
}

// GetAppPricePointQuery are query options for GetAppPricePoint
type GetAppPricePointQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    *[]string `url:"fields[territories],omitempty"`
	FilterPriceTier      *[]string `url:"filter[priceTier],omitempty"`
	FilterTerritory      *[]string `url:"filter[territory],omitempty"`
	Include              *[]string `url:"include,omitempty"`
}

// ListAppPriceTiers lists all app price tiers available in App Store Connect, including related price points.
func (s *Service) ListAppPriceTiers(params *ListAppPriceTiersQuery) (*AppPriceTiersResponse, *services.Response, error) {
	res := new(AppPriceTiersResponse)
	resp, err := s.GetWithQuery("appPriceTiers", params, res)
	return res, resp, err
}

// GetAppPriceTier reads available app price tiers.
func (s *Service) GetAppPriceTier(id string, params *GetAppPriceTierQuery) (*AppPriceTierResponse, *services.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s", id)
	res := new(AppPriceTierResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListPricePointsForAppPriceTier lists price points across all App Store territories for a specific price tier.
func (s *Service) ListPricePointsForAppPriceTier(id string, params *ListPricePointsForAppPriceTierQuery) (*AppPricePointsResponse, *services.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s/pricePoints", id)
	res := new(AppPricePointsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppPricePoints lists all app price points available in App Store Connect, including related price tier, developer proceeds, and territory.
func (s *Service) ListAppPricePoints(params *ListAppPricePointsQuery) (*AppPricePointsResponse, *services.Response, error) {
	res := new(AppPricePointsResponse)
	resp, err := s.GetWithQuery("appPricePoints", params, res)
	return res, resp, err
}

// GetTerritoryForAppPricePoint gets the territory in which a specific price point applies.
func (s *Service) GetTerritoryForAppPricePoint(id string, params *GetTerritoryForAppPricePointQuery) (*TerritoryResponse, *services.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppPricePoint reads the customer prices and your proceeds for a price tier.
func (s *Service) GetAppPricePoint(id string, params *GetAppPricePointQuery) (*AppPricePointResponse, *services.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s", id)
	res := new(AppPricePointResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
