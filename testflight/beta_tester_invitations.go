package testflight

import "github.com/aaronsky/asc-go/internal"

// BetaTesterInvitation defines model for BetaTesterInvitation.
type BetaTesterInvitation struct {
	ID    string                 `json:"id"`
	Links internal.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// BetaTesterInvitationCreateRequest defines model for BetaTesterInvitationCreateRequest.
type BetaTesterInvitationCreateRequest struct {
	Data struct {
		Relationships struct {
			App struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"app"`
			BetaTester struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"betaTester"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterInvitationResponse defines model for BetaTesterInvitationResponse.
type BetaTesterInvitationResponse struct {
	Data  BetaTesterInvitation   `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// CreateBetaTesterInvitation sends or resends an invitation to a beta tester to test a specified app.
func (s *Service) CreateBetaTesterInvitation(body *BetaTesterCreateRequest) (*BetaTesterInvitationResponse, *internal.Response, error) {
	res := new(BetaTesterInvitationResponse)
	resp, err := s.Post("betaTesterInvitations", body, res)
	return res, resp, err
}
