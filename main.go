package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/easen/deezer-boa-playlist-generator/bao"
	"github.com/easen/godeezer"

	_ "github.com/joho/godotenv/autoload"
)

const defaultTopTrackLimit = 3

var (
	deezerAccessToken = os.Getenv("DEEZER_ACCESS_TOKEN")
	deezerPlaylistID  = os.Getenv("DEEZER_PLAYLIST_ID")
	topTrackLimit     = os.Getenv("TOP_TRACK_LIMIT")
)

func main() {
	if deezerAccessToken == "" || deezerPlaylistID == "" {
		printUsage()
		return
	}

	artists, err := bao.GetAllBAOArtists()
	if err != nil {
		panic(err)
	}

	tracksIDs := getTopTrackIDsForArtists(artists)
	playlistID, _ := strconv.Atoi(deezerPlaylistID)
	err = godeezer.UpdatePlaylistTracks(deezerAccessToken, playlistID, tracksIDs)
	if err != nil {
		panic(err)
	}
	log.Printf("Updated Playlist")
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Printf("\tDEEZER_ACCESS_TOKEN=<TOKEN> DEEZER_PLAYLIST_ID=<TOKEN> %s\n", os.Args[0])
}

func getTopTrackIDsForArtists(artists []string) []int {
	var trackIDs []int
	for _, artist := range artists {
		for _, trackID := range getTopTrackIDsForArtist(artist) {
			trackIDs = append(trackIDs, trackID)
		}
	}
	return trackIDs
}

func getTopTrackIDsForArtist(artist string) []int {
	var trackIDs []int
	deezerArtist, err := godeezer.SearchForArtistViaAPI(artist)
	if err != nil {
		log.Panicf("Error occured while trying to find the artist \"%s\" via API", artist)
		return trackIDs
	}
	var artistID int
	if deezerArtist == nil {
		webArtistID, err := godeezer.SearchForArtistIDViaWeb(artist)
		if err != nil {
			log.Panicf("Error occured while trying to find the artist \"%s\" via Web", artist)
			return trackIDs
		}
		artistID = webArtistID
	} else {
		artistID = deezerArtist.ID
	}
	fmt.Printf("Artist %s --> %d\n", artist, artistID)
	if artistID == 0 {
		return trackIDs
	}

	topTracks, err := godeezer.GetTopTracksForArtistID(artistID, getTopTrackLimit())
	if err != nil {
		log.Panicf("Error occured while trying to find the top tracks for artist \"%s\"", artist)
		return trackIDs
	}
	for index, track := range topTracks {
		log.Printf("Top track for %s: %d - %d - %s", artist, index, track.ID, track.Title)
		trackIDs = append(trackIDs, track.ID)
	}
	return trackIDs
}

func getTopTrackLimit() int {
	if topTrackLimit == "" {
		return defaultTopTrackLimit
	}
	i, err := strconv.Atoi(topTrackLimit)
	if err != nil {
		panic(err)
	}
	return i
}
