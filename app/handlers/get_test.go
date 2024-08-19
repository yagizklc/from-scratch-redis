package handlers

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{"key", "value"},
			expected: "value",
		},
		{
			args:     []string{"key", "value2"},
			expected: "value2",
		},
		{
			args:     []string{"nonkey", "value2"},
			expected: "",
		},
	}
	for _, tc := range tests {
		if tc.args[0] != "nonkey" {
			_, err := Set(tc.args)
			if err != nil {
				t.Errorf("Set(%q) returned error: %v", tc.args, err)
			}
		}
		got, err := Get(tc.args)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", tc.args, err)
		}
		log.Printf("got: %v", string(got))
		exp := string(RespBulkStringEncode(tc.expected))
		if string(got) != exp {
			t.Errorf("expected %v, got %v", exp, string(got))
		}
	}
}
