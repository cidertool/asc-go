package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
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
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshotSet *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appScreenshotSet,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotCreateRequest defines model for AppScreenshotCreateRequest.
type AppScreenshotCreateRequest struct {
	Data struct {
		Attributes struct {
			FileName string `json:"fileName"`
			FileSize int    `json:"fileSize"`
		} `json:"attributes"`
		Relationships struct {
			AppScreenshotSet struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appScreenshotSet"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppScreenshotUpdateRequest defines model for AppScreenshotUpdateRequest.
type AppScreenshotUpdateRequest struct {
	Data struct {
		Attributes *struct {
			SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
			Uploaded           *bool   `json:"uploaded,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppScreenshotResponse defines model for AppScreenshotResponse.
type AppScreenshotResponse struct {
	Data  AppScreenshot        `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// AppScreenshotsResponse defines model for AppScreenshotsResponse.
type AppScreenshotsResponse struct {
	Data  []AppScreenshot           `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type GetAppScreenshotQuery struct {
	Fields *struct {
		AppScreenshots *[]string `url:"appScreenshots,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// GetAppScreenshot gets information about an app screenshot and its upload and processing status.
func (s *Service) GetAppScreenshot(id string, params *GetAppScreenshotQuery) (*AppScreenshotResponse, *common.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppScreenshot adds a new screenshot to a screenshot set.
func (s *Service) CreateAppScreenshot(body *AppScreenshotCreateRequest) (*AppScreenshotResponse, *common.Response, error) {
	res := new(AppScreenshotResponse)
	resp, err := s.Post("appScreenshots", body, res)
	return res, resp, err
}

// UpdateAppScreenshot commits an app screenshot after uploading it.
func (s *Service) UpdateAppScreenshot(id string, body *AppScreenshotUpdateRequest) (*AppScreenshotResponse, *common.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppScreenshot deletes an app screenshot that is associated with a screenshot set.
func (s *Service) DeleteAppScreenshot(id string) (*common.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	return s.Delete(url, nil)
}
