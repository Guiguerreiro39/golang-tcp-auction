package main

import (
	"bufio"
	"encoding/gob"
	"net"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
)

func addRoom(dec *gob.Decoder, enc *gob.Encoder, user *users.User, service services.Service, insideRoom *int) string {
	var room rooms.Room

	dec.Decode(&room)
	room.Owner = user.ID

	success := false

	for _, reward := range user.Rewards {
		if reward == room.Reward {
			success = true
			break
		}
	}

	if !success {
		enc.Encode(false)
		return "You don't own that reward!"
	}

	_, err := service.GetRewardByID(room.Reward)
	if err != nil {
		enc.Encode(false)
		return "That reward doesn't exist!"
	}

	roomID, err := service.AddRoom(room)
	if err != nil {
		enc.Encode(false)
		return "Failed to create room!"
	}

	*insideRoom = roomID
	enc.Encode(true)
	return "Room created!"
}

func addReward(dec *gob.Decoder, user *users.User, service services.Service) string {
	var reward rewards.Reward
	dec.Decode(&reward)

	rewardID := service.AddReward(reward)
	user.Rewards = append(user.Rewards, rewardID)
	service.UpdateUser(*user)

	return "Reward created!"
}

func addUser(in *bufio.Reader, service services.Service, con net.Conn, username string) int {
	userID := service.AddUser(users.User{
		Name: username,
		Con:  con,
	})

	return userID
}
