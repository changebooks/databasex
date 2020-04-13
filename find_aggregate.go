package databasex

import (
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"time"
)

// "AS 'aggregate'" must be contained in Query
func (x *Database) AggregateInt(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result int64, err error) {
	tag := "AggregateInt"

	start := time.Now()

	result, err, closeErr := database.AggregateInt(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("result: %d", result), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

// "AS 'aggregate'" must be contained in Query
func (x *Database) AggregateFloat(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result float64, err error) {
	tag := "AggregateFloat"

	start := time.Now()

	result, err, closeErr := database.AggregateFloat(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("result: %f", result), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

// "AS 'aggregate'" must be contained in Query
func (x *Database) Aggregate(idRegister *log.IdRegister,
	driver *database.Driver, query string, args ...interface{}) (result interface{}, err error) {
	tag := "Aggregate"

	start := time.Now()

	result, err, closeErr := database.Aggregate(driver, query, args...)

	done := time.Now()
	remark := NewRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("result: %v", result), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
