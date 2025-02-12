// Code generated by goctl. DO NOT EDIT.
package types

type CategoryVideosListReq struct {
	Category uint32 `form:"category"`
}

type CategoryVideosListResp struct {
	Videos []VideoInfo `json:"video_list"`
}

type CreateVideoReq struct {
	Url      string `json:"url" validate:"required" msg:"UrlFormat"` //视频地址
	CoverUrl string `json:"coverUrl" validate:"required" msg:"CoverUrlFormat"`
	Title    string `json:"title" validate:"required" msg:"TitleFormat"`
	Category int    `json:"category" validate:"required" msg:"CategoryFormat"`
}

type DeleteVideoReq struct {
	Vid int64 `json:"video_id" validate:"required" msg:"VidFormat"`
}

type DurationTestReq struct {
	Duration string `json:"duration" validate:"required" msg:"DurationFormat"`
	Vid      int64  `json:"video_id" validate:"required" msg:"VidFormat"`
}

type FindVideoByIdReq struct {
	Vid int64 `form:"video_id" validate:"required" msg:"VidFormat"`
}

type FindVideoByIdResp struct {
	Video VideoInfo `json:"video_info"`
}

type HistoryVideosResp struct {
	VideoList []VideoInfo `json:"video_list"`
}

type NeighborsVideoReq struct {
	Vid int64 `form:"video_id" validate:"required" msg:"VidFormat"`
}

type NeighborsVideoResp struct {
	VideoList []VideoInfo `json:"video_list"`
}

type PopularVideosListReq struct {
	Offset        int64 `json:"offset" validate:"required" msg:"OffsetFormat"`
	ReadedVideoId int64 `json:"readed_videoId" validate:"required,min=1" msg:"ReadedVideoIdFormat"`
}

type PopularVideosListResp struct {
	Videos []VideoInfo `json:"video_list"`
}

type RecommendVideosListReq struct {
	Offset        int64 `json:"offset" validate:"required" msg:"OffsetFormat"`
	ReadedVideoId int64 `json:"readed_videoId" validate:"required,min=1" msg:"ReadedVideoIdFormat"`
}

type RecommendVideosListResp struct {
	Videos []VideoInfo `json:"video_list"`
}

type SearchEsReq struct {
	Content string `json:"content" validate:"required" msg:"ContentFormat"`
}

type SearchEsResp struct {
	VideoList []VideoInfo `json:"video_list"`
}

type UserInfo struct {
	Id              uint32 `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Gender          uint32 `json:"gender"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	FollowCount     uint32 `json:"follow_count"`
	FollowerCount   uint32 `json:"follower_count"`
	TotalFavorited  uint32 `json:"total_favorited"`
	WorkCount       uint32 `json:"work_count"`
	FavoriteCount   uint32 `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
	FriendCount     int64  `json:"friend_count"`
}

type UserVideoListReq struct {
	ToUid int `json:"to_user_id" validate:"required" msg:"ToUidFormat"`
}

type UserVideoListResp struct {
	VideoList []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	VideoId       int64    `json:"video_id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	StarCount     int64    `json:"star_count"`
	IsStar        bool     `json:"is_star"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
	CreateTime    string   `json:"create_time"`
	Duration      string   `json:"duration"`
}

type VideosListResp struct {
	Videos []VideoInfo `json:"video_list"`
}
