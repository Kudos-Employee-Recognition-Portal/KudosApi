package routes

import (
	"../handlers"
	"database/sql"
	"github.com/gorilla/mux"
)

func UsersRouter(r *mux.Router, db *sql.DB) {
	r.StrictSlash(true)
	r.Handle("/", handlers.GetUsers(db)).Methods("GET")
	//r.Handle("/managers", GetManagers(db)).Methods("GET")
	//r.Handle("/managers", CreateManager(db)).Methods("POST")
	//r.Handle("/managers/{id}", GetManager(db)).Methods("GET")
	//r.Handle("/managers/{id}", UpdateManager(db)).Methods("PUT")
	//r.Handle("/managers/{id}", DeleteManager(db)).Methods("DELETE")
	//r.Handle("/admins", GetAdmins(db)).Methods("GET")
	//r.Handle("/admins", CreateAdmin(db)).Methods("POST")
	//r.Handle("/admins/{id}", GetAdmin(db)).Methods("GET")
	//r.Handle("/admins/{id}", UpdateAdmin(db)).Methods("PUT")
	//r.Handle("/admins/{id}", DeleteAdmin(db)).Methods("DELETE")
}
