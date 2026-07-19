package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
		{
		input:    "  HeLLo  worLD  ",
		expected: []string{"hello", "world"},
	}
		{
		input:    "     hello,world  ",
		expected: []string{"hello,world"},
	},
	}
}
// auto below
	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if !equal(actual, tc.expected) {
			t.Errorf("cleanInput(%q) = %v, want %v", tc.input, actual, tc.expected)
		}
	}
// above auto
	for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice
	// if they don't match, use t.Errorf and continue to the next case
	if len(actual) != len(c.expected) {
		t.Errorf("cleanInput(%q) = %v, want %v", c.input, actual, c.expected)
		// error and continue here
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("cleanInput(%q) = %v, want %v", c.input, actual, c.expected)
		}
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}	
}
}