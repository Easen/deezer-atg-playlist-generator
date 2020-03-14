package deezer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SETUP
// Importantly you need to call Run() once you've done what you need
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetTopTracksWithNoTrackURL(t *testing.T) {
	testArtist := Artist{TracklistURL: nil}

	tracks, err := testArtist.GetTopTracks(3)

	assert.Nil(t, err, "err should not be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 0, "tracks should have the length of 0")
}

func TestGetTopTracksWithUnableToParseTrackURL(t *testing.T) {
	testArtist := Artist{TracklistURL: ":INVALID"}

	tracks, err := testArtist.GetTopTracks(3)

	assert.NotNil(t, err, "err should not be `nil`")
	if err == nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 0, "tracks should have the length of 0")
}

func TestGetTopTracksWithNoTracks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
    "data": []
}`)
	}))
	defer ts.Close()

	testArtist := Artist{TracklistURL: ts.URL}

	tracks, err := testArtist.GetTopTracks(3)

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 0, "tracks should have the length of 0")
}

func TestGetTopTracksWith1Track(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
  "data": [
    {
      "id": 711332042,
      "readable": true,
      "title": "A Dawn to Fear",
      "title_short": "A Dawn to Fear",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332042",
      "duration": 533,
      "rank": 366436,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-0.dzcdn.net/stream/c-0f1acd2a53faceb4a080e6815a5d665d-4.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
		}
	],
  "total": 1,
  "next": "https://api.deezer.com/search?strict=on&q=artist%3A%22Cult%20Of%20Luna%22&index=25"
}`)
	}))
	defer ts.Close()

	testArtist := Artist{TracklistURL: ts.URL}

	tracks, err := testArtist.GetTopTracks(2)

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 1, "tracks should have the length of 1")
}

func TestGetTopTracksWith2Tracks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
  "data": [
    {
      "id": 711332042,
      "readable": true,
      "title": "A Dawn to Fear",
      "title_short": "A Dawn to Fear",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332042",
      "duration": 533,
      "rank": 366436,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-0.dzcdn.net/stream/c-0f1acd2a53faceb4a080e6815a5d665d-4.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
    },
    {
      "id": 711332022,
      "readable": true,
      "title": "The Silent Man",
      "title_short": "The Silent Man",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332022",
      "duration": 636,
      "rank": 431844,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-c.dzcdn.net/stream/c-cdb27c46e8bd41f33a161897e831cd0c-3.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
		}
	],
  "total": 2,
  "next": "https://api.deezer.com/search?strict=on&q=artist%3A%22Cult%20Of%20Luna%22&index=25"
}`)
	}))
	defer ts.Close()

	testArtist := Artist{TracklistURL: ts.URL}

	tracks, err := testArtist.GetTopTracks(2)

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 2, "tracks should have the length of 2")
}

func TestGetTopTracksWith3Tracks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
{
  "data": [
    {
      "id": 711332042,
      "readable": true,
      "title": "A Dawn to Fear",
      "title_short": "A Dawn to Fear",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332042",
      "duration": 533,
      "rank": 366436,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-0.dzcdn.net/stream/c-0f1acd2a53faceb4a080e6815a5d665d-4.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
    },
    {
      "id": 711332022,
      "readable": true,
      "title": "The Silent Man",
      "title_short": "The Silent Man",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332022",
      "duration": 636,
      "rank": 431844,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-c.dzcdn.net/stream/c-cdb27c46e8bd41f33a161897e831cd0c-3.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
    },
    {
      "id": 711332032,
      "readable": true,
      "title": "Lay Your Head to Rest",
      "title_short": "Lay Your Head to Rest",
      "title_version": "",
      "link": "https://www.deezer.com/track/711332032",
      "duration": 383,
      "rank": 337298,
      "explicit_lyrics": false,
      "explicit_content_lyrics": 0,
      "explicit_content_cover": 2,
      "preview": "https://cdns-preview-b.dzcdn.net/stream/c-bc303af11715aac52c468bed562d70cc-3.mp3",
      "artist": {
        "id": 7273,
        "name": "Cult of Luna",
        "link": "https://www.deezer.com/artist/7273",
        "picture": "https://api.deezer.com/artist/7273/image",
        "picture_small": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/56x56-000000-80-0-0.jpg",
        "picture_medium": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/250x250-000000-80-0-0.jpg",
        "picture_big": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/500x500-000000-80-0-0.jpg",
        "picture_xl": "https://e-cdns-images.dzcdn.net/images/artist/982a859d1a65d407e63ee157da52134d/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/artist/7273/top?limit=50",
        "type": "artist"
      },
      "album": {
        "id": 103408372,
        "title": "A Dawn to Fear",
        "cover": "https://api.deezer.com/album/103408372/image",
        "cover_small": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/56x56-000000-80-0-0.jpg",
        "cover_medium": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/250x250-000000-80-0-0.jpg",
        "cover_big": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/500x500-000000-80-0-0.jpg",
        "cover_xl": "https://e-cdns-images.dzcdn.net/images/cover/92d3ad878e7833dbfcaacad27fcb6c41/1000x1000-000000-80-0-0.jpg",
        "tracklist": "https://api.deezer.com/album/103408372/tracks",
        "type": "album"
      },
      "type": "track"
		}
	],
  "total": 3,
  "next": "https://api.deezer.com/search?strict=on&q=artist%3A%22Cult%20Of%20Luna%22&index=25"
}`)
	}))
	defer ts.Close()

	testArtist := Artist{TracklistURL: ts.URL}

	tracks, err := testArtist.GetTopTracks(2)

	assert.Nil(t, err, "err should be `nil`")
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, tracks, 2, "tracks should have the length of 2")
}
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
      "id": 7273,
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
