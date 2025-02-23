package handlers

import (
	"io"
	"log"
	"net"
	"strings"
)

func HandleConnection(connection net.Conn) {

	defer connection.Close()
	username := HandleLogin(connection)
	buffer := make([]byte, 1024)

	for {
		length, err := connection.Read(buffer)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		message := strings.TrimSpace(string(buffer[:length]))
		if message != "" && message[0] == '/' {
			HandleCommand(message, username)
		} else {
			HandleMessage(connection, message)
		}

	}

	return

}
