package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func createRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("Going to insert record in database for name:", name[0])
		stmt, err := db.Prepare("INSERT employee SET name=?")
		if err != nil {
			log.Print("error preparing query :: ", err)
			return
		}
		result, err := stmt.Exec(name[0])
		if err != nil {
			log.Print("error executing query :: ", err)
			return
		}
		id, err := result.LastInsertId()
		fmt.Fprintf(w, "Last Inserted Record Id is :: %s", strconv.FormatInt(id, 10))
	} else {
		fmt.Fprintf(w, "Error occurred while creating record in database for name :: %s", name[0])
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/create", createRecord).Methods("POST")
	defer db.Close()
	err := http.ListenAndServe(HOST+":"+PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
