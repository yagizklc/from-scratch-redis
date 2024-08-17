package main

import (
	"strings"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected *Command
	}{
		{
			input:    "*1\r\n$4\r\nPING\r\n",
			expected: &Command{Name: "PING", Args: []string{}},
		},
		{
			input:    "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n",
			expected: &Command{Name: "ECHO", Args: []string{"hey"}},
		},
		{
			input:    "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
			expected: &Command{Name: "SET", Args: []string{"key", "value"}},
		},
	}

	for _, tc := range tests {

		reader := strings.NewReader(tc.input)
		command, err := parseCommand(reader)
		if err != nil {
			t.Errorf("parseRESP(%q) returned error: %v", tc.input, err)
		}
		if command.Name != tc.expected.Name {
			t.Errorf("expected %v, got %v", tc.expected.Name, command.Name)
		}
		if len(command.Args) != len(tc.expected.Args) {
			t.Errorf("expected %v, got %v", tc.expected.Args, command.Args)
		}
		for i := range command.Args {
			if command.Args[i] != tc.expected.Args[i] {
				t.Errorf("expected %v, got %v", tc.expected.Args[i], command.Args[i])
			}
		}

	}
}
