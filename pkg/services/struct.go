package services

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

// AddService is a structure that stores all the add functionalities
type service struct {
	room   rooms.Storage
	reward rewards.Storage
	user   users.Storage
}

// Service is an interface for the add service
type Service interface {
	AddRoom(rooms.Room) (int, error)
	AddUser(users.User) int
	AddReward(rewards.Reward) int

	GetRoomByID(int) (rooms.Room, error)
	GetUserByID(int) (users.User, error)
	GetRewardByID(int) (rewards.Reward, error)

	GetRooms() []rooms.Room
	GetUsers() []users.User
	GetRewards() []rewards.Reward

	UpdateRoom(rooms.Room)
	UpdateUser(users.User)
	UpdateReward(rewards.Reward)
}

// NewService creates a new add service
func NewService(room *rooms.Storage, reward *rewards.Storage, user *users.Storage) Service {
	return &service{*room, *reward, *user}
}
