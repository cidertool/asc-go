package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppCategories(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppCategories(ctx, &ListAppCategoriesQuery{})
	})
}

func TestListSubcategoriesForAppCategory(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoriesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListSubcategoriesForAppCategory(ctx, "10", &ListSubcategoriesForAppCategoryQuery{})
	})
}

func TestGetAppCategory(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppCategory(ctx, "10", &GetAppCategoryQuery{})
	})
}

func TestGetAppCategoryIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"appCategories"},{"type":"appInfos"}]}`, func(ctx context.Context, client *Client) {
		category, _, err := client.Apps.GetAppCategory(ctx, "10", &GetAppCategoryQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, category.Included)

		assert.NotNil(t, category.Included[0].AppCategory())
		assert.Nil(t, category.Included[1].AppCategory())
	})
}

func TestGetParentCategoryForAppCategory(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetParentCategoryForAppCategory(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimaryCategoryForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimaryCategoryForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondaryCategoryForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondaryCategoryForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimarySubcategoryOneForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimarySubcategoryOneForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetPrimarySubcategoryTwoForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetPrimarySubcategoryTwoForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondarySubcategoryOneForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondarySubcategoryOneForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}

func TestGetSecondarySubcategoryTwoForAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppCategoryResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetSecondarySubcategoryTwoForAppInfo(ctx, "10", &GetAppCategoryForAppInfoQuery{})
	})
}
