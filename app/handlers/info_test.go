package handlers_test

import (
	"testing"

	"github.com/yagizklc/from-scratch-redis/app/handlers"
	"github.com/yagizklc/from-scratch-redis/app/pkg"
)

func TestInfo(t *testing.T) {
	tests := []struct {
		command pkg.Command
		exp     string
	}{
		{
			command: pkg.Command{Name: "INFO", Args: []string{"hello", "world"}},
			exp:     "+PONG\r\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.command.Name, func(t *testing.T) {
			got, err := handlers.Info(tc.command.Args)
			if err != nil {
				if err.Error() != tc.exp {
					t.Errorf("Ping() error = %s, exp %s", err, tc.exp)
				}
				return
			}

			if string(got) != tc.exp {
				t.Errorf("Ping() = %s, want %s", got, tc.exp)
			}
		})
	}
}
