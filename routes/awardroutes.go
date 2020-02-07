package routes

import (
	"database/sql"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/handlers"
	"github.com/gorilla/mux"
)

func AwardsRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/", handlers.CreateAward(db)).Methods("POST")
	r.Handle("/{id}", handlers.DeleteAward(db)).Methods("DELETE")
	r.Handle("/", handlers.GetAwards(db)).Methods("GET")
	r.Handle("/{id}", handlers.GetAward(db)).Methods("GET")

	// TODO: Query handling route.
}
