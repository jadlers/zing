package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var clientAccessToken string

const baseURL = "https://api.genius.com/"

func urlSafeString(s string) string {
	spacesRemoved := strings.Replace(s, " ", "%20", -1)
	return spacesRemoved
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

func getCliInput() string {
	reader := bufio.NewReader(os.Stdin)
	searchString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(searchString, "\n")
}

func makeChoice(numResults int) int {
	fmt.Printf("\nChoose a song to get more info: ")
	var intChoice int
	for intChoice == 0 {
		choice := getCliInput()
		intVer, _ := strconv.ParseInt(choice, 10, 0)
		intChoice = int(intVer)

		if intChoice < 1 || intChoice > numResults {
			intChoice = 0
			fmt.Printf("\rNot a valid number, try again: ")
		}
	}

	return intChoice - 1
}

func main() {
	loadEnvironmentVariables()

	// Make a search
	fmt.Print("Is is sing along time? I'm exited, search away: ")
	searchString := getCliInput()

	requestURL := searchRequestURL(searchString)
	res := makeRequest(requestURL).Response

	if len(res.Hits) > 0 {
		for i, hit := range res.Hits {
			fmt.Printf("%2d. %s\n", i+1, hit.Result.FullTitle)
		}
	} else {
		fmt.Println("No results :(")
	}

	// Dvelve deeper on result
	choice := makeChoice(len(res.Hits))
	fmt.Printf("You chose: %d\n", choice)
}
