package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleInput(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		fmt.Print("Input: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(input) == "/exit" {
			break
		}
		_, err = connection.Write([]byte(input))
		if err != nil {
			log.Fatal(err)
		}

		if len(input) > 1 && input[0] != '/' {
			length, err := connection.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}
			message := strings.TrimSpace(string(buffer[:length]))
			fmt.Println(message)
		}
	}
}
