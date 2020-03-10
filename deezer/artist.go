package deezer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// Artist Deezer Artist response
type Artist struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Link          string      `json:"link"`
	Picture       string      `json:"picture"`
	PictureSmall  string      `json:"picture_small"`
	PictureMedium string      `json:"picture_medium"`
	PictureBig    string      `json:"picture_big"`
	PictureXl     string      `json:"picture_xl"`
	TracklistURL  interface{} `json:"tracklist"`
	Type          string      `json:"type"`
}

type artistTracklist struct {
	Data []Track `json:"data"`
}

// GetTopTracks for the currrent Artist
func (a Artist) GetTopTracks(limit int) ([]Track, error) {
	var tracks []Track
	if a.TracklistURL == nil {
		return tracks, nil
	}

	tracklistURL, ok := a.TracklistURL.(string)
	if ok == false {
		return tracks, fmt.Errorf("Tracklist is not a string")
	}
	parsedURL, parsedURLErr := url.Parse(tracklistURL)
	if parsedURLErr != nil {
		return tracks, parsedURLErr
	}

	parsedURL.Query().Set("limit", strconv.Itoa(limit))

	requestURL := parsedURL.String()

	log.Printf("Get Tracks URL: %s\n", requestURL)
	println(requestURL)

	res, httpErr := http.Get(requestURL)
	if httpErr != nil {
		return tracks, httpErr
	}
	log.Printf("Invoked %s return status code: %d\n", requestURL, res.StatusCode)

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return tracks, fmt.Errorf("Status code was %d", res.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return tracks, bodyErr
	}

	log.Printf("Received body: %s\n", string(body))

	var result artistTracklist
	if err := json.Unmarshal(body, &result); err != nil {
		return tracks, err
	}

	log.Printf("Unmarshalled into %s\n", spew.Sdump(result))

	if len(result.Data) <= limit {
		return result.Data, nil
	}

	return result.Data[0:limit], nil
}