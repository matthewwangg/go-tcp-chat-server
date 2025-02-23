package handlers

import (
	"fmt"
	"strings"
	"sync"
)

type Room struct {
	Users map[string]bool
	Ch    chan string
}

var UserRoom = make(map[string]string)
var MuUserRoom sync.RWMutex
var Rooms = make(map[string]*Room)
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
	_, exists := UserRoom[user]
	MuUserRoom.RUnlock()
	if exists {
		LeaveRoom(user)
	}

	MuRooms.Lock()
	_, exists = Rooms[room]
	if !exists {
		Rooms[room] = &Room{
			Users: make(map[string]bool),
			Ch:    make(chan string, 10),
		}
		go StartRoomListener(room)
	}
	roomStruct := Rooms[room]
	roomStruct.Users[user] = true
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
	if roomStruct, exists := Rooms[room]; exists {
		delete(roomStruct.Users, user)
		if len(roomStruct.Users) == 0 {
			close(roomStruct.Ch)
			delete(Rooms, room)
		}
	}
	MuRooms.Unlock()

	MuUserRoom.Lock()
	delete(UserRoom, user)
	MuUserRoom.Unlock()

	return
}
