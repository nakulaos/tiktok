package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/commentsmodel"
	"tiktok/comment/errorcode"
	"tiktok/common/constant"
	errorcode2 "tiktok/common/errorcode"
	userModel "tiktok/user/model"
	"tiktok/user/rpc/user"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	userID := in.UserId
	videoID := in.VideoId

	// 1. 判断用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.CommentUserIdEmptyError, "user is not found , id = %d", userID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user failed, err = %v", err)
		}
	}

	// 2.检测视频是否存在
	_, err = l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.CommentVideoIdEmptyError, "video is not found , id = %d", videoID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
		}
	}

	comments, err := l.svcCtx.CommentsModel.FindCommentsByVid(l.ctx, videoID)
	if err != nil && err != commentsmodel.ErrNotFound {
		return nil, errors.Wrapf(errorcode2.ServerError, "find comments failed, err = %v", err)
	}

	reslist := make([]*comment.Comment, 0)
	for i := 0; i < len(comments); i++ {
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userID, ActorId: int64(comments[i].Uid)})
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		userDetail := &comment.User{
			Id:             userInfo.User.Id,
			Nickname:       userInfo.User.Name,
			FollowCount:    userInfo.User.FollowCount,
			Gender:         userInfo.User.Gender,
			FollowerCount:  userInfo.User.FollowerCount,
			IsFollow:       userInfo.User.IsFollow,
			Avatar:         userInfo.User.Avatar,
			Signature:      userInfo.User.Signature,
			TotalFavorited: userInfo.User.TotalFavorited,
			WorkCount:      userInfo.User.WorkCount,
			FavoriteCount:  userInfo.User.FavoriteCount,
			FriendCount:    userInfo.User.FriendCount,
		}
		res := &comment.Comment{
			Id:         int64(comments[i].Id),
			User:       userDetail,
			Content:    comments[i].Content,
			CreateDate: comments[i].CreateTime.Format(constant.TimeFormat),
		}

		reslist = append(reslist, res)
	}

	return &comment.CommentListResponse{CommentList: reslist}, nil

}
