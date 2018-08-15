package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/gorilla/schema"
)

const (
	HOST                   = "localhost"
	PORT                   = "8080"
	USERNAME_ERROR_MESSAGE = "Please enter a valid user name"
	PASSWORD_ERROR_MESSAGE = "Plesae enter a valid password"
	GENERIC_ERROR_MESSAGE  = "Validation error"
)

type User struct {
	Username string `valid:"alpha,required"`
	Password string `valid:"alpha,required"`
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("Error mapping parsed data to struct :", decodeErr)
	}
	return user
}

func validateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Printf("username validation error : ", usernameError)
			return valid, USERNAME_ERROR_MESSAGE
		}
		if passwordError != "" {
			log.Printf("password validation error : ", passwordError)
			return valid, PASSWORD_ERROR_MESSAGE
		}
	}
	return valid, GENERIC_ERROR_MESSAGE
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		parsedTemplate.Execute(w, nil)
	} else {
		user := readForm(r)
		valid, vaildationErrorMessage := validateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, vaildationErrorMessage)
			return
		}
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("Error starting up http server : ", err)
		return
	}
}
