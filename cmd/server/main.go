package main

import (
	"fmt"
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

	fmt.Println("Server is started and listening on port 8080")
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlers.HandleConnection(connection)
	}

	return
}
