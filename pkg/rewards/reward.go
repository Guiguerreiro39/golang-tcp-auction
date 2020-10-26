package rewards

import "errors"

// Reward is the structure of a reward
type Reward struct {
	ID    int
	Name  string
	Value float64
}

// ErrNotFound creates a not found error
var ErrNotFound = errors.New("Reward was not found")

// Storage is an interface with the methods to store and retrieve a reward
type Storage interface {
	Add(Reward) error
	Get(int) (Reward, error)
	GetAll() []Reward
}
