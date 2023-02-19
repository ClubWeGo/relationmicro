package service

import (
	"fmt"
	kitexServer "github.com/ClubWeGo/relationmicro/kitex_server"
	"log"
	"strconv"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

type FollowerList struct {
	UserList []FollowerUser
}

type FollowerUser struct {
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
func FindFollowerList(myUid int64, targetUid int64) ([]FollowerUser, error) {
	var followerList = make([]FollowerUser, 0)

	key := redisUtil.GetFollowerKey(targetUid)
	// 拿到按关注时间 从新到老 的粉丝userId
	res, err := redisUtil.FindTopVal(key)
	if err != nil {
		return followerList, fmt.Errorf("FindFollowerList: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}

	// []int64 的followerUserId 为填入昵称作预备
	var followerUserIds = make([]int64, len(res))
	// []followerUserId (string) -> followList 封装
	for i, val := range res {
		// target 的 粉丝 userId 不处理err 因为err时返回0 无影响
		followerUserId, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Printf("FindFollowerList: myUid:%d, targetUid:%d, parseInt exception:%s", myUid, targetUid, err)
			continue
		}
		followerUserIds[i] = followerUserId
	}
	// kitex 请求用户服务获取用户详细信息
	return FindFollowerUserDetailBySet(followerUserIds)
}

/*
*
根据userIds 获取 对应的 昵称集合
填入followerList中
*/
func SetFollowerNameByUserIds(followerList []FollowerUser, followerUserIds []int64) {
	// 缓存里拿到userIds对应的names
	nameMap := FindUserNameByUserIdSet(followerUserIds)
	if nameMap != nil && followerUserIds != nil && len(nameMap) == len(followerList) {
		for i, u := range followerList {
			// map 若无key 返回 ""
			if name := nameMap[u.Id]; name != "" {
				followerList[i].Name = name
			}
		}
	}

	fmt.Println("xxx↓")
	for _, follower := range followerList {
		fmt.Println(follower.Name)
	}

}

/*
*
查询粉丝的其他信息
*/
func FindFollowerOther(myId int64, followerUserId int64) FollowerUser {
	var followerUser = FollowerUser{Id: followerUserId, Name: redisUtil.USER_DEFAULT_NAME}
	// todo 查询用户名
	followCount := FindFollowCount(followerUserId)
	followerCount := FindFollowerCount(followerUserId)
	followerUser.FollowCount = &followCount
	followerUser.FollowerCount = &followerCount
	followerUser.IsFollow = FindIsFollow(myId, followerUserId)
	userInfo, err := kitexServer.GetUserInfo(followerUserId)
	if err != nil {
		log.Printf("FindFollowOther error, myId:%d, followUserId:%d, err:%s", myId, followerUserId, err)
		return followerUser
	}
	// rpc 请求
	followerUser.Name = userInfo.Name
	followerUser.Avatar = userInfo.Avatar
	followerUser.BackgroundImage = userInfo.BackgroundImage
	followerUser.Signature = userInfo.Signature
	followerUser.TotalFavorited = userInfo.TotalFavorited
	followerUser.WorkCount = userInfo.WorkCount
	followerUser.FavoriteCount = userInfo.FavoriteCount
	return followerUser
}
