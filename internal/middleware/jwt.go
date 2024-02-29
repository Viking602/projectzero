package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectzero/pkg/response"
	"projectzero/utils"
	"time"
)

// JWT 自定义中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = http.StatusOK
		token := c.Query("token")
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			// 解析token
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 410
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 410
			}
		}
		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, response.Response{
				Code: code,
				Msg:  "登录失败",
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
