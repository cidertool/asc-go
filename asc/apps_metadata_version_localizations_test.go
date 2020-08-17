package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListLocalizationsForAppStoreVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListLocalizationsForAppStoreVersion(ctx, "10", &ListLocalizationsForAppStoreVersionQuery{})
	})
}

func TestGetAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetAppStoreVersionLocalization(ctx, "10", &GetAppStoreVersionLocalizationQuery{})
	})
}

func TestGetAppStoreVersionLocalizationIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"appScreenshotSets"},{"type":"appPreviewSets"}]}`, func(ctx context.Context, client *Client) {
		localization, _, err := client.Apps.GetAppStoreVersionLocalization(ctx, "10", &GetAppStoreVersionLocalizationQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, localization.Included)

		assert.NotNil(t, localization.Included[0].AppScreenshotSet())
		assert.NotNil(t, localization.Included[1].AppPreviewSet())

		assert.Nil(t, localization.Included[0].AppPreviewSet())
		assert.Nil(t, localization.Included[1].AppScreenshotSet())
	})
}

func TestCreateAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.CreateAppStoreVersionLocalization(ctx, AppStoreVersionLocalizationCreateRequestAttributes{}, "")
	})
}

func TestUpdateAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreVersionLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAppStoreVersionLocalization(ctx, "10", &AppStoreVersionLocalizationUpdateRequestAttributes{})
	})
}

func TestDeleteAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.DeleteAppStoreVersionLocalization(ctx, "10")
	})
}

func TestListAppScreenshotSetsForAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppScreenshotSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppScreenshotSetsForAppStoreVersionLocalization(ctx, "10", &ListAppScreenshotSetsForAppStoreVersionLocalizationQuery{})
	})
}

func TestListAppPreviewSetsForAppStoreVersionLocalization(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppPreviewSetsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListAppPreviewSetsForAppStoreVersionLocalization(ctx, "10", &ListAppPreviewSetsForAppStoreVersionLocalizationQuery{})
	})
}
