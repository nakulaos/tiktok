package errorcode

import "tiktok/common/errorcode"

var (
	CommentUserIdEmptyError       = errorcode.New(600001, "comment.CommentUserIdEmptyError")
	CommentVideoIdEmptyError      = errorcode.New(600002, "comment.CommentVideoIdEmptyError")
	CommentInvalidActionTypeError = errorcode.New(600003, "comment.CommentInvalidActionTypeError")
	CommentNotExistError          = errorcode.New(600004, "comment.CommentNotExistError")
	DanMuUserIdEmptyError         = errorcode.New(600005, "comment.DanMuUserIdEmptyError")
	DanMuVideoIdEmptyError        = errorcode.New(600006, "comment.DanMuVideoIdEmptyError")
	DanMuContentEmptyError        = errorcode.New(600007, "comment.DanMuContentEmptyError")
	DanMuLimitError               = errorcode.New(600008, "comment.DanMuLimitError")
	CommentIsEmptyError           = errorcode.New(600009, "comment.CommentIsEmptyError")
	CommentIsUnSafeError          = errorcode.New(600010, "comment.CommentIsUnSafeError")
)
