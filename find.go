package databasex

import (
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"time"
)

func (x *Database) Find(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result []map[string]interface{}, err error) {
	tag := "Find"

	start := time.Now()

	result, err, closeErr := database.Find(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("result's rows: %d", len(result)), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
