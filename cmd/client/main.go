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

	username := handlers.HandleLogin(connection)

	go handlers.ListenForMessages(connection, incoming)
	go handlers.SendMessages(connection, outgoing)

	go handlers.HandleInput(connection, outgoing, username)

	for message := range incoming {
		fmt.Print("\r\033[K")
		fmt.Println(message)
		fmt.Print("Input: ")
	}
	close(outgoing)
	return
}
