package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppInfo(ctx, "10", &GetAppInfoQuery{})
	})
}

func TestGetAppInfoIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"appInfoLocalizations"},{"type":"appCategories"}]}`, func(ctx context.Context, client *Client) {
		infos, _, err := client.Apps.GetAppInfo(ctx, "10", &GetAppInfoQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, infos.Included)

		assert.NotNil(t, infos.Included[0].AppInfoLocalization())
		assert.NotNil(t, infos.Included[1].AppCategory())

		assert.Nil(t, infos.Included[0].AppCategory())
		assert.Nil(t, infos.Included[1].AppInfoLocalization())
	})
}

func TestListAppInfosForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfosResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppInfosForApp(ctx, "10", &ListAppInfosForAppQuery{})
	})
}

func TestUpdateAppInfo(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppInfo(ctx, "10", &AppInfoUpdateRequestRelationships{})
	})
}
