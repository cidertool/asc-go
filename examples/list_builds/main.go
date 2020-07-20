package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/aaronsky/asc-go/v1/asc"
	"github.com/aaronsky/asc-go/v1/asc/builds"
)

var (
	keyID          = flag.String("kid", "", "key ID")
	issuerID       = flag.String("iss", "", "issuer ID")
	privateKeyPath = flag.String("private_key", "", "private key used to sign authorization token")
	appID          = flag.String("app", "", "App to list builds for")
)

func main() {
	flag.Parse()

	// Read private key file as []byte
	secret, err := ioutil.ReadFile(*privateKeyPath)
	if err != nil {
		log.Fatalf("secret key could not be loaded: %s", err)
	}
	// Create the token using the required information
	auth, err := asc.NewTokenConfig(*keyID, *issuerID, time.Now().Add(20*time.Minute), secret)
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}
	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	var cursor string
	params := builds.ListBuildsForAppQuery{}
	for {
		if cursor != "" {
			params.Cursor = &cursor
		}
		b, _, err := client.Builds.ListBuildsForApp("com.wayfair.WayfairApp", &params)
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
