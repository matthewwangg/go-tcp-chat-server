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
	}
}
