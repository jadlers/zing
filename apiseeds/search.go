package apiseeds

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var clientApiKey string = "QDR3B5i0Ae5p9x6QFmOvslObrUCV5FE5PxXCr0l78hcLS73Bf78ivOfCpL2FVpzw"

const baseURL = "https://orion.apiseeds.com/api/music/lyric/"

// GetLyrics searches the apiseeds lyrics database for information about the song.
// If a match is found the lyrics for that song is returned, otherwise an error message.
func GetLyrics(artist, title string) (string, error) {
	requestUrl := searchRequestURL(artist, title)
	song, statusCode := makeSongRequest(requestUrl)

	switch statusCode {
	case 200:
		fmt.Printf("Showing lyrics for %s - %s\n\n", song.Result.Artist.Name, song.Result.Track.Name)
		return song.Result.Track.Text, nil
	case 404:
		return "", fmt.Errorf("No lyrics found for \"%s\" - \"%s\"", artist, title)
	default:
		return "", fmt.Errorf("Error fetching lyrics")
	}
}

func urlSafeString(s string) string {
	spacesRemoved := strings.Replace(s, " ", "%20", -1)
	return spacesRemoved
}

// makeSongRequest returns a SongResponse object and an status code
func makeSongRequest(url string) (SongResponse, int) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var body SongResponse
	if json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatalf("Couldn't convert response to struct (%s)", err)
	}

	return body, res.StatusCode
}

func searchRequestURL(artist, title string) string {
	safeArtist := urlSafeString(artist)
	safeSong := urlSafeString(title)
	return fmt.Sprintf("%s%s/%s?apikey=%s", baseURL, safeArtist, safeSong, clientApiKey)
}
