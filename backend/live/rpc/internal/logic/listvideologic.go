package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/pili"
	"strconv"
	"tiktok/common/errorcode"
	"tiktok/live/rpc/internal/svc"
	"tiktok/live/rpc/live"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看直播列表
func (l *ListVideoLogic) ListVideo(in *live.ListLiveRequest) (*live.ListLiveResponse, error) {
	man := pili.ManagerConfig{
		AccessKey: l.svcCtx.Config.QiNiu.AccessKey,
		SecretKey: l.svcCtx.Config.QiNiu.SecretKey,
		Transport: nil,
	}

	manager := pili.NewManager(man)
	list, err := manager.GetStreamsList(l.ctx, pili.GetStreamListRequest{
		Hub:      l.svcCtx.Config.QiNiu.LiveBucket,
		LiveOnly: true,
	})

	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get stream list failed :%v", err)
	}

	if len(list.Items) == 0 {
		return &live.ListLiveResponse{UserList: nil}, err
	}

	userlist := make([]*live.User, 0)

	for _, item := range list.Items {
		uid, err := strconv.Atoi(item.Key)
		if err != nil {
			return nil, err
		}
		info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(uid),
		})
		userlist = append(userlist, &live.User{
			Id:           int64(uid),
			Name:         info.User.Name,
			IsFollow:     info.User.IsFollow,
			Avatar:       *info.User.Avatar,
			Signature:    *info.User.Signature,
			LiveUrl:      fmt.Sprintf("http://%s/%s/%s.m3u8", l.svcCtx.Config.QiNiu.LiveUrl, l.svcCtx.Config.QiNiu.LiveBucket, item.Key),
			LiveCoverUrl: fmt.Sprintf("http://%s/%s/%s.jpg", l.svcCtx.Config.QiNiu.LiveCoverUrl, l.svcCtx.Config.QiNiu.LiveBucket, item.Key),
		})

	}

	return &live.ListLiveResponse{UserList: userlist}, nil

}
