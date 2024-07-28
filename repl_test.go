package main

import (
	"testing"
)

func TestFormatCommandFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "HELLO",
			expected: []string{"hello"},
		},
		{
			input:    "WoRlD",
			expected: []string{"world"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}

	for _, cs := range cases {
		actual := formatCommand(cs.input)
		if actual[0] != cs.expected[0] {
			t.Errorf("command is not formatted correctly:\n actual: %v\n expected: %v\n", actual, cs.expected)
			continue
		}
	}
}
