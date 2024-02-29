package logic

import (
	"context"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/ent/user"
	"projectzero/internal/types"
	"projectzero/pkg/response"
	"projectzero/utils"
	"time"
)

type UserUpdateLogic struct {
	client *ent.Client
	logger *zap.Logger
}

func NewUserUpdateLogic(client *ent.Client, logger *zap.Logger) *UserUpdateLogic {
	return &UserUpdateLogic{
		client: client,
		logger: logger,
	}
}

func (l *UserUpdateLogic) Update(req *types.UserUpdateRequest) response.Response {
	// 获取请求用户信息
	nickname := req.NickName
	password := req.Password
	userType := req.UserType
	// 查询用户基础信息
	userInfo, err := l.client.User.Query().Where(user.UserName(req.UserName)).First(context.Background())
	if err != nil {
		l.logger.Error("查询用户失败", zap.Error(err))
		return response.ParamErr("查询用户失败", err)
	}
	// 为空则不修改
	if nickname == "" {
		nickname = userInfo.NickName
	} else {
		nickname = req.NickName
	}
	if req.Password == "" {
		password = userInfo.Password
	} else {
		// 如果不为空则加密密码
		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			l.logger.Error("密码加密失败", zap.Error(err))
			return response.ParamErr("修改密码失败", err)
		}
		password = hashPassword
	}
	if userType == 0 {
		userType = userInfo.UserType
	}

	_, err = l.client.User.Update().Where(user.UserName(req.UserName)).
		SetNickName(req.NickName).
		SetPassword(password).
		SetUpdateAt(time.Now().Unix()).
		Save(context.Background())
	if err != nil {
		l.logger.Error("更新用户信息失败", zap.Error(err))
		return response.ParamErr("更新用户信息失败", err)
	}

	return response.Success(nil)
}
