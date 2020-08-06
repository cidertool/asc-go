package asc

// BuildBetaNotification defines model for BuildBetaNotification.
type BuildBetaNotification struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// BuildBetaNotificationCreateRequest defines model for BuildBetaNotificationCreateRequest.
type BuildBetaNotificationCreateRequest struct {
	Relationships BuildBetaNotificationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// BuildBetaNotificationCreateRequestRelationships are relationships for BuildBetaNotificationCreateRequest
type BuildBetaNotificationCreateRequestRelationships struct {
	Build struct {
		Data RelationshipsData `json:"data"`
	} `json:"build"`
}

// BuildBetaNotificationResponse defines model for BuildBetaNotificationResponse.
type BuildBetaNotificationResponse struct {
	Data  BuildBetaNotification `json:"data"`
	Links DocumentLinks         `json:"links"`
}

// CreateAvailableBuildNotification sends a notification to all assigned beta testers that a build is available for testing.
func (s *TestflightService) CreateAvailableBuildNotification(body *BuildBetaNotificationCreateRequest) (*BuildBetaNotificationResponse, *Response, error) {
	res := new(BuildBetaNotificationResponse)
	resp, err := s.client.post("buildBetaNotifications", body, res)
	return res, resp, err
}
