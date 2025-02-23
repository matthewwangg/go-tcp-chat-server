package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleInput(connection net.Conn, outgoing chan<- string, user string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("\033[A\033[K")
		if strings.TrimSpace(input) == "/exit" {
			connection.Close()
			return
		}
		outgoing <- strings.TrimSpace(input)
	}
}
