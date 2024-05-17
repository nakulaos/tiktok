package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/apimachinery/pkg/util/json"
	"strconv"
	"tiktok/common/crypt"
	"tiktok/common/errorcode"
	"tiktok/common/gorse"
	"tiktok/common/utils"
	userErrorcode "tiktok/user/errorcode"
	"tiktok/user/model"
	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	//m := make(map[string]string)
	//m["reason"] = "测试"
	//detail := errdetails.ErrorInfo{Metadata: m}
	//return &user.RegisterResponse{}, errorcode.WithDetails(errorcode.AlreadyExists, &detail)
	if in.Phone != nil {
		_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, *in.Phone)
		if err == nil {
			return nil, errors.Wrapf(userErrorcode.PhoneExistError, "phone: %s", *in.Phone)
		}
	}

	if in.Email != nil {
		_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, *in.Email)
		if err == nil {
			return nil, errors.Wrapf(userErrorcode.EmailExistError, "email: %s", *in.Email)
		}
	}

	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserName)
	if err == nil {
		return nil, errors.Wrapf(userErrorcode.UserExistError, "username: %s", in.UserName)
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Username: in.UserName,
			Email:    *in.Email,
			Gender:   in.Gender,
			Role:     "normalUser",
			Status:   0,
			Mobile:   *in.Phone,
			Password: crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Dec:      in.Dec,
			Avatar:   in.Avatar,
			DelState: 0,
			Version:  0,
			BackgroundUrl: sql.NullString{
				String: in.BackgroundImage,
				Valid:  true,
			},
		}

		sqlRes, err := l.svcCtx.UserModel.Insert(l.ctx, nil, &newUser)
		if err != nil {
			return nil, errors.Wrapf(errorcode.DatabaseError, "insert user error: %v", err)
		}
		userId, _ := sqlRes.LastInsertId()
		postbaseUrl := l.svcCtx.Config.RecommendUrl + "/api/user"
		userReq := gorse.UserGoresBody{UserId: fmt.Sprintf("%d", userId)}
		jsonData, err := json.Marshal(userReq)
		if err != nil {
			l.Logger.Errorf("JSON marshal failed: %v", err)
			return nil, err
		}

		_, err = utils.Post(postbaseUrl, jsonData)
		if err != nil {
			l.Logger.Errorf("Post failed: %v", err)
		}

		// TODO: casbin
		if ok, err := l.svcCtx.Cbn.AddRoleForUser(strconv.FormatInt(userId, 10), "normalUser"); !ok {
			return nil, errors.Wrapf(errorcode.ServerError, "add policy failed: %v", err)
		}
		return nil, nil

	}

	return nil, errors.Wrapf(errorcode.DatabaseError, "detail: %v", err)

}
