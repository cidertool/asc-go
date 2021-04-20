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

func TestListInvitations(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Users.CancelInvitation(ctx, "10")
	})
}

func TestListVisibleAppsForInvitation(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Users.ListVisibleAppsForInvitation(ctx, "10", &ListVisibleAppsQuery{})
	})
}
