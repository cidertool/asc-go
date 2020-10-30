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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/cidertool/asc-go/asc"
	"github.com/cidertool/asc-go/examples/util"
)

var (
	bundleID = flag.String("bundleid", "", "Bundle ID for an app")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())
	app, err := util.GetApp(ctx, client, &asc.ListAppsQuery{
		FilterBundleID: []string{*bundleID},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	var cursor string
	params := asc.ListBuildsQuery{
		FilterApp: []string{app.ID},
	}
	for {
		if cursor != "" {
			params.Cursor = cursor
		}
		builds, _, err := client.Builds.ListBuilds(ctx, &params)
		if err != nil {
			log.Fatal(err)
		}

		for _, build := range builds.Data {
			fmt.Println(*build.Attributes.Version)
		}

		if builds.Links.Next != nil {
			cursor = builds.Links.Next.Cursor()
		}
		if cursor == "" {
			break
		}
	}
}
