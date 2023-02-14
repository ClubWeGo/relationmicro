package service

import (
	"fmt"
	redisUtil "relationmicor/util"
	"testing"
)

// 关注
func TestFollow(t *testing.T) {
	err := Follow(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(1, 4)
	if err != nil {
		fmt.Println(err)
	}
	err = Follow(1, 5)
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
	list, err := FindFollowList(1)
	if err != nil {
		fmt.Printf("TestFindFollowList exception:%s", err)
	}
	for _, a := range list {
		fmt.Println(a)
	}
}

func TestMain(m *testing.M) {
	redisUtil.Init()
	m.Run()
}
