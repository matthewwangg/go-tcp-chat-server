package handlers

import (
	"net"
)

func ListenForMessages(connection net.Conn, incoming chan<- string) {
	buffer := make([]byte, 1024)
	for {
		length, err := connection.Read(buffer)
		if err != nil {
			close(incoming)
			return
		}
		incoming <- string(buffer[:length])
	}
}

func SendMessages(connection net.Conn, outgoing <-chan string) {
	for message := range outgoing {
		_, err := connection.Write([]byte(message + "\n"))
		if err != nil {
			return
		}
	}
}
