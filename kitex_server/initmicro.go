package kitex_server

import (
	"log"

	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

var Userclient userservice.Client

func Init(r discovery.Resolver) {
	uc, err := userservice.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	Userclient = uc
}
