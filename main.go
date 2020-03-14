package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/easen/deezer-atg-playlist-generator/atg"
	"github.com/easen/deezer-atg-playlist-generator/deezer"

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

	artists, err := atg.GetAllATGArtists()
	if err != nil {
		panic(err)
	}
	tracksIDs := getTopTrackIDsForArtists(artists)
	playlistID, _ := strconv.Atoi(deezerPlaylistID)
	err = deezer.UpdatePlaylistTracks(deezerAccessToken, playlistID, tracksIDs)
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
	deezerArtist, err := deezer.SearchForArtistViaAPI(artist)
	if err != nil {
		log.Panicf("Error occured while trying to find the artist \"%s\" via API", artist)
		return trackIDs
	}
	var artistID int
	if deezerArtist == nil {
		webArtistID, err := deezer.SearchForArtistIDViaWeb(artist)
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

	topTracks, err := deezer.GetTopTracksForArtistID(artistID, getTopTrackLimit())
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
