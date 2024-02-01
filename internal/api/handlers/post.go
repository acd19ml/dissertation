package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/acd19ml/dissertation/pkg/models"
)

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
