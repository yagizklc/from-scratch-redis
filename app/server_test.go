package main

import (
	"strings"
	"testing"
)

func TestMultipleResponse(t *testing.T) {
	tests := []struct {
		input    string
		expected *Command
	}{
		{
			input:    "PING",
			expected: &Command{Name: "PING", Args: []string{}},
		},
		{
			input:    "echo -e PING\nPING |",
			expected: &Command{Name: "ECHO", Args: []string{"hey"}},
		},
		{
			input:    "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
			expected: &Command{Name: "SET", Args: []string{"key", "value"}},
		},
	}
	runServer()

	for _, tc := range tests {

		reader := strings.NewReader(tc.input)

		command, err := parseCommand(reader)
		if err != nil {
			t.Errorf("parseCommand(%q) returned error: %v", tc.input, err)
		}
		if command != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, command)
		}
	}
}
