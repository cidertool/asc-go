package submission

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppStoreReviewAttachment defines model for AppStoreReviewAttachment.
type AppStoreReviewAttachment struct {
	Attributes *struct {
		AssetDeliveryState *apps.AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		FileName           *string                  `json:"fileName,omitempty"`
		FileSize           *int                     `json:"fileSize,omitempty"`
		SourceFileChecksum *string                  `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *[]apps.UploadOperation  `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreReviewDetail *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreReviewDetail,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreReviewAttachmentCreateRequest defines model for AppStoreReviewAttachmentCreateRequest.
type AppStoreReviewAttachmentCreateRequest struct {
	Data struct {
		Attributes struct {
			FileName string `json:"fileName"`
			FileSize int    `json:"fileSize"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreReviewDetail struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appStoreReviewDetail"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreReviewAttachmentResponse defines model for AppStoreReviewAttachmentResponse.
type AppStoreReviewAttachmentResponse struct {
	Data  AppStoreReviewAttachment `json:"data"`
	Links common.DocumentLinks     `json:"links"`
}

// AppStoreReviewAttachmentUpdateRequest defines model for AppStoreReviewAttachmentUpdateRequest.
type AppStoreReviewAttachmentUpdateRequest struct {
	Data struct {
		Attributes *struct {
			SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
			Uploaded           *bool   `json:"uploaded,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreReviewAttachmentsResponse defines model for AppStoreReviewAttachmentsResponse.
type AppStoreReviewAttachmentsResponse struct {
	Data  []AppStoreReviewAttachment `json:"data"`
	Links common.PagedDocumentLinks  `json:"links"`
	Meta  *common.PagingInformation  `json:"meta,omitempty"`
}

// ReadAttachmentOptions are query options forGetAttachment
type ReadAttachmentOptions struct {
	Fields *struct {
		AppStoreReviewAttachments *[]string `url:"appStoreReviewAttachments,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

// ListAttachmentOptions are query options for ListAttachmentsForReviewDetail
type ListAttachmentOptions struct {
	Fields *struct {
		AppStoreReviewAttachments *[]string `url:"appStoreReviewAttachments,omitempty"`
		AppStoreReviewDetails     *[]string `url:"appStoreReviewDetails,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
}

// GetAttachment gets information about an App Store review attachment and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_attachment_information
func (s *Service) GetAttachment(id string, params *ReadAttachmentOptions) (*AppStoreReviewAttachmentResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListAttachmentsForReviewDetail lists all the App Store review attachments you include with a version when you submit it for App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_review_attachments_for_an_app_store_review_detail
func (s *Service) ListAttachmentsForReviewDetail(id string, params *ListAttachmentOptions) (*AppStoreReviewAttachmentsResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s/appStoreReviewAttachments", id)
	res := new(AppStoreReviewAttachmentsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// CreateAttachment attaches a document for App Review to an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_review_attachment
func (s *Service) CreateAttachment(body *AppStoreReviewAttachmentCreateRequest) (*AppStoreReviewAttachmentResponse, *common.Response, error) {
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.Post("appStoreReviewAttachments", body, res)
	return res, resp, err
}

// CommitAttachment commits an app screenshot after uploading it to the App Store.
//
// https://developer.apple.com/documentation/appstoreconnectapi/commit_an_app_store_review_attachment
func (s *Service) CommitAttachment(id string, body *AppStoreReviewAttachmentUpdateRequest) (*AppStoreReviewAttachmentResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAttachment removes an attachment before you send your app to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_review_attachment
func (s *Service) DeleteAttachment(id string) (*common.Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	return s.Delete(url, nil)
}
