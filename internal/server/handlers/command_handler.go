package handlers

import (
	"fmt"
	"strings"
	"sync"
)

var UserRoom = make(map[string]string)
var MuUserRoom sync.Mutex
var Rooms = make(map[string]map[string]bool)
var MuRooms sync.Mutex

func HandleCommand(cmd string, user string) {
	fmt.Println("Command: " + cmd)
	command := strings.Fields(cmd)

	if command[0] == "/join" && len(command) == 2 {
		JoinRoom(command[1], user)
	} else if command[0] == "/leave" && len(command) == 1 {
		LeaveRoom(user)
	} else {
		fmt.Println("Invalid command")
	}

	return
}

func JoinRoom(room string, user string) {
	MuRooms.Lock()
	MuUserRoom.Lock()
	defer MuUserRoom.Unlock()
	defer MuRooms.Unlock()

	if oldRoom, exists := UserRoom[user]; exists && oldRoom != room {
		delete(Rooms[oldRoom], user)
		if len(Rooms[oldRoom]) == 0 {
			delete(Rooms, oldRoom)
		}
	}

	if _, exists := Rooms[room]; !exists {
		Rooms[room] = make(map[string]bool)
	}
	Rooms[room][user] = true
	UserRoom[user] = room

	return
}

func LeaveRoom(user string) {
	MuRooms.Lock()
	MuUserRoom.Lock()
	defer MuUserRoom.Unlock()
	defer MuRooms.Unlock()

	room, exists := UserRoom[user]
	if !exists {
		return
	}

	if users, exists := Rooms[room]; exists {
		delete(users, user)
		if len(users) == 0 {
			delete(Rooms, room)
		}
	}

	delete(UserRoom, user)

	return
}
