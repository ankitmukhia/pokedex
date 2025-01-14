package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  ankit mukhia   ",
			expected: []string{"ankit", "mukhia"},
		},
	}

	for _, c := range testCases {
		actual := cleanInput(c.input)

		if len(actual) !=  len(c.expected) {
			t.Errorf("Lenght is not matched, actual: %v, expected: %v", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("")
			}
		}
	}
}
