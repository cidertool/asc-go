package asc

import (
	"context"
	"testing"
)

func TestCreateBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaGroup(ctx, &BetaGroupCreateRequest{})
	})
}

func TestUpdateBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.UpdateBetaGroup(ctx, "10", &BetaGroupUpdateRequest{})
	})
}

func TestDeleteBetaGroup(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaGroup(ctx, "10")
	})
}

func TestListBetaGroups(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroups(ctx, &ListBetaGroupsQuery{})
	})
}

func TestGetBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaGroup(ctx, "10", &GetBetaGroupQuery{})
	})
}

func TestGetAppForBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetAppForBetaGroup(ctx, "10", &GetAppForBetaGroupQuery{})
	})
}

func TestListBetaGroupsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroupsForApp(ctx, "10", &ListBetaGroupsForAppQuery{})
	})
}

func TestAddBetaTestersToBetaGroup(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AddBetaTestersToBetaGroup(ctx, "10", &[]RelationshipData{})
	})
}

func TestRemoveBetaTestersFromBetaGroup(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveBetaTestersFromBetaGroup(ctx, "10", &[]RelationshipData{})
	})
}

func TestAddBuildsToBetaGroup(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AddBuildsToBetaGroup(ctx, "10", &[]RelationshipData{})
	})
}

func TestRemoveBuildsFromBetaGroup(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveBuildsFromBetaGroup(ctx, "10", &[]RelationshipData{})
	})
}

func TestListBuildsForBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildsForBetaGroup(ctx, "10", &ListBuildsForBetaGroupQuery{})
	})
}

func TestListBuildIDsForBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupBuildsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildIDsForBetaGroup(ctx, "10", &ListBuildIDsForBetaGroupQuery{})
	})
}

func TestListBetaTestersForBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTestersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaTestersForBetaGroup(ctx, "10", &ListBetaTestersForBetaGroupQuery{})
	})
}

func TestListBetaTesterIDsForBetaGroup(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupBetaTestersLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaTesterIDsForBetaGroup(ctx, "10", &ListBetaTesterIDsForBetaGroupQuery{})
	})
}
