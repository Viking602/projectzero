package logic

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"projectzero/ent/user"
	"projectzero/internal/svc"
	"projectzero/internal/types"
	"projectzero/pkg/response"
	"projectzero/utils"
)

type UserLogic struct {
	svc *svc.Service
}

func NewUserLogic(svc *svc.Service) *UserLogic {
	return &UserLogic{
		svc: svc,
	}
}

func (l *UserLogic) RegisterUser(req *types.UserRegister) response.Response {

	// 校验用户是否存在
	exist, err := l.svc.Client.User.Query().Where(user.UserName(req.UserName)).Exist(context.Background())
	if err != nil {
		l.svc.Logger.Error("查询用户失败", zap.Error(err))
		return response.DBErr("数据库查询失败", err)
	}
	// 判断 如果存在则返回错误用户已存在
	if exist {
		return response.ParamErr("用户已存在", errors.New("用户已存在"))
	} else {
		// 否则创建新用户
		// 在这里您可以对注册请求进行验证和处理
		// 创建新用户
		password, err := utils.HashPassword(req.Password)
		if err != nil {
			return response.ParamErr("创建用户失败", err)
		}
		_, err = l.svc.Client.User.Create().SetUserName(req.UserName).SetNickName(req.UserName).SetPassword(password).
			Save(context.Background())
		if err != nil {
			l.svc.Logger.Error("创建用户失败", zap.Error(err))
			return response.ErrorResponse(err)
		}

	}
	return response.Success(nil)
}
