package handlers

import (
	"fmt"
	"github.com/yagizklc/from-scratch-redis/app/pkg"
	"log"
)

func Ping(arguments []string) ([]byte, error) {
	lenArgs := len(arguments)
	if lenArgs > 1 {
		return nil, fmt.Errorf("too many arguments for PING command: expected 1 or 0, got %d", lenArgs)
	}

	if lenArgs == 0 {
		log.Println("Handling PING command with no arguments")
		return []byte(pkg.RespSimpleStringEncode("PONG")), nil
	}

	log.Printf("Handling PING command with argument: %s", arguments[0])
	return []byte(pkg.RespSimpleStringEncode(arguments[0])), nil
}
