package handlers

import (
	"sync"
	"time"
)

var LastSeen sync.Map

func UpdateLastSeen(username string) {
	LastSeen.Store(username, time.Now())
}

func GetLastSeen(username string) string {
	val, exists := LastSeen.Load(username)
	if !exists {
		return "Unknown"
	}
	return val.(time.Time).Format("15:04")
}

func RemoveUser(username string) {
	LastSeen.Delete(username)
}
