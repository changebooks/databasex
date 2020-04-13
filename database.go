package databasex

import (
	"errors"
	"github.com/changebooks/log"
)

type Database struct {
	logger *log.Logger
}

func New(logger *log.Logger) (*Database, error) {
	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}

	return &Database{logger: logger}, nil
}

func (x *Database) GetLogger() *log.Logger {
	return x.logger
}
