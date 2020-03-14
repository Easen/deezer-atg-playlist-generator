package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/easen/deezer-atg-playlist-generator/atg"
	"github.com/easen/deezer-atg-playlist-generator/deezer"

	_ "github.com/joho/godotenv/autoload"
)

const maxConcurrency = 4

var throttle = make(chan int, maxConcurrency)

func main() {

	artists, err := atg.GetAllATGArtists()
	if err != nil {
		panic(err)
	}
	trackIDChannel := make(chan int)

	var wg sync.WaitGroup
	go func() {
		for _, artist := range artists {
			throttle <- 1
			wg.Add(1)

			go func(artist string, wg *sync.WaitGroup, throttle chan int) {
				defer wg.Done()
				defer func() {
					<-throttle
				}()
				deezerArtist, err := deezer.SearchForArtist(artist)
				if err != nil {
					log.Panicf("Error occured while trying to find the artist \"%s\"", artist)
					return
				}

				if deezerArtist == nil {
					log.Printf("Cannot find artist: %s", artist)
					return
				}

				topTracks, err := deezerArtist.GetTopTracks(3)
				if err != nil {
					log.Panicf("Error occured while trying to find the top tracks for artist \"%s\"", artist)
					return
				}
				for index, track := range topTracks {
					log.Printf("Top track for %s: %d - %d - %s", artist, index, track.ID, track.Title)
					trackIDChannel <- track.ID
				}
			}(artist, &wg, throttle)
		}
		go func() {
			wg.Wait()
			close(trackIDChannel)
		}()
	}()

	var songs strings.Builder
	for trackID := range trackIDChannel {
		if songs.Len() > 0 {
			fmt.Fprintf(&songs, ",")
		}
		fmt.Fprintf(&songs, "%s", strconv.Itoa(trackID))

	}

	log.Printf("Finished songs=%s", songs.String())
}
