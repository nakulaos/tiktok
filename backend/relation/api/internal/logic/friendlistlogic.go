package logic

import (
	"context"
	"encoding/json"
	"tiktok/relation/rpc/relation"

	"tiktok/relation/api/internal/svc"
	"tiktok/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	list, err := l.svcCtx.RelationRpc.FriendList(l.ctx, &relation.FriendListReq{
		Uid:     uid,
		ActUser: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	resList := make([]types.UserInfo, 0)
	for _, val := range list.List {
		resList = append(resList, types.UserInfo{
			Id:              val.Id,
			Name:            val.NickName,
			Gender:          val.Gender,
			Mobile:          val.Mobile,
			Avatar:          val.Avatar,
			Dec:             val.Dec,
			BackgroundImage: val.BackgroundImage,
			FollowCount:     val.FollowCount,
			FollowerCount:   val.FollowerCount,
			TotalFavorited:  val.TotalFavorited,
			WorkCount:       val.WorkCount,
			FavoriteCount:   val.FavoriteCount,
			IsFollow:        val.IsFollow,
			CoverUrl:        val.CoverUrl,
			VideoId:         val.VideoId,
			FriendCount:     val.FriendCount,
		})
	}

	return &types.FriendListResp{
		List: resList,
	}, nil
}
