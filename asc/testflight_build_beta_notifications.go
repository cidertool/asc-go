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

import "context"

// BuildBetaNotification defines model for BuildBetaNotification.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotification
type BuildBetaNotification struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// buildBetaNotificationCreateRequest defines model for buildBetaNotificationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationcreaterequest/data
type buildBetaNotificationCreateRequest struct {
	Relationships buildBetaNotificationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// buildBetaNotificationCreateRequestRelationships are relationships for BuildBetaNotificationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationcreaterequest/data/relationships
type buildBetaNotificationCreateRequestRelationships struct {
	Build relationshipDeclaration `json:"build"`
}

// BuildBetaNotificationResponse defines model for BuildBetaNotificationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationresponse
type BuildBetaNotificationResponse struct {
	Data  BuildBetaNotification `json:"data"`
	Links DocumentLinks         `json:"links"`
}

// CreateAvailableBuildNotification sends a notification to all assigned beta testers that a build is available for testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/send_notification_of_an_available_build
func (s *TestflightService) CreateAvailableBuildNotification(ctx context.Context, buildID string) (*BuildBetaNotificationResponse, *Response, error) {
	req := buildBetaNotificationCreateRequest{
		Relationships: buildBetaNotificationCreateRequestRelationships{
			Build: relationshipDeclaration{
				Data: RelationshipData{
					ID:   buildID,
					Type: "builds",
				},
			},
		},
		Type: "buildBetaNotifications",
	}
	res := new(BuildBetaNotificationResponse)
	resp, err := s.client.post(ctx, "buildBetaNotifications", newRequestBody(req), res)

	return res, resp, err
}
