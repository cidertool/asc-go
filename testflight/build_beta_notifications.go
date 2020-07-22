package testflight

import "github.com/aaronsky/asc-go/internal"

// BuildBetaNotification defines model for BuildBetaNotification.
type BuildBetaNotification struct {
	ID    string                 `json:"id"`
	Links internal.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// BuildBetaNotificationCreateRequest defines model for BuildBetaNotificationCreateRequest.
type BuildBetaNotificationCreateRequest struct {
	Data struct {
		Relationships struct {
			Build struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"build"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BuildBetaNotificationResponse defines model for BuildBetaNotificationResponse.
type BuildBetaNotificationResponse struct {
	Data  BuildBetaNotification  `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// CreateAvailableBuildNotification sends a notification to all assigned beta testers that a build is available for testing.
func (s *Service) CreateAvailableBuildNotification(body *BuildBetaNotificationCreateRequest) (*BuildBetaNotificationResponse, *internal.Response, error) {
	res := new(BuildBetaNotificationResponse)
	resp, err := s.Post("buildBetaNotifications", body, res)
	return res, resp, err
}
