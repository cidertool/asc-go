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
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cidertool/asc-go/asc"
	"github.com/cidertool/asc-go/examples/util"
)

var (
	bundleID          = flag.String("bundleid", "", "Bundle ID for an app")
	platform          = flag.String("platform", "IOS", "Platform to query versions for (IOS, MAC_OS, TV_OS)")
	versionString     = flag.String("version", "", "Version string")
	locale            = flag.String("locale", "", "Locale to add previews to")
	previewTypeString = flag.String("previewtype", "", "Preview type")
	previewFile       = flag.String("previewfile", "", "Path to a file to upload as a preview")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	// 1. Create an Authorization header value with bearer token (JWT).
	//    The token is set to expire in 20 minutes, and is used for all App Store
	//    Connect API calls.
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	// 2. Look up the app by bundle ID.
	//    If the app is not found, report an error and exit.
	app, err := util.GetApp(ctx, client, &asc.ListAppsQuery{
		FilterBundleID: []string{*bundleID},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 3. Look up the version version by platform and version number.
	//    If the version is not found, report an error and exit.
	versions, _, err := client.Apps.ListAppStoreVersionsForApp(ctx, app.ID, &asc.ListAppStoreVersionsQuery{
		FilterVersionString: []string{*versionString},
		FilterPlatform:      []string{*platform},
	})
	if err != nil {
		log.Fatal(err)
	}

	var version asc.AppStoreVersion
	if len(versions.Data) > 0 {
		version = versions.Data[0]
	} else {
		log.Fatalf("No app store version found with version %s", *versionString)
	}

	// 4. Get all localizations for the version and look for the requested locale.
	localizations, _, err := client.Apps.ListLocalizationsForAppStoreVersion(ctx, version.ID, nil)
	if err != nil {
		fmt.Println(err)
	}
	var selectedLocalizations []asc.AppStoreVersionLocalization
	for _, loc := range localizations.Data {
		if *loc.Attributes.Locale == *locale {
			selectedLocalizations = append(selectedLocalizations, loc)
		}
	}

	// 5. If the requested localization does not exist, create it.
	//    Localized attributes are copied from the primary locale so there's
	//    no need to worry about them here.
	var selectedLocalization asc.AppStoreVersionLocalization
	if len(selectedLocalizations) > 0 {
		selectedLocalization = selectedLocalizations[0]
	} else {
		newLocalization, _, err := client.Apps.CreateAppStoreVersionLocalization(ctx, asc.AppStoreVersionLocalizationCreateRequestAttributes{
			Locale: *locale,
		}, version.ID)
		if err != nil {
			log.Fatal(err)
		}
		selectedLocalization = newLocalization.Data
	}

	// 6. Get all available app preview sets from the localization.
	//    If a preview set for the desired preview type already exists, use it.
	//    Otherwise, make a new one.
	var previewSets asc.AppPreviewSetsResponse
	_, err = client.FollowReference(ctx, selectedLocalization.Relationships.AppPreviewSets.Links.Related, &previewSets)
	if err != nil {
		fmt.Println(err)
	}
	previewType := asc.PreviewType(*previewTypeString)
	var selectedPreviewSets []asc.AppPreviewSet
	for _, set := range previewSets.Data {
		if *set.Attributes.PreviewType == previewType {
			selectedPreviewSets = append(selectedPreviewSets, set)
		}
	}

	// 7. If an app preview set for the requested type doesn't exist, create it.
	var selectedPreviewSet asc.AppPreviewSet
	if len(selectedPreviewSets) > 0 {
		selectedPreviewSet = selectedPreviewSets[0]
	} else {
		newPreviewSet, _, err := client.Apps.CreateAppPreviewSet(ctx, previewType, selectedLocalization.ID)
		if err != nil {
			log.Fatal(err)
		}
		selectedPreviewSet = newPreviewSet.Data
	}

	// 8. Reserve an app preview in the selected app preview set.
	//    Tell the API to create a preview before uploading the
	//    preview data.
	file, err := os.Open(*previewFile)
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}
	fmt.Println("Reserving space for a new app preview.")
	reservePreview, _, err := client.Apps.CreateAppPreview(ctx, file.Name(), stat.Size(), selectedPreviewSet.ID)
	if err != nil {
		fmt.Println(err)
	}
	preview := reservePreview.Data

	// 9. Upload each part according to the returned upload operations.
	//     The reservation returned uploadOperations, which instructs us how
	//     to split the asset into parts. Upload each part individually.
	//     Note: To speed up the process, upload multiple parts asynchronously
	//     if you have the bandwidth.
	uploadOperations := preview.Attributes.UploadOperations
	fmt.Printf("Uploading %d preview components\n", len(uploadOperations))
	err = client.Upload(ctx, uploadOperations, file)
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}

	// 10. Commit the reservation and provide a checksum.
	//     Committing tells App Store Connect the script is finished uploading parts.
	//     App Store Connect uses the checksum to ensure the parts were uploaded
	//     successfully.
	fmt.Println("Commit the reservation")
	previewURL := preview.Links.Self
	checksum, err := md5Checksum(*previewFile)
	if err != nil {
		log.Fatalf("file checksum could not be calculated: %s", err)
	}

	client.Apps.CommitAppPreview(ctx, preview.ID, asc.Bool(true), &checksum, nil)

	// Report success to the caller.
	fmt.Printf("\nApp Preview successfully uploaded to:\n%s\nYou can verify success in App Store Connect or using the API.\n\n", previewURL.String())
}

func md5Checksum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer util.Close(f)

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
