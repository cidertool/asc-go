package asc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListInvitations(t *testing.T) {
	marshaled := `{"data":[{"attributes":{"email":"me@email.com"}}]}`
	client, server := newServer(marshaled)
	defer server.Close()

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
	got, resp, err := client.Users.ListInvitations(context.Background(), &ListInvitationsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetInvitation(t *testing.T) {
	marshaled := `{"data":{"attributes":{"email":"me@email.com"}}}`
	client, server := newServer(marshaled)
	defer server.Close()

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
	got, resp, err := client.Users.GetInvitation(context.Background(), "10", &GetInvitationQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestCreateInvitation(t *testing.T) {
	marshaled := `{"data":{"attributes":{"email":"me@email.com"}}}`
	client, server := newServer(marshaled)
	defer server.Close()

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
	got, resp, err := client.Users.CreateInvitation(context.Background(), &UserInvitationCreateRequest{
		Attributes: UserInvitationCreateRequestAttributes{
			Email: email,
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestCancelInvitation(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Users.CancelInvitation(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListVisibleAppsForInvitation(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppsResponse{}
	got, resp, err := client.Users.ListVisibleAppsForInvitation(context.Background(), "10", &ListVisibleAppsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
