package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
func commandExit(cfg *Config) error {
	if len(words) > 0 && words[0] == "exit" {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
	}
	return nil
}

// Add a help command function.
func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Display 20 location areas in the Pokemon world")
	return nil
}

// commandMap command uses the PokeAPI to GET
// and display the names of 20 location areas in the
// Pokemon world. Each subsequent call to the map function
//
//	should display the next 20 locations.
func commandMap(cfg *Config) error {
	// amend the URL to use the nextPageURL if it is set
	var url string
	url = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if cfg.nextPageURL != "" {
		url = cfg.nextPageURL
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
	cfg.nextPageURL = locationAreaResponse.Next
	cfg.prevPageURL = locationAreaResponse.Previous
	//fmt.Printf("Next page URL: %s\n", cfg.nextPageURL)
	//fmt.Printf("Previous page URL: %s\n", cfg.prevPageURL)
	//fmt.Println("Location Areas:")
	for _, area := range locationAreaResponse.Results {
		fmt.Println("- " + area.Name)
	}
	return nil
}

// TODO: Implement pagination logic to handle next and
// previous pages using cfg.nextPageURL and cfg.prevPageURL

// Define a struct to represent all commands in the CLI.
type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

// Create a struct to hold the configuration for the CLI.
type Config struct {
	// Add any configuration fields you need here
	nextPageURL string
	prevPageURL string
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
		callback:    commandMap,
	},
}
