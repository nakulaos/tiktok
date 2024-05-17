package errorcode

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/protoadapt"
)

type ErrorCode struct {
	code    uint32
	msg     string
	details []any
}

// 业务错误码
func (e ErrorCode) Error() string {
	return e.msg
}

func (e ErrorCode) Code() uint32 {
	return e.code
}

func (e ErrorCode) Details() []any {
	return e.details
}
func (e ErrorCode) Message() string {
	return e.msg
}

func New(code uint32, msg string) ErrorCode {
	return ErrorCode{code: code, msg: msg}
}

func WithDetails(err ErrorCode, detail any) ErrorCode {
	err.details = append(err.details, detail)
	return err
}

func IsErrorCode(err error) bool {
	_, ok := err.(*ErrorCode)
	return ok
}

func GrpcStatusToErrorCode(grpcStatus *status.Status) ErrorCode {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return Success
	case codes.InvalidArgument:
		details := grpcStatus.Details()

		RequestError = WithDetails(RequestError, details)

	case codes.NotFound:
		NotFound = WithDetails(NotFound, grpcStatus.Details())
		return NotFound
	case codes.PermissionDenied:
		AccessDenied = WithDetails(AccessDenied, grpcStatus.Details())
		return AccessDenied
	case codes.Unauthenticated:
		Unauthorized = WithDetails(Unauthorized, grpcStatus.Details())
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceeded
	case codes.Unimplemented:
		return ErrorCode{
			code:    uint32(grpcStatus.Code()),
			msg:     grpcStatus.Message(),
			details: grpcStatus.Details(),
		}
	case codes.DeadlineExceeded:
		return DeadlineExceeded
	case codes.Unavailable:
		return ServiceUnavailable
	default:
		// 业务错误码
		return ErrorCode{
			code:    uint32(grpcStatus.Code()),
			msg:     grpcStatus.Message(),
			details: grpcStatus.Details(),
		}

	}

	return ServerError
}

func GrcpStatusFromErrorCode(err error) *status.Status {
	err = errors.Cause(err)
	if errcode, ok := err.(ErrorCode); ok {
		st := status.New(codes.Code(errcode.Code()), errcode.Message())
		if errcode.Details() != nil {
			details := errcode.Details()
			for i := len(details) - 1; i >= 0; i-- {
				detail := details[i]
				if detailsVal, ok := detail.(protoadapt.MessageV1); ok {
					st, _ = st.WithDetails(detailsVal)
					return st
				}
			}
		}
		return st
	}

	var grpcStatus *status.Status
	switch err {
	case context.Canceled:
		return status.New(codes.Code(Canceled.Code()), Canceled.Message())
	case context.DeadlineExceeded:
		return status.New(codes.Code(DeadlineExceeded.Code()), DeadlineExceeded.Message())

	default:
		grpcStatus, _ = status.FromError(err)
		return grpcStatus
	}

}
