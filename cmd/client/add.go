package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/Guiguerreiro39/go-auction-house/input"
)

func addRoom(enc *gob.Encoder, dec *gob.Decoder) {
	var success bool

	room := input.AddRoom()
	enc.Encode(room)

	dec.Decode(&success)
	if success {
		insideRoom = true
		isOwner = true
	}
}

func addReward(enc *gob.Encoder) {
	reward := input.AddReward()
	enc.Encode(reward)
}

func addUser(out *bufio.Reader) {
	fmt.Print(">> Username: ")
	username, _ := out.ReadString('\n')
	username = strings.TrimSpace(string(username))

	fmt.Fprintf(con, username+"\n")
}
