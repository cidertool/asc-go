package apps

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/internal/services"
)

// AppScreenshot defines model for AppScreenshot.
type AppScreenshot struct {
	Attributes *struct {
		AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		AssetToken         *string             `json:"assetToken,omitempty"`
		AssetType          *string             `json:"assetType,omitempty"`
		FileName           *string             `json:"fileName,omitempty"`
		FileSize           *int                `json:"fileSize,omitempty"`
		ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
		SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *[]UploadOperation  `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshotSet *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appScreenshotSet,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotCreateRequest defines model for AppScreenshotCreateRequest.
type AppScreenshotCreateRequest struct {
	Attributes    AppScreenshotCreateRequestAttributes    `json:"attributes"`
	Relationships AppScreenshotCreateRequestRelationships `json:"relationships"`
	Type          string                                  `json:"type"`
}

// AppScreenshotCreateRequestAttributes are attributes for AppScreenshotCreateRequest
type AppScreenshotCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
}

// AppScreenshotCreateRequestRelationships are relationships for AppScreenshotCreateRequest
type AppScreenshotCreateRequestRelationships struct {
	AppScreenshotSet struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"appScreenshotSet"`
}

// AppScreenshotUpdateRequest defines model for AppScreenshotUpdateRequest.
type AppScreenshotUpdateRequest struct {
	Attributes *AppScreenshotUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                `json:"id"`
	Type       string                                `json:"type"`
}

// AppScreenshotUpdateRequestAttributes are attributes for AppScreenshotUpdateRequest
type AppScreenshotUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppScreenshotResponse defines model for AppScreenshotResponse.
type AppScreenshotResponse struct {
	Data  AppScreenshot          `json:"data"`
	Links services.DocumentLinks `json:"links"`
}

// AppScreenshotsResponse defines model for AppScreenshotsResponse.
type AppScreenshotsResponse struct {
	Data  []AppScreenshot             `json:"data"`
	Links services.PagedDocumentLinks `json:"links"`
	Meta  *services.PagingInformation `json:"meta,omitempty"`
}

// GetAppScreenshotQuery are query options for GetAppScreenshot
type GetAppScreenshotQuery struct {
	FieldsAppScreenshots *[]string `url:"fields[appScreenshots],omitempty"`
	Include              *[]string `url:"include,omitempty"`
}

// GetAppScreenshot gets information about an app screenshot and its upload and processing status.
func (s *Service) GetAppScreenshot(id string, params *GetAppScreenshotQuery) (*AppScreenshotResponse, *http.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppScreenshot adds a new screenshot to a screenshot set.
func (s *Service) CreateAppScreenshot(body *AppScreenshotCreateRequest) (*AppScreenshotResponse, *http.Response, error) {
	res := new(AppScreenshotResponse)
	resp, err := s.Post("appScreenshots", body, res)
	return res, resp, err
}

// UpdateAppScreenshot commits an app screenshot after uploading it.
func (s *Service) UpdateAppScreenshot(id string, body *AppScreenshotUpdateRequest) (*AppScreenshotResponse, *http.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppScreenshot deletes an app screenshot that is associated with a screenshot set.
func (s *Service) DeleteAppScreenshot(id string) (*http.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	return s.Delete(url, nil)
}
