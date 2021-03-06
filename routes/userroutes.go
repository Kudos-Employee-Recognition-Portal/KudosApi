package routes

import (
	"database/sql"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/handlers"
	"github.com/gorilla/mux"
)

// All subroutes prefixed by [HOST]/users
func UsersRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/managers", handlers.GetManagers(db)).Methods("GET")
	r.Handle("/managers", handlers.CreateManager(db)).Methods("POST")
	r.Handle("/managers/{id:[0-9]+}", handlers.GetManager(db)).Methods("GET")
	r.Handle("/managers/{id:[0-9]+}", handlers.UpdateManager(db)).Methods("PUT")
	r.Handle("/managers/{id:[0-9]+}", handlers.DeleteUser(db)).Methods("DELETE")
	r.Handle("/managers/{id:[0-9]+}/signature", handlers.SetManagerSignature(db)).Methods("POST")
	r.Handle("/managers/{id:[0-9]+}/awards", handlers.GetManagerAwards(db)).Methods("GET")
	r.Handle("/admins", handlers.GetAdmins(db)).Methods("GET")
	r.Handle("/admins", handlers.CreateAdmin(db)).Methods("POST")
	r.Handle("/admins/{id:[0-9]+}", handlers.GetAdmin(db)).Methods("GET")
	r.Handle("/admins/{id:[0-9]+}", handlers.UpdateAdmin(db)).Methods("PUT")
	r.Handle("/admins/{id:[0-9]+}", handlers.DeleteUser(db)).Methods("DELETE")
	r.Handle("/", handlers.GetUsers(db)).Methods("GET")
	r.Handle("/{email}", handlers.GetUser(db)).Methods("GET")
}
