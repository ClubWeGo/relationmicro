package service

import (
	"fmt"
	"testing"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

// go test -v message_service_test.go message_service.go
func TestGenerateP2PRoomID(t *testing.T) {
	roomName := GenerateP2PRoomID(3, 14)
	fmt.Println(roomName)
	roomName = GenerateP2PRoomID(58, 14)
	fmt.Println(roomName)
}

func TestMsg(t *testing.T) {
	config := redisUtil.Config{
		Url:         "localhost:6379",
		Password:    "123456",
		DB:          0,
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeOut: 300,
	}

	redisUtil.Init(config)
	redisUtil.XDELALL("1_2")

	r, err := SendP2PMsg(1, 2, "testmsg1")
	if err != nil {
		println(err.Error())
	}
	println(r)
	r2, err := GetAllP2PMsg(1, 2)
	if err != nil {
		println(err.Error())
	}
	for _, item := range r2 {
		println(item.UserId, item.ToUserId, item.Create_at.Unix(), item.Content)
	}
}
