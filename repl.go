package main

import (
	"strings"
)

// need to implement the cleanInput function that takes
// a string input and returns a slice of strings
// The function should clean the input by removing leading
// and trailing whitespace, converting the string to lowercase,
// and splitting it into words based on whitespace between them.
func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
