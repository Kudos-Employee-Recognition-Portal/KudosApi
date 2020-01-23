package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
)

type App struct {
	Router	*mux.Router
	DB		*sql.DB
}

func (app *App) Init(user, password, dbname string) {}

func (app *App) Run(addr string) {}

