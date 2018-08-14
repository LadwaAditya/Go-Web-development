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

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "1", Name: "Aditya"}
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occured while executing the template or writing input", err)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}

}
