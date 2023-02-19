package kitex_server

import (
	"context"
	"fmt"
	"github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
)

// 根据userId 获取用户信息
func GetUserInfo(uid int64) (relation.User, error) {
	fmt.Println("GetUserInfo")
	fmt.Println(Userclient == nil)
	resp, err := Userclient.GetUserMethod(context.Background(), &usermicro.GetUserReq{
		Id: &uid,
	})
	fmt.Println("GetUserInfoMethod end")
	if err != nil {
		return relation.User{}, fmt.Errorf("kitex-usermicroserver exception, uid：%d, err:%s", uid, err)
	}

	if resp.Status {
		userInfo := resp.User
		return ConvertUserInfo2User(*userInfo), nil

	}
	return relation.User{}, fmt.Errorf("kitex-usermicroserver : error to get user")
}

func GetUserInfos(userIds []int64) ([]relation.User, error) {
	resp, err := Userclient.GetUserSetByIdSetMethod(context.Background(), &usermicro.GetUserSetByIdSetReq{IdSet: userIds})
	if err != nil {
		return []relation.User{}, fmt.Errorf("GetUserInfos kitex-usermicroserver exception, err:%s", err)
	}

	if resp.Status {
		userInfos := resp.UserSet
		//infoMap := ConvertUserInfoToMap(userInfos)
		users := make([]relation.User, len(userInfos))
		for i, userInfo := range userInfos {
			users[i] = ConvertUserInfo2User(*userInfo)
		}
		return users, nil
	}
	return []relation.User{}, fmt.Errorf("GetUserInfos kitex-usermicroserver : error to get userInfos by userIds")
}

func ConvertUserInfo2User(userInfo usermicro.UserInfo) relation.User {
	return relation.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     &userInfo.FollowCount,
		FollowerCount:   &userInfo.FollowerCount,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &userInfo.TotalFavorited,
		WorkCount:       &userInfo.WorkCount,
		FavoriteCount:   &userInfo.FavoriteCount,
	}
}

// userInfo -> userId - userInfo
func ConvertUserInfoToMap(userInfos []*usermicro.UserInfo) map[int64]usermicro.UserInfo {
	userInfoMap := make(map[int64]usermicro.UserInfo, len(userInfos))
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.Id] = *userInfo
	}
	return userInfoMap
}
