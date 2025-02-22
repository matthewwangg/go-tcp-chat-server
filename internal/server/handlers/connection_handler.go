package handlers

import (
	"fmt"
	"io"
	"log"
	"net"
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
		packet := buffer[:length]

		fmt.Println(username + ": " + string(packet))
	}

	return

}
