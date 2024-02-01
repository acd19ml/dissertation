package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/acd19ml/dissertation/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("root")
	dbPassword := os.Getenv("020312")
	dbHost := os.Getenv("127.0.0.1")
	dbPort := os.Getenv("9090")
	dbName := os.Getenv("MYSQL_Go")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service.SetDB(db)

	// Verify the database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Initialize the router
	r := mux.NewRouter()

	// Define your routes here, for example:
	r.HandleFunc("/users", service.getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":9090", r))

}
