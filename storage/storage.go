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
func (m *MemoryRoomStorage) Add(room rooms.Room) (int, error) {
	for _, r := range m.rooms {
		if room.Name == r.Name {
			return 0, rooms.ErrDuplicate
		}
	}

	room.ID = len(m.rooms) + 1
	m.rooms = append(m.rooms, room)

	return room.ID, nil
}

// Add is a method to add a new user
func (m *MemoryUserStorage) Add(user users.User) int {
	user.ID = len(m.users) + 1
	user.Cash = 1000.0
	m.users = append(m.users, user)

	return user.ID
}

// Add is a method to add a new reward
func (m *MemoryRewardStorage) Add(reward rewards.Reward) int {
	reward.ID = len(m.rewards) + 1
	m.rewards = append(m.rewards, reward)

	return reward.ID
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
func (m *MemoryRoomStorage) GetAll() []rooms.Room {
	return m.rooms
}

// GetAll retrieves all users
func (m *MemoryUserStorage) GetAll() []users.User {
	return m.users
}

// GetAll retrieves all rewards
func (m *MemoryRewardStorage) GetAll() []rewards.Reward {
	return m.rewards
}

// Update updates a room in the slice
func (m *MemoryRoomStorage) Update(room rooms.Room) {
	for index, r := range m.rooms {
		if r.ID == room.ID {
			m.rooms[index] = room
		}
	}
}

// Update updates a room in the slice
func (m *MemoryUserStorage) Update(user users.User) {
	for index, u := range m.users {
		if u.ID == user.ID {
			m.users[index] = user
		}
	}
}

// Update updates a reward in the slice
func (m *MemoryRewardStorage) Update(reward rewards.Reward) {
	for index, r := range m.rewards {
		if r.ID == reward.ID {
			m.rewards[index] = reward
		}
	}
}
