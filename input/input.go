package input

import (
	"fmt"
	"strconv"

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
