package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &UsersResponse{}
	got, resp, err := client.Users.ListUsers(context.Background(), &ListUsersQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetUser(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &UserResponse{}
	got, resp, err := client.Users.GetUser(context.Background(), "10", &GetUserQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdateUser(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &UserResponse{}
	got, resp, err := client.Users.UpdateUser(context.Background(), "10", &UserUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestRemoveUser(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Users.RemoveUser(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListVisibleAppsForUser(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppsResponse{}
	got, resp, err := client.Users.ListVisibleAppsForUser(context.Background(), "10", &ListVisibleAppsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListVisibleAppsByResourceIDForUser(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &UserVisibleAppsLinkagesResponse{}
	got, resp, err := client.Users.ListVisibleAppsByResourceIDForUser(context.Background(), "10", &ListVisibleAppsByResourceIDQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestAddVisibleAppsForUser(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Users.AddVisibleAppsForUser(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUpdateVisibleAppsForUser(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Users.UpdateVisibleAppsForUser(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRemoveVisibleAppsFromUser(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Users.RemoveVisibleAppsFromUser(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
