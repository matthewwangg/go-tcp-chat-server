package handlers

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func ListenForMessages(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		message := strings.TrimSpace(string(buffer[:length]))

		fmt.Println(message)
	}

	return
}
