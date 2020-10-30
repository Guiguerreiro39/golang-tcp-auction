package main

import (
	"encoding/gob"
	"fmt"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
	"github.com/Guiguerreiro39/go-auction-house/util"
)

func bid(dec *gob.Decoder, room rooms.Room, user *users.User, service services.Service) string {
	var bid float64
	dec.Decode(&bid)

	if !room.Active {
		return "This room is not active anymore!"
	}

	if bid <= room.CurrentBid {
		return "Bid is too low!"
	}

	if user.ID == room.CurrentWinner {
		return "You are the highest bid already!"
	}

	if bid > user.Cash {
		return "You don't have enough cash to place that bid!"
	}

	if bid < room.Min {
		min := fmt.Sprintf("%.2f", room.Min)
		return "Bid does not satisfy the minimum price of " + min + string('$')
	}

	user.BidValue = user.BidValue + bid
	if user.BidValue > user.Cash {
		return "Your total bids surpace the amount of cash you can afford!"
	}
	service.UpdateUser(*user)

	winner, _ := service.GetUserByID(room.CurrentWinner)
	winner.BidValue = winner.BidValue - room.CurrentBid
	service.UpdateUser(winner)

	room.CurrentBid = bid
	room.CurrentWinner = user.ID
	service.UpdateRoom(room)

	return "Bid placed!"
}

func endAuction(service services.Service, room rooms.Room, insideRoom *int) string {
	if room.CurrentWinner == 0 {
		*insideRoom = 0
		return "There are no bids, the auction has closed without selling the reward!"
	}

	winner, _ := service.GetUserByID(room.CurrentWinner)
	owner, _ := service.GetUserByID(room.Owner)
	reward, _ := service.GetRewardByID(room.Reward)

	winner.Cash = winner.Cash - room.CurrentBid
	winner.BidValue = winner.BidValue - room.CurrentBid
	winner.Rewards = append(winner.Rewards, room.Reward)
	service.UpdateUser(winner)

	reward.Value = room.CurrentBid
	service.UpdateReward(reward)

	owner.Cash = owner.Cash + room.CurrentBid
	for i, r := range owner.Rewards {
		if r == room.Reward {
			owner.Rewards = util.Remove(owner.Rewards, i)
			break
		}
	}
	service.UpdateUser(owner)

	room.Active = false
	service.UpdateRoom(room)

	*insideRoom = 0

	bid := fmt.Sprintf("%.2f", room.CurrentBid)
	return "The auction has ended! You've sold the reward '" + reward.Name + "' for " + bid + string('$') + " to the user " + winner.Name
}

func leaveRoom(service services.Service, room rooms.Room, user *users.User, insideRoom *int) string {
	for i, u := range room.Users {
		roomUser, _ := service.GetUserByID(u)
		if roomUser.ID == user.ID {
			room.Users = util.Remove(room.Users, i)
			break
		}
	}

	service.UpdateRoom(room)
	*insideRoom = 0
	return "You've left the room"
}

func joinRoom(dec *gob.Decoder, enc *gob.Encoder, service services.Service, user *users.User, insideRoom *int) string {
	var id int
	dec.Decode(&id)
	room, err := service.GetRoomByID(id)
	if err != nil {
		enc.Encode(false)
		return "Failed to enter room!"
	}
	room.Users = append(room.Users, user.ID)
	service.UpdateRoom(room)

	*insideRoom = id
	enc.Encode(true)
	return "You've just joined room - " + room.Name
}
