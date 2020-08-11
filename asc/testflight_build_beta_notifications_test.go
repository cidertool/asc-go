package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAvailableBuildNotification(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildBetaNotificationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateAvailableBuildNotification(ctx, &BuildBetaNotificationCreateRequest{})
	})
}

func TestCreateAvailableBuildNotificationApplyRequestTypes(t *testing.T) {
	var req *BuildBetaNotificationCreateRequest
	req.applyTypes()
	assert.Nil(t, req)

	req = &BuildBetaNotificationCreateRequest{}
	req.applyTypes()
	assert.Equal(t, "buildBetaNotifications", req.Type)
	assert.Nil(t, req.Relationships.Build.Data)

	req = &BuildBetaNotificationCreateRequest{
		Relationships: BuildBetaNotificationCreateRequestRelationships{
			Build: RelationshipDeclaration{
				Data: &RelationshipData{},
			},
		},
	}
	req.applyTypes()
	assert.Equal(t, "builds", req.Relationships.Build.Data.Type)

	req = &BuildBetaNotificationCreateRequest{
		Type: "dog",
		Relationships: BuildBetaNotificationCreateRequestRelationships{
			Build: RelationshipDeclaration{
				Data: &RelationshipData{
					Type: "dog",
				},
			},
		},
	}
	req.applyTypes()
	assert.Equal(t, "dog", req.Type)
	assert.Equal(t, "dog", req.Relationships.Build.Data.Type)
}
