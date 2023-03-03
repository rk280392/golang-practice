package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	username = "rajeshk"
	password = "Wifi@123"
)

type Words struct {
	Page  string   `json:"page"`
	Words []string `json:"words"`
	Input string   `json:"input"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("URL is invalid: %s", err)
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", args[1], nil)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(1)
		os.Exit(1)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP CODE %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed Json Values\n Page: %s\nWords : %v\n", words.Page, strings.Join(words.Words, ", "))

}
