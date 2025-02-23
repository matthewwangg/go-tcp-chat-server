package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func HandleLogin(connection net.Conn) string {
	buffer := make([]byte, 1024)
	stdinReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(connection)
	input := ""

	for {
		fmt.Print("Username: ")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err = connection.Write([]byte(input))
		if err != nil {
			log.Fatal(err)
		}

		length, err := connReader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		response := string(buffer[:length])
		if string(response) == "Connected!" {
			break
		}
		fmt.Print("\033[A\033[K")

	}

	return input
}
