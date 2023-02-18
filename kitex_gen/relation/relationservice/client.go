// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	relation "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FollowMethod(ctx context.Context, request *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowResp, err error)
	GetFollowInfoMethod(ctx context.Context, request *relation.GetFollowInfoReq, callOptions ...callopt.Option) (r *relation.GetFollowInfoResp, err error)
	GetFollowListMethod(ctx context.Context, request *relation.GetFollowListReq, callOptions ...callopt.Option) (r *relation.GetFollowListResp, err error)
	GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq, callOptions ...callopt.Option) (r *relation.GetFollowerListResp, err error)
	GetFriendListMethod(ctx context.Context, request *relation.GetFriendListReq, callOptions ...callopt.Option) (r *relation.GetFriendListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) FollowMethod(ctx context.Context, request *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowMethod(ctx, request)
}

func (p *kRelationServiceClient) GetFollowInfoMethod(ctx context.Context, request *relation.GetFollowInfoReq, callOptions ...callopt.Option) (r *relation.GetFollowInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowInfoMethod(ctx, request)
}

func (p *kRelationServiceClient) GetFollowListMethod(ctx context.Context, request *relation.GetFollowListReq, callOptions ...callopt.Option) (r *relation.GetFollowListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowListMethod(ctx, request)
}

func (p *kRelationServiceClient) GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq, callOptions ...callopt.Option) (r *relation.GetFollowerListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerListMethod(ctx, request)
}

func (p *kRelationServiceClient) GetFriendListMethod(ctx context.Context, request *relation.GetFriendListReq, callOptions ...callopt.Option) (r *relation.GetFriendListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendListMethod(ctx, request)
}
