package main

import (
	"fmt"

	"github.com/easen/deezer-atg-playlist-generator/atg"
)

func main() {
	artists, err := atg.GetAllATGArtists()
	if err != nil {
		panic(err)
	}
	for _, artist := range artists {
		fmt.Println(artist)
	}
	fmt.Println("Finished")
}
