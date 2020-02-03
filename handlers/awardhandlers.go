package handlers

import (
	"../models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAwards(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awards, err := models.GetAwards(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}

func CreateAward(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awards, err := models.CreateAward(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}

func GetAward(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awards, err := models.GetAward(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}

func DeleteAward(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awards, err := models.DeleteAward(db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}
