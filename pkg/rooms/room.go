package rooms

import (
	"errors"
)

// Room is the structure for a room
type Room struct {
	ID            int
	Name          string
	Users         []int
	Owner         int
	Min           float64
	CurrentBid    float64
	CurrentWinner int
	Reward        int
	Active        bool
}

// ErrDuplicate creates a duplicate error
var ErrDuplicate = errors.New("Room already exists")

// ErrNotFound creates a not found error
var ErrNotFound = errors.New("Room was not found")

// Storage is an interface with the methods to store and retrieve a room
type Storage interface {
	Add(Room) (int, error)
	Get(int) (Room, error)
	GetAll() []Room
	Update(Room)
}
