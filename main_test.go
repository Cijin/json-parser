package main

import (
	"os"
	"testing"
)

func TestParseJson(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "tests/step1/valid.json", expected: true},
		{input: "tests/step1/invalid.json", expected: false},
		{input: "tests/step2/valid.json", expected: true},
		{input: "tests/step2/invalid.json", expected: false},
		{input: "tests/step2/valid2.json", expected: true},
		{input: "tests/step2/invalid2.json", expected: false},
	}

	for i, test := range tests {
		data, err := os.ReadFile(test.input)
		if err != nil {
			t.Fatalf("test %d: unable to read file %s: %v", i, test.input, err)
		}

		got := isValidJson(data)
		if got != test.expected {
			t.Fatalf("test %d: expected %s to be %t, got=%t", i, test.input, test.expected, got)
		}
	}
}
