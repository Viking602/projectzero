package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/internal/logic"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserLoginHandler struct {
	logic *logic.UserLoginLogic
}

func NewUserLoginHandler(l *logic.UserLoginLogic) *UserLoginHandler {
	return &UserLoginHandler{
		logic: l,
	}
}

func (l *UserLoginHandler) UserLogin(c *gin.Context) {
	var req types.UserRegister

	if err := c.ShouldBind(&req); err == nil {
		res := l.logic.Login(&req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(err))
	}

}
