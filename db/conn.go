package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "102030"
	dbname   = "db_agenda"
)

func ConnectDB() (*sql.DB, error) {
	/*	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)*/
	psqlinfo := "root:102030@tcp(127.0.0.1:3306)/db_agenda"
	db, err := sql.Open("mysql", psqlinfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to " + dbname)
	return db, nil
}
