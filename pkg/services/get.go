package services

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

func (s service) GetRoomByID(id int) (rooms.Room, error) {
	return s.room.Get(id)
}

func (s service) GetRooms() []rooms.Room {
	return s.room.GetAll()
}

func (s service) GetRewardByID(id int) (rewards.Reward, error) {
	return s.reward.Get(id)
}

func (s service) GetRewards() []rewards.Reward {
	return s.reward.GetAll()
}

func (s service) GetUserByID(id int) (users.User, error) {
	return s.user.Get(id)
}

func (s service) GetUsers() []users.User {
	return s.user.GetAll()
}
