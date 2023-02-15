package service

import (
	"log"
	redisUtil "relationmicor/util"
)

/*
*
查询粉丝数
*/
func FindFollowerCount(userId int64) int64 {
	key := redisUtil.GetFollowerKey(userId)
	count, err := redisUtil.FindZSetCount(key)
	if err != nil {
		log.Printf("FindFollowerCount: userId:%d, redis findZSetCount exception:%s", userId, err)
	}
	return count
}
