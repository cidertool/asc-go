// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBuilds(t *testing.T) {
	builds, _, err := client.Builds.ListBuilds(context.Background(), nil)
	assert.NoError(t, err, "ListBuilds responded with an error")
	assert.NotEmpty(t, builds.Data, "ListBuilds returned no builds")

	firstBuild := builds.Data[0]
	build, _, err := client.Builds.GetBuild(context.Background(), firstBuild.ID, nil)
	assert.NoError(t, err, "GetBuild responded with an error")
	assert.NotNil(t, build.Data.Attributes)

	app, _, err := client.Builds.GetAppForBuild(context.Background(), firstBuild.ID, nil)
	assert.NoError(t, err, "GetAppForBuild responded with an error")
	assert.NotNil(t, app.Data.Attributes)
}
