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

// AppScreenshot defines model for AppScreenshot.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshot
type AppScreenshot struct {
	Attributes    *AppScreenshotAttributes    `json:"attributes,omitempty"`
	ID            string                      `json:"id"`
	Links         ResourceLinks               `json:"links"`
	Relationships *AppScreenshotRelationships `json:"relationships,omitempty"`
	Type          string                      `json:"type"`
}

// AppScreenshotAttributes defines model for AppScreenshot.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshot/attributes
type AppScreenshotAttributes struct {
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	AssetToken         *string             `json:"assetToken,omitempty"`
	AssetType          *string             `json:"assetType,omitempty"`
	FileName           *string             `json:"fileName,omitempty"`
	FileSize           *int64              `json:"fileSize,omitempty"`
	ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
	SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
	UploadOperations   []UploadOperation   `json:"uploadOperations,omitempty"`
}

// AppScreenshotRelationships defines model for AppScreenshot.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshot/relationships
type AppScreenshotRelationships struct {
	AppScreenshotSet *Relationship `json:"appScreenshotSet,omitempty"`
}

// AppScreenshotCreateRequest defines model for AppScreenshotCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotcreaterequest/data
type appScreenshotCreateRequest struct {
	Attributes    appScreenshotCreateRequestAttributes    `json:"attributes"`
	Relationships appScreenshotCreateRequestRelationships `json:"relationships"`
	Type          string                                  `json:"type"`
}

// AppScreenshotCreateRequestAttributes are attributes for AppScreenshotCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotcreaterequest/data/attributes
type appScreenshotCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// AppScreenshotCreateRequestRelationships are relationships for AppScreenshotCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotcreaterequest/data/relationships
type appScreenshotCreateRequestRelationships struct {
	AppScreenshotSet relationshipDeclaration `json:"appScreenshotSet"`
}

// AppScreenshotUpdateRequest defines model for AppScreenshotUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotupdaterequest/data
type appScreenshotUpdateRequest struct {
	Attributes *appScreenshotUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                `json:"id"`
	Type       string                                `json:"type"`
}

// AppScreenshotUpdateRequestAttributes are attributes for AppScreenshotUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotupdaterequest/data/attributes
type appScreenshotUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppScreenshotResponse defines model for AppScreenshotResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotresponse
type AppScreenshotResponse struct {
	Data  AppScreenshot `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppScreenshotsResponse defines model for AppScreenshotsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appscreenshotsresponse
type AppScreenshotsResponse struct {
	Data  []AppScreenshot    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GetAppScreenshotQuery are query options for GetAppScreenshot
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_information
type GetAppScreenshotQuery struct {
	FieldsAppScreenshots []string `url:"fields[appScreenshots],omitempty"`
	Include              []string `url:"include,omitempty"`
}

// GetAppScreenshot gets information about an app screenshot and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_information
func (s *AppsService) GetAppScreenshot(ctx context.Context, id string, params *GetAppScreenshotQuery) (*AppScreenshotResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateAppScreenshot adds a new screenshot to a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_screenshot
func (s *AppsService) CreateAppScreenshot(ctx context.Context, fileName string, fileSize int64, appScreenshotSetID string) (*AppScreenshotResponse, *Response, error) {
	req := appScreenshotCreateRequest{
		Attributes: appScreenshotCreateRequestAttributes{
			FileName: fileName,
			FileSize: fileSize,
		},
		Relationships: appScreenshotCreateRequestRelationships{
			AppScreenshotSet: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appScreenshotSetID,
					Type: "appScreenshotSets",
				},
			},
		},
		Type: "appScreenshots",
	}
	res := new(AppScreenshotResponse)
	resp, err := s.client.post(ctx, "appScreenshots", newRequestBody(req), res)

	return res, resp, err
}

// CommitAppScreenshot commits an app screenshot after uploading it.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_screenshot
func (s *AppsService) CommitAppScreenshot(ctx context.Context, id string, uploaded *bool, sourceFileChecksum *string) (*AppScreenshotResponse, *Response, error) {
	req := appScreenshotUpdateRequest{
		ID:   id,
		Type: "appScreenshots",
	}

	if uploaded != nil || sourceFileChecksum != nil {
		req.Attributes = &appScreenshotUpdateRequestAttributes{
			Uploaded:           uploaded,
			SourceFileChecksum: sourceFileChecksum,
		}
	}

	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteAppScreenshot deletes an app screenshot that is associated with a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_screenshot
func (s *AppsService) DeleteAppScreenshot(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)

	return s.client.delete(ctx, url, nil)
}
