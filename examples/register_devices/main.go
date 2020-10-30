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
