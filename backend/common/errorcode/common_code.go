package errorcode

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"tiktok/common/i18n/constant"
)

var (
	Success            = New(0, constant.Success)
	RequestError       = New(100001, constant.RequestError)
	Unauthorized       = New(100002, constant.Unauthorized)
	NotFound           = New(100004, constant.NotFound)
	AccessDenied       = New(100003, constant.AccessDenied)
	MethodNotAllowed   = New(100005, constant.MethodNotAllowed)
	Canceled           = New(100006, constant.Canceled)
	ServerError        = New(100007, constant.ServerError)
	ServiceUnavailable = New(100008, constant.ServiceUnavailable)
	DeadlineExceeded   = New(100009, constant.DeadlineExceeded)
	LimitExceeded      = New(100010, constant.LimitExceeded)
	AlreadyExists      = New(100011, constant.AlreadyExists)
	DatabaseError      = New(100012, constant.DatabaseError)
	PermissionDeny     = New(100013, constant.PermissionDeny)
)

var (
	unknown = 100000
)

//// NewCodeCanceledError returns code Error with custom cancel error code
//func NewCodeCanceledError(msg string) error {
//	return &Errorcode{code: 100006, msg: msg}
//}

// NewCodeInvalidArgumentError returns code Error with custom invalid argument error code
func NewCodeInvalidArgumentError(msg string) error {
	return ErrorCode{code: 100001, msg: msg}
}

// NewCodeNotFoundError returns code Error with custom not found error code
func NewCodeNotFoundError(msg string) error {
	return &ErrorCode{code: 100004, msg: msg}
}

// NewCodeAlreadyExistsError returns code Error with custom already exists error code
func NewCodeAlreadyExistsError(msg string) error {
	return &ErrorCode{code: 100006, msg: msg}
}

// NewCodeAbortedError returns code Error with custom aborted error code
func NewCodeAbortedError(msg string) error {
	return &ErrorCode{code: 10, msg: msg}
}

// NewCodeUnavailableError returns code Error with custom unavailable error code
func NewCodeUnavailableError(msg string) error {
	return &ErrorCode{code: 14, msg: msg}
}

func NewDefaultError(msg string) error {
	return New(uint32(unknown), msg)
}

func NewRequestError(details []errdetails.BadRequest_FieldViolation) ErrorCode {
	RequestErr := RequestError
	for _, val := range details {
		WithDetails(RequestErr, val)
	}
	return RequestErr
}

// NewCodeInternalError returns code Error with custom internal error code
func NewCodeInternalError(details []errdetails.ErrorInfo) error {
	serverErr := ServerError
	for _, val := range details {
		WithDetails(serverErr, val)
	}
	return serverErr
}
