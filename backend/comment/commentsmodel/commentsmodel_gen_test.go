package commentsmodel

import (
	"context"
	"database/sql"
	"fmt"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"reflect"
	"testing"
)

func Test_defaultCommentsModel_Insert(t *testing.T) {
	type fields struct {
		CachedConn sqlc.CachedConn
		table      string
	}
	sqlconn := sqlx.NewMysql("root:asdasd@tcp(0.0.0.0:3306)/tiktok?charset=utf8&parseTime=True&loc=Local")
	redis, err := redis2.NewRedis(redis2.RedisConf{
		Host: "0.0.0.0:6379",
	})
	if err != nil {
		fmt.Println(err)
	}
	conn := sqlc.NewNodeConn(sqlconn, redis)
	type args struct {
		ctx     context.Context
		session sqlx.Session
		data    *Comments
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				CachedConn: conn,
				table:      "`comments`",
			},
			args: args{
				ctx:     context.Background(),
				session: nil,
				data: &Comments{
					Uid:     1,
					Vid:     1,
					Content: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &defaultCommentsModel{
				CachedConn: tt.fields.CachedConn,
				table:      tt.fields.table,
			}
			got, err := m.Insert(tt.args.ctx, tt.args.session, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
