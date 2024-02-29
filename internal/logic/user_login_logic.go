package logic

import (
	"context"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/ent/user"
	"projectzero/internal/types"
	"projectzero/pkg/response"
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
		return response.DBErr("数据查询失败", err)
	}

	return response.Success(types.UserLoginResponse{
		UserName: userInfo.UserName,
		Token:    "asd",
	})
}
