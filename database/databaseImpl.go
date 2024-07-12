package database

import (
	"database/sql"
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type databaseImpl struct {
	db *sql.DB
}

var (
	once       sync.Once
	dbInstance *databaseImpl
)

// interface implementation (returning struct to impl)
func NewDatabase(conf *config.Config) Database {
	// idk if this is necessary (no cope)
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@%s(%s)/%s",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Protocol,
			conf.Database.Path,
			conf.Database.DBName)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic("Failed to connect to the database")
		}

		dbInstance = &databaseImpl{db: db}
	})

	return dbInstance

}

func (m *databaseImpl) GetDb() *sql.DB {
	return dbInstance.db
}
