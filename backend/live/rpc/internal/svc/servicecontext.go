package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/interceptors"
	"tiktok/live/rpc/internal/config"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
