package main

import (
	"log"
	"net/http"

	"github.com/acd19ml/dissertation/routes"

	"github.com/acd19ml/dissertation/configs"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	log.Fatal(http.ListenAndServe(":6000", router))
}
