package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/interceptors"
	videosModel "tiktok/feed/videosmodel"
	relationModel "tiktok/relation/model"
	"tiktok/relation/rpc/internal/config"
	userModel "tiktok/user/model"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel relationModel.RelationsModel
	UserRpc       usesrv.UseSrv
	UserModel     userModel.UserModel
	VideosModel   videosModel.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		RelationModel: relationModel.NewRelationsModel(sqlConn),
		UserRpc:       usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		UserModel:     userModel.NewUserModel(sqlConn, c.Cache),
		VideosModel:   videosModel.NewVideosModel(sqlConn, c.Cache),
	}
}
