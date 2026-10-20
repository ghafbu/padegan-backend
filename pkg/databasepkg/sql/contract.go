package sql

import "database/sql"

type DB interface {
	Conn() *sql.DB
}
