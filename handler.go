package main

import (
	"context"
	"fmt"

	relation "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	"github.com/ClubWeGo/relationmicro/service"
)

const(
	ERROR = "0"
	SUCCESS = "1"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}



// ActionMethod implements the RelationServiceImpl interface.
func (s *CombineServiceImpl) ActionMethod(ctx context.Context, request *relation.ActionReq) (resp *relation.ActionResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowAndFollowerMethod implements the RelationServiceImpl interface.
func (s *CombineServiceImpl) GetFollowAndFollowerMethod(ctx context.Context, request *relation.GetFollowAndFollowerReq) (resp *relation.GetFollowAndFollowerResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowListReqMethod implements the RelationServiceImpl interface.
func (s *CombineServiceImpl) GetFollowListReqMethod(ctx context.Context, request *relation.GetFollowListReq) (resp *relation.GetFollowListResp, err error) {
	myId := request.MyId
	targetId := request.TargetId
	// todo 参数校验
	//if myId != nil && *myId == targetId {
	//
	//}
	// todo 需要兼容 myId为nil的情况
	followList, err := service.FindFollowList(*myId, targetId)

	if err != nil {
		return &relation.GetFollowListResp{
			StatusCode: ERROR,
			UserList: []*relation.User{},
		}, err
	}

	// 封装响应
	respUserList := make([]*relation.User, len(followList))
	for i, followUser := range followList {
		fmt.Println(followUser)
		respUserList[i] = &relation.User{
			Id: followUser.Id,
			Name: followUser.Name,
			FollowCount: followUser.FollowCount,
			FollowerCount: followUser.FollowerCount,
			IsFollow: followUser.IsFollow,
		}
	}
	return &relation.GetFollowListResp{
		StatusCode: SUCCESS,
		UserList: respUserList,
	}, err



}

// GetFollowerListMethod implements the RelationServiceImpl interface.
func (s *CombineServiceImpl) GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq) (resp *relation.GetFollowerListResp, err error) {
	myId := request.MyId
	targetId := request.TargetId
	// todo 参数校验
	// todo 需要兼容 myId为nil的情况
	followerList, err := service.FindFollowerList(*myId, targetId)
	if err != nil {
		return &relation.GetFollowerListResp{
			StatusCode: ERROR,
			UserList: []*relation.User{},

		}, err
	}
	// 封装响应
	respUserList := make([]*relation.User, len(followerList))
	for i, followerUser := range followerList {
		respUserList[i] = &relation.User{
			Id: followerUser.Id,
			Name: followerUser.Name,
			FollowCount: followerUser.FollowCount,
			FollowerCount: followerUser.FollowerCount,
			IsFollow: followerUser.IsFollow,
		}
	}
	return &relation.GetFollowerListResp{
		StatusCode: SUCCESS,
		UserList: respUserList,

	}, nil
}

// GetAllMessageMethod implements the MessageServiceImpl interface.
func (s *CombineServiceImpl) GetAllMessageMethod(ctx context.Context, request *relation.GetAllMessageReq) (resp *relation.GetAllMessageResp, err error) {
	// TODO: Your code here...
	// service层拿数据
	msgs, err := service.GetAllP2PMsg(request.UserId, request.ToUserId)

	if err != nil {
		return &relation.GetAllMessageResp{
			Status: false,
			Msg:    []*relation.Message{}, //返回空消息
		}, nil
	}

	respMsg := make([]*relation.Message, len(msgs))
	for index, msg := range msgs {
		createTimeString := msg.Create_at.Format("2006-01-02")
		respMsg[index] = &relation.Message{
			Id:         msg.Id,
			FromUserId: msg.UserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Content,
			CreateTime: &createTimeString,
		}
	}
	return &relation.GetAllMessageResp{
		Status: true,
		Msg:    respMsg,
	}, nil
}

// SendMessageMethod implements the MessageServiceImpl interface.
func (s *CombineServiceImpl) SendMessageMethod(ctx context.Context, request *relation.SendMessageReq) (resp *relation.SendMessageResp, err error) {
	// TODO: Your code here...
	// service层拿数据
	_, err = service.SendP2PMsg(request.UserId, request.ToUserId, request.Content)
	if err != nil {
		return &relation.SendMessageResp{
			Status: false,
		}, err
	}
	return &relation.SendMessageResp{
		Status: true,
	}, nil
}
