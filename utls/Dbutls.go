package utls

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Dbconnection *sql.DB

func CreateDbConnection() (*sql.DB, bool) {

	err := Dbconnection.Ping()
	if err != nil {
		log.Println("Error in Ping", err)
	}
	_, isOpen := OpenDbConnection()

	if isOpen {
		return Dbconnection, true
	} else {
		return Dbconnection, false
	}
}

func OpenDbConnection() (*sql.DB, bool) {
	//CnnectionString := Config.ConnectionString
	//CnnectionString = "postgresql://postgres@localhost:5432/postgres?sslmode=disable"
	myDb, err := sql.Open("postgres", Config.ConnectionString)
	if err != nil {
		fmt.Println("Db Open Failed:", err)
	}
	myDb.SetMaxOpenConns(0)
	myDb.SetMaxIdleConns(100)
	x := myDb
	return x, true
}

func init() {
	log.Println("Connecting to POSTGRES DB...")
	OpenDbConnection()
}
