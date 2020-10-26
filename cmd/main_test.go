package main

import (
	"fmt"

	"github.com/Guiguerreiro39/go-auction-house/pkg/rewards"
	"github.com/Guiguerreiro39/go-auction-house/pkg/rooms"
	"github.com/Guiguerreiro39/go-auction-house/pkg/services"
	"github.com/Guiguerreiro39/go-auction-house/pkg/users"
	"github.com/Guiguerreiro39/go-auction-house/storage"
)

func main() {
	var roomStorage rooms.Storage
	var rewardStorage rewards.Storage
	var userStorage users.Storage

	roomStorage = new(storage.MemoryRoomStorage)
	rewardStorage = new(storage.MemoryRewardStorage)
	userStorage = new(storage.MemoryUserStorage)

	service := services.NewService(&roomStorage, &rewardStorage, &userStorage)

	errUser := service.AddUser(users.User{Name: "Guilherme", Password: "pass1234", Email: "gui@gui"})
	errRoom := service.AddRoom(rooms.Room{Name: "New Room", Owner: 1, Min: 60.0, Reward: 1})
	errReward := service.AddReward(rewards.Reward{Name: "Awesome reward", Value: 60.0})

	if errUser == nil && errRoom == nil && errReward == nil {
		fmt.Println(service.GetUserByID(1))
		fmt.Println(service.GetRoomByID(1))
		fmt.Println(service.GetRewardByID(1))
	}

}
