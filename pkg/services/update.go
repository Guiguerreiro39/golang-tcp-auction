package services

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

func (s service) UpdateRoom(room rooms.Room) {
	s.room.Update(room)
}

func (s service) UpdateUser(user users.User) {
	s.user.Update(user)
}

func (s service) UpdateReward(reward rewards.Reward) {
	s.reward.Update(reward)
}
