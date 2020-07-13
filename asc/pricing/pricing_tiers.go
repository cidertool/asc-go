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
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
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
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
		Territory *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
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

type ListAppPriceTiersOptions struct {
	Fields *struct {
		AppPricePoints *[]string `url:"appPricePoints,omitempty"`
		AppPriceTiers  *[]string `url:"appPriceTiers,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		ID *[]string `url:"id,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		PricePoints *int `url:"pricePoints,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAppPriceTierOptions struct {
	Fields *struct {
		AppPricePoints *[]string `url:"appPricePoints,omitempty"`
		AppPriceTiers  *[]string `url:"appPriceTiers,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		PricePoints *int `url:"pricePoints,omitempty"`
	} `url:"limit,omitempty"`
}

type ListPricePointsForAppPriceTierOptions struct {
	Fields *struct {
		AppPricePoints *[]string `url:"appPricePoints,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListAppPricePointsOptions struct {
	Fields *struct {
		AppPricePoints *[]string `url:"appPricePoints,omitempty"`
		Territories    *[]string `url:"territories,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		PriceTier *[]string `url:"priceTier,omitempty"`
		Territory *[]string `url:"territory,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
}

type GetTerritoryForAppPricePointOptions struct {
	Fields *struct {
		Territories *[]string `url:"territories,omitempty"`
	} `url:"fields,omitempty"`
}

type GetAppPricePointOptions struct {
	Fields *struct {
		AppPricePoints *[]string `url:"appPricePoints,omitempty"`
		Territories    *[]string `url:"territories,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		PriceTier *[]string `url:"priceTier,omitempty"`
		Territory *[]string `url:"territory,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// ListAppPriceTiers lists all app price tiers available in App Store Connect, including related price points.
func (s *Service) ListAppPriceTiers(params *ListAppPriceTiersOptions) (*AppPriceTiersResponse, *common.Response, error) {
	res := new(AppPriceTiersResponse)
	resp, err := s.Get("appPriceTiers", params, res)
	return res, resp, err
}

// GetAppPriceTier reads available app price tiers.
func (s *Service) GetAppPriceTier(id string, params *GetAppPriceTierOptions) (*AppPriceTierResponse, *common.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s", id)
	res := new(AppPriceTierResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListPricePointsForAppPriceTier lists price points across all App Store territories for a specific price tier.
func (s *Service) ListPricePointsForAppPriceTier(id string, params *ListPricePointsForAppPriceTierOptions) (*AppPricePointsResponse, *common.Response, error) {
	url := fmt.Sprintf("appPriceTiers/%s/pricePoints", id)
	res := new(AppPricePointsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListAppPricePoints lists all app price points available in App Store Connect, including related price tier, developer proceeds, and territory.
func (s *Service) ListAppPricePoints(params *ListAppPricePointsOptions) (*AppPricePointsResponse, *common.Response, error) {
	res := new(AppPricePointsResponse)
	resp, err := s.Get("appPricePoints", params, res)
	return res, resp, err
}

// GetTerritoryForAppPricePoint gets the territory in which a specific price point applies.
func (s *Service) GetTerritoryForAppPricePoint(id string, params *GetTerritoryForAppPricePointOptions) (*TerritoryResponse, *common.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s/territory", id)
	res := new(TerritoryResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppPricePoint reads the customer prices and your proceeds for a price tier.
func (s *Service) GetAppPricePoint(id string, params *GetAppPricePointOptions) (*AppPricePointResponse, *common.Response, error) {
	url := fmt.Sprintf("appPricePoints/%s", id)
	res := new(AppPricePointResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
