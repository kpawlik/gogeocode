package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func newClient(googleAPIKey, googleChannel, googleID, googleSignature string) (c *maps.Client, err error) {
	if len(googleAPIKey) > 0 && len(googleChannel) > 0 {
		c, err = maps.NewClient(maps.WithAPIKey(googleAPIKey), maps.WithChannel(googleChannel))
		return
	}
	if len(googleAPIKey) > 0 {
		c, err = maps.NewClient(maps.WithAPIKey(googleAPIKey))
		return
	}
	if len(googleChannel) > 0 && len(googleID) > 0 && len(googleSignature) > 0 {
		c, err = maps.NewClient(maps.WithClientIDAndSignature(googleID, googleSignature), maps.WithChannel(googleChannel))
		return
	}
	if len(googleID) > 0 && len(googleSignature) > 0 {
		c, err = maps.NewClient(maps.WithClientIDAndSignature(googleID, googleSignature))
		return
	}
	return nil, fmt.Errorf("Client was not created. Missing one of: API key, ID, signature")

}

func singleGeocode(csvIn *csv.Reader, csvOut *csv.Writer, googleAPIKey, googleChannel, googleID, googleSignature string) (apiStat *stat) {
	var (
		line     []string
		c        *maps.Client
		res      []maps.GeocodingResult
		err      error
		csvLines [][]string
		csvLine  []string
	)

	apiStat = newStat()

	ctx := context.Background()
	c, err = newClient(googleAPIKey, googleChannel, googleID, googleSignature)
	if err != nil {
		log.Printf("Cannot create client: %s\n", err)
		return
	}
	for {
		if line, err = csvIn.Read(); err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading csv line: %s\n", err)
		}

		id := strings.TrimSpace(line[0])
		address := line[1]
		req := &maps.GeocodingRequest{Address: string(line[1])}
		res, err = c.Geocode(ctx, req)
		if isQueryLimitError(err) {
			log.Printf("QUERY LIMIT")
			return
		}
		if err != nil {
			log.Printf("Geocoding error: %s. Line: %s\n", err, line)
			apiStat.addIgnored()
			continue
		}
		csvLine = make([]string, 4)
		csvLine[0] = id
		csvLine[1] = address

		csvLines = append(csvLines, csvLine)
		if success := len(res) > 0; success {
			csvLine[2] = fmt.Sprintf("%f", res[0].Geometry.Location.Lat)
			csvLine[3] = fmt.Sprintf("%f", res[0].Geometry.Location.Lng)
			apiStat.success(true)
		} else {
			apiStat.success(false)
		}
		if err = csvOut.Write(csvLine); err != nil {
			log.Fatalf("Error write to file %s. Line %s\n", outFileName, csvLine)
		}
		csvOut.Flush()
	}
	return apiStat

}
