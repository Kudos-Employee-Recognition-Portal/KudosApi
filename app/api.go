package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func ApiInfo(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		fmt.Fprintf(w, "Hello from the API!")
	})
}
