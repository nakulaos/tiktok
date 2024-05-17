package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"tiktok/common/errorcode"
)

func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			grpcStatus, _ := status.FromError(err)
			err = errorcode.GrpcStatusToErrorCode(grpcStatus)
		}
		return err
	}
}
