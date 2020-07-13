package publishing

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppPreOrder defines model for AppPreOrder.
type AppPreOrder struct {
	Attributes *struct {
		AppReleaseDate        *common.Date `json:"appReleaseDate,omitempty"`
		PreOrderAvailableDate *common.Date `json:"preOrderAvailableDate,omitempty"`
	} `json:"attributes,omitempty"`
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
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreOrderCreateRequest defines model for AppPreOrderCreateRequest.
type AppPreOrderCreateRequest struct {
	Data struct {
		Attributes *struct {
			AppReleaseDate *common.Date `json:"appReleaseDate,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			App struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"app"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreOrderUpdateRequest defines model for AppPreOrderUpdateRequest.
type AppPreOrderUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AppReleaseDate *common.Date `json:"appReleaseDate,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreOrderResponse defines model for AppPreOrderResponse.
type AppPreOrderResponse struct {
	Data  AppPreOrder          `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// GetPreOrderOptions are query options for GetPreOrder
type GetPreOrderOptions struct {
	Fields *struct {
		appPreOrders *[]string `url:"appPreOrders,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetPreOrderForAppOptions struct {
	Fields *struct {
		AppPreOrders *[]string `url:"appPreOrders,omitempty"`
	} `url:"fields,omitempty"`
}

// GetPreOrder gets information about your app's pre-order configuration.
func (s *Service) GetPreOrder(id string, params *GetPreOrderOptions) (*AppPreOrderResponse, *common.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetPreOrderForApp gets available date and release date of an app that is available for pre-order.
func (s *Service) GetPreOrderForApp(id string, params *GetPreOrderForAppOptions) (*AppPreOrderResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/preOrder", id)
	res := new(AppPreOrderResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// CreatePreOrder turns on pre-order and set the expected app release date.
func (s *Service) CreatePreOrder(body *AppPreOrderCreateRequest) (*AppPreOrderResponse, *common.Response, error) {
	res := new(AppPreOrderResponse)
	resp, err := s.Post("appPreOrders", body, res)
	return res, resp, err
}

// UpdatePreOrder updates the release date for your app pre-order.
func (s *Service) UpdatePreOrder(id string, body *AppPreOrderUpdateRequest) (*AppPreOrderResponse, *common.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeletePreOrder cancels a planned app pre-order that has not begun.
func (s *Service) DeletePreOrder(id string) (*common.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	return s.Delete(url, nil)
}
