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
	maxArgumentSize = 10   // Maximum handler args size in integer
	HOST            = "0.0.0.0"
	PORT            = "6379"
)

func main() {
	log.Println("Starting server...")
	runServer()
}

func runServer() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", HOST, PORT))
	if err != nil {
		log.Println("Failed to bind to port 6379:", err)
		os.Exit(1)
	}
	log.Println("Listening on", l.Addr())
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

	for {
		command, err := parseCommand(conn)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Client closed the connection")
				return
			}
			log.Println("Error parsing command:", err)
			conn.Write([]byte("-ERR " + err.Error() + "\r\n"))
			continue
		}
		log.Printf("Received command: %s with arguments: %v\n", command.Name, command.Args)

		var response []byte
		var respErr error
		switch command.Name {
		case "PING":
			response, respErr = handlers.Ping(command.Args)
		case "ECHO":
			response, respErr = handlers.Echo(command.Args)
		case "SET":
			response, respErr = handlers.Set(command.Args)
		case "GET":
			response, respErr = handlers.Get(command.Args)
		default:
			response = []byte(fmt.Sprintf("-ERR unsupported command %s\r\n", command.Name))
			respErr = fmt.Errorf("unsupported command: %s", command.Name)
		}

		if respErr != nil {
			log.Println("Error:", respErr)
		}
		log.Printf("Sending response: %v\n", string(response))
		_, err = conn.Write(response)
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	}
}
