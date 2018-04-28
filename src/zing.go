package main

import (
	"bufio"
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

func urlSafeString(s string) string {
	spacesRemoved := strings.Replace(s, " ", "%20", -1)
	trimmed := strings.Trim(spacesRemoved, "\n")
	return trimmed
}

func makeRequest(url string) SearchResponse {
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

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientAccessToken = os.Getenv("CLIENT_ACCESS_TOKEN")
}

func main() {
	loadEnvironmentVariables()
	fmt.Print("Is is sing along time? I'm exited, search away: ")
	reader := bufio.NewReader(os.Stdin)
	searchString, _ := reader.ReadString('\n')

	requestURL := searchRequestURL(searchString)
	res := makeRequest(requestURL).Response

	if len(res.Hits) > 0 {
		for i, hit := range res.Hits {
			fmt.Printf("%2d. %s\n", i+1, hit.Result.FullTitle)
		}
	} else {
		fmt.Println("No results :(")
	}
}
