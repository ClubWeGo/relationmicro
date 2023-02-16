package main

import (
	"log"

	relation "github.com/ClubWeGo/relationmicro/kitex_gen/relation/relationservice"
	redisUtil "github.com/ClubWeGo/relationmicro/util"
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

	svr := relation.NewServer(new(CombineServiceImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
