package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Customer struct {
	Id              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DateOfBirth     string `json:"date_of_birth"`
	LocationOfBirth string `json:"location_of_birth"`
}

type Passport struct {
	Id           string `json:"id"`
	DateOfIssue  string `json:"date_of_issue"`
	DateOfExpiry string `json:"date_of_expiry"`
	Authority    string `json:"authority"`
	CustomerId   int    `json:"customer_id"`
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Nothing to see here. #kthxbai")
}

func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HandleHealthchecks")
}

func UsersHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Handling Users")
}

func PassportsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Handling Passports")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/healthcheck", HealthcheckHandler).Methods("GET")

	router.HandleFunc("/users", UsersHandler).Methods("GET")
	router.HandleFunc("/users/{uid}", UsersHandler).Methods("GET")
	router.HandleFunc("/users", UsersHandler).Methods("POST")
	router.HandleFunc("/users/{uid}", UsersHandler).Methods("PUT")
	router.HandleFunc("/users/{uid}", UsersHandler).Methods("DELETE")

	router.HandleFunc("/users/{uid}/passports", PassportsHandler).Methods("GET")
	router.HandleFunc("/passports/{pid}", PassportsHandler).Methods("GET")
	router.HandleFunc("/users/{uid}/passports", PassportsHandler).Methods("POST")
	router.HandleFunc("/passports/{pid}", PassportsHandler).Methods("PUT")
	router.HandleFunc("/passports/{pid}", PassportsHandler).Methods("DELETE")

	n := negroni.Classic()
	n.UseHandler(router)

	fmt.Println("Starting server on :3009")
	n.Run(":3009")
}
