package pricing

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppPrice defines model for AppPrice.
type AppPrice struct {
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"app,omitempty"`
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
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPriceResponse defines model for AppPriceResponse.
type AppPriceResponse struct {
	Data  AppPrice             `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// AppPricesResponse defines model for AppPricesResponse.
type AppPricesResponse struct {
	Data  []AppPrice                `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type ListPricesQuery struct {
	Fields *struct {
		AppPrices     *[]string `url:"appPrices,omitempty"`
		Apps          *[]string `url:"apps,omitempty"`
		AppPriceTiers *[]string `url:"appPriceTiers,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Cursor  *string   `url:"cursor,omitempty"`
}

type GetPriceQuery struct {
	Fields *struct {
		AppPrices *[]string `url:"appPrices,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// ListPricesForApp gets current price tier of an app and any future planned price changes.
func (s *Service) ListPricesForApp(id string, params *ListPricesQuery) (*AppPricesResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/prices", id)
	res := new(AppPricesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPrice reads current price and scheduled price changes for an app, including price tier and start date.
func (s *Service) GetPrice(id string, params *GetPriceQuery) (*AppPriceResponse, *common.Response, error) {
	url := fmt.Sprintf("appPrices/%s", id)
	res := new(AppPriceResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
