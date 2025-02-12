// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id             uint32 `json:"id"`
	Name           string `json:"name"`
	Gender         uint32 `json:"gender"`
	Avatar         string `json:"avatar"`
	Signature      string `json:"signature"`
	FollowCount    uint32 `json:"follow_count"`
	FollowerCount  uint32 `json:"follower_count"`
	TotalFavorited uint32 `json:"total_favorited"`
	WorkCount      uint32 `json:"work_count"`
	FavoriteCount  uint32 `json:"favorite_count"`
	IsFollow       bool   `json:"is_follow"`
	FriendCount    int64  `json:"friend_count"`
}

type VideoInfo struct {
	VideoId       int64  `json:"video_id"`
	User          User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	StarCount     int64  `json:"star_count"`
	IsStar        bool   `json:"is_star"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	CreateTime    string `json:"create_time"`
	Duration      string `json:"duration"`
}

type ActionReq struct {
	VideoId    int64 `json:"video_id" validate:"required,min=1" msg:"VideoIdFormat"`
	ActionType int32 `json:"action_type" validate:"required,min=1,max=10" msg:"ActionTypeFormat"`
}

type ListReq struct {
	UserId int64 ` validate:"required,min=1" msg:"UserIdFormat" form:"to_user_id"` // 用户id
}

type ListResp struct {
	VideoList []VideoInfo `json:"video_list"`
}
