package handlers

func StartRoomListener(room string) {
	for message := range Rooms[room].Ch {
		MuRooms.RLock()
		users := Rooms[room].Users
		MuRooms.RUnlock()

		for user := range users {
			MuUsers.RLock()
			conn, exists := Users[user]
			MuUsers.RUnlock()
			if exists {
				conn.Write([]byte(message + "\n"))
			}
		}
	}
	return
}
