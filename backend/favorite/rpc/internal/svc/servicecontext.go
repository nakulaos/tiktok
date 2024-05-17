package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/interceptors"
	favoriteModel "tiktok/favorite/favoritemodel"
	"tiktok/favorite/rpc/internal/config"
	"tiktok/favorite/starmodel"
	videosModel "tiktok/feed/videosmodel"
	userModel "tiktok/user/model"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     userModel.UserModel
	VideosModel   videosModel.VideosModel
	FavoriteModel favoriteModel.FavoritesModel
	StarModel     starmodel.StarsModel
	// BizRedis 业务的redis
	BizRedis *redis.Redis
	UserRpc  usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     userModel.NewUserModel(sqlConn, c.Cache),
		VideosModel:   videosModel.NewVideosModel(sqlConn, c.Cache),
		StarModel:     starmodel.NewStarsModel(sqlConn, c.Cache),
		BizRedis:      redis.MustNewRedis(c.BizRedis),
		FavoriteModel: favoriteModel.NewFavoritesModel(sqlConn, c.Cache),
		UserRpc:       usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
