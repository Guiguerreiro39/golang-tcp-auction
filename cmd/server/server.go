package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
	"github.com/Guiguerreiro39/go-auction-house/storage"
)

var roomStorage rooms.Storage
var rewardStorage rewards.Storage
var userStorage users.Storage

func handleOwner(text string, service services.Service, user *users.User, room rooms.Room, insideRoom *int) string {
	enc := gob.NewEncoder(user.Con)

	switch text {
	case "1":
		getRoomUsers(enc, room, service)
		return ""
	case "2":
		return getCurrentBid(room)
	case "3":
		return getRoomWinner(service, room)
	case "4":
		return endAuction(service, room, insideRoom)
	default:
		return "Unknown command! " + text
	}
}

func handleRoom(text string, service services.Service, user *users.User, room rooms.Room, insideRoom *int) string {
	dec := gob.NewDecoder(user.Con)

	switch text {
	case "1":
		return bid(dec, room, user, service)
	case "2":
		return getCurrentBid(room)
	case "3":
		return getRoomWinner(service, room)
	case "4":
		return getRoomReward(service, room)
	case "5":
		return leaveRoom(service, room, user, insideRoom)
	default:
		return "Unknown command! " + text
	}
}

func handleCommands(text string, service services.Service, user *users.User, insideRoom *int) string {
	dec := gob.NewDecoder(user.Con)
	enc := gob.NewEncoder(user.Con)

	switch text {
	case "1":
		return addRoom(dec, enc, user, service, insideRoom)
	case "2":
		getActiveRooms(enc, service)
		return ""
	case "3":
		return joinRoom(dec, enc, service, user, insideRoom)
	case "4":
		getUserRewards(enc, user, service)
		return ""
	case "5":
		return addReward(dec, user, service)
	case "6":
		getAllUsers(enc, service)
		return ""
	case "7":
		return getUserCash(user)
	default:
		return "Unknown command! " + text
	}
}

func handleClose(con net.Conn) {
	con.Close()
	fmt.Println("Client connection closed!")
}

func handler(text string, service services.Service, user *users.User, insideRoom *int) string {
	switch *insideRoom {
	case 0:
		return handleCommands(text, service, user, insideRoom)
	default:
		room, _ := service.GetRoomByID(*insideRoom)
		if room.Owner == user.ID {
			return handleOwner(text, service, user, room, insideRoom)
		}

		return handleRoom(text, service, user, room, insideRoom)
	}
}

func handleConnection(con net.Conn) {
	service := services.NewService(&roomStorage, &rewardStorage, &userStorage)
	in := bufio.NewReader(con)
	var insideRoom int
	var response string

	defer handleClose(con)

	username, err := in.ReadString('\n')
	if err != nil {
		return
	}

	userID := addUser(in, service, con, username)

	for {
		temp, err := in.ReadString('\n')
		if err != nil {
			return
		}

		text := strings.TrimSpace(string(temp))
		if text == "8" {
			break
		}

		user, _ := service.GetUserByID(userID)
		response = handler(text, service, &user, &insideRoom)

		if len(response) > 0 {
			fmt.Fprintf(con, response+"\n")
		}
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
	fmt.Println("Server started in port " + args[1])

	for {
		con, err := listener.Accept()
		if err != nil {
			return
		}

		fmt.Print("-> New Client!\n")

		go handleConnection(con)
	}
}
