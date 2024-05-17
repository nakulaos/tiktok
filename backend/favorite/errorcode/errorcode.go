package errorcode

import "tiktok/common/errorcode"

var (
	FavoriteUserIdEmptyError       = errorcode.New(400001, "favorite.FavoriteUserIdEmptyError")
	FavoriteVideoIdEmptyError      = errorcode.New(400002, "favorite.FavoriteVideoIdEmptyError")
	FavoriteServiceDuplicateError  = errorcode.New(400003, "favorite.FavoriteServiceDuplicateError")
	FavoriteServiceCancelError     = errorcode.New(400004, "favorite.FavoriteServiceCancelError")
	FavoriteInvalidActionTypeError = errorcode.New(400005, "favorite.FavoriteInvalidActionTypeError")
	FavoriteLimitError             = errorcode.New(400006, "favorite.FavoriteLimitError")
	StarUserIdEmptyError           = errorcode.New(400007, "favorite.StarUserIdEmptyError")
	StarVideoIdEmptyError          = errorcode.New(400008, "favorite.StarVideoIdEmptyError")
	StarServiceDuplicateError      = errorcode.New(400009, "favorite.StarServiceDuplicateError")
	StarServiceCancelError         = errorcode.New(400010, "favorite.StarServiceCancelError")
	StarInvalidActionTypeError     = errorcode.New(400011, "favorite.StarInvalidActionTypeError")
	StarLimitError                 = errorcode.New(400012, "favorite.StarLimitError")
)
