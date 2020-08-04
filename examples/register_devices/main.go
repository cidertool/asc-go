package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	deviceName = flag.String("device_name", "", "Name of the iOS device being added")
	deviceUDID = flag.String("device_udid", "", "UDID of the iOS device being added")
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
			Name:     *deviceName,
			Platform: asc.BundleIDPlatformiOS,
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
