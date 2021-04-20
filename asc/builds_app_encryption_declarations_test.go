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

func TestListAppEncryptionDeclarations(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListAppEncryptionDeclarations(ctx, &ListAppEncryptionDeclarationsQuery{})
	})
}

func TestGetAppEncryptionDeclaration(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppEncryptionDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppEncryptionDeclaration(ctx, "10", &GetAppEncryptionDeclarationQuery{})
	})
}

func TestGetAppForAppEncryptionDeclaration(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.GetAppForAppEncryptionDeclaration(ctx, "10", &GetAppForEncryptionDeclarationQuery{})
	})
}

func TestAssignBuildsToAppEncryptionDeclaration(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Builds.AssignBuildsToAppEncryptionDeclaration(ctx, "10", []string{"10"})
	})
}
