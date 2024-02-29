package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projectzero/conf"
	"projectzero/internal/svc"
)

func main() {
	gin.SetMode(gin.DebugMode)
	logger, _ := zap.NewProduction()
	conf.Init()
	r := svc.NewServices(logger)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
