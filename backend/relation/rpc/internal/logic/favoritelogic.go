package logic

import (
	"context"
	"github.com/pkg/errors"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/relation/errorcode"
	relationModel "tiktok/relation/model"
	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"
	userModel "tiktok/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *relation.FavoriteRequest) (*relation.FavoriteResponse, error) {
	// 点赞逻辑
	// action 只能为1或2
	if in.Action != 1 && in.Action != 2 {
		return nil, errors.Wrapf(errorcode.LikeParameterError, "Like parameter error, action:%d", in.Action)
	}

	// 被关注的人不存在
	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(in.ToUid)); err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.RelationUserNotExistError, "the user is not exist , id : %d", in.ToUid)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user error, id : %d,err :%s", in.ToUid, err.Error())
		}
	}

	// 点赞操作
	if in.Action == 1 {
		if in.Uid == in.ToUid {
			// 点赞自己
			return nil, errors.Wrapf(errorcode.RelationUnableFavorSelfError, "uid = %d", in.Uid)
		}

		// 判读是否重复关注
		ok, err := l.svcCtx.RelationModel.IsFollow(l.ctx, in.Uid, in.ToUid)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "uid:%d,toUid:%d,err:%s", in.Uid, in.ToUid, err.Error())
		}
		if ok {
			return nil, errors.Wrapf(errorcode.RelationUnableFavorMoreError, "uid:%d,toUid:%d ", in.Uid, in.ToUid)
		}

		newRelation := relationModel.Relations{
			FollowerId:  in.Uid,
			FollowingId: in.ToUid,
		}

		if _, err := l.svcCtx.RelationModel.Insert(l.ctx, nil, &newRelation); err != nil {
			return nil, errors.Wrapf(errorcode2.DatabaseError, "failed to insert relation table ,uid:%d,toUid:%d,err:%s", in.Uid, in.ToUid, err.Error())
		}

		return nil, nil

	} else {
		// 判断是否关注
		res, err := l.svcCtx.RelationModel.FindRelation(l.ctx, in.Uid, in.ToUid)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "uid:%d,toUid:%d,action: %d err:%s", in.Uid, in.ToUid, in.Action, err.Error())
		}

		if res == 0 {
			return nil, errors.Wrapf(errorcode.RelationUnableUnFavorNotFavorUserError, "uid:%d,toUid:%d,action: %d", in.Uid, in.ToUid, in.Action)
		}

		if err := l.svcCtx.RelationModel.Delete(l.ctx, nil, res); err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "uid:%d,toUid:%d,action: %derr:%s", in.Uid, in.ToUid, in.Action, err.Error())
		}

		return &relation.FavoriteResponse{}, nil

	}

}
