package routes

import (
	"../handlers"
	"database/sql"
	"github.com/gorilla/mux"
)

func UsersRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/managers", handlers.GetManagers(db)).Methods("GET")
	r.Handle("/managers", handlers.CreateManager(db)).Methods("POST")
	r.Handle("/managers/{id}", handlers.GetManager(db)).Methods("GET")
	r.Handle("/managers/{id}", handlers.UpdateManager(db)).Methods("PUT")
	r.Handle("/managers/{id}", handlers.DeleteManager(db)).Methods("DELETE")
	r.Handle("/admins", handlers.GetAdmins(db)).Methods("GET")
	r.Handle("/admins", handlers.CreateAdmin(db)).Methods("POST")
	r.Handle("/admins/{id}", handlers.GetAdmin(db)).Methods("GET")
	r.Handle("/admins/{id}", handlers.UpdateAdmin(db)).Methods("PUT")
	r.Handle("/admins/{id}", handlers.DeleteAdmin(db)).Methods("DELETE")
	r.Handle("/", handlers.GetUsers(db)).Methods("GET")
	r.Handle("/{name}", handlers.GetUser(db)).Methods("GET")
}
