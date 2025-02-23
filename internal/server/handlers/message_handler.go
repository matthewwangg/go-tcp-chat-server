package handlers

import (
	"log"
	"net"
)

func HandleMessage(connection net.Conn, message string) {
	_, err := connection.Write([]byte("You: " + message))
	if err != nil {
		log.Fatal(err)
	}

	return
}
