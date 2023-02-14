package util

import "strconv"

type zsetItem struct {
	val   string
	score string
}

type zsetRes struct {
	res []zsetItem
}

const (
	SERVICE_NAME    = "relation" // 服务名
	KEY_INTERVAL    = "_"        // 键名间隔
	FOLLOW_PREFIX   = "follow"   // 关注 前缀
	FOLLOWER_PREFIX = "follower"
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
