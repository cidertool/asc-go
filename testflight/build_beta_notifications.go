package testflight

import (
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// BuildBetaNotification defines model for BuildBetaNotification.
type BuildBetaNotification struct {
	ID    string              `json:"id"`
	Links types.ResourceLinks `json:"links"`
	Type  string              `json:"type"`
}

// BuildBetaNotificationCreateRequest defines model for BuildBetaNotificationCreateRequest.
type BuildBetaNotificationCreateRequest struct {
	Relationships BuildBetaNotificationCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// BuildBetaNotificationCreateRequestRelationships are relationships for BuildBetaNotificationCreateRequest
type BuildBetaNotificationCreateRequestRelationships struct {
	Build struct {
		Data types.RelationshipsData `json:"data"`
	} `json:"build"`
}

// BuildBetaNotificationResponse defines model for BuildBetaNotificationResponse.
type BuildBetaNotificationResponse struct {
	Data  BuildBetaNotification `json:"data"`
	Links types.DocumentLinks   `json:"links"`
}

// CreateAvailableBuildNotification sends a notification to all assigned beta testers that a build is available for testing.
func (s *Service) CreateAvailableBuildNotification(body *BuildBetaNotificationCreateRequest) (*BuildBetaNotificationResponse, *services.Response, error) {
	res := new(BuildBetaNotificationResponse)
	resp, err := s.Post("buildBetaNotifications", body, res)
	return res, resp, err
}
