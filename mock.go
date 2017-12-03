package main

import (
	"fmt"

	"golang.org/x/net/context"

	"googlemaps.github.io/maps"
)

type geocoder interface {
	Geocode(context.Context, *maps.GeocodingRequest) ([]maps.GeocodingResult, error)
}

type mock struct {
	apiKey  string
	apiName string
	cnt     int
}

func newMock(api string, name string) (*mock, error) {
	return &mock{apiKey: api, apiName: name}, nil
}

func (m *mock) Geocode(ctx context.Context, req *maps.GeocodingRequest) (resp []maps.GeocodingResult, err error) {
	m.cnt++
	// if m.apiKey == "passen" && m.cnt == 2 {
	return nil, fmt.Errorf("OVER_QUERY_LIMIT")
	// }
	return
}
