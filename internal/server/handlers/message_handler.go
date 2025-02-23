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

func HandleMessage(message string, user string, outgoing chan<- string) {
	MuUserRoom.RLock()
	currentRoom, exists := UserRoom[user]
	MuUserRoom.RUnlock()
	if exists {
		MuRooms.RLock()
		roomStruct, exists := Rooms[currentRoom]
		MuRooms.RUnlock()
		if exists {
			roomStruct.Ch <- user + ": " + message
		}
	} else {
		outgoing <- "You: " + message
	}
}
