// +build integration

package integration

import (
	"fmt"
	"io/ioutil"
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
		privateKey, err = ioutil.ReadFile(keyPath)
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
