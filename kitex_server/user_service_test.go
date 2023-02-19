package kitex_server

import (
	"fmt"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"testing"
)

func TestGetUserInfos(t *testing.T) {
	userIds := make([]int64, 5)
	userIds[0] = 1
	userIds[1] = 2006
	userIds[2] = 2009
	userIds[3] = 3
	userIds[4] = 4

	userInfos, err := GetUserInfos(userIds)
	if err != nil {
		t.Error(err)
	}

	for _, userInfo := range userInfos {
		fmt.Println("id:", userInfo.Id)
		fmt.Println("name:", userInfo.Name          )
		fmt.Println("followCount:",*userInfo.FollowCount   )
		fmt.Println("followerCount:",*userInfo.FollowerCount )
		fmt.Println("isFollow:", userInfo.IsFollow      )
		fmt.Println("avatar:", *userInfo.Avatar        )
		fmt.Println("BackgroundImage:", *userInfo.BackgroundImage)
		fmt.Println("Signature:", *userInfo.Signature     )
		fmt.Println("TotalFavorited:", *userInfo.TotalFavorited)
	}
}

func TestMain(m *testing.M) {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}
	Init(r)

	m.Run()
}
