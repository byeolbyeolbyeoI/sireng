package database

import (
	"database/sql"
	"fmt"
	dbConfig "github.com/chaaaeeee/internal/config/database"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf(dbConfig.User, dbConfig.Password, dbConfig.Protocol, dbConfig.Path, dbConfig.Name)
	db, err := sql.Open(dbConfig.Driver, dsn)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

type Database struct {
	db *sql.DB
}

func (db *Database) IsConnedted() bool {
	return false
}

func (db *Database) Reconnect() {
}
