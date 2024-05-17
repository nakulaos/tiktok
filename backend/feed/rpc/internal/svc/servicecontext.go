package svc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/interceptors"
	favoriteModel "tiktok/favorite/favoritemodel"
	"tiktok/favorite/starmodel"
	"tiktok/feed/historymodel"
	"tiktok/feed/rpc/internal/config"
	videosModel "tiktok/feed/videosmodel"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config        config.Config
	VideosModel   videosModel.VideosModel
	KqJobPush     *kq.Pusher
	UserRpc       usesrv.UseSrv
	FavoriteModel favoriteModel.FavoritesModel
	StarModel     starmodel.StarsModel
	HistoryModel  historymodel.HistoryModel
	Es            *elasticsearch.TypedClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	esConf := elasticsearch.Config{
		Addresses: c.Es.Addresses,
	}

	esClient, err := elasticsearch.NewTypedClient(esConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		VideosModel:   videosModel.NewVideosModel(sqlConn, c.Cache),
		KqJobPush:     kq.NewPusher(c.KqJobPush.Brokers, c.KqJobPush.Topic, kq.WithAllowAutoTopicCreation()),
		UserRpc:       usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		StarModel:     starmodel.NewStarsModel(sqlConn, c.Cache),
		FavoriteModel: favoriteModel.NewFavoritesModel(sqlConn, c.Cache),
		HistoryModel:  historymodel.NewHistoryModel(sqlConn, c.Cache),
		Es:            esClient,
	}
}
