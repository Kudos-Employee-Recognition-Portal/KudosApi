package app

import (
	"database/sql"
	"fmt"
	"net/http"
)

func ApiInfo(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the API!")
	})
}
