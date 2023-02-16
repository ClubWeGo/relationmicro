package service

import (
	"fmt"
	"log"
	"strconv"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

type FollowerList struct {
	userList []FollowerUser
}

type FollowerUser struct {
	id            int64
	name          string // 昵称
	followCount   int64  // 关注数
	followerCount int64  // 粉丝数
	isFollow      bool   // 是否关注 true-已关注 false-未关注
}

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

/*
*
查询粉丝集合
myUid: 我的userId
targetUid: 查询目标userId
*/
func FindFollowerList(myUid int64, targetUid int64) (FollowerList, error) {
	var followerList = FollowerList{}

	key := redisUtil.GetFollowerKey(targetUid)
	// 拿到按关注时间 从新到老 的粉丝userId
	res, err := redisUtil.FindTopVal(key)
	if err != nil {
		return followerList, fmt.Errorf("FindFollowerList: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}

	// []int64 的followerUserId 为填入昵称作预备
	var followerUserIds = make([]int64, 0, 0)
	// []followerUserId (string) -> followList 封装
	for _, val := range res {
		// target 的 粉丝 userId 不处理err 因为err时返回0 无影响
		followerUserId, err := strconv.ParseInt(val, 10, 64)
		followerUserIds = append(followerUserIds, followerUserId)
		if err != nil {
			log.Printf("FindFollowerList: myUid:%d, targetUid:%d, parseInt exception:%s", myUid, targetUid, err)
			continue
		} else {
			// 查询target粉丝的 其他信息&我与target粉丝的关系
			followerUser := FindFollowerOther(myUid, followerUserId)
			followerList.userList = append(followerList.userList, followerUser)
		}
	}
	// 去缓存拿到用户名集合并填入
	SetFollowerNameByUserIds(&followerList, followerUserIds)
	return followerList, nil
}

/*
*
根据userIds 获取 对应的 昵称集合
填入followerList中
*/
func SetFollowerNameByUserIds(followerList *FollowerList, followerUserIds []int64) {
	nameMap := FindUserNameByUserIdSet(followerUserIds)

	fmt.Println("nameMap: ")
	for k, v := range nameMap {
		fmt.Println(k, v)
	}

	if nameMap != nil && followerUserIds != nil && len(nameMap) == len(followerList.userList) {
		for i, u := range followerList.userList {
			// map 若无key 返回 ""
			if name := nameMap[u.id]; name != "" {
				followerList.userList[i].name = name
			}
		}
	}
}

/*
*
查询粉丝的其他信息
*/
func FindFollowerOther(myId int64, followerUserId int64) FollowerUser {
	var followerUser = FollowerUser{id: followerUserId, name: redisUtil.USER_DEFAULT_NAME}
	// todo 查询用户名
	followerUser.followCount = FindFollowCount(followerUserId)
	followerUser.followerCount = FindFollowerCount(followerUserId)
	followerUser.isFollow = FindIsFollow(myId, followerUserId)
	return followerUser
}
