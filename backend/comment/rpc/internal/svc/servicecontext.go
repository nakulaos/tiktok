package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/comment/commentsmodel"
	"tiktok/comment/danmumodel"
	"tiktok/comment/rpc/internal/config"
	"tiktok/common/interceptors"
	videosModel "tiktok/feed/videosmodel"
	userModel "tiktok/user/model"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config        config.Config
	CommentsModel commentsmodel.CommentsModel
	DanmuModel    danmumodel.DanmuModel
	UserModel     userModel.UserModel
	VideosModel   videosModel.VideosModel
	SqlConn       sqlx.SqlConn
	UserRpc       usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		CommentsModel: commentsmodel.NewCommentsModel(sqlConn, c.Cache),
		DanmuModel:    danmumodel.NewDanmuModel(sqlConn, c.Cache),
		UserModel:     userModel.NewUserModel(sqlConn, c.Cache),
		VideosModel:   videosModel.NewVideosModel(sqlConn, c.Cache),
		SqlConn:       sqlConn,
		UserRpc:       usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
