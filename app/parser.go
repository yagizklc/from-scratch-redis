package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

// reads RESP serialized command from connection, returns parts of the command
//
// PING = *1\r\n$4\r\nPING\r\n
//
// ECHO = *2\r\n$4\r\nECHO\r\n$3\r\n
func parseCommand(conn io.Reader) (*Command, error) {
	scanner := bufio.NewScanner(conn)
	command := make([]string, 0, maxArgumentSize)

	// read the first line, which is the length of the command
	if !scanner.Scan() {
		return nil, fmt.Errorf("EOF")
	}

	t := scanner.Text()
	if t[0] != '*' {
		return nil, fmt.Errorf("expected array")
	}

	numArgs, err := strconv.Atoi(t[1:])
	if err != nil {
		return nil, err
	}

	// take the each second argument (skipping the length of the argument)
	for i := 0; i < numArgs*2 && scanner.Scan(); i++ {
		if i%2 == 1 {
			command = append(command, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from connection:", err)
		return nil, err
	}

	log.Printf("Read command: %v", command)
	return &Command{
		Name: strings.ToUpper(command[0]),
		Args: command[1:],
	}, nil
}
