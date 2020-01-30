package routes

import (
	"../models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UsersRouter(r *mux.Router, db *sql.DB) {
	r.Handle("/", GetUsers(db)).Methods("GET")
	r.HandleFunc("/managers", GetManagers).Methods("GET")
	r.HandleFunc("/managers", CreateManager).Methods("POST")
	r.HandleFunc("/managers/{id}", GetManager).Methods("GET")
	r.HandleFunc("/managers/{id}", UpdateManager).Methods("PUT")
	r.HandleFunc("/managers/{id}", DeleteManager).Methods("DELETE")
	r.HandleFunc("/admins", GetAdmins).Methods("GET")
	r.HandleFunc("/admins", CreateAdmin).Methods("POST")
	r.HandleFunc("/admins/{id}", GetAdmin).Methods("GET")
	r.HandleFunc("/admins/{id}", UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admins/{id}", DeleteAdmin).Methods("DELETE")
}

func GetUsers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		users, err := models.GetAllUsers(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, user := range users {
			fmt.Fprintf(w, "%s, %s, %s\n", user.ID, user.Name, user.Age)
		}
	})
}

func GetManagers(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetManagers")
}

func CreateManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: CreateManager")
}

func GetManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetManager")
}

func UpdateManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: UpdateManager")
}

func DeleteManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: DeleteManager")
}

func GetAdmins(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetAdmins")
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: CreateAdmin")
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetAdmin")
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: UpdateAdmin")
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: DeleteAdmin")
}
