package main

import (
	"bufio"
	"fmt"
	"log"
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

	fmt.Println(searchRequestURL(searchString))
}
