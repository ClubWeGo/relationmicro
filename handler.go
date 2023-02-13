package main

import (
	"context"
	relation "relationmicor/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// ActionMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ActionMethod(ctx context.Context, request *relation.ActionReq) (resp *relation.ActionResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowAndFollowerMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowAndFollowerMethod(ctx context.Context, request *relation.GetFollowAndFollowerReq) (resp *relation.GetFollowAndFollowerResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowListReqMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowListReqMethod(ctx context.Context, request *relation.GetFollowListReq) (resp *relation.GetFollowListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerListMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq) (resp *relation.GetFollowerListResp, err error) {
	// TODO: Your code here...
	return
}
