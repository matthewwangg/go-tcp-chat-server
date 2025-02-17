package main

import (
	"github.com/matthewwangg/go-tcp-server/internal/server/handlers"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlers.HandleConnection(connection)
	}

	return
}
