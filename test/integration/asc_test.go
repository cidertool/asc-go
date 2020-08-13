// +build integration

package integration

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/aaronsky/asc-go/asc"
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
	token, err := tokenConfig()
	if err != nil {
		log.Fatal(err)
	}
	client = asc.NewClient(token.Client())
}

// TokenConfig creates the auth transport using the required information
func tokenConfig() (auth *asc.AuthTransport, err error) {
	var secret []byte
	if key := os.Getenv(envPrivateKey); key != "" {
		secret = []byte(key)
	} else if keyPath := os.Getenv(envPrivateKeyPath); keyPath != "" {
		// Read private key file as []byte
		secret, err = ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no private key provided to either the ASC_INTEGRATION_PRIVATE_KEY or ASC_INTEGRATION_PRIVATE_KEY_PATH environment variables")
	}
	keyID := os.Getenv(envKeyID)
	issuerID := os.Getenv(envIssuerID)
	// Create the token using the required information
	auth, err = asc.NewTokenConfig(keyID, issuerID, 20*time.Minute, secret)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
