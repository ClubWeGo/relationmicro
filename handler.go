package main

import (
	"context"

	message "github.com/ClubWeGo/relationmicro/kitex_gen/message"
	relation "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	"github.com/ClubWeGo/relationmicro/service"
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
	// TODO: Your code here...
	return
}

// GetFollowerListMethod implements the RelationServiceImpl interface.
func (s *CombineServiceImpl) GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq) (resp *relation.GetFollowerListResp, err error) {
	// TODO: Your code here...
	return
}

// GetAllMessageMethod implements the MessageServiceImpl interface.
func (s *CombineServiceImpl) GetAllMessageMethod(ctx context.Context, request *relation.GetAllMessageReq) (resp *relation.GetAllMessageResp, err error) {
	// TODO: Your code here...
	// service层拿数据
	msgs, err := service.GetAllP2PMsg(request.UserId, request.ToUserId)

	if err != nil {
		return &message.GetAllMessageResp{
			Status: false,
			Msg:    []*message.Message{}, //返回空消息
		}, nil
	}

	respMsg := make([]*message.Message, len(msgs))
	for index, msg := range msgs {
		createTimeString := msg.Create_at.Format("2006-01-02")
		respMsg[index] = &message.Message{
			Id:         msg.Id,
			FromUserId: msg.UserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Content,
			CreateTime: &createTimeString,
		}
	}
	return &message.GetAllMessageResp{
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
		return &message.SendMessageResp{
			Status: false,
		}, err
	}
	return &message.SendMessageResp{
		Status: true,
	}, nil
}
