package pokeapi

import (
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
)

// Test CommandMap, checking for a valid return value.
func TestCommandMap(t *testing.T) {
	err := CommandMap(&Config{
		NextPageURL: "",
		PrevPageURL: "",
		Cache:       pokecache.NewCache(20 * time.Second),
	})
	if err != nil {
		t.Fatalf("CommandMap() returned error: %v", err)
	}
}

// Test CommandMapb, checking for a valid return value.
func TestCommandMapb(t *testing.T) {
	err := CommandMapb(&Config{
		NextPageURL: "",
		PrevPageURL: "",
		Cache:       pokecache.NewCache(20 * time.Second),
	})
	if err != nil {
		t.Fatalf("CommandMapb() returned error: %v", err)
	}
}
