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

var count int = 0
var roomStorage rooms.Storage
var rewardStorage rewards.Storage
var userStorage users.Storage

func handleCommands(text string, service services.Service, user int, con net.Conn) string {
	dec := gob.NewDecoder(con)
	enc := gob.NewEncoder(con)

	switch text {
	case "1":
		var room rooms.Room
		dec.Decode(&room)

		errRoom := service.AddRoom(room)
		if errRoom != nil {
			return "Failed to create room!"
		}

		return "Room created!"
	case "2":
		rooms := service.GetRooms()
		enc.Encode(rooms)
	}

	return "Unknown command! " + text
}

func handleClose(con net.Conn) {
	count--
	con.Close()
	fmt.Println("Client connection closed!")
}

func handleConnection(con net.Conn) {
	service := services.NewService(&roomStorage, &rewardStorage, &userStorage)
	in := bufio.NewReader(con)
	user := count

	defer handleClose(con)

	for {
		temp, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text := strings.TrimSpace(string(temp))
		if text == "STOP" {
			break
		}

		response := handleCommands(text, service, user, con)
		con.Write([]byte(response + "\n"))
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

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("-> New Client!\n")
		go handleConnection(con)
		count++
	}
}
