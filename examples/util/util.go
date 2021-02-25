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

package util

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/cidertool/asc-go/asc"
)

var (
	keyID          = flag.String("kid", "", "key ID")
	issuerID       = flag.String("iss", "", "issuer ID")
	privateKey     = flag.String("privatekey", "", "private key used to sign authorization token")
	privateKeyPath = flag.String("privatekeypath", "", "path to a private key used to sign authorization token")
)

// TokenConfig creates the auth transport using the required information
func TokenConfig() (auth *asc.AuthTransport, err error) {
	var secret []byte
	if *privateKey != "" {
		secret = []byte(*privateKey)
	} else if *privateKeyPath != "" {
		// Read private key file as []byte
		secret, err = os.ReadFile(*privateKeyPath)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no private key provided to either the -privatekey or -privatekeypath flags")
	}
	// Create the token using the required information
	auth, err = asc.NewTokenConfig(*keyID, *issuerID, 20*time.Minute, secret)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// GetApp returns a single asc.App by filtering by its full name on App Store Connect
func GetApp(ctx context.Context, client *asc.Client, params *asc.ListAppsQuery) (*asc.App, error) {
	apps, _, err := client.Apps.ListApps(ctx, params)
	if err != nil {
		return nil, err
	} else if len(apps.Data) == 0 {
		return nil, fmt.Errorf("query for app returned no data")
	}
	return &apps.Data[0], nil
}

// Close closes an open descriptor.
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
