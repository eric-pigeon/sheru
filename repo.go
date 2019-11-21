package sheru

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	db      *sql.DB
	adapter adapter
}

func Open(args ...interface{}) (db *DB, err error) {
	driver := "postgres"
	source := args[1].(string)
	var dbSQL *sql.DB
	dbSQL, err = sql.Open(driver, source)

	db = &DB{
		db: dbSQL,
	}
	if err != nil {
		return
	}

	if err = db.db.Ping(); err != nil {
		db.db.Close()
	}
	return
}

type OnConflict struct {
}

func (s *DB) Insert(changeset Changeset) {
	fmt.Printf("INSERT INTO %s", changeset.metadata.quotedTableName())
}
