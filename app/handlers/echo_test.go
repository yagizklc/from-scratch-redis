package handlers

import "testing"

func TestEcho(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "no arguments",
			args: []string{},
			want: "too few arguments for ECHO command: expected 1, got 0",
		},
		{
			name: "one argument",
			args: []string{"hey"},
			want: "$3\r\nhey\r\n",
		},
		{
			name: "too many arguments",
			args: []string{"arg1", "arg2"},
			want: "too many arguments for ECHO command: expected 1, got 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Echo(tt.args)
			if err != nil {
				if err.Error() != tt.want {
					t.Errorf("Echo() error = %s, want %s", err, tt.want)
				}
				return
			}
			if string(got) != tt.want {
				t.Errorf("Echo() = %s, want %s", got, tt.want)
			}
		})
	}
}
