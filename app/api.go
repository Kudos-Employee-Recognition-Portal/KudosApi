package app

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Use this route to list paths, objects, properties, and query parameters for the api.
func ApiInfo(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		fmt.Fprintf(w, "Hello from Kudos API!")
	})
}
