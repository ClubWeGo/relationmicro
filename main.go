package main

import (
	"log"
	"net"

	relationmicro "github.com/ClubWeGo/relationmicro/kitex_gen/relation/relationservice"
	redisUtil "github.com/ClubWeGo/relationmicro/util"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	kitexServer "github.com/ClubWeGo/relationmicro/kitex_server"
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

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatalf("etcd registry err:%s", err)
	}

	resolver, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatalf("etcd resolver err:%s", err)
	}

	kitexServer.Init(resolver)

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10002")
	svr := relationmicro.NewServer(new(CombineServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "relationservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
