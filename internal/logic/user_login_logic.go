package logic

import (
	"context"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/ent/user"
	"projectzero/internal/middleware"
	"projectzero/internal/types"
	"projectzero/pkg/response"
	"projectzero/utils"
)

type UserLoginLogic struct {
	client *ent.Client
	logger *zap.Logger
}

func NewUserLoginLogic(client *ent.Client, logger *zap.Logger) *UserLoginLogic {
	return &UserLoginLogic{
		client: client,
		logger: logger,
	}
}

func (l *UserLoginLogic) Login(req *types.UserRegister) response.Response {
	userInfo, err := l.client.User.Query().Where(user.UserName(req.UserName)).First(context.Background())
	if err != nil {
		l.logger.Error("查询用户失败", zap.Error(err))
		return response.ParamErr("用户名或密码错误", err)
	}

	if userInfo.Status == 0 {
		return response.ParamErr("用户已被禁用", nil)
	}

	if err := utils.ComparePassword(userInfo.Password, req.Password); err != nil {
		return response.ParamErr("用户名或密码错误", err)
	}

	token, err := middleware.GenToken(req.UserName, userInfo.Password)
	if err != nil {
		l.logger.Error("生成token失败", zap.Error(err))
		return response.ParamErr("登录失败", err)
	}

	return response.Success(types.UserLoginResponse{
		UserName: userInfo.UserName,
		Token:    token,
	})
}
