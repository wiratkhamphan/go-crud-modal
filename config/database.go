package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// conn database go_crud_modal
func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_crud_modal"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
