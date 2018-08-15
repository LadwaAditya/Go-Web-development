package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	}, Route{
		"getEmployee",
		"GET",
		"/employee/{id}",
		getEmployee,
	}, Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	}, Route{
		"updateEmployee",
		"PUT",
		"/employee/update",
		updateEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Print("error getting requested employee :: ", err)
			}
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employeedata :: ", err)
		return
	}
	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id,
		employee.FirstName, employee.LastName)
	employees = append(employees, Employee{Id: employee.Id,
		FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	var isUpsert = true
	for idx, emp := range employees {
		if emp.Id == employee.Id {
			isUpsert = false
			log.Printf("updating employee id :: %s with firstName as :: %s and lastName as:: %s ",
				employee.Id, employee.FirstName, employee.LastName)
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			break
		}
	}
	if isUpsert {
		log.Printf("upserting employee id :: %s with firstName as :: %s and lastName as:: %s ",
			employee.Id, employee.FirstName, employee.LastName)
		employees = append(employees, Employee{Id: employee.Id,
			FirstName: employee.FirstName, LastName: employee.LastName})
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(HOST+":"+PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}