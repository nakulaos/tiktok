package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/errorcode"

	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowingLogic {
	return &IsFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowingLogic) IsFollowing(in *relation.IsFollowingReq) (*relation.IsFollowingResp, error) {
	flag, err := l.svcCtx.RelationModel.IsFollow(l.ctx, in.ActorId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errorcode.DatabaseError, "Database query user follow flag error")
	}

	return &relation.IsFollowingResp{Flag: flag}, nil
}
