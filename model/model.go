package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Loading driver
)

var (
	// DBCon Shared DB object
	db *sql.DB
)

// GetConnection return a connection for working on
func GetConnection() (*sql.DB, error) {
	if db == nil || db.Ping() == nil {
		dbt, err := sql.Open("sqlite3", "todos.db")
		db = dbt
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)
		return db, err
	}
	return db, nil
}
