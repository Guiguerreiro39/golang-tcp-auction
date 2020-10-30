package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Guiguerreiro39/go-auction-house/input"
)

var insideRoom = false
var isOwner = false
var con net.Conn

func handleOwner(text string) {
	dec := gob.NewDecoder(con)

	switch text {
	case "1":
		getRoomUsers(dec)
	case "2":
		receiver()
	case "3":
		receiver()
	case "4":
		receiver()
		insideRoom = false
	default:
		receiver()
	}
}

func handleRoom(text string) {
	enc := gob.NewEncoder(con)

	switch text {
	case "1":
		bid(enc)
		receiver()
	case "2":
		receiver()
	case "3":
		receiver()
	case "4":
		receiver()
	case "5":
		receiver()
		insideRoom = false
	default:
		receiver()
	}

}

func handleInput(text string) {
	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)

	switch text {
	case "1":
		addRoom(enc, dec)
		receiver()
	case "2":
		getActiveRooms(dec)
	case "3":
		joinRoom(enc, dec)
	case "4":
		getUserRewards(dec)
	case "5":
		addReward(enc)
		receiver()
	case "6":
		getAllUsers(dec)
	case "7":
		receiver()
	default:
		receiver()
	}
}

func handler(text string, status int) {
	switch status {
	case -1:
		handleOwner(text)
	case 0:
		handleInput(text)
	case 1:
		handleRoom(text)
	default:
		receiver()
	}
}

func handleUI() int {
	switch insideRoom {
	case false:
		input.Home()
		return 0
	case true:
		if isOwner {
			input.RoomOwner()
			return -1
		}

		input.RoomClient()
		return 1
	}

	return 0
}

func main() {
	args := os.Args
	var err error

	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := args[1]
	con, err = net.Dial("tcp", CONNECT)

	if err != nil {
		fmt.Println(err)
		return
	}

	out := bufio.NewReader(os.Stdin)

	input.Welcome()
	addUser(out)

	for {
		status := handleUI()

		var text string
		for {
			fmt.Print(">> ")
			temp, _ := out.ReadString('\n')

			text = strings.TrimSpace(string(temp))
			if text == "8" {
				fmt.Println("TCP client exiting...")
				return
			}
			if text != "" {
				break
			}
		}

		fmt.Fprintf(con, text+"\n")
		handler(text, status)
	}

}
