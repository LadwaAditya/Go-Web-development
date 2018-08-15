package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized to view this page", http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Home page")
	} else {
		http.Error(w, "You are unauthorized to view this page", http.StatusForbidden)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "You have successfully logged in.")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "You have sucessfully logged out.")
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("Error starting server : ", err)
		return
	}
}
