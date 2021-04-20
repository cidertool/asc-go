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

// RoutingAppCoverage defines model for RoutingAppCoverage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverage
type RoutingAppCoverage struct {
	Attributes    *RoutingAppCoverageAttributes    `json:"attributes,omitempty"`
	ID            string                           `json:"id"`
	Links         ResourceLinks                    `json:"links"`
	Relationships *RoutingAppCoverageRelationships `json:"relationships,omitempty"`
	Type          string                           `json:"type"`
}

// RoutingAppCoverageAttributes defines model for RoutingAppCoverage.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverage/attributes
type RoutingAppCoverageAttributes struct {
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	FileName           *string             `json:"fileName,omitempty"`
	FileSize           *int64              `json:"fileSize,omitempty"`
	SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
	UploadOperations   []UploadOperation   `json:"uploadOperations,omitempty"`
}

// RoutingAppCoverageRelationships defines model for RoutingAppCoverage.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverage/relationships
type RoutingAppCoverageRelationships struct {
	AppStoreVersion *Relationship `json:"appStoreVersion,omitempty"`
}

// RoutingAppCoverageCreateRequest defines model for RoutingAppCoverageCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoveragecreaterequest/data
type routingAppCoverageCreateRequest struct {
	Attributes    routingAppCoverageCreateRequestAttributes    `json:"attributes"`
	Relationships routingAppCoverageCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// RoutingAppCoverageCreateRequestAttributes are attributes for RoutingAppCoverageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoveragecreaterequest/data/attributes
type routingAppCoverageCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// RoutingAppCoverageCreateRequestRelationships are relationships for RoutingAppCoverageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoveragecreaterequest/data/relationships
type routingAppCoverageCreateRequestRelationships struct {
	AppStoreVersion relationshipDeclaration `json:"appStoreVersion"`
}

// RoutingAppCoverageResponse defines model for RoutingAppCoverageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverageresponse
type RoutingAppCoverageResponse struct {
	Data  RoutingAppCoverage `json:"data"`
	Links DocumentLinks      `json:"links"`
}

// RoutingAppCoverageUpdateRequest defines model for RoutingAppCoverageUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverageupdaterequest/data
type routingAppCoverageUpdateRequest struct {
	Attributes *routingAppCoverageUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// RoutingAppCoverageUpdateRequestAttributes are attributes for RoutingAppCoverageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/routingappcoverageupdaterequest/data/attributes
type routingAppCoverageUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppMediaAssetState defines model for AppMediaAssetState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appmediastateerror
type AppMediaAssetState struct {
	Errors   []AppMediaStateError `json:"errors,omitempty"`
	State    *string              `json:"state,omitempty"`
	Warnings []AppMediaStateError `json:"warnings,omitempty"`
}

// AppMediaStateError defines model for AppMediaStateError.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appmediaassetstate
type AppMediaStateError struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// GetRoutingAppCoverageForVersionQuery are query options for GetRoutingAppCoverageForVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_routing_app_coverage_information_of_an_app_store_version
type GetRoutingAppCoverageForVersionQuery struct {
	FieldsRoutingAppCoverages []string `url:"fields[routingAppCoverages],omitempty"`
}

// GetRoutingAppCoverageQuery are query options for GetRoutingAppCoverage
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_routing_app_coverage_information
type GetRoutingAppCoverageQuery struct {
	FieldsRoutingAppCoverages []string `url:"fields[routingAppCoverages],omitempty"`
	Include                   []string `url:"include,omitempty"`
}

// GetRoutingAppCoverageForAppStoreVersion gets the routing app coverage file that is associated with a specific App Store version
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_routing_app_coverage_information_of_an_app_store_version
func (s *AppsService) GetRoutingAppCoverageForAppStoreVersion(ctx context.Context, id string, params *GetRoutingAppCoverageForVersionQuery) (*RoutingAppCoverageResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/routingAppCoverage", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetRoutingAppCoverage gets information about the routing app coverage file and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_routing_app_coverage_information
func (s *AppsService) GetRoutingAppCoverage(ctx context.Context, id string, params *GetRoutingAppCoverageQuery) (*RoutingAppCoverageResponse, *Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateRoutingAppCoverage attaches a routing app coverage file to an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_routing_app_coverage
func (s *AppsService) CreateRoutingAppCoverage(ctx context.Context, fileName string, fileSize int64, appStoreVersionID string) (*RoutingAppCoverageResponse, *Response, error) {
	req := routingAppCoverageCreateRequest{
		Attributes: routingAppCoverageCreateRequestAttributes{
			FileName: fileName,
			FileSize: fileSize,
		},
		Relationships: routingAppCoverageCreateRequestRelationships{
			AppStoreVersion: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "routingAppCoverages",
	}
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.post(ctx, "routingAppCoverages", newRequestBody(req), res)

	return res, resp, err
}

// CommitRoutingAppCoverage commits a routing app coverage file after uploading it.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_routing_app_coverage
func (s *AppsService) CommitRoutingAppCoverage(ctx context.Context, id string, uploaded *bool, sourceFileChecksum *string) (*RoutingAppCoverageResponse, *Response, error) {
	req := routingAppCoverageUpdateRequest{
		ID:   id,
		Type: "routingAppCoverages",
	}

	if uploaded != nil || sourceFileChecksum != nil {
		req.Attributes = &routingAppCoverageUpdateRequestAttributes{
			Uploaded:           uploaded,
			SourceFileChecksum: sourceFileChecksum,
		}
	}

	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteRoutingAppCoverage deletes the routing app coverage file that is associated with a version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_routing_app_coverage
func (s *AppsService) DeleteRoutingAppCoverage(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)

	return s.client.delete(ctx, url, nil)
}
