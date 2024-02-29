package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/internal/handlers"
	"projectzero/internal/logic"
	"projectzero/internal/middleware"
	"time"
)

func NewRouter(logger *zap.Logger, client *ent.Client) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.NoRoute(handlers.NoRouter)
	r.NoMethod(handlers.NoMethod)

	v1 := r.Group("/api/v1")
	{
		// 用户注册
		userLogic := logic.NewUserLogic(client, logger)
		userHandler := handlers.NewUserHandler(userLogic)
		v1.POST("/user/register", userHandler.UserRegister)
		// 用户登录
		userLoginLogic := logic.NewUserLoginLogic(client, logger)
		userLoginHandler := handlers.NewUserLoginHandler(userLoginLogic)
		v1.POST("/user/login", userLoginHandler.UserLogin)

		// 需要鉴权的路由
		auth := v1.Group("")
		auth.Use(middleware.JWTAuthMiddleware())
		{
			// 用户信息
			userInfoLogic := logic.NewUserInfoLogic(client, logger)
			userInfoHandler := handlers.NewUserInfoHandler(userInfoLogic)
			auth.GET("/user/info", userInfoHandler.UserInfo)

			// 更新用户信息
			userUpdateLogic := logic.NewUserUpdateLogic(client, logger)
			userUpdateHandler := handlers.NewUserUpdateHandler(userUpdateLogic)
			auth.POST("/user/update", userUpdateHandler.UserUpdate)

		}
	}

	return r
}
