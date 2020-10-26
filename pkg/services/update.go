package services

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
)

func (s service) UpdateRoom(room rooms.Room) {
	s.room.Update(room)
}
