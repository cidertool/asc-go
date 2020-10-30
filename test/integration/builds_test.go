// +build integration

/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

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
