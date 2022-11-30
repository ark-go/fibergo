package userdata

import (
// "log"
)

type roomName string

type Stage struct {
	Room roomName
	//Rooms map[roomName]func(user *User)
}

// переход к другой комнате
func (s *Stage) Goto(Name string) {
	// if _, ok := s.Rooms[roomName(Name)]; ok {
	// 	s.Room = roomName(Name)
	// } else {
	// 	log.Println("Нет такой комнаты:", Name)
	// }
}

// Добавление комнаты
func (s *Stage) ChangeRoom(Name string, fun func(user *User)) {
	// s.Rooms[roomName(Name)] = fun
}
