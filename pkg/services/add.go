package services

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

func (s service) AddRoom(room rooms.Room) (int, error) {
	return s.room.Add(room)
}

func (s service) AddUser(user users.User) int {
	return s.user.Add(user)
}

func (s service) AddReward(reward rewards.Reward) int {
	return s.reward.Add(reward)
}
