package routes

import (
	"fmt"
	"net/http"
)

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's how to use the api.")
	fmt.Println("/ endpoint hit")
}
