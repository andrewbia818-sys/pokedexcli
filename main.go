package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {
	// Create a cfg that has a persistent state for the REPL loop
	var cfg *pokeapi.Config
	cfg = &pokeapi.Config{
		NextPageURL: "",
		PrevPageURL: "",
		Cache:       pokecache.NewCache(20 * time.Second),
	}

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	// Start an infinite for loop that  will execute once for
	//  every command the user types in
	for {
		// Use fmt.Print to print the prompt "Pokedex >"
		//  without a newline character
		fmt.Print("Pokedex > ")
		// Read a line from standard input
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
			break
		}
		// Get the text from the scanner
		text := scanner.Text()
		// Call the cleanInput function with the text
		words := cleanInput(text)
		//Update the REPL loop to use the "command" (first word)
		if len(words) == 0 {
			continue
		}
		command, exists := commands[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
