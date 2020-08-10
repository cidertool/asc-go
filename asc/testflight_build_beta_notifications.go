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

// BuildBetaNotificationCreateRequest defines model for BuildBetaNotificationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationcreaterequest
type BuildBetaNotificationCreateRequest struct {
	Relationships BuildBetaNotificationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// BuildBetaNotificationCreateRequestRelationships are relationships for BuildBetaNotificationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationcreaterequest/data/relationships
type BuildBetaNotificationCreateRequestRelationships struct {
	Build RelationshipDeclaration `json:"build"`
}

// BuildBetaNotificationResponse defines model for BuildBetaNotificationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildbetanotificationresponse
type BuildBetaNotificationResponse struct {
	Data  BuildBetaNotification `json:"data"`
	Links DocumentLinks         `json:"links"`
}

func (r *BuildBetaNotificationCreateRequest) applyTypes() {
	if r == nil {
		return
	}
	r.Type = "buildBetaNotifications"
	r.Relationships.Build.applyType("builds")
}

// CreateAvailableBuildNotification sends a notification to all assigned beta testers that a build is available for testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/send_notification_of_an_available_build
func (s *TestflightService) CreateAvailableBuildNotification(ctx context.Context, body *BuildBetaNotificationCreateRequest) (*BuildBetaNotificationResponse, *Response, error) {
	res := new(BuildBetaNotificationResponse)
	resp, err := s.client.post(ctx, "buildBetaNotifications", body, res)
	return res, resp, err
}
