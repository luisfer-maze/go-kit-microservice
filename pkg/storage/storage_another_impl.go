package storage

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type AnotherSaluteSaver struct {
	logger log.Logger
}

func NewAnotherStorage(logger log.Logger) SaluteSaver {
	return AnotherSaluteSaver{
		logger: logger,
	}
}

func (a AnotherSaluteSaver) Save(salute string) {
	level.Info(a.logger).Log("msg", fmt.Sprintf("saving salute with another saver %s", salute))
}
