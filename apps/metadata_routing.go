package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// RoutingAppCoverage defines model for RoutingAppCoverage.
type RoutingAppCoverage struct {
	Attributes *struct {
		AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		FileName           *string             `json:"fileName,omitempty"`
		FileSize           *int                `json:"fileSize,omitempty"`
		SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *[]UploadOperation  `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
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

type RoutingAppCoverageCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
}

type RoutingAppCoverageCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data types.RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// RoutingAppCoverageResponse defines model for RoutingAppCoverageResponse.
type RoutingAppCoverageResponse struct {
	Data  RoutingAppCoverage  `json:"data"`
	Links types.DocumentLinks `json:"links"`
}

// RoutingAppCoverageUpdateRequest defines model for RoutingAppCoverageUpdateRequest.
type RoutingAppCoverageUpdateRequest struct {
	Attributes *RoutingAppCoverageUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

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
	FieldsRoutingAppCoverages *[]string `url:"fields[routingAppCoverages],omitempty"`
	Include                   *[]string `url:"include,omitempty"`
}

// GetRoutingAppCoverage gets information about the routing app coverage file and its upload and processing status.
func (s *Service) GetRoutingAppCoverage(id string, params *GetRoutingAppCoverageQuery) (*RoutingAppCoverageResponse, *services.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateRoutingAppCoverage attaches a routing app coverage file to an App Store version.
func (s *Service) CreateRoutingAppCoverage(body *RoutingAppCoverageCreateRequest) (*RoutingAppCoverageResponse, *services.Response, error) {
	res := new(RoutingAppCoverageResponse)
	resp, err := s.Post("routingAppCoverages", body, res)
	return res, resp, err
}

// UpdateRoutingAppCoverage commits a routing app coverage file after uploading it.
func (s *Service) UpdateRoutingAppCoverage(id string, body *RoutingAppCoverageUpdateRequest) (*RoutingAppCoverageResponse, *services.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteRoutingAppCoverage deletes the routing app coverage file that is associated with a version.
func (s *Service) DeleteRoutingAppCoverage(id string) (*services.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	return s.Delete(url, nil)
}
