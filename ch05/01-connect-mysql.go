package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST             = "localhost"
	PORT             = "8080"
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:rootuser@/mydb"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("Error connecting to database :: ", connectionError)
	}
}

func getCurrentDb(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATABASE() as db")
	if err != nil {
		log.Fatal("Error querying ::", err)
		return
	}
	var db string
	for rows.Next() {
		rows.Scan(&db)
	}
	fmt.Fprintf(w, "Current Database is :: %s", db)
}

func main() {
	http.HandleFunc("/", getCurrentDb)
	defer db.Close()
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
