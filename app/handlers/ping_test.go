package handlers

import "testing"

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "no arguments",
			args: []string{},
			want: "+PONG\r\n",
		},
		{
			name: "one argument",
			args: []string{"hello"},
			want: "+hello\r\n",
		},
		{
			name: "too many arguments",
			args: []string{"arg1", "arg2"},
			want: "Too much arguments: expected 1 or 0, got 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ping(tt.args)
			if err != nil {
				if err.Error() != tt.want {
					t.Errorf("Ping() error = %s, want %s", err, tt.want)
				}
				return
			}

			if string(got) != tt.want {
				t.Errorf("Ping() = %s, want %s", got, tt.want)
			}
		})
	}
}
