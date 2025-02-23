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
	fmt.Print("Input: ")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("\033[A\033[K")
		if strings.TrimSpace(input) == "/exit" {
			return
		}
		outgoing <- strings.TrimSpace(input)
	}
}
