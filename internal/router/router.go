package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projectzero/ent"
	"projectzero/internal/handlers"
	"projectzero/internal/logic"
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
		userLoginHandler := handlers.NewUserLoginHandlers(userLoginLogic)
		v1.POST("/user/login", userLoginHandler.UserLogin)
	}

	return r
}
