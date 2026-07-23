package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
)

// Create a struct to hold the configuration for the CLI.
type Config struct {
	// Add any configuration fields you need here
	NextPageURL string
	PrevPageURL string
	Cache       *pokecache.Cache
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
	// Add caching to the CommandMap function.
	var cache *pokecache.Cache
	//	cache = pokecache.NewCache(5 * time.Second)
	cache = cfg.Cache
	// check if the response for the URL is already in the cache
	if val, ok := cache.Get(url); ok {
		// if it is, use that instead of making a new request
		fmt.Println("Using cached response for URL:", url)
		var locationAreaResponse LocationAreaResponse
		err := json.Unmarshal(val, &locationAreaResponse)
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
	// Add the response to the cache
	cache.Add(url, body)
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

// Add commandMapb function. It does the same as commandMap but uses the previous page URL to get the previous 20 location areas.
func CommandMapb(cfg *Config) error {
	// amend the URL to use the prevPageURL if it is set
	var url string
	url = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if cfg.PrevPageURL != "" {
		url = cfg.PrevPageURL
	}
	var cache *pokecache.Cache
	cache = cfg.Cache
	// check if the response for the URL is already in the cache
	if val, ok := cache.Get(url); ok {
		// if it is, use that instead of making a new request
		//	fmt.Println("Using cached response for URL:", url)
		var locationAreaResponse LocationAreaResponse
		err := json.Unmarshal(val, &locationAreaResponse)
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
	// Add the response to the cache
	cache.Add(url, body)
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
