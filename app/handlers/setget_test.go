package handlers

import (
	"strconv"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{"key", "value"},
			expected: RespBulkStringEncode("value"),
		},
		{
			args:     []string{"key", "value2"},
			expected: RespBulkStringEncode("value2"),
		},
		{
			args:     []string{"key", "value3", "px", "100"},
			expected: RespSimpleError("ERR key not found"),
		},
		{
			args:     []string{"key", "value4", "px", "1000"},
			expected: RespBulkStringEncode("value4"),
		},
	}
	for _, tc := range tests {
		_, err := Set(tc.args)
		if err != nil {
			t.Errorf("Set(%q) returned error: %v", tc.args, err)
		}
		if len(tc.args) == 4 {
			v, err := strconv.Atoi(tc.args[3])
			if err != nil {
				t.Errorf("Set(%q) returned error: %v", tc.args, err)
			}
			if v < 1000 {
				time.Sleep(time.Millisecond * time.Duration(v))
			}
		}

		got, err := Get(tc.args)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", tc.args, err)
		}
		if string(got) != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, string(got))
		}

	}
}
