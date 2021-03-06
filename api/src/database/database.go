package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Open the coonection with database
func Connect() (*sql.DB, error) {

	db, error := sql.Open("mysql", config.URL_DB)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil

}
