package main

import (
	"os/exec"
	"testing"
)

func TestMultipleResponse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "redis-cli PING",
			expected: "PING",
		},
		{
			input:    "echo -e PING\nPING | redis-cli",
			expected: "PING\nPING",
		},
	}
	runServer()

	for _, tc := range tests {
		if err := exec.Command(tc.input).Run(); err != nil {
			t.Errorf("exec.Command(%q) returned error: %v", tc.input, err)
		}

	}
}
