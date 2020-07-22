package pricing

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// AppPrice defines model for AppPrice.
type AppPrice struct {
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		PriceTier *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"priceTier,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceResponse defines model for AppPriceResponse.
type AppPriceResponse struct {
	Data  AppPrice               `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// AppPricesResponse defines model for AppPricesResponse.
type AppPricesResponse struct {
	Data  []AppPrice                  `json:"data"`
	Links internal.PagedDocumentLinks `json:"links"`
	Meta  *internal.PagingInformation `json:"meta,omitempty"`
}

type ListPricesQuery struct {
	FieldsAppPrices     *[]string `url:"fields[appPrices],omitempty"`
	FieldsApps          *[]string `url:"fields[apps],omitempty"`
	FieldsAppPriceTiers *[]string `url:"fields[appPriceTiers],omitempty"`
	Include             *[]string `url:"include,omitempty"`
	Limit               *int      `url:"limit,omitempty"`
	Cursor              *string   `url:"cursor,omitempty"`
}

type GetPriceQuery struct {
	FieldsAppPrices *[]string `url:"fields[appPrices],omitempty"`
	Include         *[]string `url:"include,omitempty"`
}

// ListPricesForApp gets current price tier of an app and any future planned price changes.
func (s *Service) ListPricesForApp(id string, params *ListPricesQuery) (*AppPricesResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/prices", id)
	res := new(AppPricesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPrice reads current price and scheduled price changes for an app, including price tier and start date.
func (s *Service) GetPrice(id string, params *GetPriceQuery) (*AppPriceResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPrices/%s", id)
	res := new(AppPriceResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
