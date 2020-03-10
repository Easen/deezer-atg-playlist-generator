package deezer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"
)

type artistSearchResult struct {
	Data    []Artist `json:"data"`
	NextURL string   `json:"next,omitempty"`
}

var (
	deezerSearchURL = "https://api.deezer.com/search/artist"
)

// SearchForArtist return the matched Artist
func SearchForArtist(artistName string) (*Artist, error) {
	var result *Artist
	url := fmt.Sprintf("%s?q=%s", deezerSearchURL, url.QueryEscape(artistName))
	res, httpErr := http.Get(url)
	if httpErr != nil {
		return result, httpErr
	}
	log.Printf("Invoked %s return status code: %d\n", url, res.StatusCode)

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return result, fmt.Errorf("Status code was %d", res.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return result, bodyErr
	}

	log.Printf("Received body: %s\n", string(body))

	searchResult := artistSearchResult{}
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return result, err
	}

	log.Printf("Unmarshalled into %s\n", spew.Sdump(searchResult))

	if len(searchResult.Data) == 1 {
		return &searchResult.Data[0], nil
	}
	for _, artist := range searchResult.Data {
		if artist.Name == artistName {
			return &artist, nil
		}
	}

	return result, nil
}
