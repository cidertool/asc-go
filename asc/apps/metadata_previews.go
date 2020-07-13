package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// PreviewType defines model for PreviewType.
type PreviewType string

// List of PreviewType
const (
	AppleTV        PreviewType = "APPLE_TV"
	Desktop        PreviewType = "DESKTOP"
	IPad105        PreviewType = "IPAD_105"
	IPad97         PreviewType = "IPAD_97"
	IPadPro129     PreviewType = "IPAD_PRO_129"
	IPadPro3Gen11  PreviewType = "IPAD_PRO_3GEN_11"
	IPadPro3Gen129 PreviewType = "IPAD_PRO_3GEN_129"
	IPhone35       PreviewType = "IPHONE_35"
	IPhone40       PreviewType = "IPHONE_40"
	IPhone47       PreviewType = "IPHONE_47"
	IPhone55       PreviewType = "IPHONE_55"
	IPhone58       PreviewType = "IPHONE_58"
	IPhone65       PreviewType = "IPHONE_65"
	WatchSeries3   PreviewType = "WATCH_SERIES_3"
	WatchSeries4   PreviewType = "WATCH_SERIES_4"
)

// AppPreview defines model for AppPreview.
type AppPreview struct {
	Attributes *struct {
		AssetDeliveryState   *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		FileName             *string             `json:"fileName,omitempty"`
		FileSize             *int                `json:"fileSize,omitempty"`
		MimeType             *string             `json:"mimeType,omitempty"`
		PreviewFrameTimeCode *string             `json:"previewFrameTimeCode,omitempty"`
		PreviewImage         *ImageAsset         `json:"previewImage,omitempty"`
		SourceFileChecksum   *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations     *[]UploadOperation  `json:"uploadOperations,omitempty"`
		VideoURL             *string             `json:"videoUrl,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppPreviewSet *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appPreviewSet,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppPreviewCreateRequest defines model for AppPreviewCreateRequest.
type AppPreviewCreateRequest struct {
	Data struct {
		Attributes struct {
			FileName             string  `json:"fileName"`
			FileSize             int     `json:"fileSize"`
			MimeType             *string `json:"mimeType,omitempty"`
			PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			AppPreviewSet struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appPreviewSet"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreviewUpdateRequest defines model for AppPreviewUpdateRequest.
type AppPreviewUpdateRequest struct {
	Data struct {
		Attributes *struct {
			PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
			SourceFileChecksum   *string `json:"sourceFileChecksum,omitempty"`
			Uploaded             *bool   `json:"uploaded,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppPreviewResponse defines model for AppPreviewResponse.
type AppPreviewResponse struct {
	Data  AppPreview           `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// AppPreviewsResponse defines model for AppPreviewsResponse.
type AppPreviewsResponse struct {
	Data  []AppPreview              `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// UploadOperation defines model for UploadOperation.
type UploadOperation struct {
	Length         *int                     `json:"length,omitempty"`
	Method         *string                  `json:"method,omitempty"`
	Offset         *int                     `json:"offset,omitempty"`
	RequestHeaders *[]UploadOperationHeader `json:"requestHeaders,omitempty"`
	URL            *string                  `json:"url,omitempty"`
}

// UploadOperationHeader defines model for UploadOperationHeader.
type UploadOperationHeader struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ImageAsset defines model for ImageAsset.
type ImageAsset struct {
	Height      *int    `json:"height,omitempty"`
	TemplateURL *string `json:"templateUrl,omitempty"`
	Width       *int    `json:"width,omitempty"`
}

type GetAppPreviewQuery struct {
	Fields *struct {
		AppPreviews *[]string `url:"appPreviews,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// GetAppPreview gets information about an app preview and its upload and processing status.
func (s *Service) GetAppPreview(id string, params *GetAppPreviewQuery) (*AppPreviewResponse, *common.Response, error) {
	url := fmt.Sprintf("appPreviews/%s", id)
	res := new(AppPreviewResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppPreview adds a new preview to a preview set.
func (s *Service) CreateAppPreview(body *AppPreviewCreateRequest) (*AppPreviewResponse, *common.Response, error) {
	res := new(AppPreviewResponse)
	resp, err := s.Post("appPreviews", body, res)
	return res, resp, err
}

// UpdateAppPreview commits an app preview after uploading it.
func (s *Service) UpdateAppPreview(id string, body *AppPreviewUpdateRequest) (*AppPreviewResponse, *common.Response, error) {
	url := fmt.Sprintf("appPreviews/%s", id)
	res := new(AppPreviewResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppPreview deletes an app preview that is associated with a preview set.
func (s *Service) DeleteAppPreview(id string) (*common.Response, error) {
	url := fmt.Sprintf("appPreviews/%s", id)
	return s.Delete(url, nil)
}
