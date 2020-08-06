package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	name     = flag.String("name", "", "Name of the device")
	udid     = flag.String("udid", "", "UDID of the device")
	platform = flag.String("platform", "IOS", "Platform (IOS or MAC_OS)")
)

func main() {
	flag.Parse()

	auth, err := util.Token()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	device, _, err := client.Provisioning.CreateDevice(&asc.DeviceCreateRequest{
		Attributes: asc.DeviceCreateRequestAttributes{
			Name:     *name,
			Platform: asc.BundleIDPlatform(*platform),
			UDID:     *udid,
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
