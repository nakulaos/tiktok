// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package server

import (
	"context"

	"tiktok/relation/rpc/internal/logic"
	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"
)

type RelationServer struct {
	svcCtx *svc.ServiceContext
	relation.UnimplementedRelationServer
}

func NewRelationServer(svcCtx *svc.ServiceContext) *RelationServer {
	return &RelationServer{
		svcCtx: svcCtx,
	}
}

func (s *RelationServer) Favorite(ctx context.Context, in *relation.FavoriteRequest) (*relation.FavoriteResponse, error) {
	l := logic.NewFavoriteLogic(ctx, s.svcCtx)
	return l.Favorite(in)
}

func (s *RelationServer) FavoriteList(ctx context.Context, in *relation.FavoriteListReq) (*relation.FavoriteListResp, error) {
	l := logic.NewFavoriteListLogic(ctx, s.svcCtx)
	return l.FavoriteList(in)
}

func (s *RelationServer) FollowerList(ctx context.Context, in *relation.FollowerListReq) (*relation.FollowerListResp, error) {
	l := logic.NewFollowerListLogic(ctx, s.svcCtx)
	return l.FollowerList(in)
}

func (s *RelationServer) FriendList(ctx context.Context, in *relation.FriendListReq) (*relation.FriendListResp, error) {
	l := logic.NewFriendListLogic(ctx, s.svcCtx)
	return l.FriendList(in)
}

func (s *RelationServer) GetFollowCount(ctx context.Context, in *relation.FollowerCountReq) (*relation.FollowerCountResp, error) {
	l := logic.NewGetFollowCountLogic(ctx, s.svcCtx)
	return l.GetFollowCount(in)
}

func (s *RelationServer) GetFollowerCount(ctx context.Context, in *relation.FollowerCountReq) (*relation.FollowerCountResp, error) {
	l := logic.NewGetFollowerCountLogic(ctx, s.svcCtx)
	return l.GetFollowerCount(in)
}

func (s *RelationServer) IsFollowing(ctx context.Context, in *relation.IsFollowingReq) (*relation.IsFollowingResp, error) {
	l := logic.NewIsFollowingLogic(ctx, s.svcCtx)
	return l.IsFollowing(in)
}
