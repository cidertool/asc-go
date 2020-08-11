package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInvitations(t *testing.T) {
	email := Email("me@email.com")
	want := &UserInvitationsResponse{
		Data: []UserInvitation{
			{
				Attributes: &UserInvitationAttributes{
					Email: &email,
				},
			},
		},
	}
	testEndpointWithResponse(t, `{"data":[{"attributes":{"email":"me@email.com"}}]}`, want, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListInvitations(ctx, &ListInvitationsQuery{})
	})
}

func TestGetInvitation(t *testing.T) {
	email := Email("me@email.com")
	want := &UserInvitationResponse{
		Data: UserInvitation{
			Attributes: &UserInvitationAttributes{
				Email: &email,
			},
		},
	}
	testEndpointWithResponse(t, `{"data":{"attributes":{"email":"me@email.com"}}}`, want, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.GetInvitation(ctx, "10", &GetInvitationQuery{})
	})
}

func TestCreateInvitation(t *testing.T) {
	email := Email("me@email.com")
	want := &UserInvitationResponse{
		Data: UserInvitation{
			Attributes: &UserInvitationAttributes{
				Email: &email,
			},
		},
	}
	testEndpointWithResponse(t, `{"data":{"attributes":{"email":"me@email.com"}}}`, want, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.CreateInvitation(ctx, &UserInvitationCreateRequest{
			Attributes: UserInvitationCreateRequestAttributes{
				Email: email,
			},
		})
	})
}

func TestCreateInvitationApplyRequestTypes(t *testing.T) {
	var req *UserInvitationCreateRequest
	req.applyTypes()
	assert.Nil(t, req)

	req = &UserInvitationCreateRequest{}
	req.applyTypes()
	assert.Equal(t, "userInvitations", req.Type)
	assert.Nil(t, req.Relationships)

	req = &UserInvitationCreateRequest{
		Relationships: &UserInvitationCreateRequestRelationships{},
	}
	req.applyTypes()
	assert.Nil(t, req.Relationships.VisibleApps)

	req = &UserInvitationCreateRequest{
		Relationships: &UserInvitationCreateRequestRelationships{
			VisibleApps: &PagedRelationshipDeclaration{
				&[]RelationshipData{
					{},
				},
			},
		},
	}
	req.applyTypes()
	assert.Equal(t, "apps", (*req.Relationships.VisibleApps.Data)[0].Type)

	req = &UserInvitationCreateRequest{
		Type: "dog",
		Relationships: &UserInvitationCreateRequestRelationships{
			VisibleApps: &PagedRelationshipDeclaration{
				&[]RelationshipData{
					{Type: "dog"},
				},
			},
		},
	}
	req.applyTypes()
	assert.Equal(t, "dog", req.Type)
	assert.Equal(t, "dog", (*req.Relationships.VisibleApps.Data)[0].Type)
}

func TestCancelInvitation(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.CancelInvitation(ctx, "10")
	})
}

func TestListVisibleAppsForInvitation(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsForInvitation(ctx, "10", &ListVisibleAppsQuery{})
	})
}
