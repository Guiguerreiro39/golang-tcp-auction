package main

import (
	"encoding/gob"
	"fmt"
	"strconv"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
)

func getActiveRooms(dec *gob.Decoder) {
	var rooms []rooms.Room
	dec.Decode(&rooms)
	for _, room := range rooms {
		id := fmt.Sprintf("%d", room.ID)
		fmt.Println("-> " + room.Name + " - ID " + id)
	}
}

func getUserRewards(dec *gob.Decoder) {
	var rewards []rewards.Reward
	dec.Decode(&rewards)
	for _, reward := range rewards {
		id := strconv.Itoa(reward.ID)
		value := fmt.Sprintf("%.2f", reward.Value)
		fmt.Println("-> " + reward.Name + " - ID " + id + " - " + value + string('$'))
	}
}

func getRoomUsers(dec *gob.Decoder) {
	var users []string
	dec.Decode(&users)
	for _, u := range users {
		fmt.Println("-> " + u)
	}
}

func getAllUsers(dec *gob.Decoder) {
	var users []string
	dec.Decode(&users)
	for _, user := range users {
		fmt.Println("-> " + user)
	}
}
