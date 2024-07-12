package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	GetDb() *sql.DB
}
