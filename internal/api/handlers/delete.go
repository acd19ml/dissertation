package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
