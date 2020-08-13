package asc

import (
	"context"
	"testing"
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
		return client.Users.CreateInvitation(ctx, UserInvitationCreateRequestAttributes{Email: email}, []string{"10"})
	})
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
