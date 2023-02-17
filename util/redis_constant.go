package util

import (
	"strconv"
	"time"
)

type ZSetItem struct {
	val   string
	score string
}

func NewZSetItem() *ZSetItem { // 返回结构体ZSetItem实例的指针
	item := new(ZSetItem)
	return item
}

type ZSetRes struct {
	set []ZSetItem
}

func NewZSetRes() *ZSetRes { // 返回结构体ZSetRes实例的指针
	res := new(ZSetRes)
	return res
}

const (
	SERVICE_NAME      = "relation" // 服务名
	KEY_INTERVAL      = "_"        // 键名间隔
	FOLLOW_PREFIX     = "follow"   // 关注 前缀
	FOLLOWER_PREFIX   = "follower"
	USER_NAME_SUFFIX  = "uname"
	USER_DEFAULT_NAME = "unknow"
)

// 关注集合Key
func GetFollowKey(userId int64) string {
	return SERVICE_NAME + KEY_INTERVAL + FOLLOW_PREFIX + KEY_INTERVAL + strconv.FormatInt(userId, 10)
}

// 粉丝集合键
func GetFollowerKey(userId int64) string {
	// int64 直接强转 string 会解析成utf8
	return SERVICE_NAME + KEY_INTERVAL + FOLLOWER_PREFIX + KEY_INTERVAL + strconv.FormatInt(userId, 10)
}

// 获取当前时间 秒级格式化
func GetFollowedTimeStr() string {
	return time.Now().Format("20060102150402")
}

// 获取用户昵称集合键
func GetUserNameKey() string {
	return SERVICE_NAME + KEY_INTERVAL + USER_NAME_SUFFIX
}
