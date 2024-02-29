package svc

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projectzero/db"
	"projectzero/internal/router"
)

type Service struct {
	r      *gin.Engine
	logger *zap.Logger
}

func NewServices(logger *zap.Logger) *Service {
	client := db.Database()
	r := router.NewRouter(logger, client)
	return &Service{
		r:      r,
		logger: logger,
	}

}

func (s Service) Run() error {
	err := s.r.Run(":3000")
	return err
}
