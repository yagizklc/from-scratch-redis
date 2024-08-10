package handlers

// import (
// 	"bytes"
// 	"os"
// 	"testing"
// )

// func TestGet(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		arguments []string
// 		want      []byte
// 		err       error
// 	}{
// 		{
// 			name:      "valid input",
// 			arguments: []string{"foo"},
// 			want:      []byte("$3\r\nbar\r\n"),
// 			err:       nil,
// 		},
// 		{
// 			name:      "invalid input",
// 			arguments: []string{"invalid"},
// 			want:      nil,
// 			err:       nil, // or set an expected error if needed
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a temporary file for the database
// 			tempFile := createTempFile(t, "foo bar\r\n")
// 			defer os.Remove(tempFile)

// 			got, err := Get(tt.arguments)
// 			if err != tt.err {
// 				t.Errorf("Get(%v) error = %v, want %v", tt.arguments, err, tt.err)
// 				return
// 			}
// 			if !bytes.Equal(got, tt.want) {
// 				t.Errorf("Get(%v) = %v, want %v", tt.arguments, got, tt.want)
// 			}
// 		})
// 	}
// }

// func createTempFile(t *testing.T, content string) string {
// 	t.Helper()

// 	tempFile, err := os.CreateTemp("", "test_db_*.txt")
// 	if err != nil {
// 		t.Fatalf("failed to create temp file: %v", err)
// 	}
// 	t.Cleanup(func() {
// 		err := os.Remove(tempFile.Name())
// 		if err != nil {
// 			t.Errorf("failed to remove temp file: %v", err)
// 		}
// 	})

// 	_, err = tempFile.WriteString(content)
// 	if err != nil {
// 		t.Fatalf("failed to write temp file: %v", err)
// 	}
// 	err = tempFile.Close()
// 	if err != nil {
// 		t.Fatalf("failed to close temp file: %v", err)
// 	}

// 	return tempFile.Name()
// }
