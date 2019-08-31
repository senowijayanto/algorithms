package main

import "testing"

func TestReverse(t *testing.T) {
	if reverse("hello") != "olleh" {
		t.Error("Expected olleh from hello")
	}
}

func TestTableReverse(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"seno", "ones"},
		{"apple", "elppa"},
	}

	for _, test := range tests {
		if output := reverse(test.input); output != test.expected {
			t.Error("TestFailed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
