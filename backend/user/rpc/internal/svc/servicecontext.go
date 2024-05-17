package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiktok/common/utils"
	favoriteModel "tiktok/favorite/favoritemodel"
	videosModel "tiktok/feed/videosmodel"
	relationModel "tiktok/relation/model"
	"tiktok/user/model"
	"tiktok/user/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	Cbn            *casbin.Enforcer
	RelationsModel relationModel.RelationsModel
	VideosModel    videosModel.VideosModel
	FavorModel     favoriteModel.FavoritesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	cbn := utils.NewCasbin(c.DB.DataSource, c.Casbin.Dir, c.Casbin.Table)

	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlConn, c.Cache),
		Cbn:            cbn,
		RelationsModel: relationModel.NewRelationsModel(sqlConn),
		VideosModel:    videosModel.NewVideosModel(sqlConn, c.Cache),
		FavorModel:     favoriteModel.NewFavoritesModel(sqlConn, c.Cache),
	}
}
