package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"projectzero/internal/handlers"
	"projectzero/internal/logic"
	"projectzero/internal/middleware"
	"projectzero/internal/svc"
	"time"
)

func NewRouter(s *svc.Service) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(s.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(s.Logger, true))

	r.NoRoute(handlers.NoRouter)
	r.NoMethod(handlers.NoMethod)

	v1 := r.Group("/api/v1")
	{
		// 用户注册
		userLogic := logic.NewUserLogic(s)
		userHandler := handlers.NewUserHandler(userLogic)
		v1.POST("/user/register", userHandler.UserRegister)
		// 用户登录
		userLoginLogic := logic.NewUserLoginLogic(s)
		userLoginHandler := handlers.NewUserLoginHandler(userLoginLogic)
		v1.POST("/user/login", userLoginHandler.UserLogin)

		// 需要鉴权的路由
		auth := v1.Group("")
		auth.Use(middleware.JWTAuthMiddleware(s.Conf.JWT.Secret))
		{
			// 用户信息
			userInfoLogic := logic.NewUserInfoLogic(s)
			userInfoHandler := handlers.NewUserInfoHandler(userInfoLogic)
			auth.GET("/user/info", userInfoHandler.UserInfo)

			// 更新用户信息
			userUpdateLogic := logic.NewUserUpdateLogic(s)
			userUpdateHandler := handlers.NewUserUpdateHandler(userUpdateLogic)
			auth.POST("/user/update", userUpdateHandler.UserUpdate)

		}
	}

	return r
}
