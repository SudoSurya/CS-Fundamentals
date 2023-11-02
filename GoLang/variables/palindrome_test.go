package main

import "testing"

func TestIsPalindromes(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
		{"A Santa at NASA", true},
		{"", true}, // An empty string is considered a palindrome
		{"1 2 3 2 1", true},
		{"!@#@!A man a plan a canal Panama", true}, // Special characters and spaces should be ignored
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isPalindromes(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}

