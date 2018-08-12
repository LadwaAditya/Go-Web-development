package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, nil))
}
