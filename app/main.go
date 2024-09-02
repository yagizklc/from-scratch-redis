package main

import (
	"github.com/yagizklc/from-scratch-redis/app/pkg"
	"log"
	"os"
)

func main() {
	log.Println("Starting server...")

	port := pkg.PORT
	if len(os.Args) == 3 && os.Args[1] == "--port" {
		port = os.Args[2]
	}

	rs := pkg.NewRedisServer(pkg.HOST, port)
	rs.Start()
}
