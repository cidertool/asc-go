package asc

import (
	"context"
	"testing"
)

func TestUpdateAgeRatingDeclaration(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AgeRatingDeclarationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateAgeRatingDeclaration(ctx, "10", AgeRatingDeclarationUpdateRequest{})
	})
}
