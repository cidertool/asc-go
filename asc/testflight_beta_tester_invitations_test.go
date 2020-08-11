package asc

import (
	"context"
	"testing"
)

func TestCreateBetaTesterInvitation(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTesterInvitationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaTesterInvitation(ctx, &BetaTesterInvitationCreateRequest{})
	})
}

func TestCreateBetaTesterInvitationApplyRequestTypes(t *testing.T) {

}
