package logic

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/pkg/errors"
	"tiktok/common/constant"
	"tiktok/common/errorcode"
	errorcode2 "tiktok/feed/errorcode"
	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"
	"tiktok/feed/videosmodel"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchESLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchESLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchESLogic {
	return &SearchESLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchESLogic) SearchES(in *feed.EsSearchReq) (*feed.EsSearchResp, error) {
	content := in.Content

	fields := []string{
		"title", "content", "label",
	}
	//query := fmt.Sprintf(`
	//{
	//	"query": {
	//		"multi_match": {
	//			"query": "%s",
	//			"fields": ["title","content","label"]
	//		}
	//	}
	//}
	//`, content)

	res, err := l.svcCtx.Es.Search().Index("video-index").Request(
		&search.Request{
			Query: &types.Query{
				MultiMatch: &types.MultiMatchQuery{
					Query:  content,
					Fields: fields,
				},
			},
		}).Do(l.ctx)

	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "query index in es error:%v", err)
	}

	videoIds := make([]int, 0)

	var data []byte
	err = res.UnmarshalJSON(data)
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "unmarshal es json error:%v", err)
	}

	var responses Response

	err = json.Unmarshal(data, &responses)

	for i := 0; i < int(res.Hits.Total.Value); i++ {
		videoIds = append(videoIds, responses.Hits.Hits[i].Source.VideoID)
	}
	reslists := make([]*feed.VideoInfo, 0)

	queryIds := removeDuplicates(videoIds)
	for i := 0; i < len(queryIds); i++ {
		curVideoID := queryIds[i]
		video, err := l.svcCtx.VideosModel.FindOne(l.ctx, int64(curVideoID))
		if err != nil {
			if err == videosmodel.ErrNotFound {
				continue
			}
			return nil, errors.Wrapf(errorcode2.FeedUnableToQueryVideoError, "find video error:%v", err)
		}
		userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  int64(in.UserId),
			ActorId: int64(video.AuthorId),
		})
		if err != nil {
			return nil, errors.Wrapf(errorcode2.FeedUnableToQueryUserError, "find user error:%v", err)
		}
		IsFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.UserId), int64(curVideoID))
		IsStar, _ := l.svcCtx.StarModel.IsStar(l.ctx, int64(in.UserId), int64(curVideoID))
		userInfo := &feed.User{
			Id:              userRpcRes.User.Id,
			Nickname:        userRpcRes.User.Name,
			FollowCount:     userRpcRes.User.FollowCount,
			FollowerCount:   userRpcRes.User.FollowCount,
			IsFollow:        userRpcRes.User.IsFollow,
			Avatar:          userRpcRes.User.Avatar,
			BackgroundImage: userRpcRes.User.BackgroundImage,
			Signature:       userRpcRes.User.Signature,
			TotalFavorited:  userRpcRes.User.TotalFavorited,
			WorkCount:       userRpcRes.User.WorkCount,
			FavoriteCount:   userRpcRes.User.FavoriteCount,
			Gender:          userRpcRes.User.Gender,
			FriendCount:     userRpcRes.User.FriendCount,
		}
		videoInfo := &feed.VideoInfo{
			Id:            uint32(curVideoID),
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: uint32(video.FavoriteCount),
			CommentCount:  uint32(video.CommentCount),
			IsFavorite:    IsFavorite,
			Title:         video.Title,
			CreateTime:    video.CreateTime.Format(constant.TimeFormat),
			StarCount:     uint32(video.StarCount),
			IsStar:        IsStar,
			Duration:      video.Duration.String,
		}
		reslists = append(reslists, videoInfo)
	}

	return &feed.EsSearchResp{VideoList: reslists}, nil
}

type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				VideoID int    `json:"video_id"`
				Title   string `json:"title"`
				Content string `json:"content"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func removeDuplicates(nums []int) []int {
	// 创建一个map用于存储元素是否已经存在
	seen := make(map[int]bool)
	result := []int{}

	// 迭代原始数组
	for _, num := range nums {
		// 如果元素不存在于map中，则将其添加到结果数组中，并将其标记为已存在
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	return result
}
