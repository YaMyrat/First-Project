package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json: "id"`
	Username string `json: "username"`
	Email    string `json: "email"`
}

var users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUsers).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
