package main

import "log"

type stat struct {
	geocoded, notGeocoded, ignored int
}

func newStat() *stat {

	return &stat{}
}

func (s *stat) success(ok bool) {
	if ok {
		s.geocoded++
	} else {
		s.notGeocoded++
	}
}

func (s *stat) addIgnored() {
	s.ignored++
}

// Prints stats. How many records were geocoded by each apiKey
func (s *stat) print() {
	log.Println("Stats:")
	log.Printf("Rows geocoded: %d\n", s.geocoded)
	log.Printf("Rows not geocoded: %d\n", s.notGeocoded)
	log.Printf("Geocode errors: %d\n", s.ignored)
	log.Println("Rows geocoded by API key:")
}
