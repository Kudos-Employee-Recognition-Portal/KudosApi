package app

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Base test route should write simple welcome response string.
func ApiInfo(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		if _, err := fmt.Fprintf(w, "Hello from Kudos API!"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
