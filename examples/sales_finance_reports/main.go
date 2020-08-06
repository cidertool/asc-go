package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	vendorNumber = flag.String("vendornum", "", "Vendor number")
)

func main() {
	flag.Parse()

	if *vendorNumber == "" {
		log.Fatal("no vendor number provided to the -vendornum flag")
	}

	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	salesReport, _, err := client.Reporting.DownloadSalesAndTrendsReports(&asc.DownloadSalesAndTrendsReportsQuery{
		FilterVendorNumber: []string{*vendorNumber},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received sales report:", salesReport != nil)

	financeReport, _, err := client.Reporting.DownloadFinanceReports(&asc.DownloadFinanceReportsQuery{
		FilterVendorNumber: []string{*vendorNumber},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received finance report:", financeReport != nil)
}
