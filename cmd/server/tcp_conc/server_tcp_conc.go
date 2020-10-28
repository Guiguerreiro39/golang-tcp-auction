package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Guiguerreiro39/go-auction-house/input"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
	"github.com/Guiguerreiro39/go-auction-house/storage"
)

var count int = 0
var roomStorage rooms.Storage
var rewardStorage rewards.Storage
var userStorage users.Storage

func handleRoom(text string, service services.Service, userID int, con net.Conn, insideRoom *int) string {
	dec := gob.NewDecoder(con)
	//enc := gob.NewEncoder(con)

	room, err := service.GetRoomByID(*insideRoom)
	if err != nil {
		return "Error getting bid!"
	}

	switch text {
	case "1":
		var bid float64
		dec.Decode(&bid)

		if bid <= room.CurrentBid {
			return "Bid is too low!"
		}

		room.CurrentBid = bid
		room.CurrentWinner = userID
		service.UpdateRoom(room)

		return "Bid placed!"
	case "2":
		bid := fmt.Sprintf("%f", room.CurrentBid)
		return "Current winning bid is " + bid
	case "3":
		winner, err := service.GetUserByID(room.CurrentWinner)
		if err != nil {
			if room.CurrentWinner == 0 {
				return "There no bids yet!"
			}
			return "Error getting winner!"
		}
		return "Current winner is " + winner.Name
	case "4":
		reward, err := service.GetRewardByID(room.Reward)
		if err != nil {
			return "Error getting reward!"
		}

		return "This room reward is '" + reward.Name + "'"
	case "5":
		for i, u := range room.Users {
			user, _ := service.GetUserByID(u)
			if user.ID == userID {
				room.Users = input.RemoveUser(room.Users, i)
				break
			}
		}

		service.UpdateRoom(room)
		*insideRoom = 0
		return "You've left the room"
	}

	return "Unknown command! " + text
}

func handleCommands(text string, service services.Service, userID int, con net.Conn, insideRoom *int) string {
	dec := gob.NewDecoder(con)
	enc := gob.NewEncoder(con)

	switch text {
	case "1":
		var room rooms.Room
		dec.Decode(&room)

		_, err := service.GetRewardByID(room.Reward)
		if err != nil {
			return "That reward doesn't exist!"
		}

		err = service.AddRoom(room)
		if err != nil {
			return "Failed to create room!"
		}

		return "Room created!"
	case "2":
		rooms := service.GetRooms()
		enc.Encode(rooms)
	case "3":
		var id int
		dec.Decode(&id)
		room, err := service.GetRoomByID(id)
		if err != nil {
			return "Failed to enter room!"
		}
		room.Users = append(room.Users, userID)
		service.UpdateRoom(room)

		*insideRoom = id
		return "You've just joined room - " + room.Name
	case "4":
		user, _ := service.GetUserByID(userID)
		var rewards []string
		for _, r := range user.Rewards {
			reward, _ := service.GetRewardByID(r)
			rewards = append(rewards, reward.Name)
		}
		enc.Encode(rewards)
	case "5":
		var reward rewards.Reward
		dec.Decode(&reward)

		service.AddReward(reward)
		return "Reward created!"
	case "6":
		users := service.GetUsers()
		var names []string
		for _, user := range users {
			names = append(names, user.Name)
		}
		enc.Encode(names)
	case "7":
		user, _ := service.GetUserByID(userID)
		cash := fmt.Sprintf("%f", user.Cash)
		return "You currently have " + cash + string('$')
	}

	return "Unknown command! " + text
}

func handleClose(con net.Conn) {
	con.Close()
	fmt.Println("Client connection closed!")
}

func handleConnection(con net.Conn) {
	service := services.NewService(&roomStorage, &rewardStorage, &userStorage)
	in := bufio.NewReader(con)
	var insideRoom int

	defer handleClose(con)

	username, err := in.ReadString('\n')
	if err != nil {
		return
	}

	userID := count
	err = service.AddUser(users.User{
		ID:   count,
		Name: username,
		Con:  con,
		Cash: 1000.0,
	})
	if err != nil {
		return
	}

	for {
		var response string
		temp, err := in.ReadString('\n')
		if err != nil {
			return
		}

		text := strings.TrimSpace(string(temp))
		if text == "8" {
			break
		}

		if insideRoom == 0 {
			response = handleCommands(text, service, userID, con, &insideRoom)
		} else {
			response = handleRoom(text, service, userID, con, &insideRoom)
		}

		fmt.Fprintf(con, response+"\n")
	}
}

func main() {
	roomStorage = new(storage.MemoryRoomStorage)
	rewardStorage = new(storage.MemoryRewardStorage)
	userStorage = new(storage.MemoryUserStorage)

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + args[1]
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server start in port " + args[1])

	for {
		con, err := listener.Accept()
		if err != nil {
			return
		}

		fmt.Print("-> New Client!\n")

		go handleConnection(con)
		count++
	}
}
