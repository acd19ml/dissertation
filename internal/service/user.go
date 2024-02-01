package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// retrieve all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// retrieve a specific user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User
	err := db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)", user.FirstName, user.LastName, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user.ID = int(lastInsertID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// update an existing user by ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, id)
	if err != nil {
		log.Fatal(err)
	}

	user.ID, _ = strconv.Atoi(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// delete a user by ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	// Check if the user with the given ID exists
	_, err := getUserByID(userID)
	if err != nil {
		// If the user doesn't exist, return a 404 Not Found error
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with ID %s not found", userID)
		return
	}

	// Delete the user with the given ID
	err = deleteUserByID(userID)
	if err != nil {
		// If there was an error deleting the user, return a 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error deleting user")
		return
	}

	// If the user was deleted successfully, return a 204 No Content response
	w.WriteHeader(http.StatusNoContent)
}

func getUserByID(id string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func deleteUserByID(id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
