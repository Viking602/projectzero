package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/internal/logic"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserLoginHandlers struct {
	logic *logic.UserLoginLogic
}

func NewUserLoginHandlers(l *logic.UserLoginLogic) *UserLoginHandlers {
	return &UserLoginHandlers{
		logic: l,
	}
}

func (l *UserLoginHandlers) UserLogin(c *gin.Context) {
	var req types.UserRegister

	if err := c.ShouldBind(&req); err == nil {
		res := l.logic.Login(&req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(err))
	}

}
