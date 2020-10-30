package users

import (
	"errors"
	"net"
)

// User is the structure of an user
type User struct {
	ID       int
	Name     string
	Rewards  []int
	Con      net.Conn
	Cash     float64
	BidValue float64
}

// ErrDuplicate creates a duplicate error
var ErrDuplicate = errors.New("User already exists")

// ErrNotFound creates a not found error
var ErrNotFound = errors.New("User was not found")

// Storage is an interface with the methods to store and retrieve a user
type Storage interface {
	Add(User) int
	Get(int) (User, error)
	GetAll() []User
	Update(User)
}
