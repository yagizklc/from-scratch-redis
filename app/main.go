package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yagizklc/from-scratch-redis/app/handlers"
	"github.com/yagizklc/from-scratch-redis/app/pkg"
)

func main() {
	log.Println("Starting server...")

	// Define command-line flags
	port := flag.String("port", pkg.PORT, "Port to run the server on")
	host := flag.String("host", pkg.HOST, "Host address to bind to")

	// Parse the flags
	flag.Parse()

	// Use the parsed values
	rs := pkg.NewRedisServer(*host, *port)

	// Register Handlers
	rs.Handle("ping", handlers.Ping)
	rs.Handle("echo", handlers.Echo)
	rs.Handle("set", handlers.Set)
	rs.Handle("get", handlers.Get)
	rs.Handle("info", handlers.Info)

	// Start server
	fmt.Printf("Server starting on %s:%s\n", *host, *port)
	rs.Start()
}
