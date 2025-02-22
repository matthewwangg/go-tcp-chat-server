package handlers

import (
	"fmt"
	"sync"
)

var Rooms = make(map[string][]string)
var MuRooms sync.Mutex

func HandleCommand(cmd string) {
	fmt.Println("Command: " + cmd)
	return
}
