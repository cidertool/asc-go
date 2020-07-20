package pricing

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppPriceTier defines model for AppPriceTier.
type AppPriceTier struct {
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		PricePoints *struct {
			Data  *[]common.RelationshipsData `json:"data,omitempty"`
			Links *common.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *common.PagingInformation   `json:"meta,omitempty"`
		} `json:"pricePoints,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceTierResponse defines model for AppPriceTierResponse.
type AppPriceTierResponse struct {
	Data     AppPriceTier         `json:"data"`
	Included *[]AppPricePoint     `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// AppPriceTiersResponse defines model for AppPriceTiersResponse.
type AppPriceTiersResponse struct {
	Data     []AppPriceTier            `json:"data"`
	Included *[]AppPricePoint          `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// AppPricePoint defines model for AppPricePoint.
type AppPricePoint struct {
	Attributes *struct {
		CustomerPrice *string `json:"customerPrice,omitempty"`
		Proceeds      *string `json:"proceeds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		PriceTier *struct {
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
		Territory *struct {
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"territory,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPricePointResponse defines model for AppPricePointResponse.
type AppPricePointResponse struct {
	Data     AppPricePoint        `json:"data"`
	Included *[]Territory         `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// AppPricePointsResponse defines model for AppPricePointsResponse.
type AppPricePointsResponse struct {
	Data     []AppPricePoint           `json:"data"`
	Included *[]Territory              `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListAppPriceTiersQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  *[]string `url:"fields[appPriceTiers],omitempty"`
	FilterID             *[]string `url:"filter[id],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	LimitPricePoints     *int      `url:"limit[pricePoints],omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

type GetAppPriceTierQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsAppPriceTiers  *[]string `url:"fields[appPriceTiers],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	LimitPricePoints     *int      `url:"limit[pricePoints],omitempty"`
}

type ListPricePointsForAppPriceTierQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

type ListAppPricePointsQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    *[]string `url:"fields[territories],omitempty"`
	FilterPriceTier      *[]string `url:"filter[priceTier],omitempty"`
	FilterTerritory      *[]string `url:"filter[territory],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	Limit                *int      `url:"limit,omitempty"`
	Cursor               *string   `url:"cursor,omitempty"`
}

type GetTerritoryForAppPricePointQuery struct {
	FieldsTerritories *[]string `url:"fields[territories],omitempty"`
}

type GetAppPricePointQuery struct {
	FieldsAppPricePoints *[]string `url:"fields[appPricePoints],omitempty"`
	FieldsTerritories    *[]string `url:"fields[territories],omitempty"`
	FilterPriceTier      *[]string `url:"filter[priceTier],omitempty"`
	FilterTerritory      *[]string `url:"filter[territory],omitempty"`
	Include              *[]string `url:"include,omitempty"`
}

// ListAppPriceTiers lists all app price tiers available in App Store Connect, including related price points.
func (s *Service) ListAppPriceTiers(params *ListAppPriceTiersQuery) (*AppPriceTiersResponse, *common.Response, error) {
	res := new(AppPriceTiersResponse)
	resp, err := s.GetWithQuery("appPriceTiers", params, res)
	return res, resp, err
}

// GetAppPriceTier reads available app price tiers.
func (s *Service) GetAppPriceTier(id string, params *GetAppPriceTierQuery) (*AppPriceTierResponse, *common.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s", id)
	res := new(AppPriceTierResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListPricePointsForAppPriceTier lists price points across all App Store territories for a specific price tier.
func (s *Service) ListPricePointsForAppPriceTier(id string, params *ListPricePointsForAppPriceTierQuery) (*AppPricePointsResponse, *common.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s/pricePoints", id)
	res := new(AppPricePointsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppPricePoints lists all app price points available in App Store Connect, including related price tier, developer proceeds, and territory.
func (s *Service) ListAppPricePoints(params *ListAppPricePointsQuery) (*AppPricePointsResponse, *common.Response, error) {
	res := new(AppPricePointsResponse)
	resp, err := s.GetWithQuery("appPricePoints", params, res)
	return res, resp, err
}

// GetTerritoryForAppPricePoint gets the territory in which a specific price point applies.
func (s *Service) GetTerritoryForAppPricePoint(id string, params *GetTerritoryForAppPricePointQuery) (*TerritoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppPricePoint reads the customer prices and your proceeds for a price tier.
func (s *Service) GetAppPricePoint(id string, params *GetAppPricePointQuery) (*AppPricePointResponse, *common.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s", id)
	res := new(AppPricePointResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
