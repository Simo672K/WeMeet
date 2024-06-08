package repository

import (
	"WeMeet/pkg/utils"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RoomRepository struct {
	DBClient *redis.Client
}

func (r *RoomRepository) CreateRoom() {
	ctx := context.Background()

	randId, err := utils.RandomString(31)
	if err != nil {
		panic(err)
	}

	err = r.DBClient.Set(ctx, randId, "world", 0).Err()
	if err != nil {
		panic(err)
	}

	result, err := r.DBClient.Get(ctx, randId).Result()
	if err != nil {
		log.Fatal("An error occured while retriving data from db ", err)
	}

	fmt.Println(result)
}

func (r *RoomRepository) FindRoom() {

}

func (r *RoomRepository) ConnectToRoomById() {

}

func (r *RoomRepository) QuitRoom() {

}

func (r *RoomRepository) DestroyRoom() {

}
