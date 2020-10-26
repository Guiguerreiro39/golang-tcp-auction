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
	AddRoom(rooms.Room) error
	AddUser(users.User) error
	AddReward(rewards.Reward) error

	GetRoomByID(int) (rooms.Room, error)
	GetUserByID(int) (users.User, error)
	GetRewardByID(int) (rewards.Reward, error)

	GetRooms() []string
	GetUsers() []users.User
	GetRewards() []rewards.Reward
}

// NewService creates a new add service
func NewService(room *rooms.Storage, reward *rewards.Storage, user *users.Storage) Service {
	return &service{*room, *reward, *user}
}
