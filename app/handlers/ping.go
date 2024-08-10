package handlers

import (
	"fmt"
	"log"
)

func Ping(arguments []string) ([]byte, error) {
	lenArgs := len(arguments)
	if lenArgs > 1 || lenArgs < 0 {
		return nil, fmt.Errorf("too many arguments for PING command: expected 1 or 0, got %d", lenArgs)
	}

	if lenArgs == 0 {
		log.Println("Handling PING command with no arguments")
		return []byte(RespSimpleStringEncode("PONG")), nil
	}

	log.Printf("Handling PING command with argument: %s", arguments[0])
	return []byte(RespSimpleStringEncode(arguments[0])), nil
}
