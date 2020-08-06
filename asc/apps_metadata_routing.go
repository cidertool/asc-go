package asc

import (
	"fmt"
)

// RoutingAppCoverage defines model for RoutingAppCoverage.
type RoutingAppCoverage struct {
	Attributes *struct {
		AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		FileName           *string             `json:"fileName,omitempty"`
		FileSize           *int64              `json:"fileSize,omitempty"`
		SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *UploadOperations   `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// RoutingAppCoverageCreateRequest defines model for RoutingAppCoverageCreateRequest.
type RoutingAppCoverageCreateRequest struct {
	Attributes    RoutingAppCoverageCreateRequestAttributes    `json:"attributes"`
	Relationships RoutingAppCoverageCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// RoutingAppCoverageCreateRequestAttributes are attributes for RoutingAppCoverageCreateRequest
type RoutingAppCoverageCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// RoutingAppCoverageCreateRequestRelationships are relationships for RoutingAppCoverageCreateRequest
type RoutingAppCoverageCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// RoutingAppCoverageResponse defines model for RoutingAppCoverageResponse.
type RoutingAppCoverageResponse struct {
	Data  RoutingAppCoverage `json:"data"`
	Links DocumentLinks      `json:"links"`
}

// RoutingAppCoverageUpdateRequest defines model for RoutingAppCoverageUpdateRequest.
type RoutingAppCoverageUpdateRequest struct {
	Attributes *RoutingAppCoverageUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// RoutingAppCoverageUpdateRequestAttributes are attributes for RoutingAppCoverageCreateRequest
type RoutingAppCoverageUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppMediaAssetState defines model for AppMediaAssetState.
type AppMediaAssetState struct {
	Errors   *[]AppMediaStateError `json:"errors,omitempty"`
	State    *string               `json:"state,omitempty"`
	Warnings *[]AppMediaStateError `json:"warnings,omitempty"`
}

// AppMediaStateError defines model for AppMediaStateError.
type AppMediaStateError struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// GetRoutingAppCoverageQuery are query options for GetRoutingAppCoverage
type GetRoutingAppCoverageQuery struct {
	FieldsRoutingAppCoverages []string `url:"fields[routingAppCoverages],omitempty"`
	Include                   []string `url:"include,omitempty"`
}

// GetRoutingAppCoverageForAppStoreVersion gets the routing app coverage file that is associated with a specific App Store version
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_routing_app_coverage_information_of_an_app_store_version
func (s *AppsService) GetRoutingAppCoverageForAppStoreVersion(id string, params *GetRoutingAppCoverageForVersionQuery) (*RoutingAppCoverageResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/routingAppCoverage", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetRoutingAppCoverage gets information about the routing app coverage file and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_routing_app_coverage_information
func (s *AppsService) GetRoutingAppCoverage(id string, params *GetRoutingAppCoverageQuery) (*RoutingAppCoverageResponse, *Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateRoutingAppCoverage attaches a routing app coverage file to an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_routing_app_coverage
func (s *AppsService) CreateRoutingAppCoverage(body *RoutingAppCoverageCreateRequest) (*RoutingAppCoverageResponse, *Response, error) {
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.post("routingAppCoverages", body, res)
	return res, resp, err
}

// CommitRoutingAppCoverage commits a routing app coverage file after uploading it.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_routing_app_coverage
func (s *AppsService) CommitRoutingAppCoverage(id string, body *RoutingAppCoverageUpdateRequest) (*RoutingAppCoverageResponse, *Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteRoutingAppCoverage deletes the routing app coverage file that is associated with a version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_routing_app_coverage
func (s *AppsService) DeleteRoutingAppCoverage(id string) (*Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	return s.client.delete(url, nil)
}
