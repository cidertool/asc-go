package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	reportSales   = flag.Bool("sales", true, "Create a sales report. This is the default and takes precedence even if -finance is set.")
	reportFinance = flag.Bool("finance", false, "Create a finance report")
	vendorNumber  = flag.String("vendornum", "", "Vendor number")
	reportDate    = flag.String("reportdate", "", "Report date to fetch reports for")
	outputFile    = flag.String("output", "", "Path to output gzip file")
)

func main() {
	flag.Parse()

	if *vendorNumber == "" {
		log.Fatal("no vendor number provided to the -vendornum flag")
	}

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	var report io.Reader
	if *reportSales {
		report, _, err = client.Reporting.DownloadSalesAndTrendsReports(ctx, &asc.DownloadSalesAndTrendsReportsQuery{
			FilterVendorNumber:  []string{*vendorNumber},
			FilterReportType:    []string{"SALES"},
			FilterReportSubType: []string{"SUMMARY"},
			FilterFrequency:     []string{"WEEKLY"},
			FilterReportDate:    []string{*reportDate},
		})
	} else if *reportFinance {
		report, _, err = client.Reporting.DownloadFinanceReports(ctx, &asc.DownloadFinanceReportsQuery{
			FilterVendorNumber: []string{*vendorNumber},
			FilterRegionCode:   []string{"US"},
			FilterReportDate:   []string{*reportDate},
			FilterReportType:   []string{"FINANCIAL"},
		})
	} else {
		log.Fatal("No report will be generated, so the program is over.")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received a report!")

	outFile, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, report)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote report to %s\n", outFile.Name())
}
