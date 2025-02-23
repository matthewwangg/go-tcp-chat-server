package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleInput(outgoing chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(input) == "/exit" {
			return
		}
		outgoing <- strings.TrimSpace(input)
	}
}
