package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/internal/logic"
	"projectzero/internal/types"
	"projectzero/pkg/response"
)

type UserUpdateHandler struct {
	logic *logic.UserUpdateLogic
}

func NewUserUpdateHandler(l *logic.UserUpdateLogic) *UserUpdateHandler {
	return &UserUpdateHandler{
		logic: l,
	}
}

func (l *UserUpdateHandler) UserUpdate(c *gin.Context) {
	var req types.UserUpdateRequest

	if err := c.ShouldBind(&req); err == nil {
		user, _ := c.Get("username")
		res := l.logic.Update(user, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(err))
	}

}
