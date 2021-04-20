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

// AppStoreReviewAttachment defines model for AppStoreReviewAttachment.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachment
type AppStoreReviewAttachment struct {
	Attributes    *AppStoreReviewAttachmentAttributes    `json:"attributes,omitempty"`
	ID            string                                 `json:"id"`
	Links         ResourceLinks                          `json:"links"`
	Relationships *AppStoreReviewAttachmentRelationships `json:"relationships,omitempty"`
	Type          string                                 `json:"type"`
}

// AppStoreReviewAttachmentAttributes defines model for AppStoreReviewAttachment.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachment/attributes
type AppStoreReviewAttachmentAttributes struct {
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	FileName           *string             `json:"fileName,omitempty"`
	FileSize           *int64              `json:"fileSize,omitempty"`
	SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
	UploadOperations   []UploadOperation   `json:"uploadOperations,omitempty"`
}

// AppStoreReviewAttachmentRelationships defines model for AppStoreReviewAttachment.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachment/relationships
type AppStoreReviewAttachmentRelationships struct {
	AppStoreReviewDetail *Relationship `json:"appStoreReviewDetail,omitempty"`
}

// appStoreReviewAttachmentCreateRequest defines model for appStoreReviewAttachmentCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentcreaterequest/data
type appStoreReviewAttachmentCreateRequest struct {
	Attributes    appStoreReviewAttachmentCreateRequestAttributes    `json:"attributes"`
	Relationships appStoreReviewAttachmentCreateRequestRelationships `json:"relationships"`
	Type          string                                             `json:"type"`
}

// appStoreReviewAttachmentCreateRequestAttributes are attributes for AppStoreReviewAttachmentCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentcreaterequest/data/attributes
type appStoreReviewAttachmentCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// appStoreReviewAttachmentCreateRequestRelationships are relationships for AppStoreReviewAttachmentCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentcreaterequest/data/relationships
type appStoreReviewAttachmentCreateRequestRelationships struct {
	AppStoreReviewDetail relationshipDeclaration `json:"appStoreReviewDetail"`
}

// AppStoreReviewAttachmentResponse defines model for AppStoreReviewAttachmentResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentresponse
type AppStoreReviewAttachmentResponse struct {
	Data  AppStoreReviewAttachment `json:"data"`
	Links DocumentLinks            `json:"links"`
}

// AppStoreReviewAttachmentUpdateRequest defines model for AppStoreReviewAttachmentUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentupdaterequest/data
type appStoreReviewAttachmentUpdateRequest struct {
	Attributes *appStoreReviewAttachmentUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                           `json:"id"`
	Type       string                                           `json:"type"`
}

// AppStoreReviewAttachmentUpdateRequestAttributes are attributes for AppStoreReviewAttachmentUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentupdaterequest/data/attributes
type appStoreReviewAttachmentUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppStoreReviewAttachmentsResponse defines model for AppStoreReviewAttachmentsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewattachmentsresponse
type AppStoreReviewAttachmentsResponse struct {
	Data  []AppStoreReviewAttachment `json:"data"`
	Links PagedDocumentLinks         `json:"links"`
	Meta  *PagingInformation         `json:"meta,omitempty"`
}

// GetAttachmentQuery are query options for GetAttachment
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_attachment_information
type GetAttachmentQuery struct {
	FieldsAppStoreReviewAttachments []string `url:"fields[appStoreReviewAttachments],omitempty"`
	Include                         []string `url:"include,omitempty"`
}

// ListAttachmentQuery are query options for ListAttachmentsForReviewDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_review_attachments_for_an_app_store_review_detail
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
func (s *SubmissionService) GetAttachment(ctx context.Context, id string, params *GetAttachmentQuery) (*AppStoreReviewAttachmentResponse, *Response, error) {
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
func (s *SubmissionService) CreateAttachment(ctx context.Context, fileName string, fileSize int64, appStoreReviewDetailID string) (*AppStoreReviewAttachmentResponse, *Response, error) {
	req := appStoreReviewAttachmentCreateRequest{
		Attributes: appStoreReviewAttachmentCreateRequestAttributes{
			FileName: fileName,
			FileSize: fileSize,
		},
		Relationships: appStoreReviewAttachmentCreateRequestRelationships{
			AppStoreReviewDetail: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appStoreReviewDetailID,
					Type: "appStoreReviewDetails",
				},
			},
		},
		Type: "appStoreReviewAttachments",
	}
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.client.post(ctx, "appStoreReviewAttachments", newRequestBody(req), res)

	return res, resp, err
}

// CommitAttachment commits an app screenshot after uploading it to the App Store.
//
// https://developer.apple.com/documentation/appstoreconnectapi/commit_an_app_store_review_attachment
func (s *SubmissionService) CommitAttachment(ctx context.Context, id string, uploaded *bool, sourceFileChecksum *string) (*AppStoreReviewAttachmentResponse, *Response, error) {
	req := appStoreReviewAttachmentUpdateRequest{
		ID:   id,
		Type: "appStoreReviewAttachments",
	}

	if uploaded != nil || sourceFileChecksum != nil {
		req.Attributes = &appStoreReviewAttachmentUpdateRequestAttributes{
			Uploaded:           uploaded,
			SourceFileChecksum: sourceFileChecksum,
		}
	}

	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)
	res := new(AppStoreReviewAttachmentResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteAttachment removes an attachment before you send your app to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_review_attachment
func (s *SubmissionService) DeleteAttachment(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreReviewAttachments/%s", id)

	return s.client.delete(ctx, url, nil)
}
