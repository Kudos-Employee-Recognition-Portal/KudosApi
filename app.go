package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
	"log"
)

type App struct {
	Router	*mux.Router
	DB		*sql.DB
}

func (app *App) Init(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error

	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}

func (app *App) Run(addr string) {

}

