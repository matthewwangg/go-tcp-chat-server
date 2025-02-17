package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
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
