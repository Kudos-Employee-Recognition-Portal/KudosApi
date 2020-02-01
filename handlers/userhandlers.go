package handlers

import (
	"../models"
	"database/sql"
	"encoding/json"
	"net/http"
)

//Return the handler as a closure.
func GetUsers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetUsers(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
}

//
//func GetManagers(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetManagers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func CreateManager(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func GetManager(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func UpdateManager(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func DeleteManager(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func GetAdmins(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func CreateAdmin(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func GetAdmin(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func UpdateAdmin(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
//
//func DeleteAdmin(db *sql.DB) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, http.StatusText(405), 405)
//			return
//		}
//		users, err := handlers.GetUsers(db)
//		if err != nil {
//			http.Error(w, http.StatusText(500), 500)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(users)
//	})
//}
