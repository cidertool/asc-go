package main

import (
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

	auth, err := util.Token()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())
	app, err := util.GetApp(client, &asc.ListAppsQuery{
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
		b, _, err := client.Builds.ListBuildsForApp(app.ID, &params)
		if err != nil {
			log.Fatal(err)
			break
		}

		for _, user := range b.Data {
			fmt.Println(*user.Attributes.Version)
		}

		cursor = b.Links.Next.Cursor()
		if cursor == "" {
			break
		}
	}
}
