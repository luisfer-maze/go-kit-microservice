package service

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/luisfer-maze/go-kit-microservice/pkg/storage"
)

type service struct {
	storage storage.SaluteSaver
	logger  log.Logger
}

func NewService(storage storage.SaluteSaver, logger log.Logger) Service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s service) ProcessSalute(salute string) string {
	level.Info(s.logger).Log("msg", fmt.Sprintf("service processing for salute %s", salute))
	s.storage.Save(salute)
	return salute
}
