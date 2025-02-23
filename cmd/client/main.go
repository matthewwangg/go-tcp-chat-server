package main

import (
	"fmt"
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

	outgoing := make(chan string)
	incoming := make(chan string)

	handlers.HandleLogin(connection)

	go handlers.ListenForMessages(connection, incoming)
	go handlers.SendMessages(connection, outgoing)

	go handlers.HandleInput(outgoing)

	for message := range incoming {
		fmt.Println(message)
	}
	close(outgoing)
	return
}
