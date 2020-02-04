package handlers

import (
	"../models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

// TODO: Award search query handler.

func GetManagerAwards(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		manager := models.User{ID: id}
		awards, err := manager.GetManagerAwards(db)
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var award models.Award
		err := json.NewDecoder(r.Body).Decode(&award)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = award.CreateAward(db)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(award)
	})
}

func GetAward(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		award := models.Award{ID: id}
		err := award.GetAward(db)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(award)
	})
}

func DeleteAward(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		award := models.Award{ID: id}
		err = award.DeleteAward(db)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	})
}
