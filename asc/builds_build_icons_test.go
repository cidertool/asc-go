package asc

import (
	"context"
	"testing"
)

func TestListIconsForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildIconsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Builds.ListIconsForBuild(ctx, "10", &ListIconsQuery{})
	})
}
