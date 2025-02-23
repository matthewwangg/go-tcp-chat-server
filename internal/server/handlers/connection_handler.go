package handlers

import (
	"net"
)

func HandleConnection(connection net.Conn) {
	defer connection.Close()
	incoming := make(chan string)
	outgoing := make(chan string)

	username := HandleLogin(connection)

	go ReadFromClient(connection, incoming)
	go WriteToClient(connection, outgoing)

	for message := range incoming {
		if message != "" && message[0] == '/' {
			HandleCommand(message, username)
		} else {
			HandleMessage(outgoing, message)
		}
	}
	close(outgoing)
	return
}
