/**
 ******************************************************************************
 * @file           : server.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/4/8
 ******************************************************************************
 */

package interceptors

import (
	"context"
	"github.com/dtm-labs/dtmcli"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok/common/errorcode"
)

func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		if errors.Is(err, status.Error(codes.Aborted, dtmcli.ResultFailure)) {
			// 分布式事务错误不转化
			return resp, err
		} else {
			causeErr := errors.Cause(err)
			// log error
			logx.WithContext(ctx).Errorf("[RPC-SRV-ERR] %+v", err)
			return resp, errorcode.GrcpStatusFromErrorCode(causeErr).Err()
		}

	}
	return resp, nil
}
