package logic

import (
	"context"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/ent/user"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserInfoLogic struct {
	client *ent.Client
	logger *zap.Logger
}

func NewUserInfoLogic(client *ent.Client, logger *zap.Logger) *UserInfoLogic {
	return &UserInfoLogic{
		client: client,
		logger: logger,
	}
}

func (l *UserInfoLogic) Info(req *types.UserInfoRequest) response.Response {
	userInfo, err := l.client.User.Query().Where(user.UserName(req.UserName)).First(context.Background())
	if err != nil {
		l.logger.Error("查询用户失败", zap.Error(err))
		return response.ParamErr("查询用户失败", err)
	}

	return response.Success(types.UserInfoResponse{
		Id:       userInfo.ID,
		UserName: userInfo.UserName,
		NickName: userInfo.NickName,
		UserType: userInfo.UserType,
		Status:   userInfo.Status,
		DeleteAt: userInfo.DeleteAt,
		CreateAt: userInfo.CreateAt,
		UpdateAt: userInfo.UpdateAt,
	})
}
