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

func TestListBetaAppReviewDetails(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaAppReviewDetails(ctx, &ListBetaAppReviewDetailsQuery{})
	})
}

func TestGetBetaAppReviewDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewDetail(ctx, "10", &GetBetaAppReviewDetailQuery{})
	})
}

func TestGetAppForBetaAppReviewDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaAppReviewDetail(ctx, "10", &GetAppForBetaAppReviewDetailQuery{})
	})
}

func TestGetBetaAppReviewDetailsForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaAppReviewDetailsForApp(ctx, "10", &GetBetaAppReviewDetailsForAppQuery{})
	})
}

func TestUpdateBetaAppReviewDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BetaAppReviewDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaAppReviewDetail(ctx, "10", &BetaAppReviewDetailUpdateRequestAttributes{})
	})
}
