package util

import (
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

// Token creates the auth transport using the required information
func Token() (auth *asc.AuthTransport, err error) {
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

// GetAppByName returns a single asc.App by filtering by its full name on App Store Connect
func GetAppByName(client *asc.Client, name string) (*asc.App, error) {
	apps, _, err := client.Apps.ListApps(&asc.ListAppsQuery{
		FilterName: []string{name},
	})
	if err != nil {
		return nil, err
	} else if len(apps.Data) == 0 {
		return nil, fmt.Errorf("query for app \"%s\" returned no data", name)
	}
	return &apps.Data[0], nil
}
