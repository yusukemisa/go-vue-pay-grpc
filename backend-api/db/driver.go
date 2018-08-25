package db

import (
	"database/sql"
	"log"
	"os"
)

//Conn - sql connection handler
var Conn *sql.DB

//NewSQLHandler - init sql handler
func init() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		log.Fatal(err)
	}
	Conn = conn
}
