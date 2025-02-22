package handlers

import (
	"fmt"
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
		packet := strings.TrimSpace(string(buffer[:length]))
		if packet != "" && packet[0] == '/' {
			HandleCommand(packet)
		} else {
			fmt.Println(username + ": " + packet)
		}

	}

	return

}
