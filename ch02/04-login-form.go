package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
	parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}

}
