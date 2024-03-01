package logic

import (
	"context"
	"go.uber.org/zap"
	"projectzero/ent/user"
	"projectzero/internal/svc"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserInfoLogic struct {
	svc *svc.Service
}

func NewUserInfoLogic(svc *svc.Service) *UserInfoLogic {
	return &UserInfoLogic{
		svc: svc,
	}
}

func (l *UserInfoLogic) Info(req *types.UserInfoRequest) response.Response {
	userInfo, err := l.svc.Client.User.Query().Where(user.UserName(req.UserName)).First(context.Background())
	if err != nil {
		l.svc.Logger.Error("查询用户失败", zap.Error(err))
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
