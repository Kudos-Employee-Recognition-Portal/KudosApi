package routes

import (
	"../handlers"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AwardsRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/", handlers.GetAwards(db)).Methods("GET")
	r.Handle("/", handlers.CreateAward(db)).Methods("POST")
	r.Handle("/{id}", handlers.GetAward(db)).Methods("GET")
	r.Handle("/{id}", handlers.DeleteAward(db)).Methods("DELETE")
}

func GetAwards(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetUsers")
}
