package svc

import (
	"go.uber.org/zap"
	"projectzero/conf"
	"projectzero/db"
	"projectzero/ent"
)

type Service struct {
	Logger *zap.Logger
	Conf   *conf.Conf
	Client *ent.Client
}

func NewServices(logger *zap.Logger, c *conf.Conf) *Service {
	client := db.Database(c.MySql.DSN)
	return &Service{
		Logger: logger,
		Conf:   c,
		Client: client,
	}

}
