package storage

import (
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

// MemoryRoomStorage is a structure to store in memory the rooms, users and rewards
type MemoryRoomStorage struct {
	rooms []rooms.Room
}

// MemoryRewardStorage is a structure to store in memory the rooms, users and rewards
type MemoryRewardStorage struct {
	rewards []rewards.Reward
}

// MemoryUserStorage is a structure to store in memory the rooms, users and rewards
type MemoryUserStorage struct {
	users []users.User
}

// Add is a method to add a new room
func (m *MemoryRoomStorage) Add(room rooms.Room) error {
	for _, r := range m.rooms {
		if room.Name == r.Name {
			return rooms.ErrDuplicate
		}
	}

	room.ID = len(m.rooms) + 1
	m.rooms = append(m.rooms, room)

	return nil
}

// Add is a method to add a new user
func (m *MemoryUserStorage) Add(user users.User) error {
	for _, u := range m.users {
		if user.Name == u.Name {
			return users.ErrDuplicate
		}
	}

	user.ID = len(m.users) + 1
	m.users = append(m.users, user)

	return nil
}

// Add is a method to add a new reward
func (m *MemoryRewardStorage) Add(reward rewards.Reward) error {
	reward.ID = len(m.rewards) + 1
	m.rewards = append(m.rewards, reward)

	return nil
}

// Get retrieves the room of the given id
func (m *MemoryRoomStorage) Get(id int) (rooms.Room, error) {
	var room rooms.Room

	for _, r := range m.rooms {
		if r.ID == id {
			return r, nil
		}
	}

	return room, rooms.ErrNotFound
}

// Get retrieves the user of the given id
func (m *MemoryUserStorage) Get(id int) (users.User, error) {
	var user users.User

	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}

	return user, users.ErrNotFound
}

// Get retrieves the reward of the given id
func (m *MemoryRewardStorage) Get(id int) (rewards.Reward, error) {
	var reward rewards.Reward

	for _, r := range m.rewards {
		if r.ID == id {
			return r, nil
		}
	}

	return reward, rewards.ErrNotFound
}

// GetAll retrieves all rooms
func (m *MemoryRoomStorage) GetAll() []string {
	var allRooms []string

	for _, room := range m.rooms {
		allRooms = append(allRooms, room.Name)
	}

	return allRooms
}

// GetAll retrieves all users
func (m *MemoryUserStorage) GetAll() []users.User {
	return m.users
}

// GetAll retrieves all rewards
func (m *MemoryRewardStorage) GetAll() []rewards.Reward {
	return m.rewards
}
