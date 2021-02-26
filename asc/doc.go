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

/*
Package asc is a Go client library for accessing Apple's App Store Connect API.

Usage

Import the package as you normally would:

	import "github.com/cidertool/asc-go/asc"

Construct a new App Store Connect client, then use the various services on the client to
access different parts of the App Store Connect API. For example:

	client := asc.NewClient(nil)

	// list all apps with the bundle ID "com.sky.MyApp"
	apps, _, err := client.Apps.ListApps(&asc.ListAppsQuery{
		FilterBundleID: []string{"com.sky.MyApp"},
	})

The client is divided into logical chunks closely corresponding to the layout and structure
of Apple's own documentation at https://developer.apple.com/documentation/appstoreconnectapi.

For more sample code snippets, head over to the https://github.com/cidertool/asc-go/tree/main/examples directory.

Authentication

You may find that the code snippet above will always fail due to a lack of authorization.
The App Store Connect API has no methods that allow for unauthorized requests. To make it
easy to authenticate with App Store Connect, the asc-go library offers a solution for signing
and rotating JSON Web Tokens automatically. For example, the above snippet could be made
to look a little more like this:

	import (
		"os"
		"time"

		"github.com/cidertool/asc-go/asc"
	)

	func main() {
		// Key ID for the given private key, described in App Store Connect
		keyID := "...."
		// Issuer ID for the App Store Connect team
		issuerID := "...."
		// A duration value for the lifetime of a token. App Store Connect does not accept
		// a token with a lifetime of longer than 20 minutes
		expiryDuration = 20*time.Minute
		// The bytes of the PKCS#8 private key created on App Store Connect. Keep this key
		// safe as you can only download it once.
		privateKey = os.ReadFile("path/to/key")

		auth, err = asc.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)
		if err != nil {
			return nil, err
		}
		client := asc.NewClient(auth.Client())

		// list all apps with the bundle ID "com.sky.MyApp" in the authenticated user's team
		apps, _, err := client.Apps.ListApps(&asc.ListAppsQuery{
			FilterBundleID: []string{"com.sky.MyApp"},
		})
	}

The authenticated client created here will automatically regenerate the token if it expires.
Also note that all App Store Connect APIs are scoped to the credentials of the pre-configured key,
so you can't use this API to make queries against the entire App Store. For more information on
creating the necessary credentials for the App Store Connect API, see the documentation at
https://developer.apple.com/documentation/appstoreconnectapi/creating_api_keys_for_app_store_connect_api.

Rate Limiting

Apple imposes a rate limit on all API clients. The returned Response.Rate value contains the rate
limit information from the most recent API call. If the API produces a rate limit error, it will be
identifiable as an ErrorResponse with an error code of 429.

Learn more about rate limiting at https://developer.apple.com/documentation/appstoreconnectapi/identifying_rate_limits.

Pagination

All requests for resource collections (apps, builds, beta groups, etc.) support pagination.
Responses for paginated resources will contain a Links property of type PagedDocumentLinks,
with Reference URLs for first, next, and self. A Reference can have its cursor extracted
with the Cursor() method, and that can be passed to a query param using its Cursor field.
You can also find more information about the per-page limit and total count of resources in
the response's Meta field of type PagingInformation.

	auth, _ = asc.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)
	client := asc.NewClient(auth.Client())

	opt := &asc.ListAppsQuery{
		FilterBundleID: []string{"com.sky.MyApp"},
	}

	var allApps []asc.App
	for {
		apps, _, err := apps, _, err := client.Apps.ListApps(opt)
		if err != nil {
			return err
		}
		allApps = append(allApps, apps.Data...)
		if apps.Links.Next == nil {
			break
		}
		cursor := apps.Links.Next.Cursor()
		if cursor == "" {
			break
		}
		opt.Cursor = cursor
	}

*/
package asc
