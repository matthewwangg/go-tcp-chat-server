package main

import (
	"github.com/matthewwangg/go-tcp-server/internal/client/handlers"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	_, err = connection.Write([]byte("Hello from the client!"))
	if err != nil {
		log.Fatal(err)
	}

	handlers.HandleLogin(connection)
	handlers.HandleInput(connection)

	return
}
