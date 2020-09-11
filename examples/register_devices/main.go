package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/cidertool/asc-go/asc"
	"github.com/cidertool/asc-go/examples/util"
)

var (
	name     = flag.String("name", "", "Name of the device")
	udid     = flag.String("udid", "", "UDID of the device")
	platform = flag.String("platform", "IOS", "Platform (IOS or MAC_OS)")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	device, _, err := client.Provisioning.CreateDevice(ctx, *name, *udid, asc.BundleIDPlatform(*platform))

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
