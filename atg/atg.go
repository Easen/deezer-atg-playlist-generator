package atg

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	atgLineUp  = "https://arctangent.co.uk/line-up/"
	atgAjaxURL = "https://arctangent.co.uk/wp-admin/admin-ajax.php"
)

func getAjaxNonce() (string, error) {
	var nilResponse string
	resp, err := http.Get(atgLineUp)
	if err != nil {
		return nilResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nilResponse, err
	}
	re := regexp.MustCompile(`"ajax_nonce":"([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))
	return matches[1], nil
}

// GetAllATGArtists Get All ATG Artists
func GetAllATGArtists() ([]string, error) {
	fmt.Println("Fetching all of the ATG artists")
	artists := []string{}
	pagenum := 1
	limit := 80
	ajaxNonce, ajaxNonceErr := getAjaxNonce()
	if ajaxNonceErr != nil {
		return nil, ajaxNonceErr
	}
	for {
		log.Println("Making call to url: %s, ajax_nonce: %s, pagenum: %d, limit: %d", atgAjaxURL, ajaxNonce, pagenum, limit)
		resp, err := http.PostForm(atgAjaxURL,
			url.Values{"action": {"noisa_artists_filter"},
				"ajax_nonce":       {ajaxNonce},
				"obj[action]":      {"noisa_artists_filter"},
				"obj[filterby]":    {"taxonomy"},
				"obj[cpt]":         {"noisa_artists"},
				"obj[tax]":         {"noisa_artists_cats"},
				"obj[limit]":       {strconv.Itoa(limit)},
				"obj[filter_name]": {"all"},
				"obj[pagenum]":     {strconv.Itoa(pagenum)}})
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile(`<h2 class="grid-title">([^<]+)<\/h2>`)
		matches := re.FindAllStringSubmatch(string(body), -1)
		log.Println("Found %d artists", len(matches))
		if len(matches) == 0 {
			break
		}

		for _, v := range matches {
			artist := strings.ToLower(html.UnescapeString(v[1]))
			artists = append(artists, artist)
		}
		pagenum++
	}
	return artists, nil
}
