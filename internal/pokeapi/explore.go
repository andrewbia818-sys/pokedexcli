package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// A struct to hold the detail endpoint.
type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// CommandExplore displays the Pokemons present in the location area
// It takes as input a location area name. And prints all the Pokemons
// present in that location area to Stdout
// NEW VERSION
func CommandExplore(cfg *Config, areaName string) error {
	// Use the areaName passed from the REPL
	if areaName == "" {
		return fmt.Errorf("location area name cannot be empty")
	}

	detailURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)

	cache := cfg.Cache

	// Check cache
	if val, ok := cache.Get(detailURL); ok {
		//        fmt.Println("Cache being used for area detail")   // correct placement
		var detail LocationAreaDetail
		if err := json.Unmarshal(val, &detail); err != nil {
			return err
		}

		for _, encounter := range detail.PokemonEncounters {
			fmt.Println("- " + encounter.Pokemon.Name)
		}
		return nil
	}

	// Fetch from API
	res, err := http.Get(detailURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(detailURL, body)

	var detail LocationAreaDetail
	if err := json.Unmarshal(body, &detail); err != nil {
		return err
	}

	for _, encounter := range detail.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
