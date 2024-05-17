package logic

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
	"path/filepath"
	"strings"
	"tiktok/common/crypt"
	"tiktok/user/errorcode"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUploadLogic {
	return &UserUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUploadLogic) UserUpload(req *http.Request) (resp *types.UserUploadResponse, err error) {
	// 上传视频接口
	accessKey := l.svcCtx.Config.QiNiu.AccessKey
	secretKey := l.svcCtx.Config.QiNiu.SecretKey
	bucket := l.svcCtx.Config.QiNiu.Bucket
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()

	fileURL := ""
	file, handler, err := req.FormFile("file")

	filename := crypt.PasswordEncrypt(time.Now().String(), handler.Filename)
	//key := filename + ".mp4"
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%d__%s__%s.mp4", l.svcCtx.Config.QiNiu.Prefix, userId, now, filename)
	seveas := filename + "-1.mp4"
	seveasstr := bucket + ":" + seveas
	saceas := base64.StdEncoding.EncodeToString([]byte(seveasstr))
	mac := qbox.NewMac(accessKey, secretKey)

	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		PersistentOps: "avthumb/mp4/wmImage/aHR0cDovL3FueS5oYWxsbmFrdWxhb3MuY24vdGlrdG9rL3RleHQtaW1hZ2UlMjAlMjgzJTI5LmpwZw==/wmGravity/NorthWest|saveas/" + saceas,
		//PersistentNotifyURL: "http://fake.com/qiniu/notify",
	}
	upToken := putPolicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.Zone_z2,
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}

	resumeUploader := storage.NewResumeUploaderV2(&cfg)

	ret := storage.PutRet{}           // 上传后返回的结果
	putExtra := storage.RputV2Extra{} // 额外参数

	// 上传视频
	err = resumeUploader.Put(context.Background(), &ret, upToken, key, file, handler.Size, &putExtra)
	if err != nil {
		return nil, errors.Wrapf(errorcode.UserUploadVideoError, "upload video error: %v", err)

	}

	// 截取视频第一帧
	operationManager := storage.NewOperationManager(mac, &cfg)
	fopVframe := fmt.Sprintf("vframe/jpg/offset/1|saveas/%s",
		storage.EncodedEntry(bucket, strings.TrimSuffix(key, filepath.Ext(key))+".jpg"))
	fops := fopVframe
	_, err = operationManager.Pfop(bucket, key, fops, "", "", true)
	if err != nil {
		return nil, errors.Wrapf(errorcode.UserUploadVideoError, "upload video error: %v", err)
	}

	baseURL := l.svcCtx.Config.QiNiu.Cdn
	fileURL = baseURL + "/" + key
	coverUrl := baseURL + "/" + strings.TrimSuffix(key, filepath.Ext(key)) + ".jpg"
	l.Logger.Infof("upload video success, fileURL: %s,coverURL:%s,perisitentID: %s", fileURL, coverUrl, ret.PersistentID)

	return &types.UserUploadResponse{
		CoverUrl: coverUrl,
		Url:      fileURL,
	}, nil

}
