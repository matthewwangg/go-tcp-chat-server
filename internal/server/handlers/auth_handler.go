package handlers

import (
	"bufio"
	"log"
	"net"
	"strings"
	"sync"
)

var Users = make(map[string]net.Conn)
var MuUsers sync.RWMutex

func HandleLogin(connection net.Conn) string {
	buffer := make([]byte, 1024)
	username := ""

	reader := bufio.NewReader(connection)
	for {
		length, err := reader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		username = strings.TrimSpace(string(buffer[:length]))
		MuUsers.Lock()
		if _, exists := Users[username]; !exists && username != "" {
			Users[username] = connection
			MuUsers.Unlock()
			_, err = connection.Write([]byte("Connected!"))
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		MuUsers.Unlock()

		_, err = connection.Write([]byte("This user is already connected!"))
		if err != nil {
			log.Fatal(err)
		}
	}

	return username
}
