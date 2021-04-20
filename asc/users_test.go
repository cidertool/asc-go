/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"testing"
)

func TestListUsers(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UsersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListUsers(ctx, &ListUsersQuery{})
	})
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UserResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.GetUser(ctx, "10", &GetUserQuery{})
	})
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UserResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.UpdateUser(ctx, "10", &UserUpdateRequestAttributes{}, []string{"10"})
	})
}

func TestRemoveUser(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.RemoveUser(ctx, "10")
	})
}

func TestListVisibleAppsForUser(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsForUser(ctx, "10", &ListVisibleAppsQuery{})
	})
}

func TestListVisibleAppsByResourceIDForUser(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UserVisibleAppsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsByResourceIDForUser(ctx, "10", &ListVisibleAppsByResourceIDQuery{})
	})
}

func TestAddVisibleAppsForUser(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.AddVisibleAppsForUser(ctx, "10", []string{"10"})
	})
}

func TestUpdateVisibleAppsForUser(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.UpdateVisibleAppsForUser(ctx, "10", []string{"10"})
	})
}

func TestRemoveVisibleAppsFromUser(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.RemoveVisibleAppsFromUser(ctx, "10", []string{"10"})
	})
}
