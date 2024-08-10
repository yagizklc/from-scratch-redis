package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/yagizklc/from-scratch-redis/app/handlers"
)

const (
	maxRequestSize  = 1024 // Maximum request size in bytes
	maxArgumentSize = 10   // Maximum request size in integer
	HOST            = "0.0.0.0"
	PORT            = "6379"
)

func main() {
	log.Println("Starting server...")

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", HOST, PORT))
	if err != nil {
		log.Println("Failed to bind to port 6379:", err)
		os.Exit(1)
	}
	log.Println("Listening on port 6379")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			continue
		}
		log.Println("Accepted connection from", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	command, err := readRESP(conn)
	if err != nil {
		log.Println("Error reading command:", err)
	}

	commandName, arguments, err := parseCommand(command)
	if err != nil {
		log.Println("Error parsing command:", err)
	}

	log.Printf("Received command: %s with arguments: %v\n", commandName, arguments)

	var response []byte
	var respErr error
	switch commandName {
	case "PING":
		response, respErr = handlers.Ping(arguments)
	case "ECHO":
		response, respErr = handlers.Echo(arguments)
	case "SET":
		response, respErr = handlers.Set(arguments)
	case "GET":
		response, respErr = handlers.Get(arguments)
	default:
		respErr = fmt.Errorf("unsupported command: %s", commandName)
	}

	handleErr(respErr)
	log.Printf("Sending response: %s\n", response)
	conn.Write(response)
}

func handleErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
