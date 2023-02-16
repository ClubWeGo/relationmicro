package service

import (
	"log"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

/*
*
保存 userId-用户昵称 避免db查询，节省资源
*/
func SaveUserName(userId int64, name string) {

}

/*
*
根据userId 查询 用户昵称
*/
func FindUserNameByUserId(userId int64) {
	//redisUtil.HMGet
}

/*
*
根据userIds 查询 对应用户昵称
// todo 避免userIds数量太大 可考虑分批获取
*/
func FindUserNameByUserIdSet(userIds []int64) map[int64]string {
	//userSize := len(userIds)
	if userIds == nil {
		return nil
	}
	var nameMap = make(map[int64]string)
	key := redisUtil.GetUserNameKey()
	nameMap, err := redisUtil.HMGetI64ReturnMapI64(key, userIds...)
	if err != nil {
		for _, userId := range userIds {
			nameMap[userId] = redisUtil.USER_DEFAULT_NAME
		}
		log.Printf("FindUserNameByUserIdSet exception")

	}
	return nameMap
}
