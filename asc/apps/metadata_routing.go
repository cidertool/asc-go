package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
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
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// RoutingAppCoverageCreateRequest defines model for RoutingAppCoverageCreateRequest.
type RoutingAppCoverageCreateRequest struct {
	Data struct {
		Attributes struct {
			FileName string `json:"fileName"`
			FileSize int    `json:"fileSize"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreVersion struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// RoutingAppCoverageResponse defines model for RoutingAppCoverageResponse.
type RoutingAppCoverageResponse struct {
	Data  RoutingAppCoverage   `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// RoutingAppCoverageUpdateRequest defines model for RoutingAppCoverageUpdateRequest.
type RoutingAppCoverageUpdateRequest struct {
	Data struct {
		Attributes *struct {
			SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
			Uploaded           *bool   `json:"uploaded,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
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

type GetRoutingAppCoverageQuery struct {
	Fields *struct {
		RoutingAppCoverages *[]string `url:"routingAppCoverages,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// GetRoutingAppCoverage gets information about the routing app coverage file and its upload and processing status.
func (s *Service) GetRoutingAppCoverage(id string, params *GetRoutingAppCoverageQuery) (*RoutingAppCoverageResponse, *common.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateRoutingAppCoverage attaches a routing app coverage file to an App Store version.
func (s *Service) CreateRoutingAppCoverage(body *RoutingAppCoverageCreateRequest) (*RoutingAppCoverageResponse, *common.Response, error) {
	res := new(RoutingAppCoverageResponse)
	resp, err := s.Post("routingAppCoverages", body, res)
	return res, resp, err
}

// UpdateRoutingAppCoverage commits a routing app coverage file after uploading it.
func (s *Service) UpdateRoutingAppCoverage(id string, body *RoutingAppCoverageUpdateRequest) (*RoutingAppCoverageResponse, *common.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteRoutingAppCoverage deletes the routing app coverage file that is associated with a version.
func (s *Service) DeleteRoutingAppCoverage(id string) (*common.Response, error) {
	url := fmt.Sprintf("routingAppCoverages/%s", id)
	return s.Delete(url, nil)
}
