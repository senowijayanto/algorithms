package main

import "testing"

func TestPalindrome(t *testing.T) {
	if palindrome("abba") != true {
		t.Error("Expected true from abba")
	}
}

func TestTablePalindrome(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"abba", true},
		{"abcdefg", false},
	}

	for _, test := range tests {
		if output := palindrome(test.input); output != test.expected {
			t.Error("TestFailed: {} inputtes, {} expected, received {}", test.input, test.input, output)
		}
	}
}
