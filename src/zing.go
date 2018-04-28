package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const baseURL = "https://api.genius.com/"

func urlSafeString(s string) string {
	spacesRemoved := strings.Replace(s, " ", "%20", -1)
	trimmed := strings.Trim(spacesRemoved, "\n")
	return trimmed
}

func searchRequestURL(query string) string {
	search := "search?q=" + urlSafeString(query)
	return baseURL + search
}

func main() {
	loadEnvironmentVariables()
	fmt.Print("Is is sing along time? I'm exited, search away: ")
	reader := bufio.NewReader(os.Stdin)
	searchString, _ := reader.ReadString('\n')

	fmt.Println(searchRequestURL(searchString))
}
