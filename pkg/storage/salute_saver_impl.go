package storage

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type saluteSaver struct {
	logger log.Logger
}

func NewSaluteSaver(logger log.Logger) SaluteSaver {
	return &saluteSaver{
		logger: logger,
	}
}

func (s saluteSaver) Save(salute string) {
	level.Info(s.logger).Log("msg", fmt.Sprintf("saving salute %s", salute))
}
