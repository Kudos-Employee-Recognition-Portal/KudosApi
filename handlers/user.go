package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// Return the handler as a closure.
func GetUsers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			return
		}
		defer rows.Close()

		users := make([]*User, 0)
		for rows.Next() {
			user := new(User)
			err := rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				return
			}
			users = append(users, user)
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
