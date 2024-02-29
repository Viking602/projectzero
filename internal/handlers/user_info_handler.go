package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/internal/logic"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserInfoHandler struct {
	logic *logic.UserInfoLogic
}

func NewUserInfoHandler(l *logic.UserInfoLogic) *UserInfoHandler {
	return &UserInfoHandler{
		logic: l,
	}
}

func (l *UserInfoHandler) UserInfo(c *gin.Context) {
	var req types.UserInfoRequest

	if err := c.ShouldBind(&req); err == nil {
		res := l.logic.Info(&req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(err))
	}

}
