package main

import (
	"io"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"
)

// Mock net.Conn implementation for testing
type mockConn struct {
	reader io.Reader
}

func (m mockConn) Read(b []byte) (int, error) {
	return m.reader.Read(b)
}

func (m mockConn) Write(b []byte) (int, error) {
	return len(b), nil // Dummy implementation
}

func (m mockConn) Close() error {
	return nil // Dummy implementation
}

func (m mockConn) LocalAddr() net.Addr {
	return nil // Dummy implementation
}

func (m mockConn) RemoteAddr() net.Addr {
	return nil // Dummy implementation
}

func (m mockConn) SetDeadline(t time.Time) error {
	return nil // Dummy implementation
}

func (m mockConn) SetReadDeadline(t time.Time) error {
	return nil // Dummy implementation
}

func (m mockConn) SetWriteDeadline(t time.Time) error {
	return nil // Dummy implementation
}

func TestReadRESP(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "*1\r\n$4\r\nPING\r\n",
			expected: []string{"PING"},
		},
		{
			input:    "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n",
			expected: []string{"ECHO", "hey"},
		},
		{
			input:    "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
			expected: []string{"SET", "key", "value"},
		},
	}

	for _, test := range tests {

		mockReader := strings.NewReader(test.input)
		mockConn := mockConn{reader: mockReader}

		result, err := readRESP(mockConn)
		if err != nil {
			t.Errorf("parseRESP(%q) returned error: %v", test.input, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseRESP(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
