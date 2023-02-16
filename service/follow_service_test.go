package service

import (
	"fmt"
	"testing"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

// 关注
func TestFollow(t *testing.T) {
	err := Follow(12, 2)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(12, 3)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(12, 4)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(12, 5)
	if err != nil {
		fmt.Println(err)
	}
}

// 关注自己
func TestFollowSame(t *testing.T) {
	err := Follow(1, 1)
	if err != nil {
		fmt.Println(err)
	}
}

// 取关
func TestUnFollow(t *testing.T) {
	err := UnFollow(1, 5)
	if err != nil {
		fmt.Println(err)
	}
}

// 取关自己
func TestUnFollowSame(t *testing.T) {
	err := UnFollow(1, 1)
	if err != nil {
		fmt.Println(err)
	}
}

// 获取关注集合
func TestFindFollowList(t *testing.T) {
	list, err := FindFollowList(1, 12)
	if err != nil {
		t.Errorf("TestFindFollowList exception:%s", err)
	}
	for _, a := range list {
		fmt.Println(a)
	}
}

func TestFindIsFollow(t *testing.T) {
	isFollow := FindIsFollow(12, 5)
	fmt.Println(isFollow)
}

func TestFindNoneFollow(t *testing.T) {
	isFollow := FindIsFollow(12, -1)
	fmt.Println(isFollow)
}

func TestMain(m *testing.M) {
	config := redisUtil.Config{
		Url:         "localhost:6379",
		Password:    "123456",
		DB:          0,
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeOut: 300,
	}

	redisUtil.Init(config)
	m.Run()
}
