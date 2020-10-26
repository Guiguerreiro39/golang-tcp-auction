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

func receiver(in *bufio.Reader) {
	// Listen to the server response
	message, _ := in.ReadString('\n')
	fmt.Print("->: " + message)
}

func handleInput(text string, con net.Conn, in *bufio.Reader) {
	fmt.Fprintf(con, text+"\n")
	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)

	switch text {
	case "1":
		room := input.AddRoom()
		enc.Encode(room)
		receiver(in)
	case "2":
		var allRooms []string
		dec.Decode(&allRooms)
		for _, room := range allRooms {
			fmt.Println("-> " + room)
		}
	case "3":
		break
	case "4":
		break
	case "5":
		break
	case "6":
		break
	case "7":
		break
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
	in := bufio.NewReader(con)

	input.Welcome()

	for {
		input.Home()
		fmt.Print(">> ")
		temp, _ := out.ReadString('\n')

		text := strings.TrimSpace(string(temp))
		if text == "8" {
			fmt.Println("TCP client exiting...")
			return
		}

		handleInput(text, con, in)
	}

}
