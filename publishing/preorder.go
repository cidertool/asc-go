package publishing

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// AppPreOrder defines model for AppPreOrder.
type AppPreOrder struct {
	Attributes *struct {
		AppReleaseDate        *internal.Date `json:"appReleaseDate,omitempty"`
		PreOrderAvailableDate *internal.Date `json:"preOrderAvailableDate,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreOrderCreateRequest defines model for AppPreOrderCreateRequest.
type AppPreOrderCreateRequest struct {
	Data struct {
		Attributes *struct {
			AppReleaseDate *internal.Date `json:"appReleaseDate,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			App struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"app"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreOrderUpdateRequest defines model for AppPreOrderUpdateRequest.
type AppPreOrderUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AppReleaseDate *internal.Date `json:"appReleaseDate,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreOrderResponse defines model for AppPreOrderResponse.
type AppPreOrderResponse struct {
	Data  AppPreOrder            `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// GetPreOrderQuery are query options for GetPreOrder
type GetPreOrderQuery struct {
	FieldsAppPreOrders *[]string `url:"fields[appPreOrders],omitempty"`
	Include            *[]string `url:"include,omitempty"`
}

// GetPreOrderForAppQuery are query options for GetPreOrderForApp
type GetPreOrderForAppQuery struct {
	FieldsAppPreOrders *[]string `url:"fields[appPreOrders],omitempty"`
}

// GetPreOrder gets information about your app's pre-order configuration.
func (s *Service) GetPreOrder(id string, params *GetPreOrderQuery) (*AppPreOrderResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPreOrderForApp gets available date and release date of an app that is available for pre-order.
func (s *Service) GetPreOrderForApp(id string, params *GetPreOrderForAppQuery) (*AppPreOrderResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/preOrder", id)
	res := new(AppPreOrderResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreatePreOrder turns on pre-order and set the expected app release date.
func (s *Service) CreatePreOrder(body *AppPreOrderCreateRequest) (*AppPreOrderResponse, *internal.Response, error) {
	res := new(AppPreOrderResponse)
	resp, err := s.Post("appPreOrders", body, res)
	return res, resp, err
}

// UpdatePreOrder updates the release date for your app pre-order.
func (s *Service) UpdatePreOrder(id string, body *AppPreOrderUpdateRequest) (*AppPreOrderResponse, *internal.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeletePreOrder cancels a planned app pre-order that has not begun.
func (s *Service) DeletePreOrder(id string) (*internal.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	return s.Delete(url, nil)
}
