package service

import (
	"fmt"
	"github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	kitexServer "github.com/ClubWeGo/relationmicro/kitex_server"
)

// 根据userIds 查询用户详情信息
func FindFollowUserDetailBySet(userIds []int64) ([]FollowUser, error) {
	userInfos, err := kitexServer.GetUserInfos(userIds)
	if err != nil {
		return []FollowUser{}, fmt.Errorf("FindFollowUserDetailBySet err:%s", err)
	}
	followUsers := make([]FollowUser, len(userInfos))


	for i, userInfo := range userInfos {
		followUsers[i] = ConvertUserInfo2FollowUser(userInfo)
	}
	return followUsers, nil

}

// 根据userIds 查询用户详情信息
func FindFollowerUserDetailBySet(userIds []int64) ([]FollowerUser, error) {
	userInfos, err := kitexServer.GetUserInfos(userIds)
	if err != nil {
		return []FollowerUser{}, fmt.Errorf("FindFollowUserDetailBySet err:%s", err)
	}
	followerUsers := make([]FollowerUser, len(userInfos))

	for i, userInfo := range userInfos {
		// userInfo -> FollowerUser
		followerUsers[i] = ConvertUserInfo2FollowerUser(userInfo)
	}
	return followerUsers, nil

}

func ConvertUserInfo2FollowUser(userInfo relation.User) FollowUser {
	return FollowUser{
		Id: userInfo.Id,
		Name: userInfo.Name,
		FollowCount: userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow: userInfo.IsFollow,
		Avatar: userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:  userInfo.Signature,
		TotalFavorited: userInfo.TotalFavorited,
		WorkCount: userInfo.WorkCount,
		FavoriteCount: userInfo.FavoriteCount,
	}
}

func ConvertUserInfo2FollowerUser(userInfo relation.User) FollowerUser {
	return FollowerUser{
		Id: userInfo.Id,
		Name: userInfo.Name,
		FollowCount: userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow: userInfo.IsFollow,
		Avatar: userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:  userInfo.Signature,
		TotalFavorited: userInfo.TotalFavorited,
		WorkCount: userInfo.WorkCount,
		FavoriteCount: userInfo.FavoriteCount,
	}
}
