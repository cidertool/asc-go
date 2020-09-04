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
