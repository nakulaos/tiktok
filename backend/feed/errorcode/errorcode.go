package errorcode

import "tiktok/common/errorcode"

var (
	VideoInsertError               = errorcode.New(500001, "video.VideoInsertError")
	FeedUnableToQueryUserError     = errorcode.New(500002, "video.FeedUnableToQueryUserError")
	FeedUnableToQueryVideoError    = errorcode.New(500003, "video.FeedUnableToQueryVideoError")
	FeedRecommendServiceInnerError = errorcode.New(500004, "video.FeedRecommendServiceInnerError")
	FeedDeleteVideoDbError         = errorcode.New(500005, "video.FeedDeleteVideoDbError")
)
