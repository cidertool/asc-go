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

package asc // nolint: dupl

import (
	"context"
	"testing"
)

func TestGetAppPreviewSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppPreviewSet(ctx, "10", &GetAppPreviewSetQuery{})
	})
}

func TestCreateAppPreviewSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppPreviewSet(ctx, PreviewTypeiPadPro129, "")
	})
}

func TestDeleteAppPreviewSet(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppPreviewSet(ctx, "10")
	})
}

func TestListAppPreviewsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewsForSet(ctx, "10", &ListAppPreviewsForSetQuery{})
	})
}

func TestListAppPreviewIDsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewSetAppPreviewsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewIDsForSet(ctx, "10", &ListAppPreviewIDsForSetQuery{})
	})
}

func TestReplaceAppPreviewsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.ReplaceAppPreviewsForSet(ctx, "10", []string{"10"})
	})
}
