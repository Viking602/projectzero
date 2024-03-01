package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projectzero/conf"
	"projectzero/internal/router"
	"projectzero/internal/svc"
)

func main() {
	logger, _ := zap.NewProduction()

	var c conf.Conf

	cfg := conf.LoadConfig("etc/config.yaml", c)

	gin.SetMode(cfg.Env)
	service := svc.NewServices(logger, cfg)
	r := router.NewRouter(service)
	err := r.Run(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		panic(err)
		return
	}

}
