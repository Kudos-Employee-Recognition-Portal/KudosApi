package routes

import (
	"database/sql"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/handlers"
	"github.com/gorilla/mux"
)

// All subroutes prefixed by [HOST].com/awards
func AwardsRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/regions", handlers.GetRegions(db)).Methods("GET")
	r.Handle("/", handlers.CreateAward(db)).Methods("POST")
	r.Handle("/{id:[0-9]+}", handlers.DeleteAward(db)).Methods("DELETE")
	// Creating named route for parameter search not the most idiomatic, but
	//	much easier due to the number of possible parameter combinations.
	r.Handle("/search", handlers.QueryAwards(db)).Methods("GET")
	r.Handle("/", handlers.GetAwards(db)).Methods("GET")
	r.Handle("/{id:[0-9]+}", handlers.GetAward(db)).Methods("GET")
}
