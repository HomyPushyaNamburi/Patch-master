package db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateDbConnection() (db *sql.DB) {

	connString := fmt.Sprintf("server=%s;user ID=%s,password=%s,port=%d", mw_server, mw_user, mw_password, ms_port)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	return db

}
