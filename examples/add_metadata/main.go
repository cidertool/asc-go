package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	appName = flag.String("app", "", "App to add metadata to")
)

func main() {
	flag.Parse()

	auth, err := util.Token()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	app, err := util.GetAppByName(client, *appName)
	if err != nil {
		log.Fatalf("%s", err)
	}

	appStoreVersions, _, err := client.Apps.ListAppStoreVersionsForApp(app.ID, &asc.ListAppStoreVersionsQuery{
		Include: []string{"appStoreVersionLocalizations"},
		Limit:   1,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	latestVersion := appStoreVersions.Data[0]

	var localization *asc.AppStoreVersionLocalizationResponse

	if latestVersion.Relationships.AppStoreVersionLocalizations.Data != nil {
		localizations := *latestVersion.Relationships.AppStoreVersionLocalizations.Data
		localizationID := localizations[0].ID
		localization, _, err = client.Apps.GetAppStoreVersionLocalization(localizationID, nil)
		if err != nil {
			log.Fatalf("%s", err)
		}
	} else {
		log.Fatalf("no version localizations found")
	}

	set, _, err := client.Apps.CreateAppScreenshotSet(&asc.AppScreenshotSetCreateRequest{
		Attributes: asc.AppScreenshotSetCreateRequestAttributes{
			ScreenshotDisplayType: asc.ScreenshotDisplayTypeAppiPhone65,
		},
		Relationships: asc.AppScreenshotSetCreateRequestRelationships{
			AppStoreVersionLocalization: struct {
				Data asc.RelationshipsData "json:\"data\""
			}{
				Data: asc.RelationshipsData{
					ID:   localization.Data.ID,
					Type: "appStoreVersionLocalizations",
				},
			},
		},
		Type: "appScreenshotSets",
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("ID: %s\n", set.Data.ID)
}
