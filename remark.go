package databasex

import (
	"fmt"
	"github.com/changebooks/database"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Remark struct {
	Dsn     string        `json:"dsn"`
	Query   string        `json:"query"`
	Args    interface{}   `json:"args"`
	Command string        `json:"command"`
	Total   time.Duration `json:"total"`
	Start   time.Time     `json:"start"`
	Done    time.Time     `json:"done"`
}

func NewRemark(driver *database.Driver, start time.Time, done time.Time, query string, args ...interface{}) *Remark {
	dsn := ""
	if driver != nil {
		dsn = driver.GetDsn()
	}

	command := ReplacePlaceholder(query, args...)
	total := done.Sub(start)

	return &Remark{
		Dsn:     dsn,
		Query:   query,
		Args:    args,
		Command: command,
		Total:   total,
		Start:   start,
		Done:    done,
	}
}

func ReplacePlaceholder(query string, args ...interface{}) string {
	if query == "" {
		return ""
	}

	if len(args) == 0 {
		return query
	}

	for _, value := range args {
		if !strings.Contains(query, PlaceHolder) {
			break
		}

		s := ParseStr(value)
		query = strings.Replace(query, PlaceHolder, s, 1)
	}

	return query
}

func ParseStr(value interface{}) string {
	if value == nil {
		return QuoteStr("")
	}

	switch r := value.(type) {
	case string:
		return QuoteStr(r)
	case []byte:
		return QuoteStr(string(r))
	}

	r := reflect.ValueOf(value)
	switch r.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(r.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(r.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(r.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(r.Float(), 'g', -1, 32)
	case reflect.Bool:
		return QuoteStr(strconv.FormatBool(r.Bool()))
	}

	return QuoteStr(fmt.Sprintf("%v", value))
}

func QuoteStr(s string) string {
	return "'" + s + "'"
}
