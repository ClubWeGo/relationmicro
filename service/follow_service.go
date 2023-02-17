package service

import (
	"fmt"
	"log"
	"strconv"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

//// 关注者
//type FollowUser struct {
//	userId       int64
//	followedTime time.Time // 关注时间
//}

/*
*
"user_list": [

	    {
	        "id": 0,
	        "name": "string",
	        "follow_count": 0,
	        "follower_count": 0,
	        "is_follow": true
	    }
	]
*/
type FollowList struct {
	userList []FollowUser
}

type FollowUser struct {
	Id            int64
	Name          string // 昵称
	FollowCount   int64  // 关注数
	FollowerCount int64  // 粉丝数
	IsFollow      bool   // 是否关注 true-已关注 false-未关注
}

func Init(config redisUtil.Config) {
	redisUtil.Init(config)
}

//var redisUtil = util.

/*
*
关注
myUid: 我的userId
targetUid：关注目标userId
1、targetUid添加到关注集合
2、myUid添加到粉丝集合
3、存入两个用户的name 避免db消耗
#保证原子
*/
func Follow(myUid int64, targetUid int64) error {
	// 校验参数
	if err := CheckFollowParam(myUid, targetUid); err != nil {
		return fmt.Errorf("follow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	// 获取key
	followKey := redisUtil.GetFollowKey(myUid)
	followerKey := redisUtil.GetFollowerKey(targetUid)
	// 关注时间 精确到秒级 zset 超过17位会精度丢失
	// todo zset score值也可以使用用户等级 或者 两者的关系程度 但接口文档没提供用户等级、关系等级 只能用关注时间
	nowTimeStr := redisUtil.GetFollowedTimeStr()
	// 我关注别人的同时 也要我成为别人的粉丝 lua脚本保证原子性
	scriptStr := redisUtil.GetFollowScript()
	_, err := redisUtil.Eval(scriptStr, 2, followKey, followerKey, nowTimeStr, targetUid, myUid)
	if err != nil {
		return fmt.Errorf("follow myUid:%d, targetUid:%d redis lua eval exception: %s", myUid, targetUid, err)
	}

	return nil
}

// 取关
/**
1、targetUid从my关注集合删除
2、myUid从target粉丝集合删除
# 保证原子
*/
func UnFollow(myUid int64, targetUid int64) error {
	// 校验参数
	if err := CheckFollowParam(myUid, targetUid); err != nil {
		return fmt.Errorf("unFollow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	// 获取key
	key := redisUtil.GetFollowKey(myUid)
	// todo 取关 如果是好友的话 删除好友
	// todo 我取关别人的同时，也要从别人的粉丝中消失
	_, err := redisUtil.Zrem(key, targetUid)
	if err != nil {
		return fmt.Errorf("unFollow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	return nil
}

// 接口文档没给分页接口
/**
查询关注集合
myUid: 我的userId
targetUid: 查询目标userId
*/
func FindFollowList(myUid int64, targetUid int64) ([]FollowUser, error) {
	var followList = make([]FollowUser, 0)

	key := redisUtil.GetFollowKey(targetUid)
	// 拿到按关注时间 从新到老 的关注者userId
	res, err := redisUtil.FindTopVal(key)
	if err != nil {
		return followList, fmt.Errorf("FindFollowList: mUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}

	// []int64 的followUserId 为填入昵称作预备
	var followUserIds = make([]int64, 0, 0)
	// []followUserId (string) -> followList 封装
	for _, val := range res {
		// target 的 关注者 userId
		followUserId, err := strconv.ParseInt(val, 10, 64)
		followUserIds = append(followUserIds, followUserId)
		if err != nil {
			log.Printf("FindFollowList: mUid:%d, targetUid:%d, parseInt exception:%s", myUid, targetUid, err)
			continue
		} else {
			// 查询target关注者的 其他信息&我与target关注者的关系
			followUser := FindFollowOther(myUid, followUserId)
			followList = append(followList, followUser)
		}
	}
	// 去缓存拿到用户名集合并填入
	SetFollowNameByUserIds(followList, followUserIds)
	return followList, nil
}

/**
redis 拿到的集合串 封装成 FollowList
*/
//func PackageFollowListByRes(followList *FollowList, res []string, followUserIds []int64) {
//	for _, val := range res {
//		// target 的 关注者 userId
//		followUserId, err := strconv.ParseInt(val, 10, 64)
//		followUserIds = append(followUserIds, followUserId)
//		if err != nil {
//			log.Printf("FindFollowList: mUid:%d, targetUid:%d, parseInt exception:%s", myUid, targetUid, err)
//			continue
//		} else {
//			// 查询target关注者的 其他信息&我与target关注者的关系
//			followUser := FindFollowOther(myUid, followUserId)
//			followList.userList = append(followList.userList, followUser)
//		}
//
//	}
//}

/*
*
根据userIds 获取 对应的 昵称集合
填入followList中
*/
func SetFollowNameByUserIds(followList []FollowUser, followUserIds []int64) {
	nameMap := FindUserNameByUserIdSet(followUserIds)

	fmt.Println("nameMap: ")
	for k, v := range nameMap {
		fmt.Println(k, v)
	}

	if nameMap != nil && followUserIds != nil && len(nameMap) == len(followList) {
		for i, u := range followList {
			// map 若无key 返回 ""
			if name := nameMap[u.Id]; name != "" {
				followList[i].Name = name
			}
		}
	}
}

/*
*
查询关注用户的其他信息
*/
func FindFollowOther(myId int64, followUserId int64) FollowUser {
	var followUser = FollowUser{Id: followUserId, Name: redisUtil.USER_DEFAULT_NAME}
	// todo 查询用户名
	followUser.FollowCount = FindFollowCount(followUserId)
	followUser.FollowerCount = FindFollowerCount(followUserId)
	followUser.IsFollow = FindIsFollow(myId, followUserId)
	return followUser
}

/*
*
查询用户的关注数
*/
func FindFollowCount(userId int64) int64 {
	key := redisUtil.GetFollowKey(userId)
	count, err := redisUtil.FindZSetCount(key)
	// 异常返回0值
	if err != nil {
		log.Printf("FindFollowCount: userId:%d, redis findZSetCount exception:%s", userId, err)
	}
	return count
}

/*
*
查询我是否关注target
*/
func FindIsFollow(myUid int64, targetUid int64) bool {
	key := redisUtil.GetFollowKey(myUid)
	exists, err := redisUtil.FindZSetIsExists(key, targetUid)
	if err != nil {
		log.Printf("FindIsFollow: myUid:%d, targetUid:%d, redis FindZSetIsExists exception:%s", myUid, targetUid, err)
	}
	return exists
}

/*
*
校验参数
*/
func CheckFollowParam(myUid int64, targetUid int64) error {
	// 不能关注自己 不能取关自己
	if myUid == targetUid {
		return fmt.Errorf("param exception myUid and targetUid the same")
	}
	return nil
}

/*
*
校验userId 是否正常范围内
放在control层好一点
*/
func CheckUserId(userId int64) error {
	if userId > 0 {
		return nil
	}
	return fmt.Errorf("userId:%d out of range", userId)
}
