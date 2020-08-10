package asc

import (
	"context"
	"testing"
)

func TestListUsers(t *testing.T) {
	testEndpointWithResponse(t, "{}", &UsersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListUsers(ctx, &ListUsersQuery{})
	})
}

func TestGetUser(t *testing.T) {
	testEndpointWithResponse(t, "{}", &UserResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.GetUser(ctx, "10", &GetUserQuery{})
	})
}

func TestUpdateUser(t *testing.T) {
	testEndpointWithResponse(t, "{}", &UserResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.UpdateUser(ctx, "10", &UserUpdateRequest{})
	})
}

func TestRemoveUser(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.RemoveUser(ctx, "10")
	})
}

func TestListVisibleAppsForUser(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsForUser(ctx, "10", &ListVisibleAppsQuery{})
	})
}

func TestListVisibleAppsByResourceIDForUser(t *testing.T) {
	testEndpointWithResponse(t, "{}", &UserVisibleAppsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsByResourceIDForUser(ctx, "10", &ListVisibleAppsByResourceIDQuery{})
	})
}

func TestAddVisibleAppsForUser(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.AddVisibleAppsForUser(ctx, "10", &[]RelationshipData{})
	})
}

func TestUpdateVisibleAppsForUser(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.UpdateVisibleAppsForUser(ctx, "10", &[]RelationshipData{})
	})
}

func TestRemoveVisibleAppsFromUser(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.RemoveVisibleAppsFromUser(ctx, "10", &[]RelationshipData{})
	})
}
