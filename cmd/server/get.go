package main

import (
	"encoding/gob"
	"fmt"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

func getRoomUsers(enc *gob.Encoder, room rooms.Room, service services.Service) {
	var users []string
	for _, uid := range room.Users {
		u, _ := service.GetUserByID(uid)
		users = append(users, u.Name)
	}

	enc.Encode(users)
}

func getRoomWinner(service services.Service, room rooms.Room) string {
	winner, err := service.GetUserByID(room.CurrentWinner)
	if err != nil {
		if room.CurrentWinner == 0 {
			return "There are no bids yet!"
		}
		return "Error getting winner!"
	}
	return "Current winner is " + winner.Name
}

func getCurrentBid(room rooms.Room) string {
	var bid string

	if room.CurrentWinner == 0 {
		bid = fmt.Sprintf("%.2f", room.Min)
	} else {
		bid = fmt.Sprintf("%.2f", room.CurrentBid)
	}

	return "Current bid is " + bid + string('$')
}

func getRoomReward(service services.Service, room rooms.Room) string {
	reward, err := service.GetRewardByID(room.Reward)
	if err != nil {
		return "Error getting reward!"
	}

	value := fmt.Sprintf("%.2f", reward.Value)
	return "This room reward is '" + reward.Name + "' with a value of " + value + string('$')
}

func getActiveRooms(enc *gob.Encoder, service services.Service) {
	var availableRooms []rooms.Room
	rooms := service.GetRooms()

	for _, r := range rooms {
		if r.Active {
			availableRooms = append(availableRooms, r)
		}
	}

	enc.Encode(availableRooms)
}

func getUserRewards(enc *gob.Encoder, user *users.User, service services.Service) {
	var rewards []rewards.Reward

	for _, reward := range user.Rewards {
		r, _ := service.GetRewardByID(reward)
		rewards = append(rewards, r)
	}

	enc.Encode(rewards)
}

func getAllUsers(enc *gob.Encoder, service services.Service) {
	users := service.GetUsers()
	var names []string
	for _, u := range users {
		names = append(names, u.Name)
	}
	enc.Encode(names)
}

func getUserCash(user *users.User) string {
	cash := fmt.Sprintf("%.2f", user.Cash)
	return "You currently have " + cash + string('$')
}
