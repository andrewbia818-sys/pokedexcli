package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
