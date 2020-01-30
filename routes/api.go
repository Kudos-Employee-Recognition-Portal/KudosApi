package routes

import (
	"../models"
	"database/sql"
	"fmt"
	"net/http"
)

func ApiInfo(db *sql.DB) http.Handler {
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
