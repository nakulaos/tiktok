package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"tiktok/common/errorcode"

	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowCountLogic {
	return &GetFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowCountLogic) GetFollowCount(in *relation.FollowerCountReq) (*relation.FollowerCountResp, error) {
	sb := squirrel.Select().From("`relations`").Where("follower_id =?", in.Uid)
	count, err := l.svcCtx.UserModel.FindCount(l.ctx, sb, "id")
	if err != nil {
		return nil, errors.Wrapf(errorcode.DatabaseError, "Database query user favorite count error")
	}

	return &relation.FollowerCountResp{
		Count: count,
	}, nil
}
