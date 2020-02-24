package routes

import (
	"database/sql"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/handlers"
	"github.com/gorilla/mux"
)

// TODO: add path variable matching.
func UsersRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/managers", handlers.GetManagers(db)).Methods("GET")
	r.Handle("/managers", handlers.CreateManager(db)).Methods("POST")
	r.Handle("/managers/{id}", handlers.GetManager(db)).Methods("GET")
	r.Handle("/managers/{id}", handlers.UpdateManager(db)).Methods("PUT")
	r.Handle("/managers/{id}", handlers.DeleteUser(db)).Methods("DELETE")
	r.Handle("/managers/{id}/signature", handlers.SetManagerSignature(db)).Methods("POST")
	r.Handle("/managers/{id}/awards", handlers.GetManagerAwards(db)).Methods("GET")
	r.Handle("/admins", handlers.GetAdmins(db)).Methods("GET")
	r.Handle("/admins", handlers.CreateAdmin(db)).Methods("POST")
	r.Handle("/admins/{id}", handlers.GetAdmin(db)).Methods("GET")
	r.Handle("/admins/{id}", handlers.UpdateAdmin(db)).Methods("PUT")
	r.Handle("/admins/{id}", handlers.DeleteUser(db)).Methods("DELETE")
	r.Handle("/", handlers.GetUsers(db)).Methods("GET")
	r.Handle("/{email}", handlers.GetUser(db)).Methods("GET")
}
