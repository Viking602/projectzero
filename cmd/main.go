package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"projectzero/conf"
	"projectzero/internal/svc"
)

func main() {
	godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))
	logger, _ := zap.NewProduction()
	conf.Init()
	r := svc.NewServices(logger)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
