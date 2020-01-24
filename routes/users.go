package routes

import (
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetUsers")
}

func GetManagers(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetManagers")
}

func CreateManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: CreateManager")
}

func GetManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetManager")
}

func UpdateManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: UpdateManager")
}

func DeleteManager(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: DeleteManager")
}

func GetAdmins(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetAdmins")
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: CreateAdmin")
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: GetAdmin")
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: UpdateAdmin")
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit: DeleteAdmin")
}