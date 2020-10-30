package input

import (
	"bufio"
	"fmt"
	"os"
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

// RoomOwner is a function that shows a room options for the owner of the room
func RoomOwner() {
	fmt.Println(`
---- Room ----
1. Check room users
2. See current bid
3. Check current winner
4. End auction
	`)
}

// RoomClient is a function that shows a room options for a client
func RoomClient() {
	fmt.Println(`
---- Room ----
1. Place bid
2. See current highest bid
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Room ID: ")
		scanner.Scan()
		input = scanner.Text()
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Bid: ")
		scanner.Scan()
		input = scanner.Text()
		bid, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		}
		fmt.Println("That is not a float64!")
	}

	return bid

}

// AddReward is the input text to add a new reward
func AddReward() rewards.Reward {
	var reward rewards.Reward
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Reward name: ")
	scanner.Scan()
	reward.Name = scanner.Text()

	return reward
}

// AddRoom is the input text to add a new room
func AddRoom() rooms.Room {
	var input string
	var room rooms.Room
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Room name: ")
	scanner.Scan()
	room.Name = scanner.Text()

	for {
		fmt.Println("Minimum price: ")
		scanner.Scan()
		input = scanner.Text()
		room.Min, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		}
		fmt.Println("That is not a float64!")
	}

	for {
		fmt.Println("Reward id: ")
		scanner.Scan()
		input = scanner.Text()
		room.Reward, err = strconv.Atoi(input)
		if err == nil {
			break
		}
		fmt.Println("That is not an int!")
	}

	room.Active = true

	return room
}
