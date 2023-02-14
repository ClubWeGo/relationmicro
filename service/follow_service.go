package service

import (
	"encoding/json"
	"fmt"
	redisUtil "relationmicor/util"
	"time"
)

// 关注者
type FollowUser struct {
	userId       int64
	followedTime time.Time // 关注时间
}

func Init() {
	redisUtil.Init()
}

//var redisUtil = util.

/*
*
关注
myUid: 我的userId
targetUid：关注目标userId
*/
func Follow(myUid int64, targetUid int64) error {
	// 校验参数
	if err := CheckFollowParam(myUid, targetUid); err != nil {
		return fmt.Errorf("follow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}

	// 获取key
	key := redisUtil.GetFollowKey(myUid)
	// 关注 关注时间是now
	// todo 互关加入好友
	_, err := redisUtil.Zadd(key, time.Now().Unix(), targetUid)
	if err != nil {
		return fmt.Errorf("follow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	return nil
}

// 取关
func UnFollow(myUid int64, targetUid int64) error {
	// 校验参数
	if err := CheckFollowParam(myUid, targetUid); err != nil {
		return fmt.Errorf("unFollow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	// 获取key
	key := redisUtil.GetFollowKey(myUid)
	// todo 取关 如果是好友的话 删除好友
	_, err := redisUtil.Zrem(key, targetUid)
	if err != nil {
		return fmt.Errorf("unFollow: myUid:%d, targetUid:%d, exception:%s", myUid, targetUid, err)
	}
	return nil
}

// 接口文档没给分页接口
// 查询关注集合
func FindFollowList(userId int64) ([]FollowUser, error) {
	var followList = make([]FollowUser, 0, 0)
	var followItem FollowUser

	key := redisUtil.GetFollowKey(userId)
	res, err := redisUtil.FindTop(key)
	if err != nil {
		return nil, fmt.Errorf("FindFollowList: userId:%d, exception:%s", userId, err)
	}

	for _, v := range res {
		json.Unmarshal(v.([]byte), &followItem)
		followList = append(followList, followItem)
		//fmt.Printf("%s\n", v.([]byte))
	}
	return followList, nil
}

func CheckFollowParam(myUid int64, targetUid int64) error {
	// 不能关注自己 不能取关自己
	if myUid == targetUid {
		return fmt.Errorf("param exception myUid and targetUid the same")
	}
	return nil
}
