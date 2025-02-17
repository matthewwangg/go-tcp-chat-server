package handlers

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(connection net.Conn) {

	defer connection.Close()

	buffer := make([]byte, 1024)
	length, err := connection.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	packet := buffer[:length]

	fmt.Println(string(packet))

	return

}
