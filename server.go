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

var users []User

// 	{ID: "1", Username: "alice", Email: "alice@example.com"},
// 	{ID: "2", Username: "bob", Email: "bob@example.com"},
// 	{ID: "3", Username: "charlie", Email: "charlie@example.com"},
// }

func main() {
	populateMockData()

	router := mux.NewRouter()

	router.HandleFunc("/", handler)

	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	port := ":8080"

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func populateMockData() {
	users = []User{
		{ID: "1", Username: "nico", Email: "nico@example.com"},
		{ID: "2", Username: "benny", Email: "benny@example.com"},
		{ID: "3", Username: "charlie", Email: "charlie@example.com"},
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
	log.Printf("New User %s created", newUser.Username)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
	log.Print(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			log.Printf("User %s deleted.", id)
			return
		}
	}

	http.NotFound(w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedUser struct {
		Username *string `json:"username,omitempty"`
		Email    *string `json:"email,omitempty"`
	}

	for i, user := range users {
		if user.ID == id {
			if updatedUser.Username != nil {
				user.Username = *updatedUser.Username
			}

			if updatedUser.Email != nil {
				user.Email = *updatedUser.Email
			}

			users[i] = user
			w.WriteHeader(http.StatusNoContent)
			log.Printf("User with id %s has been updated", id)
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	i := 0
	for i < 1000 {
		fmt.Fprint(w, "Hello, World! ")
		fmt.Fprint(w, "Big boy. ")
		// fmt.Println("Hello, World! ")
		// fmt.Println("Big boy. ")

		i += 1
	}
}
