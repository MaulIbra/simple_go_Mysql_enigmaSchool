package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(driver, dataSource string) (*sql.DB, error) {
	db, _ := sql.Open(driver, dataSource)

	if err := db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
