package testflight

import (
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// BetaTesterInvitation defines model for BetaTesterInvitation.
type BetaTesterInvitation struct {
	ID    string              `json:"id"`
	Links types.ResourceLinks `json:"links"`
	Type  string              `json:"type"`
}

// BetaTesterInvitationCreateRequest defines model for BetaTesterInvitationCreateRequest.
type BetaTesterInvitationCreateRequest struct {
	Relationships BetaTesterInvitationCreateRequestRelationships `json:"relationships"`
	Type          string                                         `json:"type"`
}

// BetaTesterInvitationCreateRequestRelationships are relationships for BetaTesterInvitationCreateRequest
type BetaTesterInvitationCreateRequestRelationships struct {
	App struct {
		Data types.RelationshipsData `json:"data"`
	} `json:"app"`
	BetaTester struct {
		Data types.RelationshipsData `json:"data"`
	} `json:"betaTester"`
}

// BetaTesterInvitationResponse defines model for BetaTesterInvitationResponse.
type BetaTesterInvitationResponse struct {
	Data  BetaTesterInvitation `json:"data"`
	Links types.DocumentLinks  `json:"links"`
}

// CreateBetaTesterInvitation sends or resends an invitation to a beta tester to test a specified app.
func (s *Service) CreateBetaTesterInvitation(body *BetaTesterCreateRequest) (*BetaTesterInvitationResponse, *services.Response, error) {
	res := new(BetaTesterInvitationResponse)
	resp, err := s.Post("betaTesterInvitations", body, res)
	return res, resp, err
}
