package asc

import "context"

// BetaTesterInvitation defines model for BetaTesterInvitation.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterinvitation
type BetaTesterInvitation struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// BetaTesterInvitationCreateRequest defines model for BetaTesterInvitationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterinvitationcreaterequest/data
type betaTesterInvitationCreateRequest struct {
	Relationships betaTesterInvitationCreateRequestRelationships `json:"relationships"`
	Type          string                                         `json:"type"`
}

// BetaTesterInvitationCreateRequestRelationships are relationships for BetaTesterInvitationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterinvitationcreaterequest/data/relationships
type betaTesterInvitationCreateRequestRelationships struct {
	App        relationshipDeclaration `json:"app"`
	BetaTester relationshipDeclaration `json:"betaTester"`
}

// BetaTesterInvitationResponse defines model for BetaTesterInvitationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterinvitationresponse
type BetaTesterInvitationResponse struct {
	Data  BetaTesterInvitation `json:"data"`
	Links DocumentLinks        `json:"links"`
}

// CreateBetaTesterInvitation sends or resends an invitation to a beta tester to test a specified app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/send_an_invitation_to_a_beta_tester
func (s *TestflightService) CreateBetaTesterInvitation(ctx context.Context, appID string, betaTesterID string) (*BetaTesterInvitationResponse, *Response, error) {
	req := betaTesterInvitationCreateRequest{
		Relationships: betaTesterInvitationCreateRequestRelationships{
			App: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appID,
					Type: "apps",
				},
			},
			BetaTester: relationshipDeclaration{
				Data: RelationshipData{
					ID:   betaTesterID,
					Type: "betaTesters",
				},
			},
		},
		Type: "betaTesterInvitations",
	}
	res := new(BetaTesterInvitationResponse)
	resp, err := s.client.post(ctx, "betaTesterInvitations", req, res)
	return res, resp, err
}
