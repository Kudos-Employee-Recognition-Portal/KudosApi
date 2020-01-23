package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config := GetConfig();

	api = &api.Api{};

	api.Initialize(config);
	api.Run(":8080");
}