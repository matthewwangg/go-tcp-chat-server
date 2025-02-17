package main

import (
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	_, err = connection.Write([]byte("Hello from the client!"))
	if err != nil {
		log.Fatal(err)
	}

	return
}
