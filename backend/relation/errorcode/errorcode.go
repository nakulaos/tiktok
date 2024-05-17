package errorcode

import "tiktok/common/errorcode"

var (
	LikeParameterError                     = errorcode.New(300001, "relation.LikeParameterError")
	RelationUserNotExistError              = errorcode.New(300002, "relation.RelationUserNotExistError")
	RelationUnableFavorSelfError           = errorcode.New(300003, "relation.RelationUnableFavorSelfError")
	RelationUnableFavorMoreError           = errorcode.New(300004, "relation.RelationUnableFavorMoreError")
	RelationUnableUnFavorNotFavorUserError = errorcode.New(300005, "relation.RelationUnableUnFavorNotFavorUserError")
)
