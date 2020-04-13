package databasex

import (
	"database/sql"
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"time"
)

func (x *Database) Exec(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result sql.Result, err error) {
	tag := "Exec"

	start := time.Now()

	result, err = database.Exec(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if err == nil {
		affectedRows, affectedRowsErr := result.RowsAffected()
		if affectedRowsErr != nil {
			x.logger.E(tag, AffectedRowsFailure, remark, affectedRowsErr, "", idRegister)
		}

		x.logger.I(tag, fmt.Sprintf("affected's rows: %d", affectedRows), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
