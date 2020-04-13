package databasex

import (
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"time"
)

func (x *Database) First(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result map[string]interface{}, err error) {
	tag := "First"

	start := time.Now()

	result, err, closeErr := database.First(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("map's size: %d", len(result)), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
