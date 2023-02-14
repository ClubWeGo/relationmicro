package pojo

import "time"

// 关注者
type FollowDTO struct {
	userId       int64
	followedTime time.Time // 关注时间
}

// 粉丝
type FollowerDTO struct {
	userId         int64
	beFollowedTime time.Time // 被关注时间
}
