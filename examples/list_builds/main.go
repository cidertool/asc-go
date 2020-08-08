package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	appName = flag.String("app", "", "App to list builds for")
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
		FilterName: []string{*appName},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	var cursor string
	params := asc.ListBuildsForAppQuery{}
	for {
		if cursor != "" {
			params.Cursor = cursor
		}
		b, _, err := client.Builds.ListBuildsForApp(ctx, app.ID, &params)
		if err != nil {
			log.Fatal(err)
			break
		}

		for _, user := range b.Data {
			fmt.Println(*user.Attributes.Version)
		}

		if b.Links.Next != nil {
			cursor = b.Links.Next.Cursor()
		}
		if cursor == "" {
			break
		}
	}
}
