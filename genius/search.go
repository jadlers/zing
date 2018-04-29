package genius

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var clientAccessToken string

const baseURL = "https://api.genius.com/"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Needed for API calls to genius.")
	}
	clientAccessToken = os.Getenv("CLIENT_ACCESS_TOKEN")
}

// GetLinksFor searches the genius database for song & artist info for a query.
// It returns a string including links to more info on the best search result.
func GetLinksFor(query string) string {
	searchRes := makeSearchRequest(query)
	linkString := extractLinks(searchRes)
	return linkString
}

func extractLinks(res SearchResponse) string {
	if len(res.Response.Hits) < 1 {
		return "No result"
	}
	fstHit := res.Response.Hits[0].Result
	fullTitle := fstHit.FullTitle
	links := fmt.Sprintf(`Showing links for: %s`, fullTitle)
	return links
}

func urlSafeString(s string) string {
	spacesRemoved := strings.Replace(s, " ", "%20", -1)
	return spacesRemoved
}

func makeSearchRequest(search string) SearchResponse {
	safeString := urlSafeString(search)
	url := searchRequestURL(safeString)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var body SearchResponse
	if json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal("Couldn't convert response to struct")
	}

	return body
}

func searchRequestURL(query string) string {
	search := "search?q=" + urlSafeString(query)
	accessToken := "&access_token=" + clientAccessToken
	return baseURL + search + accessToken
}
