package common

import "errors"

var (
	NotFound   = errors.New("not found")
	BadRequest = errors.New("bad request")
	Internal   = errors.New("internal server error")

	SQLDBClosed = errors.New("sql: database is closed")
)
