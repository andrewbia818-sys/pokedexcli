package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

// Define a struct PokemonDetail to hold the name and base experience
//
//	of a Pokemon
type Base struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}

//type Types []struct {
//	Type []string `json:"typeIds"`
//}

type Profile struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type PokemonDetail struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Profile
	Base
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

// Create a map[string]PokemonDetail to hold the Pokemon names of
// Pokemon that were caught in the current session. This will be used
// to check if a Pokemon has already been caught.
var pokemonDetailsCache = make(map[string]PokemonDetail)

func CommandCatch(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("pokemon name cannot be empty")
	}

	detailURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	cache := cfg.Cache

	// Check cache
	if val, ok := cache.Get(detailURL); ok {
		//	fmt.Println("Cache being used for pokemon detail")

		var detail PokemonDetail
		if err := json.Unmarshal(val, &detail); err != nil {
			return err
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

		catchChance := 35 / float64(detail.BaseExperience)
		if rand.Float64() < catchChance {
			// add the Pokemon to the pokemonDetailsCache
			pokemonDetailsCache[pokemonName] = detail
			fmt.Printf("%s was caught!\n", pokemonName)
		} else {
			fmt.Printf("%s escaped!\n", pokemonName)
		}
		//		fmt.Printf("Pokedex contains %d Pokemon(s):", len(pokemonDetailsCache))
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

	var detail PokemonDetail
	if err := json.Unmarshal(body, &detail); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchChance := 35 / float64(detail.BaseExperience)
	if rand.Float64() < catchChance {
		//		fmt.Printf("%s was caught!\n", pokemonName)
		pokemonDetailsCache[pokemonName] = detail
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	//	fmt.Printf("Pokedex contains %d Pokemon(s):", len(pokemonDetailsCache))
	return nil
}
