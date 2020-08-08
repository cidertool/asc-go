package asc

import (
	"context"
	"fmt"
)

// AppStoreReviewAttachment defines model for AppStoreReviewAttachment.
type AppStoreReviewAttachment struct {
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
		AppStoreReviewDetail *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreReviewDetail,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreReviewAttachmentCreateRequest defines model for AppStoreReviewAttachmentCreateRequest.
type AppStoreReviewAttachmentCreateRequest struct {
	Attributes    AppStoreReviewAttachmentCreateRequestAttributes    `json:"attributes"`
	Relationships AppStoreReviewAttachmentCreateRequestRelationships `json:"relationships"`
	Type          string                                             `json:"type"`
}

// AppStoreReviewAttachmentCreateRequestAttributes are attributes for AppStoreReviewAttachmentCreateRequest
type AppStoreReviewAttachmentCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// AppStoreReviewAttachmentCreateRequestRelationships are relationships for AppStoreReviewAttachmentCreateRequest
type AppStoreReviewAttachmentCreateRequestRelationships struct {
	AppStoreReviewDetail struct {
		Data RelationshipsData `json:"data"`
	} `json:"appStoreReviewDetail"`
}

// AppStoreReviewAttachmentResponse defines model for AppStoreReviewAttachmentResponse.
type AppStoreReviewAttachmentResponse struct {
	Data  AppStoreReviewAttachment `json:"data"`
	Links DocumentLinks            `json:"links"`
}

// AppStoreReviewAttachmentUpdateRequest defines model for AppStoreReviewAttachmentUpdateRequest.
type AppStoreReviewAttachmentUpdateRequest struct {
	Attributes *AppStoreReviewAttachmentUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                           `json:"id"`
	Type       string                                           `json:"type"`
}

// AppStoreReviewAttachmentUpdateRequestAttributes are attributes for AppStoreReviewAttachmentUpdateRequest
type AppStoreReviewAttachmentUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppStoreReviewAttachmentsResponse defines model for AppStoreReviewAttachmentsResponse.
type AppStoreReviewAttachmentsResponse struct {
	Data  []AppStoreReviewAttachment `json:"data"`
	Links PagedDocumentLinks         `json:"links"`
	Meta  *PagingInformation         `json:"meta,omitempty"`
}

// ReadAttachmentQuery are query options forGetAttachment
type ReadAttachmentQuery struct {
	FieldsAppStoreReviewAttachments []string `url:"fields[appStoreReviewAttachments],omitempty"`
	Include                         []string `url:"include,omitempty"`
}

// ListAttachmentQuery are query options for ListAttachmentsForReviewDetail
type ListAttachmentQuery struct {
	FieldsAppStoreReviewAttachments []string `url:"fields[appStoreReviewAttachments],omitempty"`
	FieldsAppStoreReviewDetails     []string `url:"fields[appStoreReviewDetails],omitempty"`
	Include                         []string `url:"include,omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	Cursor                          string   `url:"cursor,omitempty"`
}

// GetAttachment gets information about an App Store review attachment and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_attachment_information
func (s *SubmissionService) GetAttachment(ctx context.Context, id string, params *ReadAttachmentQuery) (*AppStoreReviewAttachmentResponse, *Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListAttachmentsForReviewDetail lists all the App Store review attachments you include with a version when you submit it for App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_review_attachments_for_an_app_store_review_detail
func (s *SubmissionService) ListAttachmentsForReviewDetail(ctx context.Context, id string, params *ListAttachmentQuery) (*AppStoreReviewAttachmentsResponse, *Response, error) {
	url := fmt.Sprintf("appStoreReviewDetails/%s/appStoreReviewAttachments", id)
	res := new(AppStoreReviewAttachmentsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateAttachment attaches a document for App Review to an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_review_attachment
func (s *SubmissionService) CreateAttachment(ctx context.Context, body *AppStoreReviewAttachmentCreateRequest) (*AppStoreReviewAttachmentResponse, *Response, error) {
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.client.post(ctx, "appStoreReviewAttachments", body, res)
	return res, resp, err
}

// CommitAttachment commits an app screenshot after uploading it to the App Store.
//
// https://developer.apple.com/documentation/appstoreconnectapi/commit_an_app_store_review_attachment
func (s *SubmissionService) CommitAttachment(ctx context.Context, id string, body *AppStoreReviewAttachmentUpdateRequest) (*AppStoreReviewAttachmentResponse, *Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeleteAttachment removes an attachment before you send your app to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_review_attachment
func (s *SubmissionService) DeleteAttachment(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	return s.client.delete(ctx, url, nil)
}
