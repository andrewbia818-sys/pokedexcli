package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

// cleanInput function takes a string input, returns a slice of strings
// removes leading & trailing whitespace, does toLower()
// & splits into words based on whitespace between them.
var words []string

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words = strings.Fields(text)
	return words
}

// Add an exit command function.
func commandExit(cfg *pokeapi.Config) error {
	if len(words) > 0 && words[0] == "exit" {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
	}
	return nil
}

// Add a help command function.
func commandHelp(cfg *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Display 20 location areas in the Pokemon world")
	fmt.Println("mapb: Display the previous 20 location areas in the Pokemon world")
	return nil
}

// Define a struct to represent all commands in the CLI.
type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

// Create a map to hold all the commands in the CLI.
var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display a list of all commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Display 20 location areas in the Pokemon world",
		callback:    pokeapi.CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Display the previous 20 location areas in the Pokemon world",
		callback:    pokeapi.CommandMapb,
	},
}
