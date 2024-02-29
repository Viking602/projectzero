package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/pkg/response"
)

func NoRouter(c *gin.Context) {
	c.JSON(http.StatusNotFound,
		response.Response{
			Code: http.StatusNotFound,
			Msg:  "错误的资源路径",
		})
}

func NoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed,
		response.Response{
			Code: http.StatusMethodNotAllowed,
			Msg:  "错误的请求方式",
		})

}
