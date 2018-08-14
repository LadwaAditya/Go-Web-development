package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Its a post request"))
}

func PathVariableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("Hello" + name))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetRequestHandler).Methods("GET")
	router.HandleFunc("/post", PostRequestHandler).Methods("POST")
	router.HandleFunc("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")
	http.ListenAndServe(HOST+":"+PORT, router)
}
