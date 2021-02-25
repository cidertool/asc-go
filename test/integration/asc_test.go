// +build integration

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

package integration

import (
	"fmt"
	"os"
	"time"

	"github.com/cidertool/asc-go/asc"
)

const (
	envKeyID          = "ASC_INTEGRATION_KID"
	envIssuerID       = "ASC_INTEGRATION_ISS"
	envPrivateKey     = "ASC_INTEGRATION_PRIVATE_KEY"
	envPrivateKeyPath = "ASC_INTEGRATION_PRIVATE_KEY_PATH"
)

var (
	client *asc.Client
)

func init() {
	token := tokenConfig()
	if token == nil {
		return
	}
	client = asc.NewClient(token.Client())
}

// TokenConfig creates the auth transport using the required information
func tokenConfig() *asc.AuthTransport {
	var privateKey []byte
	var err error
	if key := os.Getenv(envPrivateKey); key != "" {
		privateKey = []byte(key)
	} else if keyPath := os.Getenv(envPrivateKeyPath); keyPath != "" {
		// Read private key file as []byte
		privateKey, err = os.ReadFile(keyPath)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	} else {
		fmt.Println("no private key provided to either the ASC_INTEGRATION_PRIVATE_KEY or ASC_INTEGRATION_PRIVATE_KEY_PATH environment variables")
		return nil
	}
	keyID := os.Getenv(envKeyID)
	issuerID := os.Getenv(envIssuerID)
	expiryDuration := 20 * time.Minute
	// Create the token using the required information
	auth, err := asc.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return auth
}
