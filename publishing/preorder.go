package publishing

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/schema"
	"github.com/aaronsky/asc-go/internal/services"
)

// AppPreOrder defines model for AppPreOrder.
type AppPreOrder struct {
	Attributes *struct {
		AppReleaseDate        *schema.Date `json:"appReleaseDate,omitempty"`
		PreOrderAvailableDate *schema.Date `json:"preOrderAvailableDate,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreOrderCreateRequest defines model for AppPreOrderCreateRequest.
type AppPreOrderCreateRequest struct {
	Attributes    *AppPreOrderCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships AppPreOrderCreateRequestRelationships `json:"relationships"`
	Type          string                                `json:"type"`
}

// AppPreOrderCreateRequestAttributes are attributes for AppPreOrderCreateRequest
type AppPreOrderCreateRequestAttributes struct {
	AppReleaseDate *schema.Date `json:"appReleaseDate,omitempty"`
}

// AppPreOrderCreateRequestRelationships are relationships for AppPreOrderCreateRequest
type AppPreOrderCreateRequestRelationships struct {
	App struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"app"`
}

// AppPreOrderUpdateRequest defines model for AppPreOrderUpdateRequest.
type AppPreOrderUpdateRequest struct {
	Attributes *AppPreOrderUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                              `json:"id"`
	Type       string                              `json:"type"`
}

// AppPreOrderUpdateRequestAttributes are attributes for AppPreOrderUpdateRequest
type AppPreOrderUpdateRequestAttributes struct {
	AppReleaseDate *schema.Date `json:"appReleaseDate,omitempty"`
}

// AppPreOrderResponse defines model for AppPreOrderResponse.
type AppPreOrderResponse struct {
	Data  AppPreOrder            `json:"data"`
	Links services.DocumentLinks `json:"links"`
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
func (s *Service) GetPreOrder(id string, params *GetPreOrderQuery) (*AppPreOrderResponse, *services.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetPreOrderForApp gets available date and release date of an app that is available for pre-order.
func (s *Service) GetPreOrderForApp(id string, params *GetPreOrderForAppQuery) (*AppPreOrderResponse, *services.Response, error) {
	url := fmt.Sprintf("apps/%s/preOrder", id)
	res := new(AppPreOrderResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreatePreOrder turns on pre-order and set the expected app release date.
func (s *Service) CreatePreOrder(body *AppPreOrderCreateRequest) (*AppPreOrderResponse, *services.Response, error) {
	res := new(AppPreOrderResponse)
	resp, err := s.Post("appPreOrders", body, res)
	return res, resp, err
}

// UpdatePreOrder updates the release date for your app pre-order.
func (s *Service) UpdatePreOrder(id string, body *AppPreOrderUpdateRequest) (*AppPreOrderResponse, *services.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeletePreOrder cancels a planned app pre-order that has not begun.
func (s *Service) DeletePreOrder(id string) (*services.Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	return s.Delete(url, nil)
}
