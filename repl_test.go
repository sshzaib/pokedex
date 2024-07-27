package main

import (
	"testing"
)

func TestFormatCommandFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "HELLO",
			expected: "hello",
		},
		{
			input:    "WoRlD",
			expected: "world",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, cs := range cases {
		actual := formatCommand(cs.input)
		if actual != cs.expected {
			t.Errorf("command is not formatted correctly:\n actual: %v\n expected: %v\n", actual, cs.expected)
			continue
		}
	}
}
