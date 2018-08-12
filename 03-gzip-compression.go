package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, handlers.CompressHandler(mux)))
}
