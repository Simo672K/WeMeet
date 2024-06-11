package repository

import (
	"WeMeet/pkg/utils"
	"WeMeet/services/chatservice"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RoomRepository struct {
	DBClient *redis.Client
}

func (r *RoomRepository) CreateRoom(roomName string) {
	ctx := context.Background()

	randId, err := utils.RandomString(31)
	if err != nil {
		panic(err)
	}

	newChatRoomMedata := chatservice.ChatRoomMetaData{
		Id:        randId,
		RoomName:  roomName,
		ExpiresIn: int(time.Hour),
	}
	newChatRoom := chatservice.ChatRoom{
		ChatRoomMetaData: newChatRoomMedata,
	}
	newChatRoom.InitRoomChannels()

	err = r.DBClient.Set(ctx, randId, "world", 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *RoomRepository) FindRoom() {

}

func (r *RoomRepository) ConnectToRoomById() {

}

func (r *RoomRepository) QuitRoom() {

}

func (r *RoomRepository) DestroyRoom() {

}
