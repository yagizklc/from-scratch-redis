package pkg

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMultipleResponse(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected string
	}{
		{
			name:     "Single Redis Command",
			command:  "redis-cli PING",
			expected: "PONG",
		},
		{
			name:     "Piped Commands",
			command:  "echo -e 'PING\nPING' | redis-cli",
			expected: "PONG\nPONG",
		},
	}

	rs := NewRedisServer(HOST, PORT)
	go rs.Start()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("sh", "-c", tc.command)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command execution failed: %v", err)
			}

			result := strings.TrimSpace(string(output))
			if result != tc.expected {
				t.Errorf("Expected output %q, but got %q", tc.expected, result)
			}
		})
	}
}