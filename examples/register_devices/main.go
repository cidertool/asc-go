package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/aaronsky/asc-go"
	"github.com/aaronsky/asc-go/provisioning"
)

var (
	keyID          = flag.String("kid", "", "key ID")
	issuerID       = flag.String("iss", "", "issuer ID")
	privateKeyPath = flag.String("private_key", "", "private key used to sign authorization token")
	deviceName     = flag.String("device_name", "", "Name of the iOS device being added")
	deviceUDID     = flag.String("device_udid", "", "UDID of the iOS device being added")
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

	device, _, err := client.Provisioning.CreateDevice(&provisioning.DeviceCreateRequest{
		Attributes: struct {
			Name     string                        "json:\"name\""
			Platform provisioning.BundleIDPlatform "json:\"platform\""
			UDID     string                        "json:\"udid\""
		}{
			Name:     *deviceName,
			Platform: provisioning.IOS,
			UDID:     *deviceUDID,
		},
		Type: "devices",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"Successfully registered %s (%s) at %s",
		*device.Data.Attributes.Name,
		*device.Data.Attributes.Model,
		*device.Data.Attributes.AddedDate,
	)

}
