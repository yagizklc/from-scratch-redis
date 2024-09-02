package pkg

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Handler func([]string) ([]byte, error)

type Server interface {
	Start()
}

type RedisServer struct {
	Port     string
	Host     string
	handlers map[string]Handler
}

func NewRedisServer(host, port string) RedisServer {
	rs := RedisServer{Host: host, Port: port, handlers: make(map[string]Handler, 0)}
	return rs
}

func (rs *RedisServer) Handle(command string, handler Handler) {
	rs.handlers[strings.ToUpper(command)] = handler
}

func (rs *RedisServer) Start() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", rs.Host, rs.Port))
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

		go rs.handleConnection(conn)
	}
}

func (rs *RedisServer) handleConnection(conn net.Conn) {
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
		handler, ok := rs.handlers[command.Name]
		if !ok {
			response = []byte(fmt.Sprintf("-ERR unsupported command %s\r\n", command.Name))
		} else {
			response, err = handler(command.Args)
			if err != nil {
				log.Printf("Error: %v", err)
			}
		}

		log.Printf("Sending response: %v\n", string(response))
		_, err = conn.Write(response)
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	}
}
