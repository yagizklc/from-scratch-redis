package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

/*


* = arrays
$ = bulk strings

*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n

array of 2 items
4 digit string
ECHO
3 digit string
nhey


*/

func parseCommand(command []string) (string, []string, error) {

	// Split the command by whitespace
	lenParts := len(command)

	// empty command
	if lenParts < 1 {
		return "", []string{}, fmt.Errorf("empty command")
	}

	// The first part is the command name
	commandName := strings.ToUpper(command[0])

	// check if an allowed command
	if !validateCommand(commandName) {
		return "", []string{}, fmt.Errorf("not allowed command name: %s", commandName)
	}

	// parse arguments
	arguments := make([]string, 0, maxArgumentSize)
	for _, part := range command[1:] {
		arguments = append(arguments, string(part))
	}

	log.Printf("Parsed command: %s with arguments: %v", commandName, arguments)
	return commandName, arguments, nil
}

func validateCommand(command string) bool {
	validCommands := map[string]string{"PING": "", "ECHO": "", "SET": "", "GET": ""}
	_, ok := validCommands[command]
	return ok
}

func readRESP(conn net.Conn) ([]string, error) {
	scanner := bufio.NewScanner(conn)
	command := make([]string, 0, maxArgumentSize)

	// get number of arguments
	var numArgs int
	if scanner.Scan() {
		t := scanner.Text()
		i, err := strconv.Atoi(t[1:])
		if err != nil {
			return []string{}, err
		}
		numArgs = i
	}

	// take the each second argument (first is len)
	for i := 0; i < numArgs*2 && scanner.Scan(); i++ {
		if i%2 == 1 {
			command = append(command, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from connection:", err)
		return []string{}, nil
	}

	return command, nil
}
