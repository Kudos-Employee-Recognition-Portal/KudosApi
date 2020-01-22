package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routesRouter)

	log.Fatal(http.ListenAndServe(":8080", r))
}