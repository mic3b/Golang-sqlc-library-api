package initializer

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

func InitDb() *sql.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = sql.Open("sqlite3", "db/database.sqlite3")
		if err != nil {
			panic(err)
		}
	})
	return dbInstance
}
