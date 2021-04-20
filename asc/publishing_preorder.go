/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// AppPreOrder defines model for AppPreOrder.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorder
type AppPreOrder struct {
	Attributes    *AppPreOrderAttributes    `json:"attributes,omitempty"`
	ID            string                    `json:"id"`
	Links         ResourceLinks             `json:"links"`
	Relationships *AppPreOrderRelationships `json:"relationships,omitempty"`
	Type          string                    `json:"type"`
}

// AppPreOrderAttributes defines model for AppPreOrder.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorder/attributes
type AppPreOrderAttributes struct {
	AppReleaseDate        *Date `json:"appReleaseDate,omitempty"`
	PreOrderAvailableDate *Date `json:"preOrderAvailableDate,omitempty"`
}

// AppPreOrderRelationships defines model for AppPreOrder.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorder/relationships
type AppPreOrderRelationships struct {
	App *Relationship `json:"app,omitempty"`
}

// AppPreOrderCreateRequest defines model for AppPreOrderCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreordercreaterequest/data
type appPreOrderCreateRequest struct {
	Attributes    *appPreOrderCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships appPreOrderCreateRequestRelationships `json:"relationships"`
	Type          string                                `json:"type"`
}

// AppPreOrderCreateRequestAttributes are attributes for AppPreOrderCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreordercreaterequest/data/attributes
type appPreOrderCreateRequestAttributes struct {
	AppReleaseDate *Date `json:"appReleaseDate,omitempty"`
}

// AppPreOrderCreateRequestRelationships are relationships for AppPreOrderCreateRequest.
type appPreOrderCreateRequestRelationships struct {
	App relationshipDeclaration `json:"app"`
}

// AppPreOrderUpdateRequest defines model for AppPreOrderUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorderupdaterequest/data
type appPreOrderUpdateRequest struct {
	Attributes *appPreOrderUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                              `json:"id"`
	Type       string                              `json:"type"`
}

// AppPreOrderUpdateRequestAttributes are attributes for AppPreOrderUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreorderupdaterequest/data/attributes
type appPreOrderUpdateRequestAttributes struct {
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
func (s *PublishingService) CreatePreOrder(ctx context.Context, appReleaseDate *Date, appID string) (*AppPreOrderResponse, *Response, error) {
	req := appPreOrderCreateRequest{
		Relationships: appPreOrderCreateRequestRelationships{
			App: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appID,
					Type: "apps",
				},
			},
		},
		Type: "appPreOrders",
	}

	if appReleaseDate != nil {
		req.Attributes = &appPreOrderCreateRequestAttributes{
			AppReleaseDate: appReleaseDate,
		}
	}

	res := new(AppPreOrderResponse)
	resp, err := s.client.post(ctx, "appPreOrders", newRequestBody(req), res)

	return res, resp, err
}

// UpdatePreOrder updates the release date for your app pre-order.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_pre-order
func (s *PublishingService) UpdatePreOrder(ctx context.Context, id string, appReleaseDate *Date) (*AppPreOrderResponse, *Response, error) {
	req := appPreOrderUpdateRequest{
		ID:   id,
		Type: "appPreOrders",
	}

	if appReleaseDate != nil {
		req.Attributes = &appPreOrderUpdateRequestAttributes{
			AppReleaseDate: appReleaseDate,
		}
	}

	url := fmt.Sprintf("appPreOrders/%s", id)
	res := new(AppPreOrderResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeletePreOrder cancels a planned app pre-order that has not begun.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_pre-order
func (s *PublishingService) DeletePreOrder(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appPreOrders/%s", id)

	return s.client.delete(ctx, url, nil)
}
