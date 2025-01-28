package entities

import "database/sql"

type DbQuery interface {
	RunQuery(query string, values ...interface{}) (sql.Result, error)
	GetDBData(query string, values ...interface{}) (sql.Result, error)
}
