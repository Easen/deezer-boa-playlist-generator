package bao

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	boaLineUp = "https://www.bloodstock.uk.com/events/boa-2020/stages"
)

// GetAllBAOArtists Get All Bloodstock Open Air Artists
func GetAllBAOArtists() ([]string, error) {
	artists := []string{}
	log.Printf("Making call to url: %s,", boaLineUp)
	resp, err := http.Get(boaLineUp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`"Band Logo for ([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	log.Printf("Found %d artists", len(matches))
	if len(matches) == 0 {
		return artists, nil
	}

	for _, v := range matches {
		artist := strings.ToLower(html.UnescapeString(v[1]))
		artists = append(artists, artist)
	}
	return artists, nil
}
