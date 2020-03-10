package deezer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchForArtistNoMatches(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
    "data": []
}`)
	}))
	defer ts.Close()

	deezerSearchURL = ts.URL

	artist, err := SearchForArtist("empty")

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, artist, "artist should be `nil`")
}

func TestSearchForArtistOneMatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
  "data": [
    {
      "id": "7273",
      "name": "Cult of Luna",
      "link": "https://www.deezer.com/artist/7273",
      "picture": "https://api.deezer.com/artist/7273/image",
      "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
      "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
      "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
      "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
      "nb_album": 24,
      "nb_fan": 14596,
      "radio": true,
      "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
      "type": "artist"
    }
  ],
  "total": 1
}`)
	}))
	defer ts.Close()

	deezerSearchURL = ts.URL

	artist, err := SearchForArtist("Cult of Luna")

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}

	assert.NotNil(t, artist, "artist should not be `nil`")
	assert.Equal(t, "Cult of Luna", artist.Name, "artist should not be `nil`")
}
