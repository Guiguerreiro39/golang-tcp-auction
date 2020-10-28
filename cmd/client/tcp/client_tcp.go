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

func receiver(con net.Conn) {
	// Listen to the server response
	in := bufio.NewReader(con)

	message, _ := in.ReadString('\n')
	fmt.Print("->: " + message)
}

func handleRoom(text string, con net.Conn) {
	enc := gob.NewEncoder(con)
	//dec := gob.NewDecoder(con)

	switch text {
	case "1":
		bid := input.Bid()
		enc.Encode(bid)
		receiver(con)
	case "2":
		receiver(con)
	case "3":
		receiver(con)
	case "4":
		receiver(con)
	case "5":
		receiver(con)
		insideRoom = false
	}

}

func handleInput(text string, con net.Conn) {
	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)

	switch text {
	case "1":
		room := input.AddRoom()
		enc.Encode(room)
		receiver(con)
	case "2":
		var allRooms []string
		dec.Decode(&allRooms)
		for _, room := range allRooms {
			fmt.Println("-> " + room)
		}
	case "3":
		id := input.JoinRoom()
		enc.Encode(id)
		receiver(con)
		insideRoom = true
	case "4":
		var rewards []string
		dec.Decode(&rewards)
		for _, reward := range rewards {
			fmt.Println("-> " + reward)
		}
	case "5":
		reward := input.AddReward()
		enc.Encode(reward)
		receiver(con)
	case "6":
		var users []string
		dec.Decode(&users)
		for _, user := range users {
			fmt.Println("-> " + user)
		}
	case "7":
		receiver(con)
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := args[1]
	con, err := net.Dial("tcp", CONNECT)

	if err != nil {
		fmt.Println(err)
		return
	}

	out := bufio.NewReader(os.Stdin)

	input.Welcome()

	fmt.Print(">> Username: ")
	username, _ := out.ReadString('\n')
	username = strings.TrimSpace(string(username))
	fmt.Fprintf(con, username+"\n")

	for {
		if !insideRoom {
			input.Home()
		} else {
			input.Room()
		}

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

		if !insideRoom {
			fmt.Fprintf(con, text+"\n")
			handleInput(text, con)
		} else {
			fmt.Fprintf(con, text+"\n")
			handleRoom(text, con)
		}
	}

}
