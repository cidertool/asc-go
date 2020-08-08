package util

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aaronsky/asc-go/asc"
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
		secret, err = ioutil.ReadFile(*privateKeyPath)
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
