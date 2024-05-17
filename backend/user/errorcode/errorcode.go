package errorcode

import "tiktok/common/errorcode"

var (
	UserNotExistError    = errorcode.New(200001, "user.UserNotExistError")
	UserExistError       = errorcode.New(200002, "user.UserExistError")
	UserPasswordError    = errorcode.New(200003, "user.UserPasswordError")
	UserUnExistError     = errorcode.New(200004, "user.UserUnExistError")
	PhoneExistError      = errorcode.New(200005, "user.PhoneExistError")
	EmailExistError      = errorcode.New(200006, "user.EmailExistError")
	PassWordError        = errorcode.New(200007, "user.PassWordError")
	UserUploadVideoError = errorcode.New(200008, "user.UserUploadVideoError")
)
