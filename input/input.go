package input

import (
	"fmt"
	"strconv"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
)

// Welcome is a function to welcome the client
func Welcome() {
	fmt.Println(`
#####################################
### Welcome to the auction house! ###	
#####################################
	`)
}

// Home is a function that shows the client options
func Home() {
	fmt.Println(`
---- Room ----
1. New room
2. See all available rooms
3. Join a room

---- Reward ----
4. See all my rewards
5. New reward

---- User ----
6. See all users
7. See my cash
8. Exit
	`)
}

// Room is a function that shows a room options for a client
func Room() {
	fmt.Println(`
---- Room ----
1. Place bid
2. See current price
3. Check current winner
4. Check reward
5. Quit
	`)
}

// JoinRoom receives the id of the room to join
func JoinRoom() int {
	var id int
	var input string
	var err error

	for {
		fmt.Println("Room ID: ")
		fmt.Scanln(&input)
		id, err = strconv.Atoi(input)
		if err == nil {
			break
		}
		fmt.Println("That is not an int!")
	}

	return id
}

// Bid is a function that returns the bid of the client
func Bid() float64 {
	var input string
	var bid float64
	var err error

	for {
		fmt.Println("Bid: ")
		fmt.Scanln(&input)
		bid, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		}
		fmt.Println("That is not a float64!")
	}

	return bid

}

// RemoveUser removes a user from a room using its index
func RemoveUser(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// AddReward is the input text to add a new reward
func AddReward() rewards.Reward {
	var name string
	var reward rewards.Reward

	fmt.Println("Reward name: ")
	fmt.Scanln(&name)
	reward.Name = name

	return reward
}

// AddRoom is the input text to add a new room
func AddRoom() rooms.Room {
	var input string
	var room rooms.Room
	var err error

	fmt.Println("Room name: ")
	fmt.Scanln(&input)
	room.Name = input

	for {
		fmt.Println("Minimum price: ")
		fmt.Scanln(&input)
		room.Min, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		}
		fmt.Println("That is not a float64!")
	}

	for {
		fmt.Println("Reward id: ")
		fmt.Scanln(&input)
		room.Reward, err = strconv.Atoi(input)
		if err == nil {
			break
		}
		fmt.Println("That is not an int!")
	}

	room.Active = true

	return room
}
