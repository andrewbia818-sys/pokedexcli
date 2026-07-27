package pokeapi

import (
	"encoding/json"
	"fmt"
)

// Add a CommandInspect function. It inspects a Pokemon by name and
// If the Pokemon is in the cache it prints the details of the Pokemon
// If the Pokemon is not in the cache it displays a message
// "You have not caught [pokemonName] yet. Try catching it first!" and returns an error.
func CommandInspect(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("pokemon name cannot be empty")
	}

	cache := cfg.Cache

	// Check cache
	if val, ok := cache.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)); ok {
		var detail PokemonDetail
		if err := json.Unmarshal(val, &detail); err != nil {
			return err
		}

		fmt.Printf("Name: %s\n", detail.Name)
		fmt.Printf("Height: %d\n", detail.Profile.Height)
		fmt.Printf("Weight: %d\n", detail.Profile.Weight)
		fmt.Printf("Stats:\n")
		fmt.Printf("  - hp: %d\n", detail.Base.HP)
		fmt.Printf("  - attack: %d\n", detail.Base.Attack)
		fmt.Printf("  - defense: %d\n", detail.Base.Defense)
		fmt.Printf("  - special-attack: %d\n", detail.Base.SpAttack)
		fmt.Printf("  - special-defense: %d\n", detail.Base.SpDefense)
		fmt.Printf("  - speed: %d\n", detail.Base.Speed)
		fmt.Printf("Types:\n")
		for _, t := range detail.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		//	fmt.Printf("Base Experience: %d\n", detail.BaseExperience)
		return nil
	}

	fmt.Printf("You have not caught that pokemonn")
	//	return fmt.Errorf("pokemon %s not found in cache", pokemonName)
	return nil
}
