package handlers

import (
	"fmt"
	"strings"
	"sync"
)

var UserRoom = make(map[string]string)
var MuUserRoom sync.RWMutex
var Rooms = make(map[string]map[string]bool)
var MuRooms sync.RWMutex

func HandleCommand(cmd string, user string, outgoing chan<- string) {
	fmt.Println(user + " Command: " + cmd)
	command := strings.Fields(cmd)

	if command[0] == "/join" && len(command) == 2 {
		outgoing <- "You joined " + command[1] + "!"
		JoinRoom(command[1], user)
	} else if command[0] == "/leave" && len(command) == 1 {
		prevRoom := UserRoom[user]
		LeaveRoom(user)
		outgoing <- "You left " + prevRoom + "!"
	} else {
		fmt.Println("Invalid command " + command[0] + " by " + user)
		outgoing <- "Invalid command! Please try again!"
	}

	return
}

func JoinRoom(room string, user string) {
	MuUserRoom.RLock()
	oldRoom, exists := UserRoom[user]
	MuUserRoom.RUnlock()

	if exists && oldRoom != room {
		MuRooms.Lock()
		delete(Rooms[oldRoom], user)
		if len(Rooms[oldRoom]) == 0 {
			delete(Rooms, oldRoom)
		}
		MuRooms.Unlock()
	}

	MuRooms.Lock()
	users, exists := Rooms[room]
	if !exists {
		users = make(map[string]bool)
		Rooms[room] = users
	}
	users[user] = true
	MuRooms.Unlock()

	MuUserRoom.Lock()
	UserRoom[user] = room
	MuUserRoom.Unlock()

	return
}

func LeaveRoom(user string) {
	MuUserRoom.RLock()
	room, exists := UserRoom[user]
	MuUserRoom.RUnlock()
	if !exists {
		return
	}

	MuRooms.Lock()
	if users, exists := Rooms[room]; exists {
		delete(users, user)
		if len(users) == 0 {
			delete(Rooms, room)
		}
	}
	MuRooms.Unlock()

	MuUserRoom.Lock()
	delete(UserRoom, user)
	MuUserRoom.Unlock()

	return
}
