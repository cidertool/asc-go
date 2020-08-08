package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAvailableBuildNotification(t *testing.T) {
	marshaled := `{"data":{"type":"buildBetaNotifications","id":"10","links":{"self":"https://api.appstoreconnect.apple.com/v1"}},"links":{"self":"https://api.appstoreconnect.apple.com/v1"}}`
	client, server := newServer(marshaled)
	defer server.Close()

	self := Reference{mustParseURL(t, "https://api.appstoreconnect.apple.com/v1")}
	want := &BuildBetaNotificationResponse{
		BuildBetaNotification{
			Type: "buildBetaNotifications",
			ID:   "10",
			Links: ResourceLinks{
				Self: self,
			},
		},
		DocumentLinks{
			Self: self,
		},
	}

	got, resp, err := client.TestFlight.CreateAvailableBuildNotification(context.Background(), &BuildBetaNotificationCreateRequest{
		Type: "buildBetaNotifications",
		Relationships: BuildBetaNotificationCreateRequestRelationships{
			Build: struct {
				Data RelationshipsData "json:\"data\""
			}{
				Data: RelationshipsData{
					ID:   "10",
					Type: "builds",
				},
			},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
