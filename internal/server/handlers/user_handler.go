package handlers

import (
	"bufio"
	"log"
	"net"
	"sync"
)

var Users = make(map[string]net.Conn)
var Mu sync.Mutex

func HandleLogin(connection net.Conn) {
	buffer := make([]byte, 1024)

	reader := bufio.NewReader(connection)
	length, err := reader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	username := string(buffer[:length])
	if _, exists := Users[username]; !exists {
		Mu.Lock()
		Users[username] = connection
		Mu.Unlock()
		_, err = connection.Write([]byte("Connected!"))
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}
