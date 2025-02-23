package handlers

import (
	"net"
	"strings"
)

func ReadFromClient(connection net.Conn, incoming chan<- string) {
	buffer := make([]byte, 1024)
	for {
		length, err := connection.Read(buffer)
		if err != nil {
			close(incoming)
			return
		}
		incoming <- strings.TrimSpace(string(buffer[:length]))
	}
}

func WriteToClient(connection net.Conn, outgoing <-chan string) {
	for message := range outgoing {
		_, err := connection.Write([]byte(message + "\n"))
		if err != nil {
			return
		}
	}
}

func HandleMessage(outgoing chan<- string, message string) {
	outgoing <- "You: " + message
}
