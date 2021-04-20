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

	"github.com/stretchr/testify/assert"
)

func TestListAppCategories(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppCategories(ctx, &ListAppCategoriesQuery{})
	})
}

func TestListSubcategoriesForAppCategory(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListSubcategoriesForAppCategory(ctx, "10", &ListSubcategoriesForAppCategoryQuery{})
	})
}

func TestGetAppCategory(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppCategory(ctx, "10", &GetAppCategoryQuery{})
	})
}

func TestGetAppCategoryIncludeds(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{"included":[{"type":"appCategories"},{"type":"appInfos"}]}`, func(ctx context.Context, client *Client) {
		category, _, err := client.Apps.GetAppCategory(ctx, "10", &GetAppCategoryQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, category.Included)

		assert.NotNil(t, category.Included[0].AppCategory())
		assert.Nil(t, category.Included[1].AppCategory())
	})
}

func TestGetParentCategoryForAppCategory(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetParentCategoryForAppCategory(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimaryCategoryForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimaryCategoryForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondaryCategoryForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondaryCategoryForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimarySubcategoryOneForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimarySubcategoryOneForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimarySubcategoryTwoForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimarySubcategoryTwoForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondarySubcategoryOneForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondarySubcategoryOneForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondarySubcategoryTwoForAppInfo(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondarySubcategoryTwoForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}
