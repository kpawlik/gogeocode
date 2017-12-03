package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	VERSION = "1.0.20171201"
)

var (
	googleID                string
	googleSignature         string
	googleAPIKey            string
	inFileName, outFileName string
	mockGeocoder            bool
	googleChannel           string
)

//initialize flags
func init() {
	var (
		version bool
		help    bool
	)
	flag.StringVar(&inFileName, "in", "/tmp/icoms_in.csv", "Input CSV file path")
	flag.StringVar(&outFileName, "out", "/tmp/icoms_out.csv", "Output CSV file path")
	flag.BoolVar(&mockGeocoder, "mock", false, "Mock geocoder. (for test only)")
	flag.BoolVar(&version, "version", false, "Print version")
	flag.BoolVar(&help, "help", false, "Print help")
	flag.StringVar(&googleAPIKey, "apiKey", "", "Google API key")
	flag.StringVar(&googleChannel, "channel", "", "Google channel name")
	flag.StringVar(&googleID, "id", "", "Google ID")
	flag.StringVar(&googleSignature, "signature", "", "Google signature")
	flag.Parse()
	if version {
		fmt.Printf("Version: %s\n", VERSION)
		os.Exit(0)
	}
	if help {
		fmt.Println("Geocode data from CSV file, store results in CSV file.\n" +
			"Example call:\n" +
			"   gogeocode -apiKey <YOUR_API_KEY> -in /tmp/address_in.csv -out /tmp/address_in.csv\n" +
			"   gogeocode -id <CLIENT_ID> -signature <CLIENT_SIGNATURE> -in /tmp/address_in.csv -out /tmp/address_in.csv\n" +
			"Input data format:  ID,ADDRESS. \n" +
			"Output data format: ID, ADDRESS, LAT, LNG\n\n" +
			"For more details visit: https://github.com/kpawlik/gogeocode\n\n" +
			"Parameters: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(googleAPIKey) == 0 && len(googleID) == 0 && len(googleSignature) == 0 {
		fmt.Println("Missing apiKey parameter, or ID and signature")
		os.Exit(1)
	}

}

// Checks if err is type og OVER_QUERY_LIMIT
func isQueryLimitError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "OVER_QUERY_LIMIT")
}

func main() {
	var (
		csvInFile, csvOutFile *os.File
		err                   error
	)

	ts := time.Now()
	if csvInFile, err = os.Open(inFileName); err != nil {
		log.Fatalf("Cannot open file %s\n", inFileName)
	}
	defer csvInFile.Close()
	if csvOutFile, err = os.Create(outFileName); err != nil {
		log.Fatalf("Cannot open file %s\n", outFileName)
	}
	defer csvOutFile.Close()
	csvIn, csvOut := csv.NewReader(csvInFile), csv.NewWriter(csvOutFile)

	stats := singleGeocode(csvIn, csvOut, googleAPIKey, googleChannel, googleID, googleSignature)
	stats.print()
	log.Printf("Time: %s\n", time.Now().Sub(ts))

}
