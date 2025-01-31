package handlers

import (
	"fmt"
	"github.com/yagizklc/from-scratch-redis/app/pkg"
)

func Echo(arguments []string) ([]byte, error) {
	lenArgs := len(arguments)
	if lenArgs == 0 {
		return nil, fmt.Errorf("too few arguments for ECHO command: expected 1, got %d", lenArgs)
	}
	if lenArgs > 1 {
		return nil, fmt.Errorf("too many arguments for ECHO command: expected 1, got %d", lenArgs)
	}

	return []byte(pkg.RespBulkStringEncode(arguments[0])), nil
}
