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
		UpdateLastSeen(username)
		if message != "" && message[0] == '/' {
			HandleCommand(message, username, outgoing)
		} else {
			HandleMessage(message, username, outgoing)
		}
	}

	RemoveUser(username)
	close(outgoing)
	return
}
