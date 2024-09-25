package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: "1", Username: "alice", Email: "alice@example.com"},
	{ID: "2", Username: "bob", Email: "bob@example.com"},
	{ID: "3", Username: "charlie", Email: "charlie@example.com"},
}

// 	{
// 		ID: "1",
// 		Username:"nico",
// 		Email: "nico@dann.com"
// 	}

// }

func main() {
	// populateMockData()

	router := mux.NewRouter()

	router.HandleFunc("/", handler)

	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func populateMockData() {
	users = []User{
		{ID: "4", Username: "nico", Email: "nico@example.com"},
		{ID: "5", Username: "benny", Email: "benny@example.com"},
		{ID: "6", Username: "charlie", Email: "charlie@example.com"},
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedUser User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
