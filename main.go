package main

import (
	"log"
	message "relationmicor/kitex_gen/message/messageservice"
	relation "relationmicor/kitex_gen/relation/relationservice"
	redisUtil "relationmicor/util"
)

func main() {

	config := redisUtil.Config{
		Url:         "localhost:6379",
		Password:    "123456",
		DB:          0,
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeOut: 300,
	}

	redisUtil.Init(config)

	Rsvr := relation.NewServer(new(RelationServiceImpl))

	err := Rsvr.Run()

	if err != nil {
		log.Println(err.Error())
	}

	Msvr := message.NewServer(new(MessageServiceImpl))
	err = Msvr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
