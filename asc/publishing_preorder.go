package asc

import (
	"context"
	"fmt"
)

// AppPreOrder defines model for AppPreOrder.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorder
type AppPreOrder struct {
	Attributes *struct {
		AppReleaseDate        *Date `json:"appReleaseDate,omitempty"`
		PreOrderAvailableDate *Date `json:"preOrderAvailableDate,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *Relationship `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreOrderCreateRequest defines model for AppPreOrderCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreordercreaterequest
type AppPreOrderCreateRequest struct {
	Attributes    *AppPreOrderCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships AppPreOrderCreateRequestRelationships `json:"relationships"`
	Type          string                                `json:"type"`
}

// AppPreOrderCreateRequestAttributes are attributes for AppPreOrderCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreordercreaterequest/data/attributes
type AppPreOrderCreateRequestAttributes struct {
	AppReleaseDate *Date `json:"appReleaseDate,omitempty"`
}

// AppPreOrderCreateRequestRelationships are relationships for AppPreOrderCreateRequest
type AppPreOrderCreateRequestRelationships struct {
	App RelationshipDeclaration `json:"app"`
}

// AppPreOrderUpdateRequest defines model for AppPreOrderUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorderupdaterequest
type AppPreOrderUpdateRequest struct {
	Attributes *AppPreOrderUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                              `json:"id"`
	Type       string                              `json:"type"`
}

// AppPreOrderUpdateRequestAttributes are attributes for AppPreOrderUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorderupdaterequest/data/attributes
type AppPreOrderUpdateRequestAttributes struct {
	AppReleaseDate *Date `json:"appReleaseDate,omitempty"`
}

// AppPreOrderResponse defines model for AppPreOrderResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorderresponse
type AppPreOrderResponse struct {
	Data  AppPreOrder   `json:"data"`
	Links DocumentLinks `json:"links"`
}

// GetPreOrderQuery are query options for GetPreOrder
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_pre-order_information
type GetPreOrderQuery struct {
	FieldsAppPreOrders []string `url:"fields[appPreOrders],omitempty"`
	Include            []string `url:"include,omitempty"`
}

// GetPreOrderForAppQuery are query options for GetPreOrderForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_pre-order_information_of_an_app
type GetPreOrderForAppQuery struct {
	FieldsAppPreOrders []string `url:"fields[appPreOrders],omitempty"`
}

// GetPreOrder gets information about your app's pre-order configuration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_pre-order_information
func (s *PublishingService) GetPreOrder(ctx context.Context, id string, params *GetPreOrderQuery) (*AppPreOrderResponse, *Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetPreOrderForApp gets available date and release date of an app that is available for pre-order.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_pre-order_information_of_an_app
func (s *PublishingService) GetPreOrderForApp(ctx context.Context, id string, params *GetPreOrderForAppQuery) (*AppPreOrderResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/preOrder", id)
	res := new(AppPreOrderResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreatePreOrder turns on pre-order and set the expected app release date.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_pre-order
func (s *PublishingService) CreatePreOrder(ctx context.Context, body *AppPreOrderCreateRequest) (*AppPreOrderResponse, *Response, error) {
	res := new(AppPreOrderResponse)
	resp, err := s.client.post(ctx, "appPreOrders", body, res)
	return res, resp, err
}

// UpdatePreOrder updates the release date for your app pre-order.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_pre-order
func (s *PublishingService) UpdatePreOrder(ctx context.Context, id string, body *AppPreOrderUpdateRequest) (*AppPreOrderResponse, *Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeletePreOrder cancels a planned app pre-order that has not begun.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_pre-order
func (s *PublishingService) DeletePreOrder(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)
	return s.client.delete(ctx, url, nil)
}
