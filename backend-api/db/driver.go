package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Conn - sql connection handler
var Conn *sql.DB

//NewSQLHandler - init sql handler
func init() {
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/items")
	if err != nil {
		log.Fatal(err)
	}
	Conn = conn
}
