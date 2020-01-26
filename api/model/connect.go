package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root@/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected To Db")
	con = db
	return db

}
