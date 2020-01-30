package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AwardsRouter(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", GetAwards).Methods("GET")
}

func GetAwards(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetUsers")
}
