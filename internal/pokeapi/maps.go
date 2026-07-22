package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Create a struct to hold the configuration for the CLI.
type Config struct {
	// Add any configuration fields you need here
	NextPageURL string
	PrevPageURL string
}

// Define a struct to represent the response from the PokeAPI for location areas.
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// should display the next 20 locations.
func CommandMap(cfg *Config) error {
	// amend the URL to use the nextPageURL if it is set
	var url string
	url = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if cfg.NextPageURL != "" {
		url = cfg.NextPageURL
	}
	// Make a GET request to the PokeAPI to retrieve location areas
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", body)

	// parse the JSON response to extract the next and previous URLs
	// and store them in the cfg struct for future calls.
	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		log.Fatal(err)
	}
	cfg.NextPageURL = locationAreaResponse.Next
	cfg.PrevPageURL = locationAreaResponse.Previous
	//fmt.Printf("Next page URL: %s\n", cfg.NextPageURL)
	//fmt.Printf("Previous page URL: %s\n", cfg.PrevPageURL)
	//fmt.Println("Location Areas:")
	for _, area := range locationAreaResponse.Results {
		fmt.Println("- " + area.Name)
	}
	return nil
}

// Add commandMapb function. It does the same as commandMap but uses the previous page URL to get the previous 20 location areas.
func CommandMapb(cfg *Config) error {
	// amend the URL to use the prevPageURL if it is set
	var url string
	url = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if cfg.PrevPageURL != "" {
		url = cfg.PrevPageURL
	}
	// Make a GET request to the PokeAPI to retrieve location areas
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	// parse the JSON response to extract the next and previous URLs
	// and store them in the cfg struct for future calls.
	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		log.Fatal(err)
	}
	cfg.NextPageURL = locationAreaResponse.Next
	cfg.PrevPageURL = locationAreaResponse.Previous
	for _, area := range locationAreaResponse.Results {
		fmt.Println("- " + area.Name)
	}
	return nil
}
