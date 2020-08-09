package asc

import (
	"context"
	"testing"
	"time"
)

func TestListInvitations(t *testing.T) {
	email := Email("me@email.com")
	want := &UserInvitationsResponse{
		Data: []UserInvitation{
			{
				Attributes: &struct {
					AllAppsVisible      *bool       "json:\"allAppsVisible,omitempty\""
					Email               *Email      "json:\"email,omitempty\""
					ExpirationDate      *time.Time  "json:\"expirationDate,omitempty\""
					FirstName           *string     "json:\"firstName,omitempty\""
					LastName            *string     "json:\"lastName,omitempty\""
					ProvisioningAllowed *bool       "json:\"provisioningAllowed,omitempty\""
					Roles               *[]UserRole "json:\"roles,omitempty\""
				}{
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
			Attributes: &struct {
				AllAppsVisible      *bool       "json:\"allAppsVisible,omitempty\""
				Email               *Email      "json:\"email,omitempty\""
				ExpirationDate      *time.Time  "json:\"expirationDate,omitempty\""
				FirstName           *string     "json:\"firstName,omitempty\""
				LastName            *string     "json:\"lastName,omitempty\""
				ProvisioningAllowed *bool       "json:\"provisioningAllowed,omitempty\""
				Roles               *[]UserRole "json:\"roles,omitempty\""
			}{
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
			Attributes: &struct {
				AllAppsVisible      *bool       "json:\"allAppsVisible,omitempty\""
				Email               *Email      "json:\"email,omitempty\""
				ExpirationDate      *time.Time  "json:\"expirationDate,omitempty\""
				FirstName           *string     "json:\"firstName,omitempty\""
				LastName            *string     "json:\"lastName,omitempty\""
				ProvisioningAllowed *bool       "json:\"provisioningAllowed,omitempty\""
				Roles               *[]UserRole "json:\"roles,omitempty\""
			}{
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
