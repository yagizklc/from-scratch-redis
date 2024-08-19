package handlers

import "testing"

func TestSet(t *testing.T) {
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
	}
	for _, tc := range tests {
		_, err := Set(tc.args)
		if err != nil {
			t.Errorf("Set(%q) returned error: %v", tc.args, err)
		}
		if db[tc.args[0]] != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, db[tc.args[0]])
		}
	}
}
