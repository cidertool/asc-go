package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListIconsForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildIconsResponse{}
	got, resp, err := client.Builds.ListIconsForBuild(context.Background(), "10", &ListIconsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
