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
