package testflight

import "github.com/aaronsky/asc-go/v1/asc/common"

// BetaTesterInvitation defines model for BetaTesterInvitation.
type BetaTesterInvitation struct {
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// BetaTesterInvitationCreateRequest defines model for BetaTesterInvitationCreateRequest.
type BetaTesterInvitationCreateRequest struct {
	Data struct {
		Relationships struct {
			App struct {
				Data common.RelationshipsData `json:"data"`
			} `json:"app"`
			BetaTester struct {
				Data common.RelationshipsData `json:"data"`
			} `json:"betaTester"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterInvitationResponse defines model for BetaTesterInvitationResponse.
type BetaTesterInvitationResponse struct {
	Data  BetaTesterInvitation `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// CreateBetaTesterInvitation sends or resends an invitation to a beta tester to test a specified app.
func (s *Service) CreateBetaTesterInvitation(body *BetaTesterCreateRequest) (*BetaTesterInvitationResponse, *common.Response, error) {
	res := new(BetaTesterInvitationResponse)
	resp, err := s.Post("betaTesterInvitations", body, res)
	return res, resp, err
}
