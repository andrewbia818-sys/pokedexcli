package pokeapi

import (
	"fmt"
)

// Add CommandPokedex to list all the Pokemon names in the cache.
// Each Pokemon should be printed on a new line.

func CommandPokedex(cfg *Config, pokemonName string) error {
	// Add the logic to display all pokemon names in the cahche
	// loop through all the Pokemon in the cache and print out the names
	if len(pokemonDetailsCache) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokémon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range pokemonDetailsCache {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
