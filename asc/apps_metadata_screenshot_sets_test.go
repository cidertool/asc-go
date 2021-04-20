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

func TestGetAppScreenshotSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppScreenshotSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppScreenshotSet(ctx, "10", &GetAppScreenshotSetQuery{})
	})
}

func TestCreateAppScreenshotSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppScreenshotSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppScreenshotSet(ctx, ScreenshotDisplayTypeAppiPadPro129, "")
	})
}

func TestDeleteAppScreenshotSet(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppScreenshotSet(ctx, "10")
	})
}

func TestListAppScreenshotsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppScreenshotsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotsForSet(ctx, "10", &ListAppScreenshotsForSetQuery{})
	})
}

func TestListAppScreenshotIDsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppScreenshotSetAppScreenshotsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotIDsForSet(ctx, "10", &ListAppScreenshotIDsForSetQuery{})
	})
}

func TestReplaceAppScreenshotsForSet(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.ReplaceAppScreenshotsForSet(ctx, "10", []string{"10"})
	})
}
