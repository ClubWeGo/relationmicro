package service

import (
	"fmt"
	"log"
	"strconv"

	kitexServer "github.com/ClubWeGo/relationmicro/kitex_server"
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
	Id              int64
	Name            string  // 用户昵称
	FollowCount     *int64  // 关注数
	FollowerCount   *int64  // 粉丝数
	IsFollow        bool    // 是否关注 true-已关注 false-未关注
	Avatar          *string // 头像
	BackgroundImage *string // 个人顶部大图
	Signature       *string // 个人简介
	TotalFavorited  *int64  // 获赞数
	WorkCount       *int64  // 作品数
	FavoriteCount   *int64  // 喜欢数
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
	_, err := redisUtil.EvalOptimize(scriptStr, 2, followKey, followerKey, nowTimeStr, targetUid, myUid)
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
	followKey := redisUtil.GetFollowKey(myUid)
	followerKey := redisUtil.GetFollowerKey(targetUid)
	scriptStr := redisUtil.GetUnFollowScript()
	// 将对方从自己的关注列表里删除，同时将自己从对方的粉丝列表删除
	_, err := redisUtil.EvalOptimize(scriptStr, 2, followKey, followerKey, targetUid, myUid)
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
	//log.Println("FindFollowList start")
	var followList = make([]FollowUser, 0)

	key := redisUtil.GetFollowKey(targetUid)
	// 拿到按关注时间 从新到老 的关注者userId
	res, err := redisUtil.FindTopVal(key)
	if err != nil {
		return followList, fmt.Errorf("FindFollowList: mUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}

	// []int64 的followUserId 为填入昵称作预备
	var followUserIds = make([]int64, len(res))
	// []followUserId (string) -> followList 封装
	for i, val := range res {
		// target 的 关注者 userId
		followUserId, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Printf("FindFollowList: mUid:%d, targetUid:%d, parseInt exception:%s", myUid, targetUid, err)
			continue
		}
		followUserIds[i] = followUserId
	}
	// 通过kitex 用户服务查询用户详细信息
	return FindFollowUserDetailBySet(followUserIds)
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
	// 缓存里拿到userIds 对应的 names
	nameMap := FindUserNameByUserIdSet(followUserIds)
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
	fmt.Println("FindFollowOther")
	var followUser = FollowUser{Id: followUserId, Name: redisUtil.USER_DEFAULT_NAME}
	// todo 查询用户名
	followCount := FindFollowCount(followUserId)
	followerCount := FindFollowerCount(followUserId)
	followUser.FollowCount = &followCount
	followUser.FollowerCount = &followerCount
	followUser.IsFollow = FindIsFollow(myId, followUserId)
	userInfo, err := kitexServer.GetUserInfo(followUserId)
	if err != nil {
		log.Printf("FindFollowOther error, myId:%d, followUserId:%d, err:%s", myId, followUserId, err)
		return followUser
	}
	// rpc 请求
	followUser.Name = userInfo.Name
	followUser.Avatar = userInfo.Avatar
	followUser.BackgroundImage = userInfo.BackgroundImage
	followUser.Signature = userInfo.Signature
	followUser.TotalFavorited = userInfo.TotalFavorited
	followUser.WorkCount = userInfo.WorkCount
	followUser.FavoriteCount = userInfo.FavoriteCount
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

// 根据userIds 批量查询多个用户的关注状态
func FindIsFollows(myUid int64, userIds []int64) (map[int64]int, error) {
	followKey := redisUtil.GetFollowKey(myUid)
	script := redisUtil.GetIsFollowsScript()
	keyArgvs := make([]interface{}, len(userIds)+1)
	keyArgvs[0] = followKey
	for i, id := range userIds {
		keyArgvs[i+1] = id
	}

	isFollows, err := redisUtil.EvalReturnInts(script, 1, keyArgvs...)
	if err != nil {
		return nil, fmt.Errorf("FindIsFollows lua eval error, myUid:%d, err:%s", myUid, err)
	}
	if len(userIds) != len(isFollows) {
		return nil, fmt.Errorf("FindIsFollows redis 返回 长度不一致, myUid:%d", myUid)
	}
	resMap := make(map[int64]int)
	for i, id := range userIds {
		resMap[id] = isFollows[i]
	}

	return resMap, nil
}
