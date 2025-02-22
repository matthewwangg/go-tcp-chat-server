package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func HandleLogin(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	_, err = connection.Write([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

}
