package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/internal/logic"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserRegisterHandler struct {
	logic *logic.UserLogic
}

func NewUserHandler(l *logic.UserLogic) *UserRegisterHandler {
	return &UserRegisterHandler{
		logic: l,
	}
}

func (l *UserRegisterHandler) UserRegister(c *gin.Context) {
	var req types.UserRegister

	if err := c.ShouldBind(&req); err == nil {
		res := l.logic.RegisterUser(&req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(err))
	}

}
